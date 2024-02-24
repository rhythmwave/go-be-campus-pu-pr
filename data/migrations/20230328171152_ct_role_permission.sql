-- +goose Up
-- +goose StatementBegin

CREATE TABLE "role_permission" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "role_id" uuid NOT NULL REFERENCES "roles" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "permission_id" uuid NOT NULL REFERENCES "permissions" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  UNIQUE("role_id", "permission_id")
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "role_permission";

-- +goose StatementEnd
