-- +goose Up
-- +goose StatementBegin

ALTER TABLE "students" DROP COLUMN "created_by";
ALTER TABLE "deleted_students" DROP COLUMN "created_by";

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "students" ADD COLUMN "created_by" uuid NULL REFERENCES admins(id) ON DELETE RESTRICT ON UPDATE CASCADE;
ALTER TABLE "deleted_students" ADD COLUMN "created_by" uuid NULL;

-- +goose StatementEnd
