-- +goose Up
-- +goose StatementBegin

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "admin_activity_logs" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "admin_id" uuid NOT NULL,
  "admin_name" character varying NOT NULL,
  "admin_username" character varying(100) NOT NULL,
  "module" character varying NOT NULL,
  "action" text NOT NULL,
  "ip_address" character varying NOT NULL,
  "user_agent" character varying NOT NULL,
  "execution_time" numeric(12,9) NOT NULL,
  "memory_usage" numeric(15,12) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "admin_activity_logs";

-- +goose StatementEnd
