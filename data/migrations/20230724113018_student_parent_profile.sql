-- +goose Up
-- +goose StatementBegin

ALTER TABLE "students"
  ADD COLUMN "parent_religion" character varying NULL,
  ADD COLUMN "parent_nationality" character varying NULL,
  ADD COLUMN "father_work_address" character varying NULL;

ALTER TABLE "deleted_students"
  ADD COLUMN "parent_religion" character varying NULL,
  ADD COLUMN "parent_nationality" character varying NULL,
  ADD COLUMN "father_work_address" character varying NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "students"
  DROP COLUMN "parent_religion",
  DROP COLUMN "parent_nationality",
  DROP COLUMN "father_work_address";

ALTER TABLE "deleted_students"
  DROP COLUMN "parent_religion",
  DROP COLUMN "parent_nationality",
  DROP COLUMN "father_work_address";

-- +goose StatementEnd
