-- +goose Up
-- +goose StatementBegin

ALTER TABLE "lecturer_leaves" 
  ADD COLUMN "file_path" character varying NULL,
  ADD COLUMN "file_path_type" character varying NULL;

ALTER TABLE "deleted_lecturer_leaves" 
  ADD COLUMN "file_path" character varying NULL,
  ADD COLUMN "file_path_type" character varying NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "lecturer_leaves" 
  DROP COLUMN "file_path",
  DROP COLUMN "file_path_type";

ALTER TABLE "deleted_lecturer_leaves" 
  DROP COLUMN "file_path",
  DROP COLUMN "file_path_type";

-- +goose StatementEnd
