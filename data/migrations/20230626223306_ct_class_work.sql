-- +goose Up
-- +goose StatementBegin

ALTER TABLE "classes" ADD COLUMN "total_work" integer NOT NULL DEFAULT 0;
ALTER TABLE "deleted_classes" ADD COLUMN "total_work" integer NULL;

CREATE TABLE "class_works" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "lecturer_id" uuid NOT NULL REFERENCES "lecturers" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "class_id" uuid NOT NULL REFERENCES "classes" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "title" character varying NOT NULL,
  "abstraction" text NULL, 
  "file_path" character varying NULL,
  "file_path_type" character varying(20) NULL,
  "start_time" timestamp NOT NULL,
  "end_time" timestamp NOT NULL,
  "total_submission" integer NOT NULL DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("class_id", "title"),
  CHECK("start_time" < "end_time")
);
CREATE TRIGGER "updated_at_class_works" BEFORE UPDATE ON "class_works" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_class_works" AS TABLE "class_works" WITH NO DATA;
CREATE TRIGGER "soft_delete_class_works" BEFORE DELETE ON "class_works" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

--------------------------------------------

CREATE FUNCTION classes_total_work_func()
RETURNS TRIGGER AS $$
BEGIN
  IF OLD.class_id IS NOT NULL THEN
    UPDATE classes SET total_work = total_work - 1 WHERE id = OLD.class_id;
  END IF;
  IF NEW.class_id IS NOT NULL THEN
    UPDATE classes SET total_work = total_work + 1 WHERE id = NEW.class_id;
  END IF;

  IF NEW.class_id IS NOT NULL THEN
    RETURN NEW;
  ELSE
    RETURN OLD;
  END IF;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER classes_total_work AFTER INSERT OR DELETE ON class_works FOR EACH ROW EXECUTE PROCEDURE classes_total_work_func();

--------------------------------------------

CREATE TABLE "class_work_submissions" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "class_work_id" uuid NOT NULL REFERENCES "class_works" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "student_id" uuid NOT NULL REFERENCES "students" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "file_path" character varying NOT NULL,
  "file_path_type" character varying(20) NOT NULL,
  "point" numeric(4,2) NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE("class_work_id", "student_id")
);
CREATE TABLE "deleted_class_work_submissions" AS TABLE "class_work_submissions" WITH NO DATA;
CREATE TRIGGER "soft_delete_class_work_submissions" BEFORE DELETE ON "class_work_submissions" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

--------------------------------------------

CREATE FUNCTION class_work_total_submissions_func()
RETURNS TRIGGER AS $$
BEGIN
  IF OLD.class_work_id IS NOT NULL THEN
    UPDATE class_works SET total_submission = total_submission - 1 WHERE id = OLD.class_work_id;
  END IF;
  IF NEW.class_work_id IS NOT NULL THEN
    UPDATE class_works SET total_submission = total_submission + 1 WHERE id = NEW.class_work_id;
  END IF;

  IF NEW.class_work_id IS NOT NULL THEN
    RETURN NEW;
  ELSE
    RETURN OLD;
  END IF;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER class_work_total_submissions AFTER INSERT OR DELETE ON class_work_submissions FOR EACH ROW EXECUTE PROCEDURE class_work_total_submissions_func();


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER class_work_total_submissions ON class_work_submissions;
DROP FUNCTION class_work_total_submissions_func();

DROP TABLE "class_work_submissions";
DROP TABLE "deleted_class_work_submissions";

DROP TRIGGER classes_total_work ON class_works;
DROP FUNCTION classes_total_work_func();

DROP TABLE "class_works";
DROP TABLE "deleted_class_works";

ALTER TABLE "classes" DROP COLUMN "total_work";
ALTER TABLE "deleted_classes" DROP COLUMN "total_work";

-- +goose StatementEnd
