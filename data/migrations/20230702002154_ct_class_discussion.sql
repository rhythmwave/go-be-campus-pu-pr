-- +goose Up
-- +goose StatementBegin

ALTER TABLE "classes" ADD COLUMN "total_discussion" integer NOT NULL DEFAULT 0;
ALTER TABLE "deleted_classes" ADD COLUMN "total_discussion" integer NULL;

CREATE TABLE "class_discussions" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "lecturer_id" uuid NOT NULL REFERENCES "lecturers" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "class_id" uuid NOT NULL REFERENCES "classes" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "title" character varying NOT NULL,
  "abstraction" character varying NOT NULL, 
  "total_comment" integer NOT NULL DEFAULT 0,
  "last_comment" text NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("class_id", "title")
);
CREATE TRIGGER "updated_at_class_discussions" BEFORE UPDATE ON "class_discussions" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_class_discussions" AS TABLE "class_discussions" WITH NO DATA;
CREATE TRIGGER "soft_delete_class_discussions" BEFORE DELETE ON "class_discussions" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

--------------------------------------------

CREATE FUNCTION classes_total_discussion_func()
RETURNS TRIGGER AS $$
BEGIN
  IF OLD.class_id IS NOT NULL THEN
    UPDATE classes SET total_discussion = total_discussion - 1 WHERE id = OLD.class_id;
  END IF;
  IF NEW.class_id IS NOT NULL THEN
    UPDATE classes SET total_discussion = total_discussion + 1 WHERE id = NEW.class_id;
  END IF;

  IF NEW.class_id IS NOT NULL THEN
    RETURN NEW;
  ELSE
    RETURN OLD;
  END IF;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER classes_total_discussion AFTER INSERT OR DELETE ON class_discussions FOR EACH ROW EXECUTE PROCEDURE classes_total_discussion_func();

--------------------------------------------

CREATE TABLE "class_discussion_comments" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "lecturer_id" uuid NULL REFERENCES "lecturers" ("id") ON DELETE SET NULL ON UPDATE CASCADE,
  "student_id" uuid NULL REFERENCES "students" ("id") ON DELETE SET NULL ON UPDATE CASCADE,
  "class_discussion_id" uuid NOT NULL REFERENCES "class_discussions" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "comment" text NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  CHECK(num_nonnulls(lecturer_id, student_id) IN (0,1))
);
CREATE TRIGGER "updated_at_class_discussion_comments" BEFORE UPDATE ON "class_discussion_comments" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_class_discussion_comments" AS TABLE "class_discussion_comments" WITH NO DATA;
CREATE TRIGGER "soft_delete_class_discussion_comments" BEFORE DELETE ON "class_discussion_comments" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

--------------------------------------------

CREATE FUNCTION class_discussion_comment_func()
RETURNS TRIGGER AS $$
DECLARE previousComment text;
BEGIN
  IF OLD.class_discussion_id IS NOT NULL THEN
    SELECT comment INTO previousComment FROM class_discussion_comments WHERE class_discussion_id = OLD.class_discussion_id ORDER BY created_at DESC LIMIT 1;
    UPDATE class_discussions SET total_comment = total_comment - 1, last_comment = previousComment WHERE id = OLD.class_discussion_id;
  END IF;
  IF NEW.class_discussion_id IS NOT NULL THEN
    UPDATE class_discussions SET total_comment = total_comment + 1, last_comment = NEW.comment WHERE id = NEW.class_discussion_id;
  END IF;

  IF NEW.class_discussion_id IS NOT NULL THEN
    RETURN NEW;
  ELSE
    RETURN OLD;
  END IF;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER class_discussion_comment AFTER INSERT OR DELETE ON class_discussion_comments FOR EACH ROW EXECUTE PROCEDURE class_discussion_comment_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER class_discussion_comment ON class_discussion_comments;
DROP FUNCTION class_discussion_comment_func();

DROP TABLE "deleted_class_discussion_comments";
DROP TABLE "class_discussion_comments";

DROP TRIGGER classes_total_discussion ON class_discussions;
DROP FUNCTION classes_total_discussion_func();

DROP TABLE "class_discussions";
DROP TABLE "deleted_class_discussions";

ALTER TABLE "classes" DROP COLUMN "total_discussion";
ALTER TABLE "deleted_classes" DROP COLUMN "total_discussion";

-- +goose StatementEnd
