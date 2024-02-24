-- +goose Up
-- +goose StatementBegin

ALTER TABLE "study_programs" 
  ALTER COLUMN "short_name"	SET NOT NULL,
  ALTER COLUMN "phone_number"	SET NOT NULL,
  ALTER COLUMN "email"	SET NOT NULL,
  ALTER COLUMN "contact_person"	SET NOT NULL,
  ALTER COLUMN "address" DROP DEFAULT,
  ALTER COLUMN "curiculum_review_frequency" DROP NOT NULL,
  ALTER COLUMN "curiculum_review_method" DROP NOT NULL,
  ALTER COLUMN "establishment_date" DROP NOT NULL,
  ALTER COLUMN "operational_permit_number" DROP NOT NULL,
  ALTER COLUMN "operational_permit_date" DROP NOT NULL,
  ALTER COLUMN "operational_permit_due_date" DROP NOT NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "study_programs" 
  ALTER COLUMN "short_name"	DROP NOT NULL,
  ALTER COLUMN "phone_number"	DROP NOT NULL,
  ALTER COLUMN "email"	DROP NOT NULL,
  ALTER COLUMN "contact_person"	DROP NOT NULL,
  ALTER COLUMN "address" SET DEFAULT '',
  ALTER COLUMN "curiculum_review_frequency" SET NOT NULL,
  ALTER COLUMN "curiculum_review_method" SET NOT NULL,
  ALTER COLUMN "establishment_date" SET NOT NULL,
  ALTER COLUMN "operational_permit_number" SET NOT NULL,
  ALTER COLUMN "operational_permit_date" SET NOT NULL,
  ALTER COLUMN "operational_permit_due_date" SET NOT NULL;

-- +goose StatementEnd


