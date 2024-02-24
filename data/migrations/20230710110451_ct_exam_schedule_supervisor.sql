-- +goose Up
-- +goose StatementBegin

CREATE TABLE "exam_schedule_supervisors" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "schedule_id" uuid NOT NULL REFERENCES "schedules" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "exam_supervisor_id" uuid NOT NULL REFERENCES "exam_supervisors" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "exam_supervisor_role_id" uuid NOT NULL REFERENCES "exam_supervisor_roles" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("schedule_id", "exam_supervisor_id"),
  UNIQUE("schedule_id", "exam_supervisor_role_id")
);
CREATE TRIGGER "updated_at_exam_schedule_supervisors" BEFORE UPDATE ON "exam_schedule_supervisors" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_exam_schedule_supervisors" AS TABLE "exam_schedule_supervisors" WITH NO DATA;
CREATE TRIGGER "soft_delete_exam_schedule_supervisors" BEFORE DELETE ON "exam_schedule_supervisors" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "deleted_exam_schedule_supervisors";
DROP TABLE "exam_schedule_supervisors";

-- +goose StatementEnd
