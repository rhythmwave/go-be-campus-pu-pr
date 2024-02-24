-- +goose Up
-- +goose StatementBegin

ALTER TABLE "study_levels"
  ADD COLUMN "kkni_qualification" character varying NULL,
  ADD COLUMN "acceptance_requirement" character varying NULL,
  ADD COLUMN "further_education_level" character varying NULL,
  ADD COLUMN "professional_status" character varying NULL,
  ADD COLUMN "course_language" character varying NULL,
  ADD COLUMN "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  ADD COLUMN "updated_at" timestamp NULL;

CREATE TRIGGER "updated_at_study_levels" BEFORE UPDATE ON "study_levels" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER "updated_at_study_levels" ON "study_levels";

ALTER TABLE "study_levels"
  DROP COLUMN "kkni_qualification",
  DROP COLUMN "acceptance_requirement",
  DROP COLUMN "further_education_level",
  DROP COLUMN "professional_status",
  DROP COLUMN "course_language",
  DROP COLUMN "updated_by",
  DROP COLUMN "updated_at";

-- +goose StatementEnd
