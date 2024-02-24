-- +goose Up
-- +goose StatementBegin

ALTER TABLE "students" ADD COLUMN "total_study_plan" integer NOT NULL DEFAULT 0;
ALTER TABLE "deleted_students" ADD COLUMN "total_study_plan" integer;

WITH d AS (
  SELECT student_id, COUNT(1) AS total
  FROM study_plans 
  WHERE is_approved IS true
  GROUP BY student_id
)
UPDATE students s SET total_study_plan = d.total
FROM d
WHERE s.id = d.student_id;

CREATE FUNCTION students_total_study_plan_func()
RETURNS TRIGGER AS $$
BEGIN
  WITH d AS (
    SELECT COUNT(1) AS total
    FROM study_plans 
    WHERE student_id = NEW.student_id AND is_approved IS true
  )
  UPDATE students s SET total_study_plan = d.total
  FROM d
  WHERE s.id = NEW.student_id;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER students_total_study_plan AFTER UPDATE OF is_approved ON study_plans FOR EACH ROW EXECUTE PROCEDURE students_total_study_plan_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER students_total_study_plan ON study_plans;
DROP FUNCTION students_total_study_plan_func();

ALTER TABLE "students" DROP COLUMN "total_study_plan";
ALTER TABLE "deleted_students" DROP COLUMN "total_study_plan";

-- +goose StatementEnd
