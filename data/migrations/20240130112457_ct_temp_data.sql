-- +goose Up
-- +goose StatementBegin

CREATE TABLE "temp_data" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "title" varchar NOT NULL,
  "body" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
)

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "temp_data";

-- +goose StatementEnd
