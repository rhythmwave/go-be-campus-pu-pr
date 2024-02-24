-- +goose Up
-- +goose StatementBegin

ALTER TABLE "faculties"
  ALTER COLUMN short_name SET NOT NULL,
  ALTER COLUMN address DROP DEFAULT,
  ALTER COLUMN phone_number SET NOT NULL,
  ALTER COLUMN email SET NOT NULL,
  ALTER COLUMN contact_person SET NOT NULL;

ALTER TABLE "majors"
  ALTER COLUMN short_name SET NOT NULL,
  ALTER COLUMN address DROP DEFAULT,
  ALTER COLUMN phone_number SET NOT NULL,
  ALTER COLUMN email SET NOT NULL,
  ALTER COLUMN contact_person SET NOT NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "faculties"
  ALTER COLUMN short_name DROP NOT NULL,
  ALTER COLUMN address SET DEFAULT '',
  ALTER COLUMN phone_number DROP NOT NULL,
  ALTER COLUMN email DROP NOT NULL,
  ALTER COLUMN contact_person DROP NOT NULL;

ALTER TABLE "majors"
  ALTER COLUMN short_name DROP NOT NULL,
  ALTER COLUMN address SET DEFAULT '',
  ALTER COLUMN phone_number DROP NOT NULL,
  ALTER COLUMN email DROP NOT NULL,
  ALTER COLUMN contact_person DROP NOT NULL;

-- +goose StatementEnd
