-- +goose Up
-- +goose StatementBegin

CREATE TABLE "majors" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "faculty_id" uuid NOT NULL REFERENCES "faculties" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "name" character varying NOT NULL,
  "short_name" character varying NULL,
  "english_name" character varying NULL,
  "english_short_name" character varying NULL,
  "address" character varying NOT NULL DEFAULT '',
  "phone_number" character varying(20) NULL,
  "fax" character varying(20) NULL,
  "email" character varying(100) NULL,
  "contact_person" character varying NULL,
  "experiment_building_area" numeric(6,2) NULL,
  "lecture_hall_area" numeric(6,2) NULL,
  "lecture_hall_count" integer NULL,
  "laboratorium_area" numeric(6,2) NULL,
  "laboratorium_count" integer NULL,
  "permanent_lecturer_room_area" numeric(6,2) NULL,
  "administration_room_area" numeric(6,2) NULL,
  "book_count" integer NULL,
  "book_copy_count" integer NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_majors" BEFORE UPDATE ON "majors" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_majors" AS TABLE "majors" WITH NO DATA;
CREATE TRIGGER "soft_delete_majors" BEFORE DELETE ON "majors" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "deleted_majors";
DROP TABLE "majors";

-- +goose StatementEnd
