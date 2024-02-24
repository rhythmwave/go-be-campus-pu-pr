-- +goose Up
-- +goose StatementBegin

ALTER TABLE "students" ADD COLUMN "current_semester_package" integer NOT NULL DEFAULT 0;
ALTER TABLE "deleted_students" ADD COLUMN "current_semester_package" integer NULL;

CREATE TABLE "study_plans" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "student_id" uuid NOT NULL REFERENCES "students" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "semester_id" uuid NOT NULL REFERENCES "semesters" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "semester_package" integer NOT NULL,
  "is_submitted" boolean NOT NULL DEFAULT false, 
  "is_approved" boolean NOT NULL DEFAULT false,
  "total_mandatory_credit" integer NOT NULL DEFAULT 0,
  "total_optional_credit" integer NOT NULL DEFAULT 0,
  "maximum_credit" integer NOT NULL,
  "grade_point" numeric(3,2) NOT NULL DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("student_id", "semester_id")
);
CREATE TRIGGER "updated_at_study_plans" BEFORE UPDATE ON "study_plans" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_study_plans" AS TABLE "study_plans" WITH NO DATA;
CREATE TRIGGER "soft_delete_study_plans" BEFORE DELETE ON "study_plans" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-----------------------------------------------------------------

CREATE TABLE "student_classes" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "study_plan_id" uuid NOT NULL REFERENCES "study_plans" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "student_id" uuid NOT NULL REFERENCES "students" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "subject_id" uuid NOT NULL REFERENCES "subjects" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "subject_is_mandatory" boolean NOT NULL,
  "student_equivalent_subject_id" uuid NULL REFERENCES "subjects" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "curriculum_id" uuid NOT NULL REFERENCES "curriculums" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "student_curriculum_id" uuid NOT NULL REFERENCES "curriculums" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "total_credit" integer NOT NULL,
  "grade_point" numeric(3,2) NOT NULL DEFAULT 0,
  "grade_code" character varying(2) NULL,
  "subject_repetition" integer NOT NULL,
  "class_id" uuid NOT NULL REFERENCES "classes" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("study_plan_id", "subject_id")
);
CREATE TRIGGER "updated_at_student_classes" BEFORE UPDATE ON "student_classes" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_student_classes" AS TABLE "student_classes" WITH NO DATA;
CREATE TRIGGER "soft_delete_student_classes" BEFORE DELETE ON "student_classes" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-----------------------------------------------------------------

CREATE FUNCTION "student_classes_class_total_participant_func"()
RETURNS TRIGGER AS $$
BEGIN
  IF OLD.class_id IS NOT NULL THEN
    UPDATE classes SET total_participant = total_participant - 1 WHERE id = OLD.class_id;
  END IF;

  UPDATE classes SET total_participant = total_participant + 1 WHERE id = NEW.class_id;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER student_classes_class_total_participant AFTER INSERT OR UPDATE OF class_id ON student_classes FOR EACH ROW EXECUTE PROCEDURE student_classes_class_total_participant_func();

-----------------------------------------------------------------

CREATE FUNCTION "student_classes_class_total_participant_delete_func"()
RETURNS TRIGGER AS $$
BEGIN
  UPDATE classes SET total_participant = total_participant - 1 WHERE id = OLD.class_id;

  RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER student_classes_class_total_participant_delete BEFORE DELETE ON student_classes FOR EACH ROW EXECUTE PROCEDURE student_classes_class_total_participant_delete_func();

-----------------------------------------------------------------

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
CREATE TRIGGER student_classes_predefined_column BEFORE INSERT OR UPDATE ON student_classes FOR EACH ROW EXECUTE PROCEDURE student_classes_predefined_column_func();

-----------------------------------------------------------------


CREATE FUNCTION "student_classes_grade_point_func"()
RETURNS TRIGGER AS $$
BEGIN
  WITH d AS (
    SELECT SUM(sc.total_credit * sc.grade_point) AS data_set_sum, SUM(sc.total_credit) AS total
    FROM student_classes sc
    WHERE sc.study_plan_id = OLD.study_plan_id
  )
  UPDATE study_plans SET grade_point = COALESCE(d.data_set_sum, 0) / COALESCE(d.total, 1)
  FROM d
  WHERE study_plans.id = OLD.study_plan_id;

  IF NEW.id IS NULL THEN
    RETURN OLD;
  ELSE
    RETURN NEW;
  END IF;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER student_classes_grade_point AFTER UPDATE OF grade_point OR DELETE ON student_classes FOR EACH ROW EXECUTE PROCEDURE student_classes_grade_point_func();

-----------------------------------------------------------------

CREATE FUNCTION "study_plan_total_credit_func"()
RETURNS TRIGGER AS $$
BEGIN
  IF NEW.id IS NOT NULL THEN
    IF NEW.subject_is_mandatory IS true THEN
      UPDATE study_plans SET total_mandatory_credit = total_mandatory_credit + NEW.total_credit WHERE id = NEW.study_plan_id;
    ELSE
      UPDATE study_plans SET total_optional_credit = total_optional_credit + NEW.total_credit WHERE id = NEW.study_plan_id;
    END IF;

    RETURN NEW;
  ELSIF OLD.id IS NOT NULL THEN
    IF OLD.subject_is_mandatory IS true THEN
      UPDATE study_plans SET total_mandatory_credit = total_mandatory_credit - OLD.total_credit WHERE id = OLD.study_plan_id;
    ELSE
      UPDATE study_plans SET total_optional_credit = total_optional_credit - OLD.total_credit WHERE id = OLD.study_plan_id;
    END IF;

    RETURN OLD;
  END IF;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER study_plan_total_credit AFTER INSERT OR DELETE ON student_classes FOR EACH ROW EXECUTE PROCEDURE study_plan_total_credit_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER study_plan_total_credit ON student_classes;
DROP FUNCTION "study_plan_total_credit_func"();

DROP TRIGGER student_classes_grade_point ON student_classes;
DROP FUNCTION "student_classes_grade_point_func"();

DROP TRIGGER student_classes_predefined_column ON student_classes;
DROP FUNCTION "student_classes_predefined_column_func"();

DROP TRIGGER student_classes_class_total_participant ON student_classes;
DROP FUNCTION "student_classes_class_total_participant_func"();

DROP TRIGGER student_classes_class_total_participant_delete  ON student_classes;
DROP FUNCTION "student_classes_class_total_participant_delete_func"();

DROP TABLE "deleted_student_classes";
DROP TABLE "student_classes";

DROP TABLE "deleted_study_plans";
DROP TABLE "study_plans";

ALTER TABLE "students" DROP COLUMN "current_semester_package";
ALTER TABLE "deleted_students" DROP COLUMN "current_semester_package";

-- +goose StatementEnd
