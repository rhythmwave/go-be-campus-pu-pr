-- +goose Up
-- +goose StatementBegin

ALTER TABLE "classes"
  ADD COLUMN "total_exam_plan" integer NOT NULL DEFAULT 0,
  ADD COLUMN "total_exam_done" integer NOT NULL DEFAULT 0;

ALTER TABLE "deleted_classes"
  ADD COLUMN "total_exam_plan" integer NULL,
  ADD COLUMN "total_exam_done" integer NULL;

------------------------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION lectures_conflict_func()
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

  IF NEW.is_exam IS true THEN
    WITH d AS (
      SELECT COUNT(1) AS total FROM lectures WHERE class_id = NEW.class_id AND is_exam IS true
    )
    UPDATE classes c SET total_exam_plan = d.total
    FROM d
    WHERE c.id = NEW.class_id;
  ELSE
    WITH d AS (
      SELECT COUNT(1) AS total FROM lectures WHERE class_id = NEW.class_id AND is_exam IS false
    )
    UPDATE classes c SET total_lecture_plan = d.total
    FROM d
    WHERE c.id = NEW.class_id;
  END IF;

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

------------------------------------------------------------------------------------------

CREATE FUNCTION delete_lecture_total_plan_func()
RETURNS TRIGGER AS $$
BEGIN
  IF OLD.is_exam IS true THEN
    UPDATE classes SET total_exam_plan = total_exam_plan - 1
    WHERE id = OLD.class_id;
  ELSE
    UPDATE classes SET total_lecture_plan = total_lecture_plan - 1
    WHERE id = OLD.class_id;
  END IF;

  RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER delete_lecture_total_plan BEFORE DELETE ON lectures FOR EACH ROW EXECUTE PROCEDURE delete_lecture_total_plan_func();

------------------------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION lecture_participant_func()
RETURNS TRIGGER AS $$
BEGIN
  INSERT INTO lecture_participants (
    "lecture_id",
    "student_id"
  ) SELECT
    NEW.id,
    sc.student_id
  FROM student_classes sc WHERE sc.class_id = NEW.class_id
  ON CONFLICT (lecture_id, student_id) DO NOTHING;

  NEW.lecture_actual_day_of_week = EXTRACT('isodow' FROM NEW.lecture_actual_date);

  IF NEW.is_exam IS false THEN
    UPDATE classes SET total_lecture_done = total_lecture_done + 1
    WHERE id = NEW.class_id;
  ELSE
    UPDATE classes SET total_exam_done = total_exam_done + 1
    WHERE id = NEW.class_id;
  END IF;
  

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

------------------------------------------------------------------------------------------

ALTER TABLE "subjects" ADD COLUMN "total_lesson_plan" integer NOT NULL DEFAULT 0;
ALTER TABLE "deleted_subjects" ADD COLUMN "total_lesson_plan" integer NOT NULL DEFAULT 0;

------------------------------------------------------------------------------------------

CREATE FUNCTION subjects_total_lesson_plan_func()
RETURNS TRIGGER AS $$
BEGIN
  IF OLD.id IS NOT NULL AND NEW.id IS NULL THEN
    UPDATE subjects SET total_lesson_plan = total_lesson_plan - 1
    WHERE id = OLD.subject_id;

    RETURN OLD;
  ELSIF OLD.id IS NULL AND NEW.id IS NOT NULL THEN
    UPDATE subjects SET total_lesson_plan = total_lesson_plan + 1
      WHERE id = NEW.subject_id;

    RETURN NEW;
  END IF;

END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER subjects_total_lesson_plan BEFORE INSERT OR DELETE ON lesson_plans FOR EACH ROW EXECUTE PROCEDURE subjects_total_lesson_plan_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER subjects_total_lesson_plan ON lesson_plans;
DROP FUNCTION subjects_total_lesson_plan_func();

------------------------------------------------------------------------------------------

ALTER TABLE "subjects" ADD COLUMN "total_lesson_plan" integer NOT NULL DEFAULT 0;
ALTER TABLE "deleted_subjects" ADD COLUMN "total_lesson_plan" integer NOT NULL DEFAULT 0;

------------------------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION lecture_participant_func()
RETURNS TRIGGER AS $$
DECLARE totalParticipant integer;
BEGIN
  INSERT INTO lecture_participants (
    "lecture_id",
    "student_id"
  ) SELECT
    NEW.id,
    sc.student_id
  FROM student_classes sc WHERE sc.class_id = NEW.class_id
  ON CONFLICT (lecture_id, student_id) DO NOTHING;

  NEW.lecture_actual_day_of_week = EXTRACT('isodow' FROM NEW.lecture_actual_date);

  UPDATE classes SET total_lecture_done = total_lecture_done + 1
  WHERE id = NEW.class_id;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

------------------------------------------------------------------------------------------

DROP TRIGGER delete_lecture_total_plan ON lectures;
DROP FUNCTION delete_lecture_total_plan_func();

------------------------------------------------------------------------------------------


CREATE OR REPLACE FUNCTION lectures_conflict_func()
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

------------------------------------------------------------------------------------------

ALTER TABLE "classes"
  DROP COLUMN "total_exam_plan",
  DROP COLUMN "total_exam_done";

ALTER TABLE "deleted_classes"
  DROP COLUMN "total_exam_plan",
  DROP COLUMN "total_exam_done";

-- +goose StatementEnd
