-- +goose Up
-- +goose StatementBegin

CREATE TABLE "shared_files" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "lecturer_id" uuid NOT NULL REFERENCES "lecturers" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "title" character varying NOT NULL,
  "file_path" character varying NOT NULL,
  "file_path_type" character varying(20) NOT NULL,
  "remarks" character varying NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_shared_files" BEFORE UPDATE ON "shared_files" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_shared_files" AS TABLE "shared_files" WITH NO DATA;
CREATE TRIGGER "soft_delete_shared_files" BEFORE DELETE ON "shared_files" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "shared_files";
DROP TABLE "deleted_shared_files";

-- +goose StatementEnd
