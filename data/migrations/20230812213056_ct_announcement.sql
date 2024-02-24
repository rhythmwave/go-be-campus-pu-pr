-- +goose Up
-- +goose StatementBegin

CREATE TYPE "announcement_type" AS ENUM ('academic_information', 'student_activity', 'registration', 'event');

CREATE TABLE "announcements" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "type" announcement_type NOT NULL,
  "title" character varying NOT NULL,
  "announcement_date" date NULL,
  "file_path" character varying NULL,
  "file_path_type" character varying(20) NULL,
  "file_title" character varying NULL,
  "content" text NULL,
  "for_lecturer" boolean NOT NULL,
  "for_student" boolean NOT NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  CHECK("for_lecturer" IS true OR "for_student" IS true)
);
CREATE TRIGGER "updated_at_announcements" BEFORE UPDATE ON "announcements" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_announcements" AS TABLE "announcements" WITH NO DATA;
CREATE TRIGGER "soft_delete_announcements" BEFORE DELETE ON "announcements" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

----------------------------------------------------------------

CREATE TABLE "announcement_study_programs" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "announcement_id" uuid NOT NULL REFERENCES "announcements" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "study_program_id" uuid NOT NULL REFERENCES "study_programs" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE("announcement_id", "study_program_id")
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "announcement_study_programs";
DROP TABLE "deleted_announcements";
DROP TABLE "announcements";
DROP TYPE "announcement_type";

-- +goose StatementEnd
