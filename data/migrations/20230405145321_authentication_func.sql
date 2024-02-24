-- +goose Up
-- +goose StatementBegin

CREATE FUNCTION lecturer_authentication_func()
RETURNS TRIGGER AS $$
  BEGIN
    UPDATE authentications SET username = NEW.id_national_lecturer WHERE lecturer_id = NEW.id;

    RETURN NEW;
  END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER lecturer_authentication AFTER UPDATE OF id_national_lecturer ON lecturers FOR EACH ROW EXECUTE PROCEDURE lecturer_authentication_func();

CREATE FUNCTION student_authentication_func()
RETURNS TRIGGER AS $$
  BEGIN
    UPDATE authentications SET username = NEW.nim_number WHERE student_id = NEW.id;

    RETURN NEW;
  END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER student_authentication AFTER UPDATE OF nim_number ON students FOR EACH ROW EXECUTE PROCEDURE student_authentication_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER student_authentication ON students;
DROP FUNCTION student_authentication_func();

DROP TRIGGER lecturer_authentication ON lecturers;
DROP FUNCTION lecturer_authentication_func();

-- +goose StatementEnd
