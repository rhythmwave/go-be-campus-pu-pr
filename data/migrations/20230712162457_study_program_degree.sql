-- +goose Up
-- +goose StatementBegin

ALTER TABLE "study_programs"
  ADD COLUMN "degree" character varying NULL,
  ADD COLUMN "short_degree" character varying NULL,
  ADD COLUMN "english_degree" character varying NULL;

ALTER TABLE "deleted_study_programs"
  ADD COLUMN "degree" character varying NULL,
  ADD COLUMN "short_degree" character varying NULL,
  ADD COLUMN "english_degree" character varying NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "study_programs"
  DROP COLUMN "degree",
  DROP COLUMN "short_degree",
  DROP COLUMN "english_degree";

ALTER TABLE "deleted_study_programs"
  DROP COLUMN "degree",
  DROP COLUMN "short_degree",
  DROP COLUMN "english_degree";

-- +goose StatementEnd
