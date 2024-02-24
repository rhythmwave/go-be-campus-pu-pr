-- +goose Up
-- +goose StatementBegin

ALTER TABLE "authentications" ADD COLUMN "id_sso" character varying NULL UNIQUE;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "authentications" DROP COLUMN "id_sso";

-- +goose StatementEnd
