-- +goose Up
-- +goose StatementBegin

CREATE TABLE "report_student_class_grades" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "semester_id" uuid NOT NULL REFERENCES "semesters" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "subject_id" uuid NOT NULL REFERENCES "subjects" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "grade_code" character varying NOT NULL,
  "total" integer NOT NULL DEFAULT 0,
  UNIQUE("semester_id", "subject_id", "grade_code")
);

WITH d AS (
  SELECT
    se.id AS semester_id,
    s.id AS subject_id,
    sc.grade_code,
    COUNT(sc.grade_code) AS total
  FROM student_classes sc
  JOIN subjects s ON s.id = sc.subject_id
  JOIN classes c ON c.id = sc.class_id
  JOIN semesters se ON se.id = c.semester_id
  WHERE sc.grade_code IS NOT NULL
  GROUP BY se.id, s.id, sc.grade_code
)
INSERT INTO report_student_class_grades (
  semester_id,
  subject_id,
  grade_code,
  total
) SELECT
  d.semester_id,
  d.subject_id,
  d.grade_code,
  d.total
FROM d;

CREATE FUNCTION report_student_class_grade_func()
RETURNS TRIGGER as $$
BEGIN
  IF OLD.grade_code IS NOT NULL THEN
    UPDATE report_student_class_grades rscg SET
      total = total - 1
    FROM classes c
    WHERE
      c.id = OLD.class_id AND 
      c.semester_id = rscg.semester_id AND 
      rscg.subject_id = OLD.subject_id AND
      rscg.grade_code = OLD.grade_code; 
  END IF;

  INSERT INTO report_student_class_grades (
  semester_id,
  subject_id,
  grade_code,
  total
  ) SELECT
    c.semester_id,
    NEW.subject_id,
    NEW.grade_code,
    1
  FROM classes c
  WHERE c.id = NEW.class_id
  ON CONFLICT ("semester_id", "subject_id", "grade_code") DO UPDATE SET
    total = EXCLUDED.total + 1;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER report_student_class_grade AFTER UPDATE OF grade_code ON student_classes FOR EACH ROW EXECUTE PROCEDURE report_student_class_grade_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER report_student_class_grade ON student_classes;
DROP FUNCTION report_student_class_grade_func();
DROP TABLE "report_student_class_grades";

-- +goose StatementEnd
