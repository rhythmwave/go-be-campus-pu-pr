-- +goose Up
-- +goose StatementBegin

CREATE TABLE "buildings" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "faculty_id" uuid NULL REFERENCES "faculties" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "major_id" uuid NULL REFERENCES "majors" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "code" character varying NOT NULL UNIQUE,
  "name" character varying NOT NULL UNIQUE,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  CHECK(num_nonnulls(faculty_id, major_id) = 1)
);
CREATE TRIGGER "updated_at_buildings" BEFORE UPDATE ON "buildings" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_buildings" AS TABLE "buildings" WITH NO DATA;
CREATE TRIGGER "soft_delete_buildings" BEFORE DELETE ON "buildings" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "deleted_buildings";
DROP TABLE "buildings";

-- +goose StatementEnd
