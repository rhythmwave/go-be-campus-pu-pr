-- +goose Up
-- +goose StatementBegin

ALTER TABLE "students" ADD COLUMN "guardian_id_number" character varying NULL;
ALTER TABLE "deleted_students" ADD COLUMN "guardian_id_number" character varying NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "students" DROP COLUMN "guardian_id_number";
ALTER TABLE "deleted_students" DROP COLUMN "guardian_id_number";

-- +goose StatementEnd
