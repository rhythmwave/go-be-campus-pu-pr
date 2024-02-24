-- +goose Up
-- +goose StatementBegin

CREATE TABLE "grade_types" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "study_level_id" uuid NOT NULL REFERENCES "study_levels" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "code" character varying(2) NOT NULL,
  "grade_point" numeric(3,2) NOT NULL,
  "minimum_grade" numeric(5,2) NOT NULL,
  "maximum_grade" numeric(5,2) NOT NULL,
  "grade_category" character(1) NOT NULL,
  "grade_point_category" numeric(3,2) NOT NULL,
  "label" character varying NULL,
  "english_label" character varying NULL,
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  CHECK("start_date" < "end_date")
);
CREATE TRIGGER "updated_at_grade_types" BEFORE UPDATE ON "grade_types" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_grade_types" AS TABLE "grade_types" WITH NO DATA;
CREATE TRIGGER "soft_delete_grade_types" BEFORE DELETE ON "grade_types" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "deleted_grade_types";
DROP TABLE "grade_types";

-- +goose StatementEnd
