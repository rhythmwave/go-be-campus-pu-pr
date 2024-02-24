-- +goose Up
-- +goose StatementBegin

DROP TRIGGER lecture_participant ON lectures;
CREATE OR REPLACE FUNCTION lecture_participant_func()
RETURNS TRIGGER AS $$
BEGIN
  IF NEW.thesis_defense_id IS NULL AND NEW.lecture_actual_date IS NOT NULL THEN
    INSERT INTO lecture_participants (
      "lecture_id",
      "student_id"
    ) SELECT
      NEW.id,
      sc.student_id
    FROM student_classes sc WHERE sc.class_id = NEW.class_id
    ON CONFLICT (lecture_id, student_id) DO NOTHING;

    UPDATE lectures SET lecture_actual_day_of_week = EXTRACT('isodow' FROM NEW.lecture_actual_date)
    WHERE id = NEW.id;

    IF NEW.is_exam IS false THEN
      UPDATE classes SET total_lecture_done = total_lecture_done + 1
      WHERE id = NEW.class_id;
    ELSE
      UPDATE classes SET total_exam_done = total_exam_done + 1
      WHERE id = NEW.class_id;
    END IF;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER lecture_participant AFTER INSERT OR UPDATE OF lecture_actual_date ON lectures FOR EACH ROW EXECUTE PROCEDURE lecture_participant_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER lecture_participant ON lectures;
CREATE OR REPLACE FUNCTION lecture_participant_func()
RETURNS TRIGGER AS $$
BEGIN
  IF NEW.thesis_defense_id IS NULL THEN
    INSERT INTO lecture_participants (
      "lecture_id",
      "student_id"
    ) SELECT
      NEW.id,
      sc.student_id
    FROM student_classes sc WHERE sc.class_id = NEW.class_id
    ON CONFLICT (lecture_id, student_id) DO NOTHING;

    NEW.lecture_actual_day_of_week = EXTRACT('isodow' FROM NEW.lecture_actual_date);

    IF NEW.is_exam IS false THEN
      UPDATE classes SET total_lecture_done = total_lecture_done + 1
      WHERE id = NEW.class_id;
    ELSE
      UPDATE classes SET total_exam_done = total_exam_done + 1
      WHERE id = NEW.class_id;
    END IF;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER lecture_participant BEFORE UPDATE OF lecture_actual_date ON lectures FOR EACH ROW EXECUTE PROCEDURE lecture_participant_func();

-- +goose StatementEnd
