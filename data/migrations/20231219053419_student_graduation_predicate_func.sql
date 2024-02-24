-- +goose Up
-- +goose StatementBegin

CREATE OR REPLACE FUNCTION "student_graduation_predicate_func"()
RETURNS TRIGGER AS $$
  DECLARE 
    r record;
    studySemesterCount integer;
    repeatCourse integer;
    cGradePoint numeric(3,2);
    belowCCount integer;
    diktiStudyProgramCode text;
    diplomaPrefix text;
    diplomaSequence integer;
BEGIN
  IF (OLD.status IS NULL OR OLD.status != 'LULUS') AND NEW.status = 'LULUS' THEN
    SELECT COUNT(1) INTO studySemesterCount FROM study_plans WHERE student_id = NEW.id;
    SELECT COUNT(1) INTO repeatCourse FROM student_classes WHERE student_id = NEW.id AND subject_repetition > 1;
    SELECT grade_point INTO cGradePoint FROM grade_types WHERE code = 'C';
    SELECT COUNT(1) INTO belowCCount FROM student_subjects WHERE student_id = NEW.id AND grade_point < cGradePoint;
    SELECT dsp.code INTO diktiStudyProgramCode 
    FROM study_programs sp
    JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
    WHERE sp.id = NEW.study_program_id;

    diplomaPrefix = CONCAT(diktiStudyProgramCode, EXTRACT('year' from now()));

    SELECT COUNT(1) INTO diplomaSequence FROM students WHERE diploma_number LIKE CONCAT(diplomaPrefix, '%');

    FOR r IN SELECT * FROM graduation_predicates ORDER BY minimum_gpa DESC
    LOOP
      IF NEW.gpa >= r.minimum_gpa AND studySemesterCount <= r.maximum_study_semester AND repeatCourse <= r.repeat_course_limit AND belowCCount <= r.below_minimum_grade_point_limit THEN
        NEW.graduation_predicate_id = r.id;
        exit;
      END IF;
    END LOOP;

    NEW.graduation_date = now();
    NEW.diploma_number = CONCAT(diplomaPrefix, diplomaSequence + 1);

  ELSIF NEW.status != 'LULUS' THEN
    NEW.graduation_predicate_id = NULL;
    NEW.graduation_date = NULL;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

CREATE OR REPLACE FUNCTION "student_graduation_predicate_func"()
RETURNS TRIGGER AS $$
  DECLARE 
    r record;
    studySemesterCount integer;
    repeatCourse integer;
    cGradePoint numeric(3,2);
    belowCCount integer;
    diktiStudyProgramCode text;
    diplomaPrefix text;
    diplomaSequence integer;
BEGIN
  IF (OLD.status IS NULL OR OLD.status != 'LULUS') AND NEW.status = 'LULUS' THEN
    SELECT COUNT(1) INTO studySemesterCount FROM study_plans WHERE student_id = NEW.id;
    SELECT COUNT(1) INTO repeatCourse FROM student_classes WHERE student_id = NEW.id AND subject_repetition > 1;
    SELECT grade_point INTO cGradePoint WHERE code = 'C';
    SELECT COUNT(1) INTO belowCCount FROM student_subjects WHERE student_id = NEW.id AND grade_point < cGradePoint;
    SELECT dsp.code INTO diktiStudyProgramCode 
    FROM study_programs sp
    JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
    WHERE sp.id = NEW.study_program_id;

    diplomaPrefix = CONCAT(diktiStudyProgramCode, EXTRACT('year' from now()));

    SELECT COUNT(1) INTO diplomaSequence FROM students WHERE diploma_number LIKE CONCAT(diplomaPrefix, '%');

    FOR r IN SELECT * FROM graduation_predicates ORDER BY minimum_gpa DESC
    LOOP
      IF NEW.gpa >= r.minimum_gpa AND studySemesterCount <= r.maximum_study_semester AND repeatCourse <= r.repeat_course_limit AND belowCCount <= r.below_minimum_grade_point_limit THEN
        NEW.graduation_predicate_id = r.id;
        exit;
      END IF;
    END LOOP;

    NEW.graduation_date = now();
    NEW.diploma_number = CONCAT(diplomaPrefix, diplomaSequence + 1);

  ELSIF NEW.status != 'LULUS' THEN
    NEW.graduation_predicate_id = NULL;
    NEW.graduation_date = NULL;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- +goose StatementEnd
