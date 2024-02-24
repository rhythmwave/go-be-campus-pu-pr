-- +goose Up
-- +goose StatementBegin

CREATE TABLE "officer_actions" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "document_type_id" uuid NOT NULL REFERENCES "document_types" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "document_action_id" uuid NOT NULL REFERENCES "document_actions" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "officer_id" uuid NOT NULL REFERENCES "officers" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("document_type_id", "document_action_id", "officer_id")
);
CREATE TRIGGER "updated_at_officer_actions" BEFORE UPDATE ON "officer_actions" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_officer_actions" AS TABLE "officer_actions" WITH NO DATA;
CREATE TRIGGER "soft_delete_officer_actions" BEFORE DELETE ON "officer_actions" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "deleted_officer_actions";
DROP TABLE "officer_actions";

-- +goose StatementEnd
