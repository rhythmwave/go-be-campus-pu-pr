-- +goose Up
-- +goose StatementBegin

CREATE TABLE "subjects" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "curriculum_id" uuid NOT NULL REFERENCES "curriculums" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "code" character varying NOT NULL UNIQUE,
  "name" character varying NOT NULL,
  "short_name" character varying NULL,
  "english_name" character varying NULL,
  "english_short_name" character varying NULL,
  "is_mandatory" boolean NOT NULL,
  "trait" character varying NOT NULL,
  "type" character varying NULL,
  "subject_category_id" uuid NOT NULL REFERENCES "subject_categories" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "curriculum_type" character varying NOT NULL,
  "theory_credit" integer NOT NULL DEFAULT 0,
  "practicum_credit" integer NOT NULL DEFAULT 0,
  "field_practicum_credit" integer NOT NULL DEFAULT 0,
  "semester_package" integer NOT NULL CHECK("semester_package" > 0),
  "repeat_course_limit" integer NOT NULL DEFAULT 0 CHECK (repeat_course_limit BETWEEN 0 AND 6),
  "is_active" boolean NOT NULL DEFAULT true,
  "has_lecture_unit" boolean NOT NULL DEFAULT false,
  "has_teaching_material" boolean NOT NULL DEFAULT false,
  "has_lecture_summary" boolean NOT NULL DEFAULT false,
  "supporting_lecturer_id" uuid NULL REFERENCES "lecturers" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "start_date" date NULL,
  "end_date" date NULL,
  "minimum_passing_grade_point" numeric(3,2) NOT NULL DEFAULT 0,
  "minimum_mandatory_credit_taken" integer NULL,
  "minimum_optional_credit_taken" integer NULL,
  "minimum_total_credit_taken" integer NULL,
  "minimum_mandatory_credit_passed" integer NULL,
  "minimum_optional_credit_passed" integer NULL,
  "minimum_total_credit_passed" integer NULL,
  "minimum_gpa" numeric(3,2) NULL,
  "abstraction" text NULL,
  "syllabus_path" character varying NULL,
  "syllabus_path_type" character varying NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  CHECK("theory_credit" + "practicum_credit" + "field_practicum_credit" > 0)
);
CREATE TRIGGER "updated_at_subjects" BEFORE UPDATE ON "subjects" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_subjects" AS TABLE "subjects" WITH NO DATA;
CREATE TRIGGER "soft_delete_subjects" BEFORE DELETE ON "subjects" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "deleted_subjects";
DROP TABLE "subjects";

-- +goose StatementEnd
