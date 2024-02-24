-- +goose Up
-- +goose StatementBegin

ALTER TABLE "lectures" ALTER COLUMN "class_id" DROP NOT NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "lectures" ALTER COLUMN "class_id" SET NOT NULL;

-- +goose StatementEnd
