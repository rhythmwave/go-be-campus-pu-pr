-- +goose Up
-- +goose StatementBegin

ALTER TABLE "schedules"
  ADD COLUMN "is_midterm_exam" boolean NULL,
  ADD COLUMN "is_endterm_exam" boolean NULL,
  DROP CONSTRAINT "schedules_check1",
  ADD CONSTRAINT "schedules_check1" CHECK(("is_exam" IS true AND num_nonnulls("is_theory_exam", "is_practicum_exam", "is_field_practicum_exam") = 1 AND num_nonnulls("is_midterm_exam", "is_endterm_exam") = 1) OR ("is_exam" IS false AND num_nonnulls("is_theory_exam", "is_practicum_exam", "is_field_practicum_exam") = 0 AND num_nonnulls("is_midterm_exam", "is_endterm_exam") = 0)),
  ADD CONSTRAINT "schedules_unique_exam" UNIQUE("class_id", "is_midterm_exam", "is_endterm_exam");

ALTER TABLE "deleted_schedules"
  ADD COLUMN "is_midterm_exam" boolean NULL,
  ADD COLUMN "is_endterm_exam" boolean NULL;

CREATE OR REPLACE FUNCTION schedules_conflict_func()
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

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

CREATE OR REPLACE FUNCTION schedules_conflict_func()
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

ALTER TABLE "schedules"
  DROP COLUMN "is_midterm_exam",
  DROP COLUMN "is_endterm_exam",
  ADD CONSTRAINT "schedules_check1" CHECK(("is_exam" IS true AND num_nonnulls("is_theory_exam", "is_practicum_exam", "is_field_practicum_exam") = 1) OR ("is_exam" IS false AND num_nonnulls("is_theory_exam", "is_practicum_exam", "is_field_practicum_exam") = 0));

ALTER TABLE "deleted_schedules"
  ADD COLUMN "is_midterm_exam" boolean NULL,
  ADD COLUMN "is_endterm_exam" boolean NULL;


-- +goose StatementEnd
