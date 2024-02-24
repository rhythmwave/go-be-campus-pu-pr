-- +goose Up
-- +goose StatementBegin

ALTER TABLE "curriculums" ADD COLUMN "is_active" BOOLEAN NOT NULL DEFAULT false;
ALTER TABLE "deleted_curriculums" ADD COLUMN "is_active" BOOLEAN NULL;

---------------------------------------------------------------------------------

CREATE FUNCTION curriculum_is_active_func()
RETURNS TRIGGER AS $$
BEGIN
  IF OLD.is_active IS false AND NEW.is_active IS true THEN
    UPDATE curriculums SET is_active = false WHERE is_active IS true AND study_program_id = NEW.study_program_id;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER curriculum_is_active_update BEFORE INSERT OR UPDATE OF is_active ON curriculums FOR EACH ROW EXECUTE PROCEDURE curriculum_is_active_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER curriculum_is_active_update ON curriculums;
DROP FUNCTION curriculum_is_active_func();

ALTER TABLE "curriculums" DROP COLUMN "is_active";
ALTER TABLE "deleted_curriculums" DROP COLUMN "is_active";

-- +goose StatementEnd
