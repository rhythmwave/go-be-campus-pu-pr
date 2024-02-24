-- +goose Up
-- +goose StatementBegin

ALTER TABLE "students"
  ADD COLUMN "gpa" numeric(3,2) NULL,
  ADD COLUMN "total_credit" integer NULL,
  ADD COLUMN "transcript_is_archived" boolean NOT NULL DEFAULT false;

ALTER TABLE "deleted_students"
  ADD COLUMN "gpa" numeric(3,2) NULL,
  ADD COLUMN "total_credit" integer NULL,
  ADD COLUMN "transcript_is_archived" boolean NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "students"
  DROP COLUMN "gpa",
  DROP COLUMN "total_credit",
  DROP COLUMN "transcript_is_archived";

ALTER TABLE "deleted_students"
  DROP COLUMN "gpa",
  DROP COLUMN "total_credit",
  DROP COLUMN "transcript_is_archived";

-- +goose StatementEnd
