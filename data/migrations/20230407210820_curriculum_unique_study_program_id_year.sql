-- +goose Up
-- +goose StatementBegin

ALTER TABLE "curriculums" ADD CONSTRAINT "curriculums_study_program_id_year" UNIQUE ("study_program_id", "year");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "curriculums" DROP CONSTRAINT "curriculums_study_program_id_year";

-- +goose StatementEnd
