-- +goose Up
-- +goose StatementBegin

ALTER TABLE "students" ADD CONSTRAINT "students_nim_number" UNIQUE ("nim_number");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "students" DROP CONSTRAINT "students_nim_number";

-- +goose StatementEnd
