-- +goose Up
-- +goose StatementBegin

ALTER TABLE "study_programs" ADD COLUMN "minimum_thesis_credit" integer NOT NULL DEFAULT 0;
ALTER TABLE "deleted_study_programs" ADD COLUMN "minimum_thesis_credit" integer;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "study_programs" DROP COLUMN "minimum_thesis_credit";
ALTER TABLE "deleted_study_programs" DROP COLUMN "minimum_thesis_credit";

-- +goose StatementEnd
