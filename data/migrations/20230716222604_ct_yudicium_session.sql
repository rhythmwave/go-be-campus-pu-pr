-- +goose Up
-- +goose StatementBegin

CREATE TABLE "yudicium_sessions" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "semester_id" uuid NOT NULL REFERENCES "semesters" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "name" character varying NOT NULL,
  "session_date" date NOT NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_yudicium_sessions" BEFORE UPDATE ON "yudicium_sessions" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_yudicium_sessions" AS TABLE "yudicium_sessions" WITH NO DATA;
CREATE TRIGGER "soft_delete_yudicium_sessions" BEFORE DELETE ON "yudicium_sessions" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "deleted_yudicium_sessions";
DROP TABLE "yudicium_sessions";

-- +goose StatementEnd
