-- +goose Up
-- +goose StatementBegin

ALTER TABLE "subjects" ADD COLUMN "has_equivalence" boolean NOT NULL DEFAULT false;
ALTER TABLE "deleted_subjects" ADD COLUMN "has_equivalence" boolean NULL;

ALTER TABLE "curriculums" ADD COLUMN "total_subject_with_equivalence" integer NOT NULL DEFAULT 0;
ALTER TABLE "deleted_curriculums" ADD COLUMN "total_subject_with_equivalence" integer NULL;

CREATE FUNCTION curriculum_subjects_has_equivalence_func()
RETURNS TRIGGER AS $$
DECLARE
  oldSubjectEquivalenceCount integer := 0;
  oldCurriculumId uuid;
  newCurriculumId uuid;
BEGIN
  IF OLD.subject_id IS NOT NULL AND OLD.subject_id != NEW.subject_id THEN
    SELECT COUNT(1) INTO oldSubjectEquivalenceCount FROM subject_equivalences WHERE subject_id = OLD.subject_id;
    IF oldSubjectEquivalenceCount = 0 THEN
      UPDATE subjects SET has_equivalence = false WHERE id = OLD.subject_id RETURNING curriculum_id INTO oldCurriculumId;
      UPDATE curriculums SET total_subject_with_equivalence = total_subject_with_equivalence - 1 WHERE id = oldCurriculumId;
    END IF;
  END IF;

  UPDATE subjects SET has_equivalence = true WHERE id = NEW.subject_id RETURNING curriculum_id INTO newCurriculumId;

  WITH d AS (
    SELECT COUNT(1) FROM subjects
    WHERE curriculum_id = newCurriculumId AND has_equivalence IS true
  )
  UPDATE curriculums SET total_subject_with_equivalence = d.count
  FROM d
  WHERE id = newCurriculumId;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER curriculum_subjects_has_equivalence AFTER INSERT OR UPDATE ON subject_equivalences FOR EACH ROW EXECUTE PROCEDURE curriculum_subjects_has_equivalence_func();

CREATE FUNCTION curriculum_subjects_has_equivalence_delete_func()
RETURNS TRIGGER AS $$
DECLARE
  oldSubjectEquivalenceCount integer := 0;
  oldCurriculumId uuid;
BEGIN
  SELECT COUNT(1) INTO oldSubjectEquivalenceCount FROM subject_equivalences WHERE subject_id = OLD.subject_id;
  IF oldSubjectEquivalenceCount = 0 THEN
    UPDATE subjects SET has_equivalence = false WHERE id = OLD.subject_id RETURNING curriculum_id INTO oldCurriculumId;
    UPDATE curriculums SET total_subject_with_equivalence = total_subject_with_equivalence - 1 WHERE id = oldCurriculumId;
  END IF;

  RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER curriculum_subjects_has_equivalence_delete AFTER DELETE ON subject_equivalences FOR EACH ROW EXECUTE PROCEDURE curriculum_subjects_has_equivalence_delete_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER curriculum_subjects_has_equivalence_delete ON subject_equivalences;
DROP FUNCTION curriculum_subjects_has_equivalence_delete_func();

DROP TRIGGER curriculum_subjects_has_equivalence ON subject_equivalences;
DROP FUNCTION curriculum_subjects_has_equivalence_func();

ALTER TABLE "subjects" DROP COLUMN "has_equivalence";
ALTER TABLE "deleted_subjects" DROP COLUMN "has_equivalence";
ALTER TABLE "curriculums" DROP COLUMN "total_subject_with_equivalence";
ALTER TABLE "deleted_curriculums" DROP COLUMN "total_subject_with_equivalence";

-- +goose StatementEnd
