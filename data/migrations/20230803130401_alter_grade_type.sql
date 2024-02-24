-- +goose Up
-- +goose StatementBegin

ALTER TABLE "grade_types"
  ALTER COLUMN "minimum_grade" TYPE numeric(4,1),
  ALTER COLUMN "maximum_grade" TYPE numeric(4,1);

ALTER TABLE "deleted_grade_types"
  ALTER COLUMN "minimum_grade" TYPE numeric(4,1),
  ALTER COLUMN "maximum_grade" TYPE numeric(4,1);

ALTER TABLE "student_classes" 
  ALTER COLUMN "final_grade" TYPE numeric(4,1);
ALTER TABLE "deleted_student_classes" 
  ALTER COLUMN "final_grade" TYPE numeric(4,1);


ALTER TABLE "student_class_grades" 
  ALTER COLUMN "initial_grade" TYPE numeric(4,1),
  ALTER COLUMN "final_grade" TYPE numeric(4,1);
ALTER TABLE "deleted_student_class_grades" 
  ALTER COLUMN "initial_grade" TYPE numeric(4,1),
  ALTER COLUMN "final_grade" TYPE numeric(4,1);

------------------------------------------------------------------

CREATE OR REPLACE FUNCTION student_class_grades_initial_column_func()
RETURNS TRIGGER AS $$
DECLARE
  subjectId uuid;
  gradePercentage numeric(4,1);
BEGIN
  SELECT subject_id, percentage INTO subjectId, gradePercentage FROM subject_grade_components WHERE id = NEW.subject_grade_component_id;
  NEW.subject_id = subjectId;
  NEW.final_grade = NEW.initial_grade * (gradePercentage / 100);

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

------------------------------------------------------------------

CREATE OR REPLACE FUNCTION student_classes_final_grade_func()
RETURNS TRIGGER AS $$
DECLARE
  finalGradeDiff numeric(4,1) := NEW.final_grade - COALESCE(OLD.final_grade, 0);
  classGrade numeric(4,1);
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

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "grade_types"
  ALTER COLUMN "minimum_grade" TYPE numeric(5,2),
  ALTER COLUMN "maximum_grade" TYPE numeric(5,2);

ALTER TABLE "deleted_grade_types"
  ALTER COLUMN "minimum_grade" TYPE numeric(5,2),
  ALTER COLUMN "maximum_grade" TYPE numeric(5,2);

ALTER TABLE "student_classes" 
  ALTER COLUMN "final_grade" TYPE numeric(5,2);
ALTER TABLE "deleted_student_classes" 
  ALTER COLUMN "final_grade" TYPE numeric(5,2);


ALTER TABLE "student_class_grades" 
  ALTER COLUMN "initial_grade" TYPE numeric(5,2),
  ALTER COLUMN "final_grade" TYPE numeric(5,2);
ALTER TABLE "deleted_student_class_grades" 
  ALTER COLUMN "initial_grade" TYPE numeric(5,2),
  ALTER COLUMN "final_grade" TYPE numeric(5,2);

------------------------------------------------------------------

CREATE OR REPLACE FUNCTION student_class_grades_initial_column_func()
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

------------------------------------------------------------------

CREATE OR REPLACE FUNCTION student_classes_final_grade_func()
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

-- +goose StatementEnd
