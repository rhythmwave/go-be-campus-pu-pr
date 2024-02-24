-- +goose Up
-- +goose StatementBegin

ALTER TABLE "lectures" ADD COLUMN "thesis_defense_id" uuid NULL REFERENCES "thesis_defenses"("id") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "deleted_lectures" ADD COLUMN "thesis_defense_id" uuid NULL;

ALTER TABLE "thesis_defenses" ADD COLUMN "thesis_defense_count" integer NOT NULL;
ALTER TABLE "deleted_thesis_defenses" ADD COLUMN "thesis_defense_count" integer NULL;

CREATE TABLE "thesis_defense_requests" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "thesis_id" uuid NOT NULL UNIQUE REFERENCES "theses" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "thesis_defense_id" uuid NULL REFERENCES "thesis_defenses" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "thesis_defense_count" integer NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_thesis_defense_requests" BEFORE UPDATE ON "thesis_defense_requests" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

------------------------------------------------------------------------------------------

CREATE FUNCTION thesis_defense_request_predefined_column_func()
RETURNS TRIGGER AS $$
DECLARE currentDefenseCount integer;
BEGIN
  IF OLD.id IS NULL OR (OLD.thesis_defense_id IS NOT NULL AND NEW.thesis_defense_id IS NULL) THEN
    SELECT thesis_defense_count INTO currentDefenseCount FROM theses WHERE id = NEW.thesis_id;

    NEW.thesis_defense_count = currentDefenseCount + 1;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER thesis_defense_request_predefined_column BEFORE INSERT OR UPDATE OF thesis_defense_id ON thesis_defense_requests FOR EACH ROW EXECUTE PROCEDURE thesis_defense_request_predefined_column_func();

------------------------------------------------------------------------------------------

CREATE FUNCTION thesis_defense_predefined_column_func()
RETURNS TRIGGER AS $$
DECLARE currentDefenseCount integer;
BEGIN
  SELECT thesis_defense_count INTO currentDefenseCount FROM theses WHERE id = NEW.thesis_id;
  NEW.thesis_defense_count = currentDefenseCount + 1;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER thesis_defense_predefined_column BEFORE INSERT ON thesis_defenses FOR EACH ROW EXECUTE PROCEDURE thesis_defense_predefined_column_func();


------------------------------------------------------------------------------------------

CREATE FUNCTION insert_thesis_defense_schedule_func()
RETURNS TRIGGER AS $$
DECLARE 
  activeSemesterId uuid;
  classId uuid;
  studentId uuid;
BEGIN
    SELECT id INTO activeSemesterId FROM semesters WHERE is_active IS true;

    SELECT student_id INTO studentId FROM theses WHERE id = NEW.thesis_id;

    SELECT c.id INTO classId
    FROM classes c
    JOIN student_classes sc ON sc.class_id = c.id AND sc.student_id = studentId
    JOIN subjects s ON s.id = c.subject_id
    WHERE semester_id = activeSemesterId AND s.is_thesis IS true;

    INSERT INTO lectures (
			class_id,
			room_id,
			lecture_plan_date,
			lecture_plan_start_time,
			lecture_plan_end_time,
      thesis_defense_id
		) VALUES (
			classId,
			NEW.room_id,
			NEW.plan_date,
			NEW.plan_start_time,
			NEW.plan_end_time,
      NEW.id
		);

    UPDATE thesis_defense_requests SET thesis_defense_id = NEW.id WHERE thesis_id = NEW.thesis_id;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER insert_thesis_defense_schedule AFTER INSERT ON thesis_defenses FOR EACH ROW EXECUTE PROCEDURE insert_thesis_defense_schedule_func();

------------------------------------------------------------------------------------------

CREATE FUNCTION update_thesis_defense_schedule_func()
RETURNS TRIGGER AS $$
BEGIN
  UPDATE lectures SET 
    room_id = NEW.room_id,
    lecture_plan_date = NEW.plan_date,
    lecture_plan_start_time = NEW.plan_start_time,
    lecture_plan_end_time = NEW.plan_end_time
  WHERE thesis_defense_id = NEW.id;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER update_thesis_defense_schedule AFTER UPDATE ON thesis_defenses FOR EACH ROW EXECUTE PROCEDURE update_thesis_defense_schedule_func();

------------------------------------------------------------------------------------------

CREATE FUNCTION update_actual_thesis_defense_func()
RETURNS TRIGGER AS $$
BEGIN
  IF OLD.actual_date IS NULL AND NEW.actual_date IS NOT NULL THEN
    UPDATE lectures SET 
      room_id = NEW.room_id,
      lecture_actual_date = NEW.actual_date,
      lecture_actual_start_time = NEW.actual_start_time,
      lecture_actual_end_time = NEW.actual_end_time
    WHERE thesis_defense_id = NEW.id;

    UPDATE theses SET thesis_defense_count = thesis_defense_count + 1 WHERE id = NEW.thesis_id;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER update_actual_thesis_defense AFTER UPDATE OF actual_date ON thesis_defenses FOR EACH ROW EXECUTE PROCEDURE update_actual_thesis_defense_func();

------------------------------------------------------------------------------------------

CREATE FUNCTION delete_thesis_defense_schedule_func()
RETURNS TRIGGER AS $$
BEGIN
  DELETE FROM lectures WHERE thesis_defense_id = NEW.id;

  RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER delete_thesis_defense_schedule BEFORE DELETE ON thesis_defenses FOR EACH ROW EXECUTE PROCEDURE delete_thesis_defense_schedule_func();

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

  IF NEW.thesis_defense_id IS NULL THEN
    IF (NEW.lecture_plan_date BETWEEN semesterMidtermStartDate AND semesterMidtermEndDate OR NEW.lecture_plan_date BETWEEN semesterEndtermStartDate AND semesterEndtermEndDate) AND NEW.is_exam IS false THEN
      RAISE EXCEPTION 'cannot add non-exam schedule in the exam period';
    ELSIF ((NEW.is_midterm_exam IS true AND NEW.lecture_plan_date NOT BETWEEN semesterMidtermStartDate AND semesterMidtermEndDate) OR (NEW.is_endterm_exam IS true AND NEW.lecture_plan_date NOT BETWEEN semesterEndtermStartDate AND semesterEndtermEndDate)) AND NEW.is_exam IS true THEN
      RAISE EXCEPTION 'cannot add exam schedule outside the exam period';
    END IF;
  END IF;

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

CREATE OR REPLACE FUNCTION delete_lecture_total_plan_func()
RETURNS TRIGGER AS $$
BEGIN
  IF OLD.thesis_defense_id IS NULL THEN
    IF OLD.is_exam IS true THEN
      UPDATE classes SET total_exam_plan = total_exam_plan - 1
      WHERE id = OLD.class_id;
    ELSE
      UPDATE classes SET total_lecture_plan = total_lecture_plan - 1
      WHERE id = OLD.class_id;
    END IF;
  END IF;

  RETURN OLD;
END;
$$ LANGUAGE plpgsql;

------------------------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION lecture_participant_func()
RETURNS TRIGGER AS $$
BEGIN
  IF NEW.thesis_defense_id IS NULL THEN
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
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

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

CREATE OR REPLACE FUNCTION delete_lecture_total_plan_func()
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

DROP TRIGGER update_actual_thesis_defense ON thesis_defenses;
DROP FUNCTION update_actual_thesis_defense_func();

DROP TRIGGER insert_thesis_defense_schedule ON thesis_defenses;
DROP FUNCTION insert_thesis_defense_schedule_func();
DROP TRIGGER update_thesis_defense_schedule ON thesis_defenses;
DROP FUNCTION update_thesis_defense_schedule_func();
DROP TRIGGER delete_thesis_defense_schedule ON thesis_defenses;
DROP FUNCTION delete_thesis_defense_schedule_func();

DROP TRIGGER thesis_defense_predefined_column ON thesis_defenses;
DROP FUNCTION thesis_defense_predefined_column_func();

DROP TRIGGER thesis_defense_request_predefined_column ON thesis_defense_requests;
DROP FUNCTION thesis_defense_request_predefined_column_func();

DROP TABLE "thesis_defense_requests";

ALTER TABLE "thesis_defenses" DROP COLUMN "thesis_defense_count";
ALTER TABLE "deleted_thesis_defenses" DROP COLUMN "thesis_defense_count";

ALTER TABLE "lectures" DROP COLUMN "thesis_defense_id";
ALTER TABLE "deleted_lectures" DROP COLUMN "thesis_defense_id";

-- +goose StatementEnd
