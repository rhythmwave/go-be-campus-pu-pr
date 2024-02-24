-- +goose Up
-- +goose StatementBegin

ALTER TABLE "shared_files" ADD COLUMN "is_approved" boolean NOT NULL DEFAULT false;
ALTER TABLE "deleted_shared_files" ADD COLUMN "is_approved" boolean NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "shared_files" DROP COLUMN "is_approved";
ALTER TABLE "deleted_shared_files" DROP COLUMN "is_approved";

-- +goose StatementEnd
