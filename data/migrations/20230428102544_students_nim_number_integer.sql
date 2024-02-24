-- +goose Up
-- +goose StatementBegin

DROP TRIGGER student_authentication ON students;

ALTER TABLE "students" DROP CONSTRAINT "students_student_force_check";

ALTER TABLE "students" 
  ALTER COLUMN "nim_number" TYPE integer USING ("nim_number"::integer),
  ALTER COLUMN "student_force" TYPE integer USING ("student_force"::integer);

ALTER TABLE "deleted_students" 
  ALTER COLUMN "nim_number" TYPE integer USING ("nim_number"::integer),
  ALTER COLUMN "student_force" TYPE integer USING ("student_force"::integer);

CREATE OR REPLACE FUNCTION student_authentication_func()
RETURNS TRIGGER AS $$
  BEGIN
    UPDATE authentications SET username = NEW.nim_number::text WHERE student_id = NEW.id;

    RETURN NEW;
  END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER student_authentication AFTER UPDATE OF nim_number ON students FOR EACH ROW EXECUTE PROCEDURE student_authentication_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER student_authentication ON students;

ALTER TABLE "students" 
  ALTER COLUMN "nim_number" TYPE character varying,
  ALTER COLUMN "student_force" TYPE character(4);

ALTER TABLE "deleted_students" 
  ALTER COLUMN "nim_number" TYPE character varying,
  ALTER COLUMN "student_force" TYPE character(4);

ALTER TABLE "students" ADD CONSTRAINT "students_student_force_check" CHECK ("student_force" ~ '^[0-9\.]+$');

CREATE OR REPLACE FUNCTION student_authentication_func()
RETURNS TRIGGER AS $$
  BEGIN
    UPDATE authentications SET username = NEW.nim_number WHERE student_id = NEW.id;

    RETURN NEW;
  END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER student_authentication AFTER UPDATE OF nim_number ON students FOR EACH ROW EXECUTE PROCEDURE student_authentication_func();

-- +goose StatementEnd
