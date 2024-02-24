-- +goose Up
-- +goose StatementBegin

CREATE TABLE "grade_components" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "study_program_id" uuid NOT NULL REFERENCES "study_programs" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "subject_category_id" uuid NOT NULL REFERENCES "subject_categories" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "name" character varying NOT NULL,
  "default_percentage" numeric(5,2) NOT NULL DEFAULT 0,
  "is_active" boolean NOT NULL DEFAULT true,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("study_program_id", "subject_category_id", "name")
);
CREATE TRIGGER "updated_at_grade_components" BEFORE UPDATE ON "grade_components" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_grade_components" AS TABLE "grade_components" WITH NO DATA;
CREATE TRIGGER "soft_delete_grade_components" BEFORE DELETE ON "grade_components" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

------------------------------------------------------------------------

CREATE FUNCTION grade_components_default_percentage_func()
RETURNS TRIGGER AS $$
DECLARE defaultPercentageSum numeric(5,2);
BEGIN
  SELECT SUM(default_percentage) INTO defaultPercentageSum
  FROM grade_components
  WHERE study_program_id = NEW.study_program_id AND subject_category_id = NEW.subject_category_id AND is_active IS true;

  IF defaultPercentageSum > 100 THEN
    RAISE EXCEPTION 'grade component percentage is more than 100%%';
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER grade_components_default_percentage AFTER INSERT OR UPDATE ON grade_components FOR EACH ROW EXECUTE PROCEDURE grade_components_default_percentage_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER grade_components_default_percentage ON grade_components;
DROP FUNCTION grade_components_default_percentage_func();

DROP TABLE "deleted_grade_components";
DROP TABLE "grade_components";

-- +goose StatementEnd
