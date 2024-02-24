-- +goose Up
-- +goose StatementBegin

ALTER TABLE "students" ADD COLUMN "status_change_semester_id" uuid NULL REFERENCES semesters (id) ON DELETE RESTRICT ON UPDATE CASCADE;
ALTER TABLE "deleted_students" ADD COLUMN "status_change_semester_id" uuid NULL;

CREATE FUNCTION "student_status_change_semester_func"()
RETURNS TRIGGER AS $$
DECLARE activeSemesterId uuid;
BEGIN
  IF OLD.status != NEW.status THEN
    SELECT id INTO activeSemesterId FROM semesters WHERE is_active IS true;

    NEW.status_change_semester_id = activeSemesterId;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER student_status_change_semester BEFORE UPDATE OF status ON students FOR EACH ROW EXECUTE PROCEDURE student_status_change_semester_func();

CREATE TABLE "report_student_statuses" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "semester_id" uuid NOT NULL REFERENCES "semesters" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "study_program_id" uuid NOT NULL REFERENCES "study_programs" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "status" character varying NOT NULL,
  "total" integer NOT NULL DEFAULT 0,
  UNIQUE("semester_id", "study_program_id", "status")
);

WITH 
  st AS (
    SELECT UNNEST(ARRAY['AKTIF', 'CUTI', 'NON-AKTIF', 'MENUNGGU-UKOM', 'MBKM', 'KELUAR', 'LULUS', 'DROP-OUT']) AS status
  ),
  d AS (
    SELECT COUNT(s.id) AS total, sp.id AS study_program_id, st.status 
    FROM st
    CROSS JOIN study_programs sp
    LEFT JOIN students s ON s.status = st.status AND s.study_program_id = sp.id
    GROUP BY st.status, sp.id
  )
INSERT INTO report_student_statuses (
  semester_id,
  study_program_id,
  status,
  total
)
SELECT
  se.id,
  d.study_program_id,
  d.status,
  d.total
FROM d
JOIN semesters se ON se.is_active IS false
ON CONFLICT (semester_id, study_program_id, status) DO UPDATE SET total = EXCLUDED.total;

CREATE FUNCTION report_student_status_func()
RETURNS TRIGGER AS $$
BEGIN
  IF OLD.is_active IS true AND NEW.is_active IS false THEN
    WITH 
      st AS (
        SELECT UNNEST(ARRAY['AKTIF', 'CUTI', 'NON-AKTIF', 'MENUNGGU-UKOM', 'MBKM']) AS status
      ),
      d AS (
        SELECT COUNT(s.id) AS total, sp.id AS study_program_id, st.status 
        FROM st
        CROSS JOIN study_programs sp
        LEFT JOIN students s ON s.status = st.status AND s.study_program_id = sp.id
        GROUP BY st.status, sp.id
      )
    INSERT INTO report_student_statuses (
      semester_id,
      study_program_id,
      status,
      total
    )
    SELECT
      NEW.id,
      d.study_program_id,
      d.status,
      d.total
    FROM d
    ON CONFLICT (semester_id, study_program_id, status) DO UPDATE SET total = EXCLUDED.total;

    WITH 
      st AS (
        SELECT UNNEST(ARRAY['KELUAR', 'LULUS', 'DROP-OUT']) AS status
      ),
      d AS (
        SELECT COUNT(s.id) AS total, sp.id AS study_program_id, st.status 
        FROM st
        CROSS JOIN study_programs sp
        LEFT JOIN students s ON s.status = st.status AND s.study_program_id = sp.id AND s.status_change_semester_id = NEW.id
        GROUP BY st.status, sp.id
      )
    INSERT INTO report_student_statuses (
      semester_id,
      study_program_id,
      status,
      total
    )
    SELECT
      NEW.id,
      d.study_program_id,
      d.status,
      d.total
    FROM d
    ON CONFLICT (semester_id, study_program_id, status) DO UPDATE SET total = EXCLUDED.total;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER report_student_status AFTER UPDATE OF is_active ON semesters FOR EACH ROW EXECUTE PROCEDURE report_student_status_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER report_student_status ON semesters;
DROP FUNCTION report_student_status_func();

DROP TABLE "report_student_statuses";

DROP TRIGGER student_status_change_semester ON students;
DROP FUNCTION "student_status_change_semester_func"();

ALTER TABLE "students" DROP COLUMN "status_change_semester_id";
ALTER TABLE "deleted_students" DROP COLUMN "status_change_semester_id";

-- +goose StatementEnd
