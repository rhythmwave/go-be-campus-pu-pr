-- +goose Up
-- +goose StatementBegin

ALTER TABLE "semesters"
  ADD COLUMN "midterm_start_date" date NULL,
  ADD COLUMN "midterm_end_date" date NULL,
  ADD COLUMN "endterm_start_date" date NULL,
  ADD COLUMN "endterm_end_date" date NULL,
  ADD CONSTRAINT "semesters_midterm_start_date_midterm_end_date" CHECK("midterm_start_date" <= "midterm_end_date"),
  ADD CONSTRAINT "semesters_endterm_start_date_endterm_end_date" CHECK("endterm_start_date" <= "endterm_end_date"),
  ADD CONSTRAINT "semesters_midterm_start_date_end_date" CHECK("start_date" <= "midterm_start_date" AND "midterm_end_date" <= "end_date"),
  ADD CONSTRAINT "semesters_endterm_start_date_end_date" CHECK("start_date" <= "endterm_start_date" AND "endterm_end_date" <= "end_date");

ALTER TABLE "deleted_semesters"
  ADD COLUMN "midterm_start_date" date NULL,
  ADD COLUMN "midterm_end_date" date NULL,
  ADD COLUMN "endterm_start_date" date NULL,
  ADD COLUMN "endterm_end_date" date NULL;

------------------------------------------------------------------------

CREATE FUNCTION semester_conflict_date_func()
RETURNS TRIGGER AS $$
DECLARE conflictId uuid;
BEGIN
  SELECT s.id INTO conflictId
  FROM semesters s WHERE NEW.start_date <= s.end_date AND NEW.end_date >= s.start_date;

  IF conflictId IS NOT NULL THEN
    RAISE EXCEPTION 'conflicting semester date is exist.';
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER semester_conflict_date BEFORE INSERT OR UPDATE ON semesters FOR EACH ROW EXECUTE PROCEDURE semester_conflict_date_func();

------------------------------------------------------------------------

CREATE TABLE "schedules" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "class_id" uuid NOT NULL REFERENCES "classes" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "room_id" uuid NOT NULL REFERENCES "rooms" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "single_day_lecturer_id" uuid NULL REFERENCES "lecturers" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "single_day_date" date NULL,
  "day_of_week" integer NOT NULL,
  "start_time" INTEGER NOT NULL,
  "end_time" INTEGER NOT NULL,
  "theory_credit" integer NOT NULL DEFAULT 0,
  "practicum_credit" integer NOT NULL DEFAULT 0,
  "field_practicum_credit" integer NOT NULL DEFAULT 0,
  "is_regular_schedule" boolean NOT NULL,
  "is_exam" boolean NOT NULL DEFAULT false,
  "is_theory_exam" boolean NULL,
  "is_practicum_exam" boolean NULL,
  "is_field_practicum_exam" boolean NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  CHECK(num_nonnulls("single_day_date", "single_day_lecturer_id") IN (0, 2)),
  CHECK(("is_exam" IS true AND num_nonnulls("is_theory_exam", "is_practicum_exam", "is_field_practicum_exam") = 1) OR ("is_exam" IS false AND num_nonnulls("is_theory_exam", "is_practicum_exam", "is_field_practicum_exam") = 0))
);
CREATE TRIGGER "updated_at_schedules" BEFORE UPDATE ON "schedules" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_schedules" AS TABLE "schedules" WITH NO DATA;
CREATE TRIGGER "soft_delete_schedules" BEFORE DELETE ON "schedules" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

------------------------------------------------------------------------

CREATE FUNCTION schedules_conflict_func()
RETURNS TRIGGER AS $$
DECLARE 
  existId uuid;
  duplicateRoom boolean;
  duplicateClass boolean;
  duplicateLecturer boolean;
  semesterStartDate date;
  semesterEndDate date;
  semesterMidtermStartDate date;
  semesterMidtermEndDate date;
  semesterEndtermStartDate date;
  semesterEndtermEndDate date;
BEGIN
  SELECT s.start_date, s.end_date, s.midterm_start_date, s.midterm_end_date, s.endterm_start_date, s.endterm_end_date
  INTO semesterStartDate, semesterEndDate, semesterMidtermStartDate, semesterMidtermEndDate, semesterEndtermStartDate, semesterEndtermEndDate
  FROM classes c
  JOIN semesters s ON s.id = c.semester_id
  WHERE c.id = NEW.class_id;

  IF NEW.single_day_date IS NOT NULL THEN
    IF NEW.single_day_date NOT BETWEEN semesterStartDate AND semesterEndDate THEN
      RAISE EXCEPTION 'date is outside semester';
    END IF;
    IF (NEW.single_day_date BETWEEN semesterMidtermStartDate AND semesterMidtermEndDate OR NEW.single_day_date BETWEEN semesterEndtermStartDate AND semesterEndtermEndDate) AND NEW.is_exam IS false THEN
      RAISE EXCEPTION 'cannot add non-exam schedule in the exam period';
    ELSIF (NEW.single_day_date NOT BETWEEN semesterMidtermStartDate AND semesterMidtermEndDate OR NEW.single_day_date NOT BETWEEN semesterEndtermStartDate AND semesterEndtermEndDate) AND NEW.is_exam IS true THEN
      RAISE EXCEPTION 'cannot add exam schedule outside the exam period';
    END IF;

    NEW.is_regular_schedule = false;
    NEW.day_of_week = EXTRACT('isodow' FROM NEW.single_day_date);

    SELECT 
      s.id, 
      s.room_id = NEW.room_id,
      s.class_id = NEW.class_id,
      s.single_day_lecturer_id = NEW.single_day_lecturer_id
    INTO existId, duplicateRoom, duplicateClass, duplicateLecturer
    FROM schedules s
    WHERE 
      id != NEW.id AND
      (NEW.single_day_date = s.single_day_date OR (s.single_day_date IS NULL AND NEW.is_exam IS false AND NEW.day_of_week = s.day_of_week)) AND
      NEW.start_time < s.end_time AND 
      NEW.end_time > s.start_time AND 
      (s.room_id = NEW.room_id OR s.class_id = NEW.class_id OR s.single_day_lecturer_id = NEW.single_day_lecturer_id);
  ELSE
    IF NEW.is_exam IS true THEN
      RAISE EXCEPTION 'cannot add exam as regular schedule';
    END IF;

    NEW.is_regular_schedule = true;

    SELECT 
      s.id,
      s.room_id = NEW.room_id,
      s.class_id = NEW.class_id
    INTO existId, duplicateRoom, duplicateClass
    FROM schedules s
    WHERE 
      id != NEW.id AND
      NEW.day_of_week = s.day_of_week AND 
      NEW.start_time < s.end_time AND 
      NEW.end_time > s.start_time AND 
      (s.room_id = NEW.room_id OR s.class_id = NEW.class_id);
  END IF;

  IF existId IS NOT NULL THEN
    IF duplicateRoom IS true THEN
      RAISE EXCEPTION 'conflicting room schedule is exist.';
    ELSIF duplicateClass IS true THEN
      RAISE EXCEPTION 'conflicting class schedule is exist.';
    ELSIF duplicateLecturer IS true THEN
      RAISE EXCEPTION 'conflicting lecturer schedule is exist.';
    ELSE
      RAISE EXCEPTION 'conflicting schedule is exist.';
    END IF;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER schedules_conflict BEFORE INSERT OR UPDATE ON schedules FOR EACH ROW EXECUTE PROCEDURE schedules_conflict_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER schedules_conflict ON schedules;
DROP FUNCTION schedules_conflict_func();

DROP TABLE "schedules";
DROP TABLE "deleted_schedules";

DROP TRIGGER semester_conflict_date ON semesters;
DROP FUNCTION semester_conflict_date_func();

ALTER TABLE "semesters"
  DROP CONSTRAINT "semesters_midterm_start_date_midterm_end_date",
  DROP CONSTRAINT "semesters_endterm_start_date_endterm_end_date",
  DROP COLUMN "midterm_start_date",
  DROP COLUMN "midterm_end_date",
  DROP COLUMN "endterm_start_date",
  DROP COLUMN "endterm_end_date";

ALTER TABLE "deleted_semesters"
  DROP COLUMN "midterm_start_date",
  DROP COLUMN "midterm_end_date",
  DROP COLUMN "endterm_start_date",
  DROP COLUMN "endterm_end_date";

-- +goose StatementEnd
