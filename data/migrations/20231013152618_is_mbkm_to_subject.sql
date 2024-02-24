-- +goose Up
-- +goose StatementBegin

ALTER TABLE "classes" DROP COLUMN "is_mbkm";
ALTER TABLE "deleted_classes" DROP COLUMN "is_mbkm";

ALTER TABLE "subjects" ADD COLUMN "is_mbkm" boolean NOT NULL DEFAULT false;
ALTER TABLE "deleted_subjects" ADD COLUMN "is_mbkm" boolean NULL;

CREATE OR REPLACE FUNCTION "student_subject_auto_insert_func"()
RETURNS TRIGGER AS $$
BEGIN
  IF OLD.is_approved IS false AND NEW.is_approved IS true THEN
    INSERT INTO student_subjects (
      student_id,
      subject_id,
      grade_semester_id,
      semester_package
    ) SELECT
      NEW.student_id,
      sc.subject_id,
      NEW.semester_id,
      NEW.semester_package
    FROM student_classes sc
    JOIN subjects s ON s.id = sc.subject_id
    WHERE sc.study_plan_id = NEW.id AND s.is_mbkm IS false
    ON CONFLICT (student_id, subject_id) DO NOTHING;

    WITH d AS (
      WITH e AS (
        SELECT DISTINCT(subject_id)
        FROM student_classes 
        WHERE student_id = NEW.student_id
      )
      SELECT SUM(s.theory_credit + s.practicum_credit + s.field_practicum_credit) AS total
      FROM subjects s
      JOIN e ON e.subject_id = s.id
      WHERE s.is_mbkm IS false
    )
    UPDATE students s SET total_credit = d.total
    FROM d
    WHERE s.id = NEW.student_id;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION "student_subject_grading_func"()
RETURNS TRIGGER AS $$
DECLARE
  isMbkm boolean;
  finalScoreDeterminant curriculums_final_score_determinant;
  gradePoint numeric(3,2);
  gradeCode character varying(2);
BEGIN
  SELECT is_mbkm INTO isMbkm FROM subjects WHERE id = NEW.subject_id; 

  IF COALESCE(OLD.grade_code, '') != NEW.grade_code AND isMbkm IS false THEN
    SELECT c.final_score_determinant, ss.grade_point, ss.grade_code
    INTO finalScoreDeterminant, gradePoint, gradeCode
    FROM student_subjects ss
    JOIN students s ON s.id = ss.student_id
    LEFT JOIN curriculums c ON c.id = s.curriculum_id
    WHERE student_id = NEW.student_id AND subject_id = NEW.subject_id;

    IF finalScoreDeterminant::text = 'last' OR gradePoint < NEW.grade_point THEN
      UPDATE student_subjects ss SET
        grade_semester_id = sp.semester_id,
        grade_point = NEW.grade_point,
        grade_code = NEW.grade_code,
        semester_package = NEW.semester_package
      FROM study_plans sp
      WHERE sp.id = NEW.study_plan_id AND ss.student_id = NEW.student_id AND ss.subject_id = NEW.subject_id;

      WITH d AS (
        SELECT SUM(grade_point) as grade_point_sum, COUNT(1) AS total
        FROM student_subjects
        WHERE student_id = NEW.student_id
      )
      UPDATE students s SET gpa = COALESCE(d.grade_point_sum, 0) / COALESCE(d.total, 1)
      FROM d
      WHERE s.id = NEW.student_id;
    END IF;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

CREATE OR REPLACE FUNCTION "student_subject_auto_insert_func"()
RETURNS TRIGGER AS $$
BEGIN
  IF OLD.is_approved IS false AND NEW.is_approved IS true THEN
    INSERT INTO student_subjects (
      student_id,
      subject_id,
      grade_semester_id,
      semester_package
    ) SELECT
      NEW.student_id,
      sc.subject_id,
      NEW.semester_id,
      NEW.semester_package
    FROM student_classes sc
    WHERE sc.study_plan_id = NEW.id
    ON CONFLICT (student_id, subject_id) DO NOTHING;

    WITH d AS (
      WITH e AS (
        SELECT DISTINCT(subject_id)
        FROM student_classes 
        WHERE student_id = NEW.student_id
      )
      SELECT SUM(s.theory_credit + s.practicum_credit + s.field_practicum_credit) AS total
      FROM subjects s
      JOIN e ON e.subject_id = s.id
    )
    UPDATE students s SET total_credit = d.total
    FROM d
    WHERE s.id = NEW.student_id;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION "student_subject_grading_func"()
RETURNS TRIGGER AS $$
DECLARE
  finalScoreDeterminant curriculums_final_score_determinant;
  gradePoint numeric(3,2);
  gradeCode character varying(2);
BEGIN
  IF COALESCE(OLD.grade_code, '') != NEW.grade_code THEN
    SELECT c.final_score_determinant, ss.grade_point, ss.grade_code
    INTO finalScoreDeterminant, gradePoint, gradeCode
    FROM student_subjects ss
    JOIN students s ON s.id = ss.student_id
    LEFT JOIN curriculums c ON c.id = s.curriculum_id
    WHERE student_id = NEW.student_id AND subject_id = NEW.subject_id;

    IF finalScoreDeterminant::text = 'last' OR gradePoint < NEW.grade_point THEN
      UPDATE student_subjects ss SET
        grade_semester_id = sp.semester_id,
        grade_point = NEW.grade_point,
        grade_code = NEW.grade_code,
        semester_package = NEW.semester_package
      FROM study_plans sp
      WHERE sp.id = NEW.study_plan_id AND ss.student_id = NEW.student_id AND ss.subject_id = NEW.subject_id;

      WITH d AS (
        SELECT SUM(grade_point) as grade_point_sum, COUNT(1) AS total
        FROM student_subjects
        WHERE student_id = NEW.student_id
      )
      UPDATE students s SET gpa = COALESCE(d.grade_point_sum, 0) / COALESCE(d.total, 1)
      FROM d
      WHERE s.id = NEW.student_id;
    END IF;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

ALTER TABLE "subjects" DROP COLUMN "is_mbkm";
ALTER TABLE "deleted_subjects" DROP COLUMN "is_mbkm";

ALTER TABLE "classes" ADD COLUMN "is_mbkm" boolean NOT NULL DEFAULT false;
ALTER TABLE "deleted_classes" ADD COLUMN "is_mbkm" boolean NULL;

-- +goose StatementEnd
