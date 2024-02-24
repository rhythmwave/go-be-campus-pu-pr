-- +goose Up
-- +goose StatementBegin

ALTER TABLE "student_classes" ADD COLUMN "mbkm_used_credit" integer NOT NULL DEFAULT 0;
ALTER TABLE "deleted_student_classes" ADD COLUMN "mbkm_used_credit" integer NULL;

ALTER TABLE "student_subjects" ADD COLUMN "mbkm_subject_id" uuid NULL REFERENCES "subjects" ("id") ON DELETE RESTRICT ON UPDATE CASCADE;
ALTER TABLE "deleted_student_subjects" ADD COLUMN "mbkm_subject_id" uuid NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "student_subjects" DROP COLUMN "mbkm_subject_id";
ALTER TABLE "deleted_student_subjects" DROP COLUMN "mbkm_subject_id";

ALTER TABLE "student_classes" DROP COLUMN "mbkm_used_credit";
ALTER TABLE "deleted_student_classes" DROP COLUMN "mbkm_used_credit";

-- +goose StatementEnd
