-- +goose Up
-- +goose StatementBegin

CREATE TABLE "document_types" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" character varying NOT NULL UNIQUE,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_document_types" BEFORE UPDATE ON "document_types" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_document_types" AS TABLE "document_types" WITH NO DATA;
CREATE TRIGGER "soft_delete_document_types" BEFORE DELETE ON "document_types" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

------------------------------------------

CREATE TABLE "document_actions" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "action" character varying NOT NULL UNIQUE,
  "english_action" character varying NOT NULL UNIQUE,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_document_actions" BEFORE UPDATE ON "document_actions" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_document_actions" AS TABLE "document_actions" WITH NO DATA;
CREATE TRIGGER "soft_delete_document_actions" BEFORE DELETE ON "document_actions" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "deleted_document_actions";
DROP TABLE "document_actions";

DROP TABLE "deleted_document_types";
DROP TABLE "document_types";

-- +goose StatementEnd
