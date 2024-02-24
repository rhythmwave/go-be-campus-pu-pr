-- +goose Up
-- +goose StatementBegin

CREATE TABLE "class_grade_components" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "class_id" uuid NOT NULL REFERENCES "classes" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "name" character varying NOT NULL,
  "percentage" numeric(5,2) NOT NULL,
  "is_active" boolean NOT NULL DEFAULT true,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("class_id", "name")
);
CREATE TRIGGER "updated_at_class_grade_components" BEFORE UPDATE ON "class_grade_components" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_class_grade_components" AS TABLE "class_grade_components" WITH NO DATA;
CREATE TRIGGER "soft_delete_class_grade_components" BEFORE DELETE ON "class_grade_components" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

------------------------------------------------------------------------

CREATE FUNCTION class_grade_components_percentage_func()
RETURNS TRIGGER AS $$
DECLARE percentageSum numeric(5,2);
BEGIN
  SELECT SUM(percentage) INTO percentageSum
  FROM class_grade_components
  WHERE class_id = NEW.class_id AND is_active IS true;

  IF percentageSum > 100 THEN
    RAISE EXCEPTION 'grade component percentage is more than 100%%';
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER class_grade_components_percentage AFTER INSERT OR UPDATE ON class_grade_components FOR EACH ROW EXECUTE PROCEDURE class_grade_components_percentage_func();

------------------------------------------------------------------------

CREATE FUNCTION class_grade_components_insert_func()
RETURNS TRIGGER AS $$
BEGIN
  INSERT INTO class_grade_components (
    "class_id",
    "name",
    "percentage"
  ) SELECT
    NEW.id,
    sgc.name,
    sgc.percentage
  FROM subject_grade_components sgc
  WHERE sgc.subject_id = NEW.subject_id;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER class_grade_components_insert AFTER INSERT ON classes FOR EACH ROW EXECUTE PROCEDURE class_grade_components_insert_func();

------------------------------------------------------------------------

ALTER TABLE "student_class_grades" ADD COLUMN "class_grade_component_id" uuid NOT NULL REFERENCES class_grade_components (id) ON DELETE RESTRICT ON UPDATE CASCADE;
ALTER TABLE "deleted_student_class_grades" ADD COLUMN "class_grade_component_id" uuid NULL;
ALTER TABLE "student_class_grades" ADD CONSTRAINT "student_class_grades_student_id_class_id_class_grade_component_id" UNIQUE ("student_id", "class_id", "class_grade_component_id");

CREATE OR REPLACE FUNCTION student_class_grades_initial_column_func()
RETURNS TRIGGER AS $$
DECLARE
  subjectId uuid;
  gradePercentage numeric(4,1);
BEGIN
  SELECT c.subject_id, cgc.percentage INTO subjectId, gradePercentage
  FROM class_grade_components cgc
  JOIN classes c ON c.id = cgc.class_id
  WHERE cgc.id = NEW.class_grade_component_id;

  NEW.subject_id = subjectId;
  NEW.final_grade = NEW.initial_grade * (gradePercentage / 100);

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

ALTER TABLE "student_class_grades" DROP COLUMN "subject_grade_component_id";
ALTER TABLE "deleted_student_class_grades" DROP COLUMN "subject_grade_component_id";

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "student_class_grades" ADD COLUMN "subject_grade_component_id" uuid NOT NULL REFERENCES subject_grade_components (id) ON DELETE RESTRICT ON UPDATE CASCADE;
ALTER TABLE "deleted_student_class_grades" ADD COLUMN "subject_grade_component_id" uuid NULL;
ALTER TABLE "student_class_grades" ADD CONSTRAINT "student_class_grades_student_id_class_id_subject_grade_component_id" UNIQUE ("student_id", "class_id", "subject_grade_component_id");

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

ALTER TABLE "student_class_grades" DROP COLUMN "class_grade_component_id";
ALTER TABLE "deleted_student_class_grades" DROP COLUMN "class_grade_component_id";

DROP TRIGGER class_grade_components_insert ON classes;
DROP FUNCTION class_grade_components_insert_func();

DROP TRIGGER class_grade_components_percentage ON class_grade_components;
DROP FUNCTION class_grade_components_percentage_func();

DROP TABLE "deleted_class_grade_components";
DROP TABLE "class_grade_components";

-- +goose StatementEnd
