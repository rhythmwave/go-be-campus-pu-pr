-- +goose Up
-- +goose StatementBegin

ALTER TABLE "class_lecturers" ADD COLUMN "total_attendance" integer NOT NULL DEFAULT 0;
ALTER TABLE "deleted_class_lecturers" ADD COLUMN "total_attendance" integer NULL;

CREATE FUNCTION lecture_lecturer_attendance_func()
RETURNS TRIGGER AS $$
BEGIN
  IF OLD.lecturer_id IS NOT NULL THEN
    UPDATE class_lecturers SET total_attendance = total_attendance - 1 WHERE id = OLD.lecturer_id;
  END IF;
  IF NEW.lecturer_id IS NOT NULL THEN
    UPDATE class_lecturers SET total_attendance = total_attendance + 1 WHERE id = NEW.lecturer_id;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER lecture_lecturer_attendance BEFORE UPDATE OF lecturer_id ON lectures FOR EACH ROW EXECUTE PROCEDURE lecture_lecturer_attendance_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER lecture_lecturer_attendance ON lectures;
DROP FUNCTION lecture_lecturer_attendance_func();

ALTER TABLE "class_lecturers" ADD COLUMN "total_attendance" integer NOT NULL DEFAULT 0;
ALTER TABLE "deleted_class_lecturers" ADD COLUMN "total_attendance" integer NULL;

-- +goose StatementEnd
