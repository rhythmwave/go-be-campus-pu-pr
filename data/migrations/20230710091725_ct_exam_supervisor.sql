-- +goose Up
-- +goose StatementBegin
CREATE TABLE "exam_supervisor_roles" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" character varying NOT NULL UNIQUE,
  "sort" integer NOT NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_exam_supervisor_roles" BEFORE UPDATE ON "exam_supervisor_roles" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_exam_supervisor_roles" AS TABLE "exam_supervisor_roles" WITH NO DATA;
CREATE TRIGGER "soft_delete_exam_supervisor_roles" BEFORE DELETE ON "exam_supervisor_roles" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

--------------------------------------------------------

CREATE TABLE "exam_supervisors" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "id_national_lecturer" character varying NOT NULL,
  "study_program_id" uuid NULL REFERENCES "study_programs" ("id") ON DELETE SET NULL ON UPDATE CASCADE,
  "name" character varying NOT NULL,
  "front_title" character varying NULL,
  "back_degree" character varying NULL,
  "id_number" character(16) NULL,
  "birth_date" date NULL,
  "birth_regency_id" integer NULL REFERENCES "regencies" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
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
  "employee_status" character varying NULL,
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
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_exam_supervisors" BEFORE UPDATE ON "exam_supervisors" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_exam_supervisors" AS TABLE "exam_supervisors" WITH NO DATA;
CREATE TRIGGER "soft_delete_exam_supervisors" BEFORE DELETE ON "exam_supervisors" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin


DROP TABLE "exam_supervisors";
DROP TABLE "deleted_exam_supervisors";

DROP TABLE "exam_supervisor_roles";
DROP TABLE "deleted_exam_supervisor_roles";

-- +goose StatementEnd
