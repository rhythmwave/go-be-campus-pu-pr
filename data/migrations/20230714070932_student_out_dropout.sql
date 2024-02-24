-- +goose Up
-- +goose StatementBegin

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

---------------------------------------------------------------------------

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

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

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

-- +goose StatementEnd
