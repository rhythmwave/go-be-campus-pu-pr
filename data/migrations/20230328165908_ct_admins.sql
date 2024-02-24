-- +goose Up
-- +goose StatementBegin

CREATE TABLE "admins"
(
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" character varying NOT NULL,
  "username" character varying(100) NOT NULL UNIQUE,
  "password" character varying NOT NULL,
  "role_id" uuid NULL REFERENCES "roles" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_admins" BEFORE UPDATE ON "admins" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_admins" AS TABLE "admins" WITH NO DATA;
CREATE TRIGGER "soft_delete_admins" BEFORE DELETE ON "admins" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

------------------------------------------

ALTER TABLE "roles" 
  ADD COLUMN "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  ADD COLUMN "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE;

ALTER TABLE "deleted_roles" 
  ADD COLUMN "created_by" uuid NULL,
  ADD COLUMN "updated_by" uuid NULL;

------------------------------------------

INSERT INTO "admins" ("name", "username", "password") VALUES
('super_admin',	'akdroot',	'$2a$10$7dBcJCNxbl12LXBhiBtJsO/cgQ6IEn.qdej8kJneSoqggEbE/YKNK');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "roles" 
  DROP COLUMN "created_by",
  DROP COLUMN "updated_by";

ALTER TABLE "deleted_roles" 
  DROP COLUMN "created_by",
  DROP COLUMN "updated_by";

---------------------------------------------

DROP TABLE "deleted_admins";
DROP TABLE "admins";

-- +goose StatementEnd
