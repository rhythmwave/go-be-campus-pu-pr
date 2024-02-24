-- +goose Up
-- +goose StatementBegin

ALTER TABLE "students"
  ADD COLUMN "profile_photo_path" character varying NULL,
  ADD COLUMN "profile_photo_path_type" character varying NULL,
  ADD COLUMN "blood_type" character varying NULL,
  ADD COLUMN "height" numeric(5,2) NULL,
  ADD COLUMN "weight" numeric(5,2) NULL,
  ADD COLUMN "is_color_blind" boolean NULL,
  ADD COLUMN "use_glasses" boolean NULL,
  ADD COLUMN "has_complete_teeth" boolean NULL,
  ADD COLUMN "address_type" character varying NULL,
  ADD COLUMN "is_kps_recipient" boolean NULL,
  ADD COLUMN "work_address" character varying NULL,
  ADD COLUMN "assurance_number" character varying NULL;

ALTER TABLE "deleted_students"
  ADD COLUMN "profile_photo_path" character varying NULL,
  ADD COLUMN "profile_photo_path_type" character varying NULL,
  ADD COLUMN "blood_type" character varying NULL,
  ADD COLUMN "height" numeric(5,2) NULL,
  ADD COLUMN "weight" numeric(5,2) NULL,
  ADD COLUMN "is_color_blind" boolean NULL,
  ADD COLUMN "use_glasses" boolean NULL,
  ADD COLUMN "has_complete_teeth" boolean NULL,
  ADD COLUMN "address_type" character varying NULL,
  ADD COLUMN "is_kps_recipient" boolean NULL,
  ADD COLUMN "work_address" character varying NULL,
  ADD COLUMN "assurance_number" character varying NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "students"
  DROP COLUMN "profile_photo_path",
  DROP COLUMN "profile_photo_path_type",
  DROP COLUMN "blood_type",
  DROP COLUMN "height",
  DROP COLUMN "weight",
  DROP COLUMN "is_color_blind",
  DROP COLUMN "use_glasses",
  DROP COLUMN "has_complete_teeth",
  DROP COLUMN "address_type",
  DROP COLUMN "is_kps_recipient",
  DROP COLUMN "work_address",
  DROP COLUMN "assurance_number";

ALTER TABLE "deleted_students"
  DROP COLUMN "profile_photo_path",
  DROP COLUMN "profile_photo_path_type",
  DROP COLUMN "blood_type",
  DROP COLUMN "height",
  DROP COLUMN "weight",
  DROP COLUMN "is_color_blind",
  DROP COLUMN "use_glasses",
  DROP COLUMN "has_complete_teeth",
  DROP COLUMN "address_type",
  DROP COLUMN "is_kps_recipient",
  DROP COLUMN "work_address",
  DROP COLUMN "assurance_number";


-- +goose StatementEnd
