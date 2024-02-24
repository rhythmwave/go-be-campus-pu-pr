-- +goose Up
-- +goose StatementBegin

CREATE TYPE "thesis_status" AS ENUM ('DIAJUKAN', 'SEDANG DIKERJAKAN', 'BERHASIL DISELESAIKAN', 'DIBATALKAN');

CREATE TABLE "theses" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "student_id" uuid NOT NULL REFERENCES "students" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "start_semester_id" uuid NOT NULL REFERENCES "semesters" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "finish_semester_id" uuid NULL REFERENCES "semesters" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "topic" character varying NOT NULL,
  "title" character varying NOT NULL,
  "english_title" character varying NULL,
  "start_date" date NOT NULL,
  "finish_date" date NULL,
  "remarks" character varying NULL,
  "status" thesis_status NOT NULL DEFAULT 'DIAJUKAN'::thesis_status,
  "is_joint_thesis" boolean NOT NULL,
  "proposal_seminar_date" date NULL,
  "proposal_certificate_number" character varying NULL,
  "proposal_certificate_date" date NULL,
  "thesis_defense_count" integer NOT NULL DEFAULT 0,
  "grade_point" numeric(3,2) NOT NULL DEFAULT 0,
  "grade_code" character varying(2) NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("student_id", "title")
);
CREATE TRIGGER "updated_at_theses" BEFORE UPDATE ON "theses" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_theses" AS TABLE "theses" WITH NO DATA;
CREATE TRIGGER "soft_delete_theses" BEFORE DELETE ON "theses" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

------------------------------------------------------

CREATE TABLE "thesis_files" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "thesis_id" uuid NOT NULL REFERENCES "theses" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "file_path" character varying NOT NULL,
  "file_path_type" character varying(20) NOT NULL,
  "description" character varying NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("thesis_id", "file_path")
);
CREATE TRIGGER "updated_at_thesis_files" BEFORE UPDATE ON "thesis_files" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_thesis_files" AS TABLE "thesis_files" WITH NO DATA;
CREATE TRIGGER "soft_delete_thesis_files" BEFORE DELETE ON "thesis_files" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

------------------------------------------------------

CREATE TABLE "thesis_supervisor_roles" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" character varying NOT NULL UNIQUE,
  "sort" INTEGER NOT NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_thesis_supervisor_roles" BEFORE UPDATE ON "thesis_supervisor_roles" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_thesis_supervisor_roles" AS TABLE "thesis_supervisor_roles" WITH NO DATA;
CREATE TRIGGER "soft_delete_thesis_supervisor_roles" BEFORE DELETE ON "thesis_supervisor_roles" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

------------------------------------------------------

CREATE TABLE "thesis_supervisors" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "thesis_id" uuid NOT NULL REFERENCES "theses" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "lecturer_id" uuid NOT NULL REFERENCES "lecturers" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "thesis_supervisor_role_id" uuid NOT NULL REFERENCES "thesis_supervisor_roles" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("thesis_id", "thesis_supervisor_role_id"),
  UNIQUE("thesis_id", "lecturer_id")
);
CREATE TRIGGER "updated_at_thesis_supervisors" BEFORE UPDATE ON "thesis_supervisors" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_thesis_supervisors" AS TABLE "thesis_supervisors" WITH NO DATA;
CREATE TRIGGER "soft_delete_thesis_supervisors" BEFORE DELETE ON "thesis_supervisors" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

------------------------------------------------------

CREATE TABLE "thesis_examiner_roles" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" character varying NOT NULL UNIQUE,
  "remarks" character varying NOT NULL UNIQUE,
  "sort" INTEGER NOT NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_thesis_examiner_roles" BEFORE UPDATE ON "thesis_examiner_roles" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_thesis_examiner_roles" AS TABLE "thesis_examiner_roles" WITH NO DATA;
CREATE TRIGGER "soft_delete_thesis_examiner_roles" BEFORE DELETE ON "thesis_examiner_roles" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

------------------------------------------------------

CREATE TABLE "thesis_defenses" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "thesis_id" uuid NOT NULL REFERENCES "theses" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "plan_date" date NOT NULL,
  "plan_start_time" integer NOT NULL,
  "plan_end_time" integer NOT NULL,
  "actual_date" date NULL,
  "actual_start_time" integer NULL,
  "actual_end_time" integer NULL,
  "room_id" uuid NOT NULL REFERENCES "rooms" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "revision" text NULL,
  "is_passed" boolean NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_thesis_defenses" BEFORE UPDATE ON "thesis_defenses" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_thesis_defenses" AS TABLE "thesis_defenses" WITH NO DATA;
CREATE TRIGGER "soft_delete_thesis_defenses" BEFORE DELETE ON "thesis_defenses" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

CREATE TABLE "thesis_examiners" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "thesis_defense_id" uuid NOT NULL REFERENCES "thesis_defenses" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "lecturer_id" uuid NOT NULL REFERENCES "lecturers" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "thesis_examiner_role_id" uuid NOT NULL REFERENCES "thesis_examiner_roles" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("thesis_defense_id", "thesis_examiner_role_id"),
  UNIQUE("thesis_defense_id", "lecturer_id")
);
CREATE TRIGGER "updated_at_thesis_examiners" BEFORE UPDATE ON "thesis_examiners" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_thesis_examiners" AS TABLE "thesis_examiners" WITH NO DATA;
CREATE TRIGGER "soft_delete_thesis_examiners" BEFORE DELETE ON "thesis_examiners" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "deleted_thesis_examiners";
DROP TABLE "thesis_examiners";

DROP TABLE "deleted_thesis_defenses";
DROP TABLE "thesis_defenses";

DROP TABLE "deleted_thesis_examiner_roles";
DROP TABLE "thesis_examiner_roles";

DROP TABLE "deleted_thesis_supervisors";
DROP TABLE "thesis_supervisors";

DROP TABLE "deleted_thesis_supervisor_roles";
DROP TABLE "thesis_supervisor_roles";

DROP TABLE "deleted_thesis_files";
DROP TABLE "thesis_files";

DROP TABLE "deleted_theses";
DROP TABLE "theses";

DROP TYPE "thesis_status";

-- +goose StatementEnd
