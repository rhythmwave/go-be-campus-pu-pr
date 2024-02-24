-- +goose Up
-- +goose StatementBegin

ALTER TABLE "study_plans"  ADD COLUMN "is_thesis" boolean NOT NULL DEFAULT false;
ALTER TABLE "deleted_study_plans" ADD COLUMN "is_thesis" boolean NULL;

ALTER TABLE "subjects" 
  ADD COLUMN "is_thesis" boolean NULL,
  ADD CONSTRAINT "subjects_curriculum_id_is_thesis_key" UNIQUE ("curriculum_id", "is_thesis");
ALTER TABLE "deleted_subjects" ADD COLUMN "is_thesis" boolean NULL;

ALTER TABLE "students" ADD COLUMN "has_thesis_study_plan" boolean NOT NULL DEFAULT false;
ALTER TABLE "deleted_students" ADD COLUMN "has_thesis_study_plan" boolean NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "study_plans" DROP COLUMN "is_thesis";
ALTER TABLE "deleted_study_plans" DROP COLUMN "is_thesis";

ALTER TABLE "subjects" DROP COLUMN "is_thesis";
ALTER TABLE "deleted_subjects" DROP COLUMN "is_thesis";

ALTER TABLE "students" DROP COLUMN "has_thesis_study_plan";
ALTER TABLE "deleted_students" DROP COLUMN "has_thesis_study_plan";

-- +goose StatementEnd
