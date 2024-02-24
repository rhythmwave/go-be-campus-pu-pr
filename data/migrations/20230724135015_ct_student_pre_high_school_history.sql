-- +goose Up
-- +goose StatementBegin

CREATE TABLE "student_pre_high_school_histories" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "student_id" uuid NOT NULL REFERENCES "students" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "level" character varying NOT NULL,
  "name" character varying NOT NULL,
  "graduation_year" character(4) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("student_id", "level")
);
CREATE TRIGGER "updated_at_student_pre_high_school_histories" BEFORE UPDATE ON "student_pre_high_school_histories" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_student_pre_high_school_histories" AS TABLE "student_pre_high_school_histories" WITH NO DATA;
CREATE TRIGGER "soft_delete_student_pre_high_school_histories" BEFORE DELETE ON "student_pre_high_school_histories" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "deleted_student_pre_high_school_histories";
DROP TABLE "student_pre_high_school_histories";

-- +goose StatementEnd
