-- +goose Up
-- +goose StatementBegin

ALTER TABLE "lectures" ALTER COLUMN "lecture_plan_day_of_week" DROP NOT NULL;

CREATE FUNCTION lecture_total_plan_func()
RETURNS TRIGGER AS $$
BEGIN
  IF NEW.thesis_defense_id IS NULL THEN
    IF NEW.is_exam IS true THEN
      UPDATE classes SET total_exam_plan = total_exam_plan + 1
      WHERE id = NEW.class_id;
    ELSE
      UPDATE classes SET total_lecture_plan = total_lecture_plan + 1
      WHERE id = NEW.class_id;
    END IF;
  END IF;

  UPDATE lectures SET lecture_plan_day_of_week = EXTRACT('isodow' FROM NEW.lecture_plan_date)
  WHERE id = NEW.id;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER lecture_total_plan AFTER INSERT ON lectures FOR EACH ROW EXECUTE PROCEDURE lecture_total_plan_func();

CREATE OR REPLACE FUNCTION delete_lecture_total_plan_func()
RETURNS TRIGGER AS $$
BEGIN
  IF OLD.thesis_defense_id IS NULL THEN
    IF OLD.is_exam IS true THEN
      UPDATE classes SET total_exam_plan = total_exam_plan - 1
      WHERE id = OLD.class_id;
      IF OLD.lecture_actual_date IS NOT NULL THEN
        UPDATE classes SET total_exam_done = total_exam_done - 1
        WHERE id = OLD.class_id;
      END IF;
    ELSE
      UPDATE classes SET total_lecture_plan = total_lecture_plan - 1
      WHERE id = OLD.class_id;
      IF OLD.lecture_actual_date IS NOT NULL THEN
        UPDATE classes SET total_lecture_done = total_lecture_done - 1
        WHERE id = OLD.class_id;
      END IF;
    END IF;
  END IF;

  RETURN OLD;
END;
$$ LANGUAGE plpgsql;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

CREATE OR REPLACE FUNCTION delete_lecture_total_plan_func()
RETURNS TRIGGER AS $$
BEGIN
  IF OLD.thesis_defense_id IS NULL THEN
    IF OLD.is_exam IS true THEN
      UPDATE classes SET total_exam_plan = total_exam_plan - 1
      WHERE id = OLD.class_id;
    ELSE
      UPDATE classes SET total_lecture_plan = total_lecture_plan - 1
      WHERE id = OLD.class_id;
    END IF;
  END IF;

  RETURN OLD;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER lecture_total_plan ON lectures;
DROP FUNCTION lecture_total_plan_func();

ALTER TABLE "lectures" ALTER COLUMN "lecture_plan_day_of_week" SET NOT NULL;

-- +goose StatementEnd
