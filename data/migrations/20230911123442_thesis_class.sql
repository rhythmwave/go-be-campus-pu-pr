-- +goose Up
-- +goose StatementBegin

ALTER TABLE "classes"
  ALTER COLUMN scope DROP NOT NULL,
  ALTER COLUMN is_online DROP NOT NULL,
  ALTER COLUMN is_offline DROP NOT NULL,
  ALTER COLUMN minimum_participant DROP NOT NULL,
  ALTER COLUMN maximum_participant DROP NOT NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "classes"
  ALTER COLUMN scope SET NOT NULL,
  ALTER COLUMN is_online SET NOT NULL,
  ALTER COLUMN is_offline SET NOT NULL,
  ALTER COLUMN minimum_participant SET NOT NULL,
  ALTER COLUMN maximum_participant SET NOT NULL;


-- +goose StatementEnd
