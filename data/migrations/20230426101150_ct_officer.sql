-- +goose Up
-- +goose StatementBegin

CREATE TABLE "officers" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "id_national_lecturer" character varying NULL,
  "name" character varying NOT NULL,
  "title" character varying NULL,
  "english_title" character varying NULL,
  "study_program_id" uuid NULL REFERENCES "study_programs" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "signature_path" character varying NULL,
  "signature_path_type" character varying(20) NULL,
  "show_signature" boolean NOT NULL DEFAULT true,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_officers" BEFORE UPDATE ON "officers" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_officers" AS TABLE "officers" WITH NO DATA;
CREATE TRIGGER "soft_delete_officers" BEFORE DELETE ON "officers" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "deleted_officers";
DROP TABLE "officers";

-- +goose StatementEnd
