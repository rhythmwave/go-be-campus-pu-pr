-- +goose Up
-- +goose StatementBegin

CREATE TABLE "yudicium_terms" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "curriculum_id" uuid NOT NULL REFERENCES "curriculums" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "term" character varying NOT NULL,
  "remarks" character varying NOT NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_yudicium_terms" BEFORE UPDATE ON "yudicium_terms" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_yudicium_terms" AS TABLE "yudicium_terms" WITH NO DATA;
CREATE TRIGGER "soft_delete_yudicium_terms" BEFORE DELETE ON "yudicium_terms" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "deleted_yudicium_terms";
DROP TABLE "yudicium_terms";

-- +goose StatementEnd
