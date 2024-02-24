-- +goose Up
-- +goose StatementBegin

CREATE TABLE "classes" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "subject_id" uuid NOT NULL REFERENCES "subjects" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "semester_id" uuid NOT NULL REFERENCES "semesters" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "name" character varying NOT NULL,
  "scope" character varying NOT NULL,
  "is_online" boolean NOT NULL,
  "is_offline" boolean NOT NULL,
  "minimum_participant" integer NOT NULL,
  "maximum_participant" integer NOT NULL,
  "total_participant" integer NOT NULL DEFAULT 0,
  "remarks" character varying NULL,
  "is_active" boolean NOT NULL DEFAULT true,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("subject_id", "semester_id", "name"),
  CHECK("minimum_participant" < "maximum_participant"),
  CHECK("total_participant" <= "maximum_participant")
);
CREATE TRIGGER "updated_at_classes" BEFORE UPDATE ON "classes" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_classes" AS TABLE "classes" WITH NO DATA;
CREATE TRIGGER "soft_delete_classes" BEFORE DELETE ON "classes" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

------------------------------------------------------------------------

CREATE TABLE "class_lecturers" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "class_id" uuid NOT NULL REFERENCES "classes" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "lecturer_id" uuid NOT NULL REFERENCES "lecturers" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "is_grading_responsible" boolean NOT NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("class_id", "lecturer_id")
);
CREATE TRIGGER "updated_at_class_lecturers" BEFORE UPDATE ON "class_lecturers" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_class_lecturers" AS TABLE "class_lecturers" WITH NO DATA;
CREATE TRIGGER "soft_delete_class_lecturers" BEFORE DELETE ON "class_lecturers" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "class_lecturers";
DROP TABLE "deleted_class_lecturers";

DROP TABLE "classes";
DROP TABLE "deleted_classes";

-- +goose StatementEnd
