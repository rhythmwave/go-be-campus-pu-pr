-- +goose Up
-- +goose StatementBegin

ALTER TABLE "semesters" 
  ADD COLUMN "grading_start_date" date NULL,
  ADD COLUMN "grading_end_date" date NULL,
  ADD CONSTRAINT "semesters_grading_date_null" CHECK(num_nonnulls("grading_start_date", "grading_end_date") IN (0,2)),
  ADD CONSTRAINT "semesters_grading_date_unique" UNIQUE("grading_start_date", "grading_end_date");

ALTER TABLE "deleted_semesters" 
  ADD COLUMN "grading_start_date" date NULL,
  ADD COLUMN "grading_end_date" date NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "deleted_semesters" 
  DROP COLUMN "grading_start_date",
  DROP COLUMN "grading_end_date";

ALTER TABLE "semesters" 
  DROP COLUMN "grading_start_date",
  DROP COLUMN "grading_end_date";

-- +goose StatementEnd
