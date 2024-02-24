-- +goose Up
-- +goose StatementBegin

ALTER TABLE "subjects" ADD COLUMN "has_prerequisite" boolean NOT NULL DEFAULT false;
ALTER TABLE "deleted_subjects" ADD COLUMN "has_prerequisite" boolean NULL;

ALTER TABLE "curriculums" ADD COLUMN "total_subject_with_prerequisite" integer NOT NULL DEFAULT 0;
ALTER TABLE "deleted_curriculums" ADD COLUMN "total_subject_with_prerequisite" integer NULL;

CREATE FUNCTION curriculum_subjects_has_prerequisite_func()
RETURNS TRIGGER AS $$
DECLARE
  oldSubjectPrerequisiteCount integer := 0;
  oldCurriculumId uuid;
  newCurriculumId uuid;
BEGIN
  IF OLD.subject_id IS NOT NULL AND OLD.subject_id != NEW.subject_id THEN
    SELECT COUNT(1) INTO oldSubjectPrerequisiteCount FROM subject_prerequisites WHERE subject_id = OLD.subject_id;
    IF oldSubjectPrerequisiteCount = 0 THEN
      UPDATE subjects SET has_prerequisite = false WHERE id = OLD.subject_id RETURNING curriculum_id INTO oldCurriculumId;
      UPDATE curriculums SET total_subject_with_prerequisite = total_subject_with_prerequisite - 1 WHERE id = oldCurriculumId;
    END IF;
  END IF;

  UPDATE subjects SET has_prerequisite = true WHERE id = NEW.subject_id RETURNING curriculum_id INTO newCurriculumId;

  WITH d AS (
    SELECT COUNT(1) FROM subjects
    WHERE curriculum_id = newCurriculumId AND has_prerequisite IS true
  )
  UPDATE curriculums SET total_subject_with_prerequisite = d.count
  FROM d
  WHERE id = newCurriculumId;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER curriculum_subjects_has_prerequisite AFTER INSERT OR UPDATE ON subject_prerequisites FOR EACH ROW EXECUTE PROCEDURE curriculum_subjects_has_prerequisite_func();

CREATE FUNCTION curriculum_subjects_has_prerequisite_delete_func()
RETURNS TRIGGER AS $$
DECLARE
  oldSubjectPrerequisiteCount integer := 0;
  oldCurriculumId uuid;
BEGIN
  SELECT COUNT(1) INTO oldSubjectPrerequisiteCount FROM subject_prerequisites WHERE subject_id = OLD.subject_id;
  IF oldSubjectPrerequisiteCount = 0 THEN
    UPDATE subjects SET has_prerequisite = false WHERE id = OLD.subject_id RETURNING curriculum_id INTO oldCurriculumId;
    UPDATE curriculums SET total_subject_with_prerequisite = total_subject_with_prerequisite - 1 WHERE id = oldCurriculumId;
  END IF;

  RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER curriculum_subjects_has_prerequisite_delete AFTER DELETE ON subject_prerequisites FOR EACH ROW EXECUTE PROCEDURE curriculum_subjects_has_prerequisite_delete_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER curriculum_subjects_has_prerequisite_delete ON subject_prerequisites;
DROP FUNCTION curriculum_subjects_has_prerequisite_delete_func();

DROP TRIGGER curriculum_subjects_has_prerequisite ON subject_prerequisites;
DROP FUNCTION curriculum_subjects_has_prerequisite_func();

ALTER TABLE "subjects" DROP COLUMN "has_prerequisite";
ALTER TABLE "deleted_subjects" DROP COLUMN "has_prerequisite";
ALTER TABLE "curriculums" DROP COLUMN "total_subject_with_prerequisite";
ALTER TABLE "deleted_curriculums" DROP COLUMN "total_subject_with_prerequisite";

-- +goose StatementEnd
