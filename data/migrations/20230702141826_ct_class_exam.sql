-- +goose Up
-- +goose StatementBegin

ALTER TABLE "classes" ADD COLUMN "total_exam" integer NOT NULL DEFAULT 0;
ALTER TABLE "deleted_classes" ADD COLUMN "total_exam" integer NULL;

CREATE TABLE "class_exams" (
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
CREATE TRIGGER "updated_at_class_exams" BEFORE UPDATE ON "class_exams" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_class_exams" AS TABLE "class_exams" WITH NO DATA;
CREATE TRIGGER "soft_delete_class_exams" BEFORE DELETE ON "class_exams" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

--------------------------------------------

CREATE FUNCTION classes_total_exam_func()
RETURNS TRIGGER AS $$
BEGIN
  IF OLD.class_id IS NOT NULL THEN
    UPDATE classes SET total_exam = total_exam - 1 WHERE id = OLD.class_id;
  END IF;
  IF NEW.class_id IS NOT NULL THEN
    UPDATE classes SET total_exam = total_exam + 1 WHERE id = NEW.class_id;
  END IF;

  IF NEW.class_id IS NOT NULL THEN
    RETURN NEW;
  ELSE
    RETURN OLD;
  END IF;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER classes_total_exam AFTER INSERT OR DELETE ON class_exams FOR EACH ROW EXECUTE PROCEDURE classes_total_exam_func();

--------------------------------------------

CREATE TABLE "class_exam_submissions" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "class_exam_id" uuid NOT NULL REFERENCES "class_exams" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "student_id" uuid NOT NULL REFERENCES "students" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "file_path" character varying NOT NULL,
  "file_path_type" character varying(20) NOT NULL,
  "point" numeric(4,2) NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE("class_exam_id", "student_id")
);
CREATE TABLE "deleted_class_exam_submissions" AS TABLE "class_exam_submissions" WITH NO DATA;
CREATE TRIGGER "soft_delete_class_exam_submissions" BEFORE DELETE ON "class_exam_submissions" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

--------------------------------------------

CREATE FUNCTION class_exam_total_submissions_func()
RETURNS TRIGGER AS $$
BEGIN
  IF OLD.class_exam_id IS NOT NULL THEN
    UPDATE class_exams SET total_submission = total_submission - 1 WHERE id = OLD.class_exam_id;
  END IF;
  IF NEW.class_exam_id IS NOT NULL THEN
    UPDATE class_exams SET total_submission = total_submission + 1 WHERE id = NEW.class_exam_id;
  END IF;

  IF NEW.class_exam_id IS NOT NULL THEN
    RETURN NEW;
  ELSE
    RETURN OLD;
  END IF;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER class_exam_total_submissions AFTER INSERT OR DELETE ON class_exam_submissions FOR EACH ROW EXECUTE PROCEDURE class_exam_total_submissions_func();


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER class_exam_total_submissions ON class_exam_submissions;
DROP FUNCTION class_exam_total_submissions_func();

DROP TABLE "class_exam_submissions";
DROP TABLE "deleted_class_exam_submissions";

DROP TRIGGER classes_total_exam ON class_exams;
DROP FUNCTION classes_total_exam_func();

DROP TABLE "class_exams";
DROP TABLE "deleted_class_exams";

ALTER TABLE "classes" DROP COLUMN "total_exam";
ALTER TABLE "deleted_classes" DROP COLUMN "total_exam";

-- +goose StatementEnd
