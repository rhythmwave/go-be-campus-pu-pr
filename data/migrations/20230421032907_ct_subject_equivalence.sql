-- +goose Up
-- +goose StatementBegin


CREATE TABLE "subject_equivalences" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "subject_id" uuid NOT NULL REFERENCES "subjects" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "equivalent_subject_id" uuid NOT NULL REFERENCES "subjects" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "equivalent_curriculum_id" uuid NOT NULL REFERENCES "curriculums" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE(subject_id, equivalent_curriculum_id),
  CHECK(subject_id != equivalent_subject_id)
);
CREATE TRIGGER "updated_at_subject_equivalences" BEFORE UPDATE ON "subject_equivalences" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_subject_equivalences" AS TABLE "subject_equivalences" WITH NO DATA;
CREATE TRIGGER "soft_delete_subject_equivalences" BEFORE DELETE ON "subject_equivalences" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

----------------------------------------------------------------------

CREATE FUNCTION subject_equivalences_validation_func()
RETURNS TRIGGER AS $$
DECLARE 
  targetCurriculum uuid;
  equivalentCurriculum uuid;
BEGIN
  SELECT curriculum_id INTO targetCurriculum FROM subjects WHERE id = NEW.subject_id;
  SELECT curriculum_id INTO equivalentCurriculum FROM subjects WHERE id = NEW.equivalent_subject_id;

  IF targetCurriculum = equivalentCurriculum THEN
    RAISE EXCEPTION 'equivalent subject should in the different curriculum';
  END IF;

  NEW.equivalent_curriculum_id = equivalentCurriculum;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER subject_equivalences_validation BEFORE INSERT OR UPDATE ON subject_equivalences FOR EACH ROW EXECUTE PROCEDURE subject_equivalences_validation_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER subject_equivalences_validation ON subject_equivalences;
DROP FUNCTION subject_equivalences_validation_func();

DROP TABLE "deleted_subject_equivalences";
DROP TABLE "subject_equivalences";

-- +goose StatementEnd
