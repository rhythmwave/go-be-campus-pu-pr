-- +goose Up
-- +goose StatementBegin

ALTER TABLE "yudicium_sessions" 
  ADD COLUMN "actual_date" date NULL,
  ADD COLUMN "yudicium_number" character varying NULL UNIQUE;

ALTER TABLE "deleted_yudicium_sessions" 
  ADD COLUMN "actual_date" date NULL,
  ADD COLUMN "yudicium_number" character varying NULL;

-------------------------------------------------

CREATE TABLE yudicium_students (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "student_id" uuid NOT NULL REFERENCES "students" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "application_date" date NOT NULL,
  "with_thesis" boolean NOT NULL,
  "done_yudicium" boolean NOT NULL DEFAULT false,
  "yudicium_session_id" uuid NULL REFERENCES "yudicium_sessions" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("student_id")
);
CREATE TRIGGER "updated_at_yudicium_students" BEFORE UPDATE ON "yudicium_students" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

-------------------------------------------------

ALTER TABLE "students"
  ADD COLUMN "done_yudicium" boolean NOT NULL DEFAULT false,
  ADD COLUMN "diploma_number" character varying NULL,
  ADD COLUMN "graduation_date" date NULL;
ALTER TABLE "deleted_students"
  ADD COLUMN "done_yudicium" boolean NULL,
  ADD COLUMN "diploma_number" character varying NULL,
  ADD COLUMN "graduation_date" date NULL;

-------------------------------------------------

CREATE FUNCTION student_done_yudicium_func()
RETURNS TRIGGER AS $$
BEGIN
  IF OLD.yudicium_session_id IS NULL AND NEW.yudicium_session_id IS NOT NULL THEN
    NEW.done_yudicium = true;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER student_done_yudicium BEFORE UPDATE OF yudicium_session_id ON yudicium_students FOR EACH ROW EXECUTE PROCEDURE student_done_yudicium_func();

-------------------------------------------------

DROP TRIGGER student_graduation_predicate ON students;
CREATE OR REPLACE FUNCTION "student_graduation_predicate_func"()
RETURNS TRIGGER AS $$
  DECLARE 
    r record;
    studySemesterCount integer;
    repeatCourse integer;
    cGradePoint numeric(3,2);
    belowCCount integer;
    diktiStudyProgramCode text;
    diplomaPrefix text;
    diplomaSequence integer;
BEGIN
  IF (OLD.status IS NULL OR OLD.status != 'LULUS') AND NEW.status = 'LULUS' THEN
    SELECT COUNT(1) INTO studySemesterCount FROM study_plans WHERE student_id = NEW.id;
    SELECT COUNT(1) INTO repeatCourse FROM student_classes WHERE student_id = NEW.id AND subject_repetition > 1;
    SELECT grade_point INTO cGradePoint WHERE code = 'C';
    SELECT COUNT(1) INTO belowCCount FROM student_subjects WHERE student_id = NEW.id AND grade_point < cGradePoint;
    SELECT dsp.code INTO diktiStudyProgramCode 
    FROM study_programs sp
    JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
    WHERE sp.id = NEW.study_program_id;

    diplomaPrefix = CONCAT(diktiStudyProgramCode, EXTRACT('year' from now()));

    SELECT COUNT(1) INTO diplomaSequence FROM students WHERE diploma_number LIKE CONCAT(diplomaPrefix, '%');

    FOR r IN SELECT * FROM graduation_predicates ORDER BY minimum_gpa DESC
    LOOP
      IF NEW.gpa >= r.minimum_gpa AND studySemesterCount <= r.maximum_study_semester AND repeatCourse <= r.repeat_course_limit AND belowCCount <= r.below_minimum_grade_point_limit THEN
        NEW.graduation_predicate_id = r.id;
        exit;
      END IF;
    END LOOP;

    NEW.graduation_date = now();
    NEW.diploma_number = CONCAT(diplomaPrefix, diplomaSequence + 1);

  ELSIF NEW.status != 'LULUS' THEN
    NEW.graduation_predicate_id = NULL;
    NEW.graduation_date = NULL;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER student_graduation_predicate BEFORE INSERT OR UPDATE OF status ON students FOR EACH ROW EXECUTE PROCEDURE student_graduation_predicate_func();



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER student_graduation_predicate ON students;
CREATE OR REPLACE FUNCTION "student_graduation_predicate_func"()
RETURNS TRIGGER AS $$
  DECLARE 
    r record;
    studySemesterCount integer;
    repeatCourse integer;
    cGradePoint numeric(3,2);
    belowCCount integer;
BEGIN
  IF NEW.status = 'LULUS' THEN
    SELECT COUNT(1) INTO studySemesterCount FROM study_plans WHERE student_id = NEW.id;
    SELECT COUNT(1) INTO repeatCourse FROM student_classes WHERE student_id = NEW.id AND subject_repetition > 1;
    SELECT grade_point INTO cGradePoint WHERE code = 'C';
    SELECT COUNT(1) INTO belowCCount FROM student_subjects WHERE student_id = NEW.id AND grade_point < cGradePoint;

    FOR r IN SELECT * FROM graduation_predicates ORDER BY minimum_gpa DESC
    LOOP
      IF NEW.gpa >= r.minimum_gpa AND studySemesterCount <= r.maximum_study_semester AND repeatCourse <= r.repeat_course_limit AND belowCCount <= r.below_minimum_grade_point_limit THEN
        NEW.graduation_predicate_id = r.id;
        exit;
      END IF;
    END LOOP;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER student_graduation_predicate BEFORE INSERT OR UPDATE ON students FOR EACH ROW EXECUTE PROCEDURE student_graduation_predicate_func();

DROP TRIGGER student_done_yudicium ON yudicium_students;
DROP FUNCTION student_done_yudicium_func();

ALTER TABLE "students"
  DROP COLUMN "done_yudicium",
  DROP COLUMN "diploma_number",
  DROP COLUMN "graduation_date";
ALTER TABLE "deleted_students"
  DROP COLUMN "done_yudicium",
  DROP COLUMN "diploma_number",
  DROP COLUMN "graduation_date";

ALTER TABLE "yudicium_sessions" 
  DROP COLUMN "actual_date",
  DROP COLUMN "yudicium_number";

ALTER TABLE "deleted_yudicium_sessions" 
  DROP COLUMN "actual_date",
  DROP COLUMN "yudicium_number";


-- +goose StatementEnd
