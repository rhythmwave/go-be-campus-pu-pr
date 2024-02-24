-- +goose Up
-- +goose StatementBegin

ALTER TABLE "lectures"
  ADD COLUMN "autonomous_participation_start_time" timestamp NULL,
  ADD COLUMN "autonomous_participation_end_time" timestamp NULL;

ALTER TABLE "deleted_lectures"
  ADD COLUMN "autonomous_participation_start_time" timestamp NULL,
  ADD COLUMN "autonomous_participation_end_time" timestamp NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "lectures"
  DROP COLUMN "autonomous_participation_start_time",
  DROP COLUMN "autonomous_participation_end_time";

ALTER TABLE "deleted_lectures"
  DROP COLUMN "autonomous_participation_start_time",
  DROP COLUMN "autonomous_participation_end_time";

-- +goose StatementEnd
