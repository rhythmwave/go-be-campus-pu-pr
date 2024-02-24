-- +goose Up
-- +goose StatementBegin

ALTER TABLE "students" 
  ADD COLUMN "is_registered" boolean NOT NULL DEFAULT true,
  ADD COLUMN "has_paid" boolean NOT NULL DEFAULT false;
ALTER TABLE "deleted_students" 
  ADD COLUMN "is_registered" boolean NULL,
  ADD COLUMN "has_paid" boolean NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "students" 
  DROP COLUMN "is_registered",
  DROP COLUMN "has_paid";
ALTER TABLE "deleted_students" 
  DROP COLUMN "is_registered",
  DROP COLUMN "has_paid";

-- +goose StatementEnd
