-- +goose Up
-- +goose StatementBegin

ALTER TABLE "student_leave_requests" ADD COLUMN "request_semester_id" uuid NOT NULL REFERENCES "semesters" ("id") ON DELETE RESTRICT ON UPDATE CASCADE;
ALTER TABLE "deleted_student_leave_requests" ADD COLUMN "request_semester_id" uuid NULL;

CREATE FUNCTION student_leave_semester_func()
RETURNS TRIGGER AS $$
DECLARE semesterId uuid;
BEGIN
  SELECT id INTO semesterId FROM semesters WHERE is_active IS true;

  NEW.request_semester_id = semesterId;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER student_leave_semester BEFORE INSERT ON student_leave_requests FOR EACH ROW EXECUTE PROCEDURE student_leave_semester_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER student_leave_semester ON student_leave_requests;
DROP FUNCTION student_leave_semester_func();

ALTER TABLE "student_leave_requests" DROP COLUMN "request_semester_id";
ALTER TABLE "deleted_student_leave_requests" DROP COLUMN "request_semester_id";

-- +goose StatementEnd
