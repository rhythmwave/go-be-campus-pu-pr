-- +goose Up
-- +goose StatementBegin

ALTER TABLE "students" ADD COLUMN "paid_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE;
ALTER TABLE "deleted_students" ADD COLUMN "paid_by" uuid NULL;

CREATE TABLE "student_payment_logs" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "student_id" uuid NOT NULL REFERENCES "students" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "semester_id" uuid NOT NULL REFERENCES "semesters" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE("student_id", "semester_id")
);

CREATE FUNCTION "student_payment_log_insert_func"()
RETURNS TRIGGER AS $$
DECLARE semesterId uuid;
BEGIN
  IF NEW.has_paid IS true THEN
    SELECT id INTO semesterId FROM semesters WHERE is_active IS true;

    INSERT INTO student_payment_logs (
      student_id,
      semester_id,
      created_by
    ) VALUES (
      NEW.id,
      semesterId,
      NEW.paid_by
    ) ON CONFLICT ("student_id", "semester_id") DO NOTHING;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER student_payment_log_insert AFTER UPDATE OF has_paid ON students FOR EACH ROW EXECUTE PROCEDURE student_payment_log_insert_func();

------------------------------------------------------------------------------------

CREATE FUNCTION "student_semester_payment_status_func"()
RETURNS TRIGGER AS $$
BEGIN
  IF NEW.is_active IS true THEN
    UPDATE students s SET has_paid = (spl.id IS NOT NULL), paid_by = spl.created_by
    FROM students ss
    LEFT JOIN student_payment_logs spl ON spl.student_id = ss.id AND spl.semester_id = NEW.id
    WHERE ss.id = s.id;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER student_semester_payment_status AFTER UPDATE OF is_active ON semesters FOR EACH ROW EXECUTE PROCEDURE student_semester_payment_status_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER student_payment_log_insert ON students;
DROP FUNCTION "student_payment_log_insert_func"();

DROP TRIGGER student_semester_payment_status ON semesters;
DROP FUNCTION "student_semester_payment_status_func"();

DROP TABLE "student_payment_logs";

ALTER TABLE "students" DROP COLUMN "paid_by";
ALTER TABLE "deleted_students" DROP COLUMN "paid_by";

-- +goose StatementEnd
