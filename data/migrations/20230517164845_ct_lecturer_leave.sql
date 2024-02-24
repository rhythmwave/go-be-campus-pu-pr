-- +goose Up
-- +goose StatementBegin

CREATE TABLE "lecturer_leaves" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "lecturer_id" uuid NOT NULL REFERENCES "lecturers" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "semester_id" uuid NOT NULL REFERENCES "semesters" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "start_date" date NOT NULL,
  "end_date" date NULL,
  "permit_number" character varying NOT NULL UNIQUE,
  "purpose" character varying NOT NULL,
  "remarks" character varying NOT NULL, 
  "is_active" boolean NOT NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_lecturer_leaves" BEFORE UPDATE ON "lecturer_leaves" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_lecturer_leaves" AS TABLE "lecturer_leaves" WITH NO DATA;
CREATE TRIGGER "soft_delete_lecturer_leaves" BEFORE DELETE ON "lecturer_leaves" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

CREATE FUNCTION lecturer_leaves_is_active_func()
RETURNS TRIGGER AS $$
BEGIN
  IF NEW.start_date <= now() THEN
    NEW.is_active = true;
    UPDATE lecturers SET status = 'CUTI' WHERE id = NEW.lecturer_id;
  ELSE
    NEW.is_active = false;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER lecturer_leaves_is_active BEFORE INSERT ON lecturer_leaves FOR EACH ROW EXECUTE PROCEDURE lecturer_leaves_is_active_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER lecturer_leaves_is_active ON lecturer_leaves;
DROP FUNCTION lecturer_leaves_is_active_func();

DROP TABLE "deleted_lecturer_leaves";
DROP TABLE "lecturer_leaves";

-- +goose StatementEnd
