-- +goose Up
-- +goose StatementBegin

CREATE TABLE "expertise_groups" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" character varying NOT NULL UNIQUE,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_expertise_groups" BEFORE UPDATE ON "expertise_groups" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_expertise_groups" AS TABLE "expertise_groups" WITH NO DATA;
CREATE TRIGGER "soft_delete_expertise_groups" BEFORE DELETE ON "expertise_groups" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-------------------------------------------------

ALTER TABLE "lecturers" ADD COLUMN "expertise_group_id" uuid NULL REFERENCES "expertise_groups" ("id") ON DELETE SET NULL ON UPDATE CASCADE;
ALTER TABLE "deleted_lecturers" ADD COLUMN "expertise_group_id" uuid NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "lecturers" DROP COLUMN "expertise_group_id";
ALTER TABLE "deleted_lecturers" DROP COLUMN "expertise_group_id";

DROP TABLE "expertise_groups";
DROP TABLE "deleted_expertise_groups";

-- +goose StatementEnd
