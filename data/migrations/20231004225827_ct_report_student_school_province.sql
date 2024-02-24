-- +goose Up
-- +goose StatementBegin

CREATE TABLE "report_student_school_provinces" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "province_id" integer NOT NULL REFERENCES "provinces" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "study_program_id" uuid NOT NULL REFERENCES "study_programs" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "student_force" integer NOT NULL,
  "total" integer NOT NULL DEFAULT 0,
  UNIQUE("province_id", "study_program_id", "student_force")
);

WITH d AS (
  SELECT
    s.school_province_id AS province_id,
    s.study_program_id AS study_program_id,
    s.student_force AS student_force,
    COUNT(1) AS total
  FROM students s
  WHERE s.study_program_id IS NOT NULL AND s.student_force IS NOT NULL AND s.school_province_id IS NOT NULL
  GROUP BY s.school_province_id, s.study_program_id, s.student_force
)
INSERT INTO report_student_school_provinces (
  province_id,
  study_program_id,
  student_force,
  total
) SELECT
  d.province_id,
  d.study_program_id,
  d.student_force,
  d.total
FROM d;

CREATE FUNCTION report_student_school_province_func()
RETURNS TRIGGER as $$
BEGIN
  IF num_nulls(OLD.school_province_id, OLD.study_program_id, OLD.student_force) = 0 THEN
    UPDATE report_student_school_provinces rsp SET
      total = total - 1
    WHERE
      rsp.province_id = OLD.school_province_id AND
      rsp.study_program_id = OLD.study_program_id AND
      rsp.student_force = OLD.student_force;
  END IF;

  IF num_nulls(NEW.school_province_id, NEW.study_program_id, NEW.student_force) = 0 THEN
    INSERT INTO report_student_school_provinces (
    province_id,
    study_program_id,
    student_force,
    total
    ) VALUES (
      NEW.school_province_id,
      NEW.study_program_id,
      NEW.student_force,
      1
    ) ON CONFLICT ("province_id", "study_program_id", "student_force") DO UPDATE SET
      total = EXCLUDED.total + 1;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER report_student_school_province AFTER UPDATE OF school_province_id, study_program_id, student_force ON students FOR EACH ROW EXECUTE PROCEDURE report_student_school_province_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER report_student_school_province ON students;
DROP FUNCTION report_student_school_province_func();
DROP TABLE "report_student_school_provinces";

-- +goose StatementEnd
