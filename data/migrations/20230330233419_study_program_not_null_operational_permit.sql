-- +goose Up
-- +goose StatementBegin

ALTER TABLE "study_programs"
  ALTER COLUMN operational_permit_number SET NOT NULL,
  ALTER COLUMN operational_permit_date SET NOT NULL,
  ALTER COLUMN operational_permit_due_date SET NOT NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "study_programs"
  ALTER COLUMN operational_permit_number DROP NOT NULL,
  ALTER COLUMN operational_permit_date DROP NOT NULL,
  ALTER COLUMN operational_permit_due_date DROP NOT NULL;

-- +goose StatementEnd
