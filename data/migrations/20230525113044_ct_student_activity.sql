-- +goose Up
-- +goose StatementBegin

CREATE TABLE "student_activities" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "study_program_id" uuid NOT NULL REFERENCES "study_programs" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "semester_id" uuid NOT NULL REFERENCES "semesters" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "title" character varying NOT NULL,
  "activity_type" character varying NOT NULL,
  "location" character varying NULL,
  "decision_number" character varying NULL,
  "decision_date" date NULL,
  "remarks" character varying NULL,
  "is_group_activity" boolean NOT NULL DEFAULT false,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_student_activities" BEFORE UPDATE ON "student_activities" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_student_activities" AS TABLE "student_activities" WITH NO DATA;
CREATE TRIGGER "soft_delete_student_activities" BEFORE DELETE ON "student_activities" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

------------------------------------------------------------------

CREATE TABLE "student_activity_participants" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "student_activity_id" uuid NOT NULL REFERENCES "student_activities" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "student_id" uuid NOT NULL REFERENCES "students" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "role" character varying NOT NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("student_activity_id", "student_id")
);
CREATE TRIGGER "updated_at_student_activity_participants" BEFORE UPDATE ON "student_activity_participants" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_student_activity_participants" AS TABLE "student_activity_participants" WITH NO DATA;
CREATE TRIGGER "soft_delete_student_activity_participants" BEFORE DELETE ON "student_activity_participants" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

------------------------------------------------------------------

CREATE TYPE "student_activity_lecturers_role" AS ENUM ('mentor', 'examiner');

CREATE TABLE "student_activity_lecturers" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "student_activity_id" uuid NOT NULL REFERENCES "student_activities" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "lecturer_id" uuid NOT NULL REFERENCES "lecturers" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "role" student_activity_lecturers_role NOT NULL,
  "activity_category" character varying NOT NULL,
  "sort" integer NOT NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("student_activity_id", "lecturer_id", "role")
);
CREATE TRIGGER "updated_at_student_activity_lecturers" BEFORE UPDATE ON "student_activity_lecturers" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_student_activity_lecturers" AS TABLE "student_activity_lecturers" WITH NO DATA;
CREATE TRIGGER "soft_delete_student_activity_lecturers" BEFORE DELETE ON "student_activity_lecturers" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "deleted_student_activity_participants";
DROP TABLE "student_activity_participants";

DROP TABLE "deleted_student_activity_lecturers";
DROP TABLE "student_activity_lecturers";

DROP TYPE "student_activity_lecturers_role";

DROP TABLE "deleted_student_activities";
DROP TABLE "student_activities";

-- +goose StatementEnd
