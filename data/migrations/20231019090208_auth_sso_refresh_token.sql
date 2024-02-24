-- +goose Up
-- +goose StatementBegin

ALTER TABLE "authentications" ADD COLUMN "sso_refresh_token" text NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "authentications" DROP COLUMN "sso_refresh_token";

-- +goose StatementEnd
