-- +goose Up
-- +goose StatementBegin

ALTER TABLE "classes" ADD COLUMN "unapproved_study_plan" integer NOT NULL DEFAULT 0;
ALTER TABLE "deleted_classes" ADD COLUMN "unapproved_study_plan" integer NULL;

ALTER TABLE "student_classes" ADD COLUMN "is_approved" boolean NOT NULL DEFAULT false;
ALTER TABLE "deleted_student_classes" ADD COLUMN "is_approved" boolean NULL;

--------------------------------------------------------

CREATE FUNCTION "student_classes_is_approved_func"()
RETURNS TRIGGER AS $$
BEGIN
  UPDATE student_classes SET is_approved = NEW.is_approved WHERE study_plan_id = NEW.id;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER student_classes_is_approved BEFORE INSERT OR UPDATE OF is_approved ON study_plans FOR EACH ROW EXECUTE PROCEDURE student_classes_is_approved_func();

--------------------------------------------------------

CREATE FUNCTION "classes_unapproved_study_plan_func"()
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

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER student_classes_is_approved ON study_plans;
DROP FUNCTION "student_classes_is_approved_func"();

DROP TRIGGER classes_unapproved_study_plan ON student_classes;
DROP FUNCTION "classes_unapproved_study_plan_func"();

ALTER TABLE "deleted_classes" DROP COLUMN "unapproved_study_plan";
ALTER TABLE "classes" DROP COLUMN "unapproved_study_plan";

ALTER TABLE "deleted_student_classes" DROP COLUMN "is_approved";
ALTER TABLE "student_classes" DROP COLUMN "is_approved";

-- +goose StatementEnd
