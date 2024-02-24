-- +goose Up
-- +goose StatementBegin

CREATE TABLE "rooms" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "building_id" uuid NOT NULL REFERENCES "buildings" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "study_program_id" uuid NULL REFERENCES "study_programs" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "code" character varying NOT NULL UNIQUE,
  "name" character varying NULL UNIQUE,
  "capacity" integer NULL,
  "exam_capacity" integer NULL,
  "is_usable" boolean NOT NULL DEFAULT true,
  "area" numeric(6,2) NULL,
  "phone_number" character varying(20) NULL,
  "facility" text NULL,
  "remarks" text NULL,
  "is_laboratory" boolean NOT NULL DEFAULT false,
  "purpose" character varying NOT NULL,
  "owner" character varying NULL,
  "location" character varying NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  CHECK((is_laboratory IS true AND study_program_id IS NOT NULL) OR (is_laboratory IS false AND study_program_id IS NULL))
);
CREATE TRIGGER "updated_at_rooms" BEFORE UPDATE ON "rooms" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_rooms" AS TABLE "rooms" WITH NO DATA;
CREATE TRIGGER "soft_delete_rooms" BEFORE DELETE ON "rooms" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "deleted_rooms";
DROP TABLE "rooms";

-- +goose StatementEnd
