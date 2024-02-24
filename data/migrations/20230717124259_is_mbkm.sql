-- +goose Up
-- +goose StatementBegin

ALTER TABLE "classes" 
  ADD COLUMN "is_mbkm" boolean NOT NULL DEFAULT false,
  ADD COLUMN "application_deadline" date NULL;
ALTER TABLE "deleted_classes" 
  ADD COLUMN "is_mbkm" boolean NULL,
  ADD COLUMN "application_deadline" date NULL;

ALTER TABLE "student_activities" ADD COLUMN "is_mbkm" boolean NOT NULL DEFAULT false;
ALTER TABLE "deleted_student_activities" ADD COLUMN "is_mbkm" boolean NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "classes"
  DROP COLUMN "is_mbkm",
  DROP COLUMN "application_deadline";
ALTER TABLE "deleted_classes"
  DROP COLUMN "is_mbkm",
  DROP COLUMN "application_deadline";

ALTER TABLE "student_activities" DROP COLUMN "is_mbkm";
ALTER TABLE "deleted_student_activities" DROP COLUMN "is_mbkm";

-- +goose StatementEnd
