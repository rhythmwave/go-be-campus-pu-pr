-- +goose Up
-- +goose StatementBegin

ALTER TABLE "accreditations" ADD CONSTRAINT "accreditations_decree_date_decree_due_date" CHECK ("decree_date" < "decree_due_date");
ALTER TABLE "study_programs" ADD CONSTRAINT "study_programs_operational_permit_date_operational_permit_due_date" CHECK ("operational_permit_date" < "operational_permit_due_date");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "accreditations" DROP CONSTRAINT "accreditations_decree_date_decree_due_date";
ALTER TABLE "study_programs" DROP CONSTRAINT "study_programs_operational_permit_date_operational_permit_due_date";

-- +goose StatementEnd
