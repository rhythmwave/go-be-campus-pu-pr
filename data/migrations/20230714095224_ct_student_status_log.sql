-- +goose Up
-- +goose StatementBegin

ALTER TABLE "students"
  ADD COLUMN "status_reference_number" character varying NULL,
  ADD COLUMN "status_date" date NULL,
  ADD COLUMN "status_purpose" character varying NULL,
  ADD COLUMN "status_remarks" character varying NULL;

ALTER TABLE "deleted_students"
  ADD COLUMN "status_reference_number" character varying NULL,
  ADD COLUMN "status_date" date NULL,
  ADD COLUMN "status_purpose" character varying NULL,
  ADD COLUMN "status_remarks" character varying NULL;

CREATE TABLE "student_status_logs" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "student_id" uuid NOT NULL REFERENCES "students" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "semester_id" uuid NOT NULL REFERENCES "semesters" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "study_program_id" uuid NULL REFERENCES "study_programs" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "status" character varying NOT NULL,
  "reference_number" character varying NULL,
  "status_date" date NULL,
  "purpose" character varying NULL,
  "remarks" character varying NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("student_id", "semester_id")
);
CREATE TRIGGER "updated_at_student_status_logs" BEFORE UPDATE ON "student_status_logs" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

------------------------------------------------------------------------------------

INSERT INTO "student_status_logs" (
  "student_id",
  "semester_id",
  "study_program_id",
  "status",
  "reference_number",
  "status_date",
  "purpose",
  "remarks"
) SELECT 
  s.id,
  se.id,
  s.study_program_id,
  s.status,
  s.status_reference_number,
  s.status_date,
  s.status_purpose,
  s.status_remarks
FROM students s
JOIN semesters se ON se.is_active IS true
WHERE s.status NOT IN ('DROP-OUT', 'KELUAR', 'LULUS')
ON CONFLICT ("student_id", "semester_id") DO UPDATE SET
  status = EXCLUDED.status,
  reference_number = EXCLUDED.reference_number,
  status_date = EXCLUDED.status_date,
  purpose = EXCLUDED.purpose,
  remarks = EXCLUDED.remarks;

------------------------------------------------------------------------------------

CREATE FUNCTION student_status_logs_new_semester_func()
RETURNS TRIGGER AS $$
DECLARE activeSemesterId uuid;
BEGIN
  IF OLD.is_active IS false AND NEW.is_active IS true THEN
    SELECT id INTO activeSemesterId FROM semesters WHERE is_active IS true;
    INSERT INTO "student_status_logs" (
      "student_id",
      "semester_id",
      "study_program_id",
      "status",
      "reference_number",
      "status_date",
      "purpose",
      "remarks"
    ) SELECT 
      s.id,
      NEW.id,
      s.study_program_id,
      s.status,
      s.status_reference_number,
      s.status_date,
      s.status_purpose,
      s.status_remarks
    FROM students s
    WHERE s.status NOT IN ('DROP-OUT', 'KELUAR', 'LULUS')
    ON CONFLICT ("student_id", "semester_id") DO UPDATE SET
      status = EXCLUDED.status,
      reference_number = EXCLUDED.reference_number,
      status_date = EXCLUDED.status_date,
      purpose = EXCLUDED.purpose,
      remarks = EXCLUDED.remarks;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER student_status_logs_new_semester AFTER UPDATE OF is_active ON semesters FOR EACH ROW EXECUTE PROCEDURE student_status_logs_new_semester_func();

------------------------------------------------------------------------------------

CREATE FUNCTION student_status_logs_insert_func()
RETURNS TRIGGER AS $$
DECLARE activeSemesterId uuid;
BEGIN
  IF OLD.status != NEW.status THEN
    SELECT id INTO activeSemesterId FROM semesters WHERE is_active IS true;
    INSERT INTO "student_status_logs" (
      "student_id",
      "semester_id",
      "study_program_id",
      "status",
      "reference_number",
      "status_date",
      "purpose",
      "remarks"
    ) VALUES (
      NEW.id,
      activeSemesterId,
      NEW.study_program_id,
      NEW.status,
      NEW.status_reference_number,
      NEW.status_date,
      NEW.status_purpose,
      NEW.status_remarks
    ) ON CONFLICT ("student_id", "semester_id") DO UPDATE SET
      status = EXCLUDED.status,
      reference_number = EXCLUDED.reference_number,
      status_date = EXCLUDED.status_date,
      purpose = EXCLUDED.purpose,
      remarks = EXCLUDED.remarks;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER student_status_logs_insert AFTER UPDATE OF status ON students FOR EACH ROW EXECUTE PROCEDURE student_status_logs_insert_func();

------------------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION student_leaves_is_active_func()
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

    UPDATE students SET 
      status = 'CUTI',
      status_reference_number = NEW.permit_number,
      status_date = NEW.start_date,
      status_purpose = NEW.purpose,
      status_remarks = NEW.remarks
    WHERE id = NEW.student_id;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

------------------------------------------------------------------------------------

DROP TRIGGER students_out ON students;
DROP FUNCTION "students_out_func"();
DROP TRIGGER students_drop_out ON students;
DROP FUNCTION "students_drop_out_func"();

ALTER TABLE "students"
  DROP COLUMN "out_semester_id",
  DROP COLUMN "out_date",
  DROP COLUMN "out_number",
  DROP COLUMN "out_cause",
  DROP COLUMN "out_remarks",
  DROP COLUMN "drop_out_semester_id",
  DROP COLUMN "drop_out_date",
  DROP COLUMN "drop_out_number",
  DROP COLUMN "drop_out_remarks";

ALTER TABLE "deleted_students"
  DROP COLUMN "out_semester_id",
  DROP COLUMN "out_date",
  DROP COLUMN "out_number",
  DROP COLUMN "out_cause",
  DROP COLUMN "out_remarks",
  DROP COLUMN "drop_out_semester_id",
  DROP COLUMN "drop_out_date",
  DROP COLUMN "drop_out_number",
  DROP COLUMN "drop_out_remarks";

------------------------------------------------------------------------------------

DROP TRIGGER student_inactive_date ON students;
DROP FUNCTION student_inactive_date_func();

ALTER TABLE "students" 
  DROP COLUMN "inactive_date",
  DROP COLUMN "inactive_semester_id";
ALTER TABLE "deleted_students" 
  DROP COLUMN "inactive_date",
  DROP COLUMN "inactive_semester_id";

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin


ALTER TABLE "students" 
  ADD COLUMN "inactive_date" date NULL,
  ADD COLUMN "inactive_semester_id" uuid NULL REFERENCES "semesters" ("id") ON DELETE SET NULL ON UPDATE CASCADE;
ALTER TABLE "deleted_students" 
  ADD COLUMN "inactive_date" date NULL,
  ADD COLUMN "inactive_semester_id" uuid NULL;

CREATE FUNCTION student_inactive_date_func()
RETURNS TRIGGER AS $$
DECLARE activeSemesterId uuid;
BEGIN
  IF OLD.status != 'NON-AKTIF' AND NEW.status = 'NON-AKTIF' THEN
    SELECT id INTO activeSemesterId FROM semesters WHERE is_active IS true;

    NEW.inactive_date = DATE(now());
    NEW.inactive_semester_id = activeSemesterId;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER student_inactive_date BEFORE UPDATE OF status ON students FOR EACH ROW EXECUTE PROCEDURE student_inactive_date_func();

------------------------------------------------------------------------------------

ALTER TABLE "students"
  ADD COLUMN "out_semester_id" uuid NULL REFERENCES "semesters"("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  ADD COLUMN "out_date" date NULL,
  ADD COLUMN "out_number" character varying NULL,
  ADD COLUMN "out_cause" character varying NULL,
  ADD COLUMN "out_remarks" character varying NULL,
  ADD COLUMN "drop_out_semester_id" uuid NULL REFERENCES "semesters"("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  ADD COLUMN "drop_out_date" date NULL,
  ADD COLUMN "drop_out_number" character varying NULL,
  ADD COLUMN "drop_out_remarks" character varying NULL;

ALTER TABLE "deleted_students"
  ADD COLUMN "out_semester_id" uuid NULL,
  ADD COLUMN "out_date" date NULL,
  ADD COLUMN "out_number" character varying NULL,
  ADD COLUMN "out_cause" character varying NULL,
  ADD COLUMN "out_remarks" character varying NULL,
  ADD COLUMN "drop_out_semester_id" uuid NULL,
  ADD COLUMN "drop_out_date" date NULL,
  ADD COLUMN "drop_out_number" character varying NULL,
  ADD COLUMN "drop_out_remarks" character varying NULL;

CREATE FUNCTION "students_out_func"()
RETURNS TRIGGER AS $$
DECLARE activeSemesterId uuid;
BEGIN
  IF OLD.out_date IS NULL AND NEW.out_date IS NOT NULL THEN
    SELECT id INTO activeSemesterId FROM semesters WHERE is_active IS true;
    NEW.out_semester_id = activeSemesterId;

    NEW.status = 'KELUAR';
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER students_out BEFORE UPDATE OF out_date ON students FOR EACH ROW EXECUTE PROCEDURE students_out_func();

CREATE FUNCTION "students_drop_out_func"()
RETURNS TRIGGER AS $$
DECLARE activeSemesterId uuid;
BEGIN
  IF OLD.drop_out_date IS NULL AND NEW.drop_out_date IS NOT NULL THEN
    SELECT id INTO activeSemesterId FROM semesters WHERE is_active IS true;
    NEW.drop_out_semester_id = activeSemesterId;

    NEW.status = 'DROP-OUT';
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER students_drop_out BEFORE UPDATE OF drop_out_date ON students FOR EACH ROW EXECUTE PROCEDURE students_drop_out_func();

------------------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION student_leaves_is_active_func()
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

------------------------------------------------------------------------------------

DROP TRIGGER student_status_logs_insert ON students;
DROP FUNCTION student_status_logs_insert_func();

DROP TRIGGER student_status_logs_new_semester ON semesters;
DROP FUNCTION student_status_logs_new_semester_func();

DROP TABLE "student_status_logs";

ALTER TABLE "students"
  DROP COLUMN "status_reference_number",
  DROP COLUMN "status_date",
  DROP COLUMN "status_purpose",
  DROP COLUMN "status_remarks";

ALTER TABLE "deleted_students"
  DROP COLUMN "status_reference_number",
  DROP COLUMN "status_date",
  DROP COLUMN "status_purpose",
  DROP COLUMN "status_remarks";

-- +goose StatementEnd
