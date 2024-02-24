-- +goose Up
-- +goose StatementBegin

ALTER TABLE students ALTER COLUMN "village_id" TYPE bigint;
ALTER TABLE deleted_students ALTER COLUMN "village_id" TYPE bigint;

CREATE TABLE "report_student_provinces" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "province_id" integer NOT NULL REFERENCES "provinces" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "study_program_id" uuid NOT NULL REFERENCES "study_programs" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "student_force" integer NOT NULL,
  "total" integer NOT NULL DEFAULT 0,
  UNIQUE("province_id", "study_program_id", "student_force")
);

WITH d AS (
  SELECT
    r.province_id AS province_id,
    s.study_program_id AS study_program_id,
    s.student_force AS student_force,
    COUNT(1) AS total
  FROM students s
  JOIN villages v ON v.id = s.village_id
  JOIN districts d ON d.id = v.district_id
  JOIN regencies r ON r.id = d.regency_id
  WHERE s.study_program_id IS NOT NULL AND s.student_force IS NOT NULL
  GROUP BY r.province_id, s.study_program_id, s.student_force
)
INSERT INTO report_student_provinces (
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

CREATE FUNCTION report_student_province_func()
RETURNS TRIGGER as $$
BEGIN
  IF num_nulls(OLD.village_id, OLD.study_program_id, OLD.student_force) = 0 THEN
    UPDATE report_student_provinces rsp SET
      total = total - 1
    FROM villages v
    JOIN districts d ON d.id = v.district_id
    JOIN regencies r ON r.id = d.regency_id
    WHERE
      v.id = OLD.village_id AND 
      rsp.province_id = r.province_id AND
      rsp.study_program_id = OLD.study_program_id AND
      rsp.student_force = OLD.student_force;
  END IF;

  IF num_nulls(NEW.village_id, NEW.study_program_id, NEW.student_force) = 0 THEN
    INSERT INTO report_student_provinces (
    province_id,
    study_program_id,
    student_force,
    total
    ) SELECT
      r.province_id,
      NEW.study_program_id,
      NEW.student_force,
      1
    FROM villages v
    JOIN districts d ON d.id = v.district_id
    JOIN regencies r ON r.id = d.regency_id
    WHERE v.id = NEW.village_id
    ON CONFLICT ("province_id", "study_program_id", "student_force") DO UPDATE SET
      total = EXCLUDED.total + 1;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER report_student_province AFTER UPDATE OF village_id, study_program_id, student_force ON students FOR EACH ROW EXECUTE PROCEDURE report_student_province_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER report_student_province ON students;
DROP FUNCTION report_student_province_func();
DROP TABLE "report_student_provinces";

ALTER TABLE students ALTER COLUMN "village_id" TYPE integer;
ALTER TABLE deleted_students ALTER COLUMN "village_id" TYPE integer;

-- +goose StatementEnd
