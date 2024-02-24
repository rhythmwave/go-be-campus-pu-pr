-- +goose Up
-- +goose StatementBegin

CREATE TABLE "academic_guidance_sessions" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "academic_guidance_id" uuid NOT NULL REFERENCES "academic_guidances" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "subject" character varying NOT NULL,
  "session_date" date NOT NULL,
  "summary" text NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_academic_guidance_sessions" BEFORE UPDATE ON "academic_guidance_sessions" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_academic_guidance_sessions" AS TABLE "academic_guidance_sessions" WITH NO DATA;
CREATE TRIGGER "soft_delete_academic_guidance_sessions" BEFORE DELETE ON "academic_guidance_sessions" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

----------------------------------------------------------------------------------------

CREATE TABLE "academic_guidance_session_files" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "academic_guidance_session_id" uuid NOT NULL REFERENCES "academic_guidance_sessions" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "title" character varying NOT NULL,
  "file_path" character varying NOT NULL,
  "file_path_type" character varying(20) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE(file_path, file_path_type)
);
CREATE TABLE "deleted_academic_guidance_session_files" AS TABLE "academic_guidance_session_files" WITH NO DATA;
CREATE TRIGGER "soft_delete_academic_guidance_session_files" BEFORE DELETE ON "academic_guidance_session_files" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

----------------------------------------------------------------------------------------

CREATE TABLE "academic_guidance_session_students" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "academic_guidance_session_id" uuid NOT NULL REFERENCES "academic_guidance_sessions" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "student_id" uuid NOT NULL REFERENCES "students" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE(academic_guidance_session_id, student_id)
);
CREATE TABLE "deleted_academic_guidance_session_students" AS TABLE "academic_guidance_session_students" WITH NO DATA;
CREATE TRIGGER "soft_delete_academic_guidance_session_students" BEFORE DELETE ON "academic_guidance_session_students" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "deleted_academic_guidance_session_students";
DROP TABLE "academic_guidance_session_students";

DROP TABLE "deleted_academic_guidance_session_files";
DROP TABLE "academic_guidance_session_files";

DROP TABLE "deleted_academic_guidance_sessions";
DROP TABLE "academic_guidance_sessions";

-- +goose StatementEnd
