-- +goose Up
-- +goose StatementBegin

CREATE TABLE "subject_categories" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "code" character varying NOT NULL UNIQUE,
  "name" character varying NOT NULL UNIQUE,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_subject_categories" BEFORE UPDATE ON "subject_categories" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_subject_categories" AS TABLE "subject_categories" WITH NO DATA;
CREATE TRIGGER "soft_delete_subject_categories" BEFORE DELETE ON "subject_categories" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "deleted_subject_categories";
DROP TABLE "subject_categories";

-- +goose StatementEnd
