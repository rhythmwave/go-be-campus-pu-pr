-- +goose Up
-- +goose StatementBegin

ALTER TABLE "officers" ADD COLUMN "employee_no" character varying NULL;
ALTER TABLE "deleted_officers" ADD COLUMN "employee_no" character varying NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "officers" DROP COLUMN "employee_no";
ALTER TABLE "deleted_officers" DROP COLUMN "employee_no";

-- +goose StatementEnd
