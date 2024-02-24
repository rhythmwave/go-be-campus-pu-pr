-- +goose Up
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

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER student_inactive_date ON students;
DROP FUNCTION student_inactive_date_func();

ALTER TABLE "students" 
  DROP COLUMN "inactive_date",
  DROP COLUMN "inactive_semester_id";
ALTER TABLE "deleted_students" 
  DROP COLUMN "inactive_date",
  DROP COLUMN "inactive_semester_id";

-- +goose StatementEnd
