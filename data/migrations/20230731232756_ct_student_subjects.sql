-- +goose Up
-- +goose StatementBegin

CREATE TABLE "student_subjects" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "student_id" uuid NOT NULL REFERENCES "students" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "subject_id" uuid NOT NULL REFERENCES "subjects" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "grade_semester_id" uuid NOT NULL REFERENCES "semesters" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "grade_point" numeric(3,2) NOT NULL DEFAULT 0,
  "grade_code" character varying(2) NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("student_id", "subject_id")
);
CREATE TRIGGER "updated_at_student_subjects" BEFORE UPDATE ON "student_subjects" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_student_subjects" AS TABLE "student_subjects" WITH NO DATA;
CREATE TRIGGER "soft_delete_student_subjects" BEFORE DELETE ON "student_subjects" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

----------------------------------------------------------------------

CREATE FUNCTION "student_subject_auto_insert_func"()
RETURNS TRIGGER AS $$
BEGIN
  IF OLD.is_approved IS false AND NEW.is_approved IS true THEN
    INSERT INTO student_subjects (
      student_id,
      subject_id,
      grade_semester_id
    ) SELECT
      NEW.student_id,
      sc.subject_id,
      NEW.semester_id
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
CREATE TRIGGER student_subject_auto_insert AFTER UPDATE OF is_approved ON study_plans FOR EACH ROW EXECUTE PROCEDURE student_subject_auto_insert_func();

----------------------------------------------------------------------

CREATE FUNCTION "student_subject_grading_func"()
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
        grade_code = NEW.grade_code
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
CREATE TRIGGER student_subject_grading AFTER UPDATE OF grade_code ON student_classes FOR EACH ROW EXECUTE PROCEDURE student_subject_grading_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER student_subject_auto_insert ON study_plans;
DROP FUNCTION "student_subject_auto_insert_func"();

DROP TRIGGER student_subject_grading ON student_classes;
DROP FUNCTION "student_subject_grading_func"();

DROP TABLE "student_subjects";
DROP TABLE "deleted_student_subjects";

-- +goose StatementEnd
