-- +goose Up
-- +goose StatementBegin

CREATE TABLE "faculties" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
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
CREATE TRIGGER "updated_at_faculties" BEFORE UPDATE ON "faculties" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_faculties" AS TABLE "faculties" WITH NO DATA;
CREATE TRIGGER "soft_delete_faculties" BEFORE DELETE ON "faculties" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "deleted_faculties";
DROP TABLE "faculties";

-- +goose StatementEnd
