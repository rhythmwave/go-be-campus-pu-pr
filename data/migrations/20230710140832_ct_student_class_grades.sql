-- +goose Up
-- +goose StatementBegin

ALTER TABLE "student_classes" 
  ADD COLUMN "final_grade" numeric(5,2) NULL,
  ADD COLUMN "graded_by_admin_id" uuid NULL REFERENCES "admins" ("id") ON DELETE SET NULL ON UPDATE CASCADE,
  ADD COLUMN "graded_by_lecturer_id" uuid NULL REFERENCES "lecturers" ("id") ON DELETE SET NULL ON UPDATE CASCADE,
  ADD COLUMN "graded_at" timestamp NULL;
ALTER TABLE "deleted_student_classes" 
  ADD COLUMN "final_grade" numeric(5,2) NULL,
  ADD COLUMN "graded_by_admin_id" uuid NULL,
  ADD COLUMN "graded_by_lecturer_id" uuid NULL,
  ADD COLUMN "graded_at" timestamp NULL;

ALTER TABLE "classes" ADD COLUMN "total_graded_participant" integer NOT NULL DEFAULT 0;
ALTER TABLE "deleted_classes" ADD COLUMN "total_graded_participant" integer NULL;

CREATE TABLE "student_class_grades" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "student_id" uuid NOT NULL REFERENCES "students" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "class_id" uuid NOT NULL REFERENCES "classes" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "subject_grade_component_id" uuid NOT NULL REFERENCES "subject_grade_components" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "subject_id" uuid NOT NULL REFERENCES "subjects" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "initial_grade" numeric(5,2) NOT NULL,
  "final_grade" numeric(5,2) NOT NULL,
  "graded_by_admin_id" uuid NULL REFERENCES "admins" ("id") ON DELETE SET NULL ON UPDATE CASCADE,
  "graded_by_lecturer_id" uuid NULL REFERENCES "lecturers" ("id") ON DELETE SET NULL ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("student_id", "class_id", "subject_grade_component_id")
);
CREATE TRIGGER "updated_at_student_class_grades" BEFORE UPDATE ON "student_class_grades" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_student_class_grades" AS TABLE "student_class_grades" WITH NO DATA;
CREATE TRIGGER "soft_delete_student_class_grades" BEFORE DELETE ON "student_class_grades" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

------------------------------------------------------------------

CREATE FUNCTION student_class_grades_initial_column_func()
RETURNS TRIGGER AS $$
DECLARE
  subjectId uuid;
  gradePercentage numeric(5,2);
BEGIN
  SELECT subject_id, percentage INTO subjectId, gradePercentage FROM subject_grade_components WHERE id = NEW.subject_grade_component_id;
  NEW.subject_id = subjectId;
  NEW.final_grade = NEW.initial_grade * (gradePercentage / 100);

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER student_class_grading BEFORE INSERT OR UPDATE ON "student_class_grades" FOR EACH ROW EXECUTE PROCEDURE student_class_grades_initial_column_func();

------------------------------------------------------------------

CREATE FUNCTION student_classes_final_grade_func()
RETURNS TRIGGER AS $$
DECLARE
  finalGradeDiff numeric(5,2) := NEW.final_grade - COALESCE(OLD.final_grade, 0);
  classGrade numeric(5,2);
  gradeCode character varying(2);
  gradePoint numeric(3,2);
  studyPlanId uuid;
BEGIN
  SELECT COALESCE(final_grade, 0) + finalGradeDiff, study_plan_id INTO classGrade, studyPlanId FROM student_classes WHERE student_id = NEW.student_id AND class_id = NEW.class_id;

  SELECT gt.code, gt.grade_point 
  INTO gradeCode, gradePoint
  FROM grade_types gt
  JOIN students s ON s.id = NEW.student_id
  JOIN study_programs sp ON sp.id = s.study_program_id
  JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
  WHERE dsp.study_level_id = gt.study_level_id AND classGrade BETWEEN gt.minimum_grade AND gt.maximum_grade
  ORDER BY gt.maximum_grade DESC LIMIT 1;

  IF gradeCode IS NULL OR gradePoint IS NULL THEN
    RAISE EXCEPTION 'cannot convert grade to grade code';
  END IF;

  UPDATE student_classes SET 
    final_grade = classGrade,
    grade_point = gradePoint,
    grade_code = gradeCode,
    graded_by_admin_id = NEW.graded_by_admin_id,
    graded_by_lecturer_id = NEW.graded_by_lecturer_id,
    graded_at = now()
  WHERE student_id = NEW.student_id AND class_id = NEW.class_id;

  WITH d AS (
    SELECT COUNT(1) AS total
    FROM student_classes sc
    WHERE sc.class_id = NEW.class_id AND graded_at IS NOT NULL
  )
  UPDATE classes c SET total_graded_participant = d.total
  FROM d
  WHERE c.id = NEW.class_id;

  WITH d AS (
    SELECT SUM(total_credit * grade_point) AS points, SUM(total_credit) AS total
    FROM student_classes
    WHERE study_plan_id = studyPlanId
  )
  UPDATE study_plans sp
  SET grade_point = d.points / d.total
  FROM d
  WHERE sp.id = studyPlanId;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER student_classes_final_grade AFTER INSERT OR UPDATE ON student_class_grades FOR EACH ROW EXECUTE PROCEDURE student_classes_final_grade_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER student_class_grading ON "student_class_grades";
DROP FUNCTION student_class_grades_initial_column_func();
DROP TRIGGER student_classes_final_grade ON student_class_grades;
DROP FUNCTION student_classes_final_grade_func();


DROP TABLE "deleted_student_class_grades";
DROP TABLE "student_class_grades";

ALTER TABLE "deleted_classes" DROP COLUMN "total_graded_participant";
ALTER TABLE "classes" DROP COLUMN "total_graded_participant";

ALTER TABLE "student_classes"
  DROP COLUMN "final_grade",
  DROP COLUMN "graded_by_admin_id",
  DROP COLUMN "graded_by_lecturer_id",
  DROP COLUMN "graded_at";
ALTER TABLE "deleted_student_classes" 
  DROP COLUMN "final_grade",
  DROP COLUMN "graded_by_admin_id",
  DROP COLUMN "graded_by_lecturer_id",
  DROP COLUMN "graded_at";

-- +goose StatementEnd
