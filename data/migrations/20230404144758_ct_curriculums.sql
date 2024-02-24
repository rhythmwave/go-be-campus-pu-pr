-- +goose Up
-- +goose StatementBegin

CREATE TABLE "curriculums" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "study_program_id" uuid NOT NULL REFERENCES "study_programs" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "name" character varying NOT NULL,
  "year" character(4) NOT NULL CHECK ("year" ~ '^[0-9\.]+$'),
  "rector_decision_number" character varying NULL,
  "rector_decision_date" date NULL,
  "aggreeing_party" text NULL,
  "aggreement_date" date NULL,
  "ideal_study_period" integer NOT NULL,
  "maximum_study_period" integer NOT NULL,
  "remarks" text NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_curriculums" BEFORE UPDATE ON "curriculums" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_curriculums" AS TABLE "curriculums" WITH NO DATA;
CREATE TRIGGER "soft_delete_curriculums" BEFORE DELETE ON "curriculums" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "deleted_curriculums";
DROP TABLE "curriculums";

-- +goose StatementEnd
