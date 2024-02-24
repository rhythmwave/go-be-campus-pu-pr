-- +goose Up
-- +goose StatementBegin

ALTER TABLE "study_programs" ADD COLUMN "minimum_graduation_credit" integer NOT NULL DEFAULT 0;
ALTER TABLE "deleted_study_programs" ADD COLUMN "minimum_graduation_credit" integer NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "study_programs" DROP COLUMN "minimum_graduation_credit";
ALTER TABLE "deleted_study_programs" DROP COLUMN "minimum_graduation_credit";

-- +goose StatementEnd
