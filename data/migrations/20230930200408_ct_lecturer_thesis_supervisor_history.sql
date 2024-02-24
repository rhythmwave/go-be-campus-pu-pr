-- +goose Up
-- +goose StatementBegin

ALTER TABLE "lecturers" ADD COLUMN "total_supervised_thesis" integer NOT NULL DEFAULT 0;
ALTER TABLE "deleted_lecturers" ADD COLUMN "total_supervised_thesis" integer NULL;

CREATE FUNCTION lecturers_total_supervised_thesis_func()
RETURNS TRIGGER AS $$
BEGIN
  IF NEW.id IS NULL THEN
    UPDATE lecturers SET total_supervised_thesis = total_supervised_thesis - 1
    WHERE id = OLD.lecturer_id;

    RETURN OLD;
  ELSE
    UPDATE lecturers SET total_supervised_thesis = total_supervised_thesis + 1
    WHERE id = NEW.lecturer_id;

    RETURN NEW;
  END IF;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER lecturers_total_supervised_thesis AFTER INSERT OR DELETE ON thesis_supervisors FOR EACH ROW EXECUTE PROCEDURE lecturers_total_supervised_thesis_func();

CREATE TABLE thesis_supervisor_logs (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "lecturer_id" uuid NOT NULL REFERENCES "lecturers" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "semester_id" uuid NOT NULL REFERENCES "semesters" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "thesis_supervisor_role_id" uuid NOT NULL REFERENCES "thesis_supervisor_roles" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "total" integer NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("lecturer_id", "semester_id", "thesis_supervisor_role_id")
);
CREATE TRIGGER "updated_at_thesis_supervisor_logs" BEFORE UPDATE ON "thesis_supervisor_logs" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE FUNCTION thesis_supervisor_log_func()
RETURNS TRIGGER AS $$
BEGIN
  IF OLD.is_active IS true AND NEW.is_active IS false THEN
    WITH d AS (
      SELECT ts.lecturer_id, ts.thesis_supervisor_role_id, COUNT(1) AS total
      FROM thesis_supervisors ts
      JOIN theses t ON t.id = ts.thesis_id
      WHERE t.status = 'SEDANG DIKERJAKAN'
      GROUP BY ts.lecturer_id, ts.thesis_supervisor_role_id
    ) INSERT INTO thesis_supervisor_logs (
      lecturer_id,
      semester_id,
      thesis_supervisor_role_id,
      total
    ) SELECT
      d.lecturer_id,
      NEW.id,
      d.thesis_supervisor_role_id,
      d.total
    FROM d
    ON CONFLICT ("lecturer_id", "semester_id", "thesis_supervisor_role_id") DO UPDATE SET
    total = EXCLUDED.total;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER thesis_supervisor_log AFTER UPDATE OF is_active ON semesters FOR EACH ROW EXECUTE PROCEDURE thesis_supervisor_log_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER thesis_supervisor_log ON semesters;
DROP FUNCTION thesis_supervisor_log_func();

DROP TABLE thesis_supervisor_logs;

DROP TRIGGER lecturers_total_supervised_thesis ON thesis_supervisors;
DROP FUNCTION lecturers_total_supervised_thesis_func();

ALTER TABLE "lecturers" DROP COLUMN "total_supervised_thesis";
ALTER TABLE "deleted_lecturers" DROP COLUMN "total_supervised_thesis";

-- +goose StatementEnd
