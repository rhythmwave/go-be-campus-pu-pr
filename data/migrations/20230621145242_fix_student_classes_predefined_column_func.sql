-- +goose Up
-- +goose StatementBegin

CREATE OR REPLACE FUNCTION "student_classes_predefined_column_func"()
RETURNS TRIGGER AS $$
DECLARE
  equivalentSubjectId uuid;
  studentId uuid;
  subjectId uuid;
  totalCredit integer;
  isMandatory boolean;
  subjectRepetition integer;
BEGIN
  SELECT student_id INTO studentId FROM study_plans WHERE id = NEW.study_plan_id;
  NEW.student_id = studentId;

  SELECT subject_id INTO subjectId FROM classes WHERE id = NEW.class_id;
  NEW.subject_id = subjectId;

  SELECT (theory_credit + practicum_credit + field_practicum_credit), is_mandatory
  INTO totalCredit, isMandatory
  FROM subjects
  WHERE id = NEW.subject_id;

  NEW.subject_is_mandatory = isMandatory;
  NEW.total_credit = totalCredit;

  SELECT COUNT(1) INTO subjectRepetition FROM student_classes
  WHERE id != NEW.id AND student_id = NEW.student_id AND subject_id = NEW.subject_id;

  NEW.subject_repetition = subjectRepetition + 1;

  IF NEW.curriculum_id != NEW.student_curriculum_id THEN
    SELECT equivalent_subject_id INTO equivalentSubjectId FROM subject_equivalences WHERE subject_id = NEW.subject_id AND equivalent_curriculum_id = NEW.student_curriculum_id;
    NEW.student_equivalent_subject_id = equivalentSubjectId;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

CREATE FUNCTION "student_classes_predefined_column_func"()
RETURNS TRIGGER AS $$
DECLARE
  equivalentSubjectId uuid;
  studentId uuid;
  totalCredit integer;
  isMandatory boolean;
  subjectRepetition integer;
BEGIN
  SELECT student_id INTO studentId FROM study_plans WHERE id = NEW.study_plan_id;
  NEW.student_id = studentId;

  SELECT (theory_credit + practicum_credit + field_practicum_credit), is_mandatory
  INTO totalCredit, isMandatory
  FROM subjects
  WHERE id = NEW.subject_id;

  NEW.subject_is_mandatory = isMandatory;
  NEW.total_credit = totalCredit;

  SELECT COUNT(1) INTO subjectRepetition FROM student_classes
  WHERE id != NEW.id AND student_id = NEW.student_id AND subject_id = NEW.subject_id;

  NEW.subject_repetition = subjectRepetition + 1;

  IF NEW.curriculum_id != NEW.student_curriculum_id THEN
    SELECT equivalent_subject_id INTO equivalentSubjectId FROM subject_equivalences WHERE subject_id = NEW.subject_id AND equivalent_curriculum_id = NEW.student_curriculum_id;
    NEW.student_equivalent_subject_id = equivalentSubjectId;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- +goose StatementEnd
