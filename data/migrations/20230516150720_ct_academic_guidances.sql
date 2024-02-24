-- +goose Up
-- +goose StatementBegin

CREATE TABLE "academic_guidances" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "semester_id" uuid NOT NULL REFERENCES "semesters" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "lecturer_id" uuid NOT NULL REFERENCES "lecturers" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "decision_number" character varying NULL UNIQUE,
  "decision_date" date NULL,
  "total_student" integer NOT NULL DEFAULT 0,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("semester_id", "lecturer_id")
);
CREATE TRIGGER "updated_at_academic_guidances" BEFORE UPDATE ON "academic_guidances" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_academic_guidances" AS TABLE "academic_guidances" WITH NO DATA;
CREATE TRIGGER "soft_delete_academic_guidances" BEFORE DELETE ON "academic_guidances" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

------------------------------------------------------------------

CREATE TABLE "academic_guidance_students" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "academic_guidance_id" uuid NOT NULL REFERENCES "academic_guidances" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "semester_id" uuid NOT NULL REFERENCES "semesters" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "student_id" uuid NOT NULL REFERENCES "students" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("semester_id", "student_id")
);
CREATE TRIGGER "updated_at_academic_guidance_students" BEFORE UPDATE ON "academic_guidance_students" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_academic_guidance_students" AS TABLE "academic_guidance_students" WITH NO DATA;
CREATE TRIGGER "soft_delete_academic_guidance_students" BEFORE DELETE ON "academic_guidance_students" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

CREATE FUNCTION academic_guidance_students_semester_func()
RETURNS TRIGGER AS $$
DECLARE semesterId uuid;
BEGIN
  SELECT semester_id INTO semesterId FROM academic_guidances WHERE id = NEW.academic_guidance_id;

  NEW.semester_id = semesterId;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER academic_guidance_students_semester BEFORE INSERT OR UPDATE ON academic_guidance_students FOR EACH ROW EXECUTE PROCEDURE academic_guidance_students_semester_func();

CREATE FUNCTION academic_guidances_total_student_func()
RETURNS TRIGGER AS $$
BEGIN
  WITH d AS (
    SELECT COUNT(1) FROM academic_guidance_students
    WHERE academic_guidance_id = NEW.academic_guidance_id
  )
  UPDATE academic_guidances SET total_student = d.count
  FROM d
  WHERE academic_guidances.id = NEW.academic_guidance_id;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER academic_guidances_total_student AFTER INSERT OR UPDATE ON academic_guidance_students FOR EACH ROW EXECUTE PROCEDURE academic_guidances_total_student_func();

CREATE FUNCTION academic_guidances_total_student_delete_func()
RETURNS TRIGGER AS $$
BEGIN
  WITH d AS (
    SELECT COUNT(1) FROM academic_guidance_students
    WHERE academic_guidance_id = NEW.academic_guidance_id
  )
  UPDATE academic_guidances SET total_student = d.count
  FROM d
  WHERE academic_guidances.id = NEW.academic_guidance_id;

  RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER academic_guidances_total_student_delete AFTER DELETE ON academic_guidance_students FOR EACH ROW EXECUTE PROCEDURE academic_guidances_total_student_delete_func();


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER academic_guidances_total_student AFTER INSERT OR UPDATE ON academic_guidance_students;
DROP FUNCTION academic_guidances_total_student_func();
DROP TRIGGER academic_guidances_total_student_delete AFTER DELETE ON academic_guidance_students;
DROP FUNCTION academic_guidances_total_student_delete_func();

DROP TRIGGER academic_guidance_students_semester ON academic_guidance_students;
DROP FUNCTION academic_guidance_students_semester_func();

DROP TABLE "academic_guidance_students";
DROP TABLE "deleted_academic_guidance_students";

DROP TABLE "deleted_academic_guidances";
DROP TABLE "academic_guidances";

-- +goose StatementEnd
