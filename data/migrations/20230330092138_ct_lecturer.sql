-- +goose Up
-- +goose StatementBegin

CREATE TABLE "lecturers" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "id_national_lecturer" character varying NOT NULL,
  "name" character varying NOT NULL,
  "front_title" character varying NULL,
  "back_degree" character varying NULL,
  "study_program_id" uuid NULL REFERENCES "study_programs" ("id") ON DELETE SET NULL ON UPDATE CASCADE,
  "id_number" character(16) NULL,
  "birth_date" date NULL,
  "birth_regency_id" integer NULL REFERENCES "regencies" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "id_employee" character varying NULL,
  "stambuk" character varying NULL,
  "sex" character(1) NULL,
  "blood_type" character varying(3) NULL,
  "religion" character varying(20) NULL,
  "marital_status" character varying NULL,
  "address" character varying NULL,
  "regency_id" integer NULL REFERENCES "regencies" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "postal_code" character varying(10) NULL,
  "phone_number" character varying(20) NULL,
  "fax" character varying(20) NULL,
  "mobile_phone_number" character varying(20) NULL,
  "office_phone_number" character varying(20) NULL,
  "employee_type" character varying(100) NULL,
  "sk_cpns_number" character varying NULL,
  "sk_cpns_date" date NULL,
  "tmt_cpns_date" date NULL,
  "cpns_category" character varying NULL,
  "cpns_duration_month" integer NULL,
  "pre_position_date" date NULL,
  "sk_pns_number" character varying NULL,
  "sk_pns_date" date NULL,
  "tmt_pns_date" date NULL,
  "pns_category" character varying NULL,
  "pns_oath_date" date NULL,
  "join_date" date NULL,
  "end_date" date NULL,
  "taspen_number" character varying NULL,
  "former_instance" character varying NULL,
  "remarks" character varying NULL,
  "lecturer_number" character varying NULL,
  "academic_position" character varying NULL,
  "employee_status" character varying NULL,
  "expertise" character varying NULL,
  "highest_degree" character varying NULL,
  "instance_code" character varying NULL,
  "teaching_certificate_number" character varying NULL,
  "teaching_permit_number" character varying NULL,
  "status" character varying NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_lecturers" BEFORE UPDATE ON "lecturers" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_lecturers" AS TABLE "lecturers" WITH NO DATA;
CREATE TRIGGER "soft_delete_lecturers" BEFORE DELETE ON "lecturers" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

----------------

ALTER TABLE "study_programs" ADD COLUMN "head_lecturer_id" uuid NULL REFERENCES "lecturers" ("id") ON DELETE SET NULL ON UPDATE CASCADE;
ALTER TABLE "deleted_study_programs" ADD COLUMN "head_lecturer_id" uuid NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "study_programs" DROP COLUMN "head_lecturer_id";
ALTER TABLE "deleted_study_programs" DROP COLUMN "head_lecturer_id";

DROP TABLE "deleted_lecturers";
DROP TABLE "lecturers";

-- +goose StatementEnd
