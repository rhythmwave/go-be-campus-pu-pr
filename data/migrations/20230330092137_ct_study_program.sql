-- +goose Up
-- +goose StatementBegin

CREATE TABLE "study_programs" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "major_id" uuid NOT NULL REFERENCES "majors" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "dikti_study_program_id" uuid NOT NULL REFERENCES "dikti_study_programs" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "name" character varying NOT NULL,
  "short_name" character varying NULL,
  "english_name" character varying NULL,
  "english_short_name" character varying NULL,
  "address" character varying NOT NULL DEFAULT '',
  "phone_number" character varying(20) NULL,
  "fax" character varying(20) NULL,
  "email" character varying(100) NULL,
  "website" character varying NULL,
  "curiculum_review_frequency" character varying NOT NULL,
  "curiculum_review_method" character varying NOT NULL,
  "establishment_date" date NOT NULL,
  "start_semester" character varying NULL,
  "operational_permit_number" character varying NULL,
  "operational_permit_date" date NULL,
  "operational_permit_due_date" date NULL,
  "operator_name" character varying NULL,
  "operator_phone_number" character varying(20) NULL,
  "is_active" boolean NOT NULL DEFAULT false,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_study_programs" BEFORE UPDATE ON "study_programs" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_study_programs" AS TABLE "study_programs" WITH NO DATA;
CREATE TRIGGER "soft_delete_study_programs" BEFORE DELETE ON "study_programs" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "deleted_study_programs";
DROP TABLE "study_programs";

-- +goose StatementEnd
