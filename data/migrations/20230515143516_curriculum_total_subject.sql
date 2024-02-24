-- +goose Up
-- +goose StatementBegin

ALTER TABLE "curriculums" ADD COLUMN "total_subject" integer NOT NULL DEFAULT 0;
ALTER TABLE "deleted_curriculums" ADD COLUMN "total_subject" integer NULL;

WITH d AS (
  SELECT curriculum_id, COUNT(1) FROM subjects
  GROUP BY curriculum_id
)
UPDATE curriculums SET total_subject = d.count
FROM d
WHERE id = d.curriculum_id;

CREATE FUNCTION curriculums_total_subject_func()
RETURNS TRIGGER AS $$
BEGIN
  IF OLD.curriculum_id IS NOT NULL AND OLD.curriculum_id != NEW.curriculum_id THEN
    UPDATE curriculums SET total_subject = total_subject - 1 WHERE id = OLD.curriculum_id;
  END IF;

  WITH d AS (
    SELECT COUNT(1) FROM subjects
    WHERE curriculum_id = NEW.curriculum_id
  )
  UPDATE curriculums SET total_subject = d.count
  FROM d
  WHERE id = NEW.curriculum_id;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER curriculums_total_subject AFTER INSERT OR UPDATE ON subjects FOR EACH ROW EXECUTE PROCEDURE curriculums_total_subject_func();

CREATE FUNCTION curriculums_total_subject_delete_func()
RETURNS TRIGGER AS $$
BEGIN
  UPDATE curriculums SET total_subject = total_subject - 1 WHERE id = OLD.curriculum_id;

  RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER curriculums_total_subject_delete AFTER DELETE ON subjects FOR EACH ROW EXECUTE PROCEDURE curriculums_total_subject_delete_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER curriculums_total_subject ON subjects;
DROP FUNCTION curriculums_total_subject_func();

DROP TRIGGER curriculums_total_subject_delete ON subjects;
DROP FUNCTION curriculums_total_subject_delete_func();

ALTER TABLE "curriculums" DROP COLUMN "total_subject";
ALTER TABLE "deleted_curriculums" DROP COLUMN "total_subject";

-- +goose StatementEnd
