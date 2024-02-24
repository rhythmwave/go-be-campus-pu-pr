-- +goose Up
-- +goose StatementBegin

ALTER TABLE "students" ADD COLUMN "graduation_predicate_id" uuid NULL REFERENCES graduation_predicates (id) ON DELETE RESTRICT ON UPDATE CASCADE;
ALTER TABLE "deleted_students" ADD COLUMN "graduation_predicate_id" uuid NULL;

CREATE FUNCTION "student_graduation_predicate_func"()
RETURNS TRIGGER AS $$
  DECLARE 
    r record;
    studySemesterCount integer;
    repeatCourse integer;
    cGradePoint numeric(3,2);
    belowCCount integer;
BEGIN
  IF NEW.status = 'LULUS' THEN
    SELECT COUNT(1) INTO studySemesterCount FROM study_plans WHERE student_id = NEW.id;
    SELECT COUNT(1) INTO repeatCourse FROM student_classes WHERE student_id = NEW.id AND subject_repetition > 1;
    SELECT grade_point INTO cGradePoint FROM grade_types WHERE code = 'C';
    SELECT COUNT(1) INTO belowCCount FROM student_subjects WHERE student_id = NEW.id AND grade_point < cGradePoint;

    FOR r IN SELECT * FROM graduation_predicates ORDER BY minimum_gpa DESC
    LOOP
      IF NEW.gpa >= r.minimum_gpa AND studySemesterCount <= r.maximum_study_semester AND repeatCourse <= r.repeat_course_limit AND belowCCount <= r.below_minimum_grade_point_limit THEN
        NEW.graduation_predicate_id = r.id;
        exit;
      END IF;
    END LOOP;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER student_graduation_predicate BEFORE INSERT OR UPDATE ON students FOR EACH ROW EXECUTE PROCEDURE student_graduation_predicate_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER student_graduation_predicate ON students;
DROP FUNCTION "student_graduation_predicate_func"();

ALTER TABLE "students" DROP COLUMN "graduation_predicate";
ALTER TABLE "deleted_students" DROP COLUMN "graduation_predicate";

-- +goose StatementEnd
