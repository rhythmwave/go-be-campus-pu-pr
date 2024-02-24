-- +goose Up
-- +goose StatementBegin

CREATE TYPE subject_prerequisites_prerequisite_type AS ENUM ('lulus', 'ambil');

CREATE TABLE "subject_prerequisites" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "subject_id" uuid NOT NULL REFERENCES "subjects" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "prerequisite_subject_id" uuid NOT NULL REFERENCES "subjects" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "prerequisite_type" subject_prerequisites_prerequisite_type NOT NULL,
  "minimum_grade_point" numeric(3,2) NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE(subject_id, prerequisite_subject_id),
  CHECK((prerequisite_type::text = 'ambil' AND minimum_grade_point IS NOT NULL) OR (prerequisite_type::text != 'ambil' AND minimum_grade_point IS NULL)),
  CHECK(subject_id != prerequisite_subject_id)
);
CREATE TRIGGER "updated_at_subject_prerequisites" BEFORE UPDATE ON "subject_prerequisites" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_subject_prerequisites" AS TABLE "subject_prerequisites" WITH NO DATA;
CREATE TRIGGER "soft_delete_subject_prerequisites" BEFORE DELETE ON "subject_prerequisites" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

----------------------------------------------------------------------

CREATE FUNCTION subject_prerequisites_validation_func()
RETURNS TRIGGER AS $$
DECLARE 
  targetCurriculum uuid;
  prerequisiteCurriculum uuid;
  isCircularDependent boolean;
BEGIN
  SELECT true INTO isCircularDependent FROM subject_prerequisites WHERE prerequisite_subject_id = NEW.subject_id AND subject_id = NEW.prerequisite_subject_id;

  IF isCircularDependent IS true THEN
    RAISE EXCEPTION 'you cannot have circular dependent subjects';
  END IF;

  SELECT curriculum_id INTO targetCurriculum FROM subjects WHERE id = NEW.subject_id;
  SELECT curriculum_id INTO prerequisiteCurriculum FROM subjects WHERE id = NEW.prerequisite_subject_id;

  IF targetCurriculum != prerequisiteCurriculum THEN
    RAISE EXCEPTION 'prerequisite subject should in the same curriculum';
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER subject_prerequisites_validation BEFORE INSERT OR UPDATE ON subject_prerequisites FOR EACH ROW EXECUTE PROCEDURE subject_prerequisites_validation_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER subject_prerequisites_validation ON subject_prerequisites;
DROP FUNCTION subject_prerequisites_validation_func();

DROP TABLE "deleted_subject_prerequisites";
DROP TABLE "subject_prerequisites";

DROP TYPE subject_prerequisites_prerequisite_type;

-- +goose StatementEnd
