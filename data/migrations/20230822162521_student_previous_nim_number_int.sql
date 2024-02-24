-- +goose Up
-- +goose StatementBegin

ALTER TABLE "students" ALTER COLUMN "previous_nim_number" TYPE integer USING previous_nim_number::integer;
ALTER TABLE "deleted_students" ALTER COLUMN "previous_nim_number" TYPE integer USING previous_nim_number::integer;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "students" ALTER COLUMN "previous_nim_number" TYPE character varying USING previous_nim_number::text;
ALTER TABLE "deleted_students" ALTER COLUMN "previous_nim_number" TYPE character varying USING previous_nim_number::text;

-- +goose StatementEnd
