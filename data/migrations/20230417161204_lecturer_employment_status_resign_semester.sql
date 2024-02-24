-- +goose Up
-- +goose StatementBegin

ALTER TABLE "lecturers" 
  ADD COLUMN "employment_status" character varying NULL,
  ADD COLUMN "resign_semester" character varying NULL;
ALTER TABLE "deleted_lecturers" 
  ADD COLUMN "employment_status" character varying NULL,
  ADD COLUMN "resign_semester" character varying NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "lecturers" 
  DROP COLUMN "employment_status",
  DROP COLUMN "resign_semester";
ALTER TABLE "deleted_lecturers" 
  DROP COLUMN "employment_status",
  DROP COLUMN "resign_semester";

-- +goose StatementEnd
