-- +goose Up
-- +goose StatementBegin

CREATE TABLE "lecturer_mutations" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "lecturer_id" uuid NOT NULL REFERENCES "lecturers" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "semester_id" uuid NOT NULL REFERENCES "semesters" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "mutation_date" date NOT NULL,
  "decision_number" character varying NOT NULL UNIQUE,
  "destination" character varying NOT NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_lecturer_mutations" BEFORE UPDATE ON "lecturer_mutations" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_lecturer_mutations" AS TABLE "lecturer_mutations" WITH NO DATA;
CREATE TRIGGER "soft_delete_lecturer_mutations" BEFORE DELETE ON "lecturer_mutations" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

CREATE FUNCTION lecturer_mutations_status_func()
RETURNS TRIGGER AS $$
BEGIN
  UPDATE lecturers SET
    status = 'TUGAS DI INSTANSI LAIN',
    updated_by = NEW.created_by
  WHERE id = NEW.lecturer_id;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER lecturer_mutations_status AFTER INSERT ON lecturer_mutations FOR EACH ROW EXECUTE PROCEDURE lecturer_mutations_status_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER lecturer_mutations_status ON lecturer_mutations;
DROP FUNCTION lecturer_mutations_status_func();

DROP TABLE "deleted_lecturer_mutations";
DROP TABLE "lecturer_mutations";

-- +goose StatementEnd
