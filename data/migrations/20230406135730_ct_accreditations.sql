-- +goose Up
-- +goose StatementBegin

CREATE TABLE "accreditations" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "study_program_id" uuid NOT NULL REFERENCES "study_programs" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "decree_number" character varying NOT NULL,
  "decree_date" date NOT NULL,
  "decree_due_date" date NOT NULL,
  "accreditation" character(1) NOT NULL,
  "is_active" boolean NOT NULL DEFAULT false,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_accreditations" BEFORE UPDATE ON "accreditations" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_accreditations" AS TABLE "accreditations" WITH NO DATA;
CREATE TRIGGER "soft_delete_accreditations" BEFORE DELETE ON "accreditations" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

--------------------------------------------------------------------------------

CREATE FUNCTION accreditation_is_active_func()
RETURNS TRIGGER AS $$
BEGIN
  IF OLD.is_active IS false AND NEW.is_active IS true THEN
    UPDATE accreditations SET is_active = false WHERE is_active IS true AND study_program_id = NEW.study_program_id;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER accreditation_is_active_update BEFORE INSERT OR UPDATE OF is_active ON accreditations FOR EACH ROW EXECUTE PROCEDURE accreditation_is_active_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER accreditation_is_active ON accreditations;
DROP FUNCTION accreditation_is_active_func();

DROP TABLE "deleted_accreditations";
DROP TABLE "accreditations";

-- +goose StatementEnd
