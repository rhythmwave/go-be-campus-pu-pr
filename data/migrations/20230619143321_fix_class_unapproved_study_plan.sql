-- +goose Up
-- +goose StatementBegin

DROP TRIGGER classes_unapproved_study_plan ON student_classes;
CREATE OR REPLACE FUNCTION "classes_unapproved_study_plan_func"()
RETURNS TRIGGER AS $$
DECLARE unapprovedStudyPlan integer;
BEGIN
  IF NEW.id IS NOT NULL THEN
    SELECT COUNT(1) INTO unapprovedStudyPlan 
    FROM student_classes sc
    JOIN study_plans sp ON sp.id = sc.study_plan_id AND sp.is_approved IS false
    WHERE sc.class_id = NEW.class_id;

    UPDATE classes SET unapproved_study_plan = unapprovedStudyPlan WHERE id = NEW.class_id;
  END IF;
  IF OLD.id IS NOT NULL THEN
    SELECT COUNT(1) INTO unapprovedStudyPlan 
    FROM student_classes sc
    JOIN study_plans sp ON sp.id = sc.study_plan_id AND sp.is_approved IS false
    WHERE sc.class_id = OLD.class_id;

    UPDATE classes SET unapproved_study_plan = unapprovedStudyPlan WHERE id = OLD.class_id;
  END IF;

  IF NEW.id IS NOT NULL THEN
    RETURN NEW;
  ELSE
    RETURN OLD;
  END IF;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER classes_unapproved_study_plan AFTER INSERT OR UPDATE OR DELETE ON student_classes FOR EACH ROW EXECUTE PROCEDURE classes_unapproved_study_plan_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER classes_unapproved_study_plan ON student_classes;
CREATE OR REPLACE FUNCTION "classes_unapproved_study_plan_func"()
RETURNS TRIGGER AS $$
DECLARE unapprovedStudyPlan integer;
BEGIN

  IF NEW.id IS NOT NULL THEN
    SELECT COUNT(1) INTO unapprovedStudyPlan 
    FROM student_classes sc
    JOIN study_plans sp ON sp.id = sc.study_plan_id AND sp.is_approved IS false
    WHERE sc.class_id = NEW.class_id;

    UPDATE classes SET unapproved_study_plan = unapprovedStudyPlan WHERE id = NEW.class_id;

    RETURN NEW;
  ELSE
    SELECT COUNT(1) INTO unapprovedStudyPlan 
    FROM student_classes sc
    JOIN study_plans sp ON sp.id = sc.study_plan_id AND sp.is_approved IS false
    WHERE sc.class_id = OLD.class_id;

    UPDATE classes SET unapproved_study_plan = unapprovedStudyPlan WHERE id = OLD.class_id;

    RETURN OLD;
  END IF;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER classes_unapproved_study_plan AFTER INSERT OR UPDATE OF is_approved OR DELETE ON student_classes FOR EACH ROW EXECUTE PROCEDURE classes_unapproved_study_plan_func();

-- +goose StatementEnd
