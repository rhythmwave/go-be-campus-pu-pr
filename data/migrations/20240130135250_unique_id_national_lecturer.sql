-- +goose Up
-- +goose StatementBegin

ALTER TABLE "lecturers" ADD CONSTRAINT "lecturers_id_national_lecturer" UNIQUE ("id_national_lecturer");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "lecturers" DROP CONSTRAINT "lecturers_id_national_lecturer";

-- +goose StatementEnd
