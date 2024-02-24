-- +goose Up
-- +goose StatementBegin

DROP TRIGGER student_authentication ON students;

ALTER TABLE "students" 
  ALTER COLUMN "nim_number" TYPE bigint USING nim_number::bigint,
  ALTER COLUMN "previous_nim_number" TYPE bigint USING previous_nim_number::bigint;
ALTER TABLE "deleted_students" 
  ALTER COLUMN "nim_number" TYPE bigint USING nim_number::bigint,
  ALTER COLUMN "previous_nim_number" TYPE bigint USING previous_nim_number::bigint;

CREATE TRIGGER student_authentication AFTER UPDATE OF nim_number ON students FOR EACH ROW EXECUTE PROCEDURE student_authentication_func();


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER student_authentication ON students;

ALTER TABLE "students" 
  ALTER COLUMN "nim_number" TYPE integer USING nim_number::integer,
  ALTER COLUMN "previous_nim_number" TYPE integer USING previous_nim_number::integer;
ALTER TABLE "deleted_students" 
  ALTER COLUMN "nim_number" TYPE integer USING nim_number::integer,
  ALTER COLUMN "previous_nim_number" TYPE integer USING previous_nim_number::integer;

CREATE TRIGGER student_authentication AFTER UPDATE OF nim_number ON students FOR EACH ROW EXECUTE PROCEDURE student_authentication_func();

-- +goose StatementEnd
