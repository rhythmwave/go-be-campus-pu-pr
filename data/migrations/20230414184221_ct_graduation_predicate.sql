-- +goose Up
-- +goose StatementBegin

CREATE TABLE "graduation_predicates" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "predicate" character varying NOT NULL UNIQUE,
  "minimum_gpa" numeric(3,2) NOT NULL,
  "maximum_study_semester" integer NOT NULL,
  "repeat_course_limit" integer NOT NULL,
  "minimum_grade_point" numeric(3,2) NOT NULL DEFAULT 2.0,
  "below_minimum_grade_point_limit" integer NOT NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_graduation_predicates" BEFORE UPDATE ON "graduation_predicates" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_graduation_predicates" AS TABLE "graduation_predicates" WITH NO DATA;
CREATE TRIGGER "soft_delete_graduation_predicates" BEFORE DELETE ON "graduation_predicates" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "deleted_graduation_predicates";
DROP TABLE "graduation_predicates";

-- +goose StatementEnd
