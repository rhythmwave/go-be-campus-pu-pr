-- +goose Up
-- +goose StatementBegin

CREATE TABLE "graduation_sessions" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "session_year" integer NOT NULL,
  "session_number" integer NOT NULL,
  "session_date" date NOT NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE(session_year, session_number, session_date)
);
CREATE TRIGGER "updated_at_graduation_sessions" BEFORE UPDATE ON "graduation_sessions" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_graduation_sessions" AS TABLE "graduation_sessions" WITH NO DATA;
CREATE TRIGGER "soft_delete_graduation_sessions" BEFORE DELETE ON "graduation_sessions" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

CREATE TABLE "graduation_students" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "student_id" uuid NOT NULL REFERENCES "students" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "application_date" date NOT NULL,
  "graduation_session_id" uuid NOT NULL REFERENCES "graduation_sessions" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("student_id")
);
CREATE TRIGGER "updated_at_graduation_students" BEFORE UPDATE ON "graduation_students" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_graduation_students" AS TABLE "graduation_students" WITH NO DATA;
CREATE TRIGGER "soft_delete_graduation_students" BEFORE DELETE ON "graduation_students" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "graduation_students";
DROP TABLE "deleted_graduation_students";

DROP TABLE "deleted_graduation_sessions";
DROP TABLE "graduation_sessions";

-- +goose StatementEnd
