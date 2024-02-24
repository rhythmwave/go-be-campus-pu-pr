-- +goose Up
-- +goose StatementBegin

ALTER TABLE "lecture_participants" 
  ADD COLUMN "photo_path" character varying NULL,
  ADD COLUMN "photo_path_type" character varying(20) NULL;

ALTER TABLE "deleted_lecture_participants" 
  ADD COLUMN "photo_path" character varying NULL,
  ADD COLUMN "photo_path_type" character varying(20) NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "lecture_participants" 
  DROP COLUMN "photo_path",
  DROP COLUMN "photo_path_type";

ALTER TABLE "deleted_lecture_participants" 
  DROP COLUMN "photo_path",
  DROP COLUMN "photo_path_type";

-- +goose StatementEnd
