-- +goose Up
-- +goose StatementBegin

CREATE TABLE "lecturer_resignations" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "lecturer_id" uuid NOT NULL REFERENCES "lecturers" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "semester_id" uuid NOT NULL REFERENCES "semesters" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "resign_date" date NOT NULL,
  "resignation_number" character varying NOT NULL UNIQUE,
  "purpose" character varying NOT NULL,
  "remarks" character varying NOT NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_lecturer_resignations" BEFORE UPDATE ON "lecturer_resignations" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_lecturer_resignations" AS TABLE "lecturer_resignations" WITH NO DATA;
CREATE TRIGGER "soft_delete_lecturer_resignations" BEFORE DELETE ON "lecturer_resignations" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

CREATE FUNCTION lecturer_resignations_status_func()
RETURNS TRIGGER AS $$
DECLARE semesterYear integer;
BEGIN
  SELECT semester_start_year INTO semesterYear FROM semesters WHERE id = NEW.semester_id;

  UPDATE lecturers SET
    end_date = NEW.resign_date,
    status = NEW.purpose,
    resign_semester = semesterYear::text,
    updated_by = NEW.created_by
  WHERE id = NEW.lecturer_id;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER lecturer_resignations_status AFTER INSERT ON lecturer_resignations FOR EACH ROW EXECUTE PROCEDURE lecturer_resignations_status_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER lecturer_resignations_status ON lecturer_resignations;
DROP FUNCTION lecturer_resignations_status_func();

DROP TABLE "deleted_lecturer_resignations";
DROP TABLE "lecturer_resignations";

-- +goose StatementEnd
