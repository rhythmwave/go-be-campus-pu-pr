-- +goose Up
-- +goose StatementBegin

DROP TABLE "deleted_permissions";

DROP TRIGGER "updated_at_permissions" ON "permissions";
DROP TRIGGER "soft_delete_permissions" ON "permissions";

ALTER TABLE "permissions"
  DROP COLUMN "updated_by",
  DROP COLUMN "updated_at";

CREATE RULE "protect_permissions_update" AS
ON UPDATE TO "permissions"
DO INSTEAD NOTHING;

CREATE RULE "protect_permissions_delete" AS
ON DELETE TO "permissions"
DO INSTEAD NOTHING;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP RULE "protect_permissions_update" ON "permissions";

DROP RULE "protect_permissions_delete" ON "permissions";

ALTER TABLE "permissions"
  ADD COLUMN "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  ADD COLUMN "updated_at" timestamp NULL;

CREATE TRIGGER "updated_at_permissions" BEFORE UPDATE ON "permissions" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_permissions" AS TABLE "permissions" WITH NO DATA;
CREATE TRIGGER "soft_delete_permissions" BEFORE DELETE ON "permissions" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd
