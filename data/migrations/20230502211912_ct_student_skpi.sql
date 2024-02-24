-- +goose Up
-- +goose StatementBegin

CREATE TABLE "student_skpi" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "student_id" uuid NOT NULL REFERENCES "students" ("id") ON DELETE CASCADE ON UPDATE CASCADE UNIQUE,
  "skpi_number" character varying NULL UNIQUE,
  "is_approved" boolean NOT NULL DEFAULT false,
  "achievement_path" character varying NULL,
  "achievement_path_type" character varying(20) NULL,
  "organization_path" character varying NULL,
  "organization_path_type" character varying(20) NULL,
  "certificate_path" character varying NULL,
  "certificate_path_type" character varying(20) NULL,
  "language_path" character varying NULL,
  "language_path_type" character varying(20) NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_student_skpi" BEFORE UPDATE ON "student_skpi" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_student_skpi" AS TABLE "student_skpi" WITH NO DATA;
CREATE TRIGGER "soft_delete_student_skpi" BEFORE DELETE ON "student_skpi" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

--------------------------------------------------------------------------------

CREATE TABLE "student_skpi_achievements" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "student_skpi_id" uuid NOT NULL REFERENCES "student_skpi" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "name" character varying NOT NULL,
  "year" integer NOT NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("student_skpi_id", "name")
);
CREATE TRIGGER "updated_at_student_skpi_achievements" BEFORE UPDATE ON "student_skpi_achievements" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_student_skpi_achievements" AS TABLE "student_skpi_achievements" WITH NO DATA;
CREATE TRIGGER "soft_delete_student_skpi_achievements" BEFORE DELETE ON "student_skpi_achievements" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

--------------------------------------------------------------------------------

CREATE TABLE "student_skpi_organizations" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "student_skpi_id" uuid NOT NULL REFERENCES "student_skpi" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "name" character varying NOT NULL,
  "position" character varying NOT NULL,
  "service_length" integer NOT NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("student_skpi_id", "name")
);
CREATE TRIGGER "updated_at_student_skpi_organizations" BEFORE UPDATE ON "student_skpi_organizations" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_student_skpi_organizations" AS TABLE "student_skpi_organizations" WITH NO DATA;
CREATE TRIGGER "soft_delete_student_skpi_organizations" BEFORE DELETE ON "student_skpi_organizations" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

--------------------------------------------------------------------------------

CREATE TABLE "student_skpi_certificates" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "student_skpi_id" uuid NOT NULL REFERENCES "student_skpi" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "name" character varying NOT NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("student_skpi_id", "name")
);
CREATE TRIGGER "updated_at_student_skpi_certificates" BEFORE UPDATE ON "student_skpi_certificates" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_student_skpi_certificates" AS TABLE "student_skpi_certificates" WITH NO DATA;
CREATE TRIGGER "soft_delete_student_skpi_certificates" BEFORE DELETE ON "student_skpi_certificates" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

--------------------------------------------------------------------------------

CREATE TABLE "student_skpi_character_buildings" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "student_skpi_id" uuid NOT NULL REFERENCES "student_skpi" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "name" character varying NOT NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("student_skpi_id", "name")
);
CREATE TRIGGER "updated_at_student_skpi_character_buildings" BEFORE UPDATE ON "student_skpi_character_buildings" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_student_skpi_character_buildings" AS TABLE "student_skpi_character_buildings" WITH NO DATA;
CREATE TRIGGER "soft_delete_student_skpi_character_buildings" BEFORE DELETE ON "student_skpi_character_buildings" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

--------------------------------------------------------------------------------

CREATE TABLE "student_skpi_internships" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "student_skpi_id" uuid NOT NULL REFERENCES "student_skpi" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "name" character varying NOT NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("student_skpi_id", "name")
);
CREATE TRIGGER "updated_at_student_skpi_internships" BEFORE UPDATE ON "student_skpi_internships" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_student_skpi_internships" AS TABLE "student_skpi_internships" WITH NO DATA;
CREATE TRIGGER "soft_delete_student_skpi_internships" BEFORE DELETE ON "student_skpi_internships" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

--------------------------------------------------------------------------------

CREATE TABLE "student_skpi_languages" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "student_skpi_id" uuid NOT NULL REFERENCES "student_skpi" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "name" character varying NOT NULL,
  "score" character varying NOT NULL,
  "date" date NOT NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("student_skpi_id", "name")
);
CREATE TRIGGER "updated_at_student_skpi_languages" BEFORE UPDATE ON "student_skpi_languages" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_student_skpi_languages" AS TABLE "student_skpi_languages" WITH NO DATA;
CREATE TRIGGER "soft_delete_student_skpi_languages" BEFORE DELETE ON "student_skpi_languages" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin



DROP TABLE "deleted_student_skpi_achievements";
DROP TABLE "student_skpi_achievements";

DROP TABLE "deleted_student_skpi_organizations";
DROP TABLE "student_skpi_organizations";

DROP TABLE "deleted_student_skpi_certificates";
DROP TABLE "student_skpi_certificates";

DROP TABLE "deleted_student_skpi_character_buildings";
DROP TABLE "student_skpi_character_buildings";

DROP TABLE "deleted_student_skpi_internships";
DROP TABLE "student_skpi_internships";

DROP TABLE "deleted_student_skpi_languages";
DROP TABLE "student_skpi_languages";

DROP TABLE "deleted_student_skpi";
DROP TABLE "student_skpi";


-- +goose StatementEnd
