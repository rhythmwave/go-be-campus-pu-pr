-- +goose Up
-- +goose StatementBegin

CREATE TABLE "subject_grade_components" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "subject_id" uuid NOT NULL REFERENCES "subjects" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "name" character varying NOT NULL,
  "percentage" numeric(5,2) NOT NULL,
  "is_active" boolean NOT NULL DEFAULT true,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("subject_id", "name")
);
CREATE TRIGGER "updated_at_subject_grade_components" BEFORE UPDATE ON "subject_grade_components" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_subject_grade_components" AS TABLE "subject_grade_components" WITH NO DATA;
CREATE TRIGGER "soft_delete_subject_grade_components" BEFORE DELETE ON "subject_grade_components" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

------------------------------------------------------------------------

CREATE FUNCTION subject_grade_components_percentage_func()
RETURNS TRIGGER AS $$
DECLARE percentageSum numeric(5,2);
BEGIN
  SELECT SUM(percentage) INTO percentageSum
  FROM subject_grade_components
  WHERE subject_id = NEW.subject_id AND is_active IS true;

  IF percentageSum > 100 THEN
    RAISE EXCEPTION 'grade component percentage is more than 100%%';
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER subject_grade_components_percentage AFTER INSERT OR UPDATE ON subject_grade_components FOR EACH ROW EXECUTE PROCEDURE subject_grade_components_percentage_func();

------------------------------------------------------------------------

CREATE FUNCTION subject_grade_components_insert_func()
RETURNS TRIGGER AS $$
BEGIN
  INSERT INTO subject_grade_components (
    "subject_id",
    "name",
    "percentage",
    "created_by"
  ) SELECT
    NEW.id,
    gc.name,
    gc.default_percentage,
    NEW.created_by
  FROM curriculums c
  JOIN study_programs sp ON sp.id = c.study_program_id
  JOIN grade_components gc ON gc.study_program_id = sp.id
  WHERE gc.subject_category_id = NEW.subject_category_id AND c.id = NEW.curriculum_id;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER subject_grade_components_insert AFTER INSERT ON subjects FOR EACH ROW EXECUTE PROCEDURE subject_grade_components_insert_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER subject_grade_components_insert ON subjects;
DROP FUNCTION subject_grade_components_insert_func();

DROP TRIGGER subject_grade_components_percentage ON subject_grade_components;
DROP FUNCTION subject_grade_components_percentage_func();

DROP TABLE "deleted_subject_grade_components";
DROP TABLE "subject_grade_components";

-- +goose StatementEnd
