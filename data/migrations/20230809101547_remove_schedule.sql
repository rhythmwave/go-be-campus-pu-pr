-- +goose Up
-- +goose StatementBegin

ALTER TABLE "lectures" 
  DROP COLUMN "schedule_id",
  ALTER COLUMN "room_id" SET NOT NULL;
ALTER TABLE "deleted_lectures" 
  DROP COLUMN "schedule_id";

---------------------------------------------------------------------------------

DROP TABLE "exam_schedule_supervisors";
DROP TABLE "deleted_exam_schedule_supervisors";

CREATE TABLE "exam_lecture_supervisors" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "lecture_id" uuid NOT NULL REFERENCES "lectures" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "exam_supervisor_id" uuid NOT NULL REFERENCES "exam_supervisors" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "exam_supervisor_role_id" uuid NOT NULL REFERENCES "exam_supervisor_roles" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("lecture_id", "exam_supervisor_id"),
  UNIQUE("lecture_id", "exam_supervisor_role_id")
);
CREATE TRIGGER "updated_at_exam_lecture_supervisors" BEFORE UPDATE ON "exam_lecture_supervisors" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_exam_lecture_supervisors" AS TABLE "exam_lecture_supervisors" WITH NO DATA;
CREATE TRIGGER "soft_delete_exam_lecture_supervisors" BEFORE DELETE ON "exam_lecture_supervisors" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

---------------------------------------------------------------------------------

ALTER TABLE "lectures" 
  ADD COLUMN "is_exam" boolean NOT NULL DEFAULT false,
  ADD COLUMN "is_theory_exam" boolean NULL,
  ADD COLUMN "is_practicum_exam" boolean NULL,
  ADD COLUMN "is_field_practicum_exam" boolean NULL,
  ADD COLUMN "is_midterm_exam" boolean NULL,
  ADD COLUMN "is_endterm_exam" boolean NULL,
  ALTER COLUMN "class_id" SET NOT NULL,
  ALTER COLUMN "lecture_plan_start_time" SET NOT NULL,
  ALTER COLUMN "lecture_plan_end_time" SET NOT NULL,
  ADD CONSTRAINT "lectures_exam" CHECK(("is_exam" IS true AND num_nonnulls("is_theory_exam", "is_practicum_exam", "is_field_practicum_exam") = 1 AND num_nonnulls("is_midterm_exam", "is_endterm_exam") = 1) OR ("is_exam" IS false AND num_nonnulls("is_theory_exam", "is_practicum_exam", "is_field_practicum_exam") = 0 AND num_nonnulls("is_midterm_exam", "is_endterm_exam") = 0)),
  ADD CONSTRAINT "lectures_unique_exam" UNIQUE("class_id", "is_midterm_exam", "is_endterm_exam");

ALTER TABLE "deleted_lectures" 
  ADD COLUMN "is_exam" boolean NULL,
  ADD COLUMN "is_theory_exam" boolean NULL,
  ADD COLUMN "is_practicum_exam" boolean NULL,
  ADD COLUMN "is_field_practicum_exam" boolean NULL,
  ADD COLUMN "is_midterm_exam" boolean NULL,
  ADD COLUMN "is_endterm_exam" boolean NULL;

--------------------------------------------------------------------------------

DROP TRIGGER lecture_total_plan ON lectures;
DROP FUNCTION lecture_total_plan_func();

--------------------------------------------------------------------------------

CREATE FUNCTION lectures_conflict_func()
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

  IF NEW.lecture_plan_date NOT BETWEEN semesterStartDate AND semesterEndDate THEN
    RAISE EXCEPTION 'date is outside semester';
  END IF;
  IF (NEW.lecture_plan_date BETWEEN semesterMidtermStartDate AND semesterMidtermEndDate OR NEW.lecture_plan_date BETWEEN semesterEndtermStartDate AND semesterEndtermEndDate) AND NEW.is_exam IS false THEN
    RAISE EXCEPTION 'cannot add non-exam schedule in the exam period';
  ELSIF ((NEW.is_midterm_exam IS true AND NEW.lecture_plan_date NOT BETWEEN semesterMidtermStartDate AND semesterMidtermEndDate) OR (NEW.is_endterm_exam IS true AND NEW.lecture_plan_date NOT BETWEEN semesterEndtermStartDate AND semesterEndtermEndDate)) AND NEW.is_exam IS true THEN
    RAISE EXCEPTION 'cannot add exam schedule outside the exam period';
  END IF;

  WITH d AS (
    SELECT COUNT(1) AS total FROM lectures WHERE class_id = NEW.class_id
  )
  UPDATE classes c SET total_lecture_plan = d.total
  FROM d
  WHERE c.id = NEW.class_id;
  NEW.lecture_plan_day_of_week = EXTRACT('isodow' FROM NEW.lecture_plan_date);

  SELECT 
    s.id, 
    s.room_id = NEW.room_id,
    s.class_id = NEW.class_id,
    s.lecturer_id = NEW.lecturer_id
  INTO existId, duplicateRoom, duplicateClass, duplicateLecturer
  FROM lectures s
  WHERE 
    id != NEW.id AND
    NEW.lecture_plan_date = s.lecture_plan_date AND
    NEW.lecture_plan_start_time < s.lecture_plan_end_time AND 
    NEW.lecture_plan_end_time > s.lecture_plan_start_time AND 
    (s.room_id = NEW.room_id OR s.class_id = NEW.class_id OR s.lecturer_id::text = COALESCE(NEW.lecturer_id::text, ''));

  IF existId IS NOT NULL THEN
    IF duplicateRoom IS true THEN
      RAISE EXCEPTION 'conflicting room schedule is exist.';
    ELSIF duplicateClass IS true THEN
      RAISE EXCEPTION 'conflicting class schedule is exist.';
    ELSIF duplicateLecturer IS true AND NEW.lecturer_id IS NOT NULL THEN
      RAISE EXCEPTION 'conflicting lecturer schedule is exist.';
    ELSE
      RAISE EXCEPTION 'conflicting schedule is exist.';
    END IF;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER lectures_conflict BEFORE INSERT OR UPDATE OF lecture_plan_date ON lectures FOR EACH ROW EXECUTE PROCEDURE lectures_conflict_func();

------------------------------------------------------------------------

DROP TRIGGER schedules_conflict ON schedules;
DROP FUNCTION schedules_conflict_func();

DROP TABLE "schedules";
DROP TABLE "deleted_schedules";

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

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
  "is_midterm_exam" boolean NULL,
  "is_endterm_exam" boolean NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  CHECK(num_nonnulls("single_day_date", "single_day_lecturer_id") IN (0, 2)),
  CHECK(("is_exam" IS true AND num_nonnulls("is_theory_exam", "is_practicum_exam", "is_field_practicum_exam") = 1 AND num_nonnulls("is_midterm_exam", "is_endterm_exam") = 1) OR ("is_exam" IS false AND num_nonnulls("is_theory_exam", "is_practicum_exam", "is_field_practicum_exam") = 0 AND num_nonnulls("is_midterm_exam", "is_endterm_exam") = 0)),
  UNIQUE("class_id", "is_midterm_exam", "is_endterm_exam")
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
    ELSIF ((NEW.is_midterm_exam IS true AND NEW.single_day_date NOT BETWEEN semesterMidtermStartDate AND semesterMidtermEndDate) OR (NEW.is_endterm_exam IS true AND NEW.single_day_date NOT BETWEEN semesterEndtermStartDate AND semesterEndtermEndDate)) AND NEW.is_exam IS true THEN
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

---------------------------------------------------------------------------------------

DROP TRIGGER lectures_conflict ON lectures;
DROP FUNCTION lectures_conflict_func();

---------------------------------------------------------------------------------------

CREATE FUNCTION lecture_total_plan_func()
RETURNS TRIGGER AS $$
DECLARE totalParticipant integer;
BEGIN
  UPDATE classes SET total_lecture_plan = total_lecture_plan + 1
  WHERE id = NEW.class_id;

  NEW.lecture_plan_day_of_week = EXTRACT('isodow' FROM NEW.lecture_plan_date);

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER lecture_total_plan BEFORE INSERT ON lectures FOR EACH ROW EXECUTE PROCEDURE lecture_total_plan_func();

--------------------------------------------------------------------------------------

ALTER TABLE "lectures" 
  DROP COLUMN "is_exam",
  DROP COLUMN "is_theory_exam",
  DROP COLUMN "is_practicum_exam",
  DROP COLUMN "is_field_practicum_exam",
  DROP COLUMN "is_midterm_exam",
  DROP COLUMN "is_endterm_exam";

ALTER TABLE "deleted_lectures" 
  DROP COLUMN "is_exam",
  DROP COLUMN "is_theory_exam",
  DROP COLUMN "is_practicum_exam",
  DROP COLUMN "is_field_practicum_exam",
  DROP COLUMN "is_midterm_exam",
  DROP COLUMN "is_endterm_exam";

---------------------------------------------------------------------------------

DROP TABLE "exam_lecture_supervisors";
DROP TABLE "deleted_exam_lecture_supervisors";

CREATE TABLE "exam_schedule_supervisors" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "schedule_id" uuid NOT NULL REFERENCES "schedules" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "exam_supervisor_id" uuid NOT NULL REFERENCES "exam_supervisors" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "exam_supervisor_role_id" uuid NOT NULL REFERENCES "exam_supervisor_roles" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("schedule_id", "exam_supervisor_id"),
  UNIQUE("schedule_id", "exam_supervisor_role_id")
);
CREATE TRIGGER "updated_at_exam_schedule_supervisors" BEFORE UPDATE ON "exam_schedule_supervisors" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_exam_schedule_supervisors" AS TABLE "exam_schedule_supervisors" WITH NO DATA;
CREATE TRIGGER "soft_delete_exam_schedule_supervisors" BEFORE DELETE ON "exam_schedule_supervisors" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

---------------------------------------------------------------------------------

ALTER TABLE "lectures" 
  ADD COLUMN "schedule_id" uuid NULL REFERENCES "schedules" ("id") ON DELETE SET NULL ON UPDATE CASCADE,
  ALTER COLUMN "room_id" DROP NOT NULL;
ALTER TABLE "deleted_lectures" 
  ADD COLUMN "schedule_id" uuid NULL;

-- +goose StatementEnd
