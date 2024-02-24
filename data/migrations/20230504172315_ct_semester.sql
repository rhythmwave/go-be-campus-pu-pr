-- +goose Up
-- +goose StatementBegin

CREATE TABLE "semesters" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "semester_type" character varying NOT NULL,
  "semester_start_year" integer NOT NULL,
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "study_plan_input_start_date" date NOT NULL,
  "study_plan_input_end_date" date NOT NULL,
  "study_plan_approval_start_date" date NOT NULL,
  "study_plan_approval_end_date" date NOT NULL,
  "reference_semester_id" uuid NULL REFERENCES "semesters" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "check_minimum_gpa" boolean NOT NULL DEFAULT true,
  "check_passed_credit" boolean NOT NULL DEFAULT true,
  "default_credit" integer NOT NULL,
  "is_active" boolean NOT NULL DEFAULT false,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("semester_type", "semester_start_year"),
  CHECK("reference_semester_id" != "id"),
  CHECK("start_date" < "end_date"),
  CHECK("study_plan_input_start_date" < "study_plan_input_end_date"),
  CHECK("study_plan_approval_start_date" < "study_plan_approval_end_date")
);
CREATE TRIGGER "updated_at_semesters" BEFORE UPDATE ON "semesters" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_semesters" AS TABLE "semesters" WITH NO DATA;
CREATE TRIGGER "soft_delete_semesters" BEFORE DELETE ON "semesters" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

--------------------------------------------------------------------------------

CREATE FUNCTION semesters_single_is_active_func()
RETURNS TRIGGER AS $$
BEGIN
  IF OLD.is_active IS false AND NEW.is_active IS true THEN
    UPDATE semesters SET is_active = false WHERE id != NEW.id;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER semesters_single_is_active BEFORE UPDATE OF is_active ON semesters FOR EACH ROW EXECUTE PROCEDURE semesters_single_is_active_func();

--------------------------------------------------------------------------------

CREATE TABLE "semester_curriculum" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "semester_id" uuid NOT NULL REFERENCES "semesters" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "curriculum_id" uuid NOT NULL REFERENCES "curriculums" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE("semester_id", "curriculum_id")
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "semester_curriculum";

DROP TRIGGER semesters_single_is_active ON semesters;
DROP FUNCTION semesters_single_is_active_func();

DROP TABLE "deleted_semesters";
DROP TABLE "semesters";

-- +goose StatementEnd
