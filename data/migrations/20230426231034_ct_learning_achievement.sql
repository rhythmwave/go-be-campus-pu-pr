-- +goose Up
-- +goose StatementBegin

CREATE TABLE "learning_achievement_categories" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "curriculum_id" uuid NOT NULL REFERENCES "curriculums" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "name" character varying NOT NULL,
  "english_name" character varying NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_learning_achievement_categories" BEFORE UPDATE ON "learning_achievement_categories" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_learning_achievement_categories" AS TABLE "learning_achievement_categories" WITH NO DATA;
CREATE TRIGGER "soft_delete_learning_achievement_categories" BEFORE DELETE ON "learning_achievement_categories" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

--------------------------------------------------------------------------------

CREATE TABLE "learning_achievements" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "learning_achievement_category_id" uuid NOT NULL REFERENCES "learning_achievement_categories" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "name" character varying NOT NULL,
  "english_name" character varying NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_learning_achievements" BEFORE UPDATE ON "learning_achievements" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_learning_achievements" AS TABLE "learning_achievements" WITH NO DATA;
CREATE TRIGGER "soft_delete_learning_achievements" BEFORE DELETE ON "learning_achievements" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "deleted_learning_achievements";
DROP TABLE "learning_achievements";

DROP TABLE "deleted_learning_achievement_categories";
DROP TABLE "learning_achievement_categories";

-- +goose StatementEnd
