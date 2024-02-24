-- +goose Up
-- +goose StatementBegin

CREATE TABLE "student_leave_requests" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "student_id" uuid NOT NULL REFERENCES "students" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "start_date" date NOT NULL,
  "total_leave_duration_semester" integer NOT NULL,
  "current_leave_duration_semester" integer NOT NULL DEFAULT 0,
  "permit_number" character varying NULL UNIQUE,
  "purpose" character varying NOT NULL,
  "remarks" character varying NOT NULL,
  "is_approved" boolean NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_student_leave_requests" BEFORE UPDATE ON "student_leave_requests" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_student_leave_requests" AS TABLE "student_leave_requests" WITH NO DATA;
CREATE TRIGGER "soft_delete_student_leave_requests" BEFORE DELETE ON "student_leave_requests" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

------------------------------------------------------------

CREATE TABLE "student_leaves" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "student_leave_request_id" uuid NOT NULL REFERENCES "student_leave_requests" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "semester_id" uuid NOT NULL REFERENCES "semesters" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE(student_leave_request_id,semester_id)
);
CREATE TRIGGER "updated_at_student_leaves" BEFORE UPDATE ON "student_leaves" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_student_leaves" AS TABLE "student_leaves" WITH NO DATA;
CREATE TRIGGER "soft_delete_student_leaves" BEFORE DELETE ON "student_leaves" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

------------------------------------------------------------

CREATE FUNCTION student_leaves_is_active_func()
RETURNS TRIGGER AS $$
DECLARE activeSemesterId uuid;
BEGIN
  IF DATE(NEW.start_date) <= DATE(now()) AND OLD.is_approved IS NULL AND NEW.is_approved IS true THEN
    SELECT id INTO activeSemesterId FROM semesters WHERE is_active IS true;
    INSERT INTO "student_leaves" (
      student_leave_request_id,
      semester_id
    ) VALUES (
      NEW.id,
      activeSemesterId
    ) ON CONFLICT (student_leave_request_id, semester_id) DO NOTHING;

    NEW.current_leave_duration_semester = 1;
    UPDATE students SET status = 'CUTI' WHERE id = NEW.student_id;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER student_leaves_is_active BEFORE UPDATE OF is_approved ON student_leave_requests FOR EACH ROW EXECUTE PROCEDURE student_leaves_is_active_func();

------------------------------------------------------------

CREATE FUNCTION student_leaves_new_semester_func()
RETURNS TRIGGER AS $$
BEGIN
  IF NEW.is_active IS true THEN
    WITH d AS (
      SELECT slr.id
      FROM student_leave_requests slr
      LEFT JOIN student_leaves sl ON sl.student_leave_request_id = slr.id AND sl.semester_id = NEW.id
      WHERE sl.id IS NULL AND slr.is_approved IS true AND slr.total_leave_duration_semester > slr.current_leave_duration_semester
    )
    UPDATE student_leave_requests slr SET current_leave_duration_semester = slr.current_leave_duration_semester + 1
    FROM d
    WHERE d.id = slr.id;

    INSERT INTO student_leaves(
      student_leave_request_id,
      semester_id
    ) SELECT id, NEW.id
    FROM student_leave_requests
    WHERE is_approved IS true AND total_leave_duration_semester >= current_leave_duration_semester
    ON CONFLICT (student_leave_request_id, semester_id) DO NOTHING;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER student_leaves_new_semester AFTER UPDATE OF is_active ON semesters FOR EACH ROW EXECUTE PROCEDURE student_leaves_new_semester_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER student_leaves_new_semester ON semesters;
DROP FUNCTION student_leaves_new_semester_func();

DROP TRIGGER student_leaves_is_active ON student_leave_requests;
DROP FUNCTION student_leaves_is_active_func();

DROP TABLE "deleted_student_leaves";
DROP TABLE "student_leaves";

DROP TABLE "deleted_student_leave_requests";
DROP TABLE "student_leave_requests";

-- +goose StatementEnd
