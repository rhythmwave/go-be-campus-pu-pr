-- +goose Up
-- +goose StatementBegin

CREATE TABLE "lesson_plans" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "subject_id" uuid NOT NULL REFERENCES "subjects" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "meeting_order" integer NOT NULL,
  "lesson" text NOT NULL,
  "english_lesson" text NOT NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("subject_id", "meeting_order")
);
CREATE TRIGGER "updated_at_lesson_plans" BEFORE UPDATE ON "lesson_plans" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_lesson_plans" AS TABLE "lesson_plans" WITH NO DATA;
CREATE TRIGGER "soft_delete_lesson_plans" BEFORE DELETE ON "lesson_plans" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "deleted_lesson_plans";
DROP TABLE "lesson_plans";

-- +goose StatementEnd
