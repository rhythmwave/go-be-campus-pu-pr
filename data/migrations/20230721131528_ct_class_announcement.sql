-- +goose Up
-- +goose StatementBegin

ALTER TABLE "classes" ADD COLUMN "total_announcement" integer NOT NULL DEFAULT 0;
ALTER TABLE "deleted_classes" ADD COLUMN "total_announcement" integer NULL;

CREATE TABLE "class_announcements" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "lecturer_id" uuid NOT NULL REFERENCES "lecturers" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "class_id" uuid NOT NULL REFERENCES "classes" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "title" character varying NOT NULL,
  "content" text NOT NULL,
  "file_path" character varying NULL,
  "file_path_type" character varying NULL,
  "start_time" timestamp NULL,
  "end_time" timestamp NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("class_id", "title")
);
CREATE TRIGGER "updated_at_class_announcements" BEFORE UPDATE ON "class_announcements" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_class_announcements" AS TABLE "class_announcements" WITH NO DATA;
CREATE TRIGGER "soft_delete_class_announcements" BEFORE DELETE ON "class_announcements" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

--------------------------------------------

CREATE FUNCTION classes_total_announcement_func()
RETURNS TRIGGER AS $$
BEGIN
  IF OLD.class_id IS NOT NULL THEN
    UPDATE classes SET total_announcement = total_announcement - 1 WHERE id = OLD.class_id;
  END IF;
  IF NEW.class_id IS NOT NULL THEN
    UPDATE classes SET total_announcement = total_announcement + 1 WHERE id = NEW.class_id;
  END IF;

  IF NEW.class_id IS NOT NULL THEN
    RETURN NEW;
  ELSE
    RETURN OLD;
  END IF;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER classes_total_announcement AFTER INSERT OR DELETE ON class_announcements FOR EACH ROW EXECUTE PROCEDURE classes_total_announcement_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER classes_total_announcement ON class_announcements;
DROP FUNCTION classes_total_announcement_func();

DROP TABLE "class_announcements";
DROP TABLE "deleted_class_announcements";

ALTER TABLE "classes" DROP COLUMN "total_announcement";
ALTER TABLE "deleted_classes" DROP COLUMN "total_announcement";

-- +goose StatementEnd
