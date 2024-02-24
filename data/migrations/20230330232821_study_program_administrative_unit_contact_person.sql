-- +goose Up
-- +goose StatementBegin

ALTER TABLE "study_programs"
  ADD COLUMN "contact_person" character varying NULL,
  ADD COLUMN "administrative_unit" character varying NULL;

ALTER TABLE "deleted_study_programs"
  ADD COLUMN "contact_person" character varying NULL,
  ADD COLUMN "administrative_unit" character varying NULL;


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "study_programs"
  DROP COLUMN "contact_person",
  DROP COLUMN "administrative_unit";
ALTER TABLE "deleted_study_programs"
  DROP COLUMN "contact_person",
  DROP COLUMN "administrative_unit";

-- +goose StatementEnd
