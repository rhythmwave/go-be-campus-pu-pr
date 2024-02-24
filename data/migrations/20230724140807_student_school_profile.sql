-- +goose Up
-- +goose StatementBegin

ALTER TABLE "students"
  ADD COLUMN "school_province_id" integer NULL REFERENCES "provinces"("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  ADD COLUMN "school_type" character varying NULL,
  ADD COLUMN "school_enrollment_year" character(4) NULL,
  ADD COLUMN "school_enrollment_class" character varying NULL,
  ADD COLUMN "school_status" character varying NULL,
  ADD COLUMN "school_mathematics_final_exam_score" numeric(4,2) NULL,
  ADD COLUMN "school_indonesian_final_exam_score" numeric(4,2) NULL,
  ADD COLUMN "school_english_final_exam_score" numeric(4,2) NULL,
  ADD COLUMN "school_mathematics_report_score" numeric(4,2) NULL,
  ADD COLUMN "school_indonesian_report_score" numeric(4,2) NULL,
  ADD COLUMN "school_english_report_score" numeric(4,2) NULL;

ALTER TABLE "deleted_students"
  ADD COLUMN "school_province_id" integer NULL,
  ADD COLUMN "school_type" character varying NULL,
  ADD COLUMN "school_enrollment_year" character(4) NULL,
  ADD COLUMN "school_enrollment_class" character varying NULL,
  ADD COLUMN "school_status" character varying NULL,
  ADD COLUMN "school_mathematics_final_exam_score" numeric(4,2) NULL,
  ADD COLUMN "school_indonesian_final_exam_score" numeric(4,2) NULL,
  ADD COLUMN "school_english_final_exam_score" numeric(4,2) NULL,
  ADD COLUMN "school_mathematics_report_score" numeric(4,2) NULL,
  ADD COLUMN "school_indonesian_report_score" numeric(4,2) NULL,
  ADD COLUMN "school_english_report_score" numeric(4,2) NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "students"
  DROP COLUMN "school_province_id",
  DROP COLUMN "school_type",
  DROP COLUMN "school_enrollment_year",
  DROP COLUMN "school_enrollment_class",
  DROP COLUMN "school_status",
  DROP COLUMN "school_mathematics_final_exam_score",
  DROP COLUMN "school_indonesian_final_exam_score",
  DROP COLUMN "school_english_final_exam_score",
  DROP COLUMN "school_mathematics_report_score",
  DROP COLUMN "school_indonesian_report_score",
  DROP COLUMN "school_english_report_score";

ALTER TABLE "deleted_students"
  DROP COLUMN "school_province_id",
  DROP COLUMN "school_type",
  DROP COLUMN "school_enrollment_year",
  DROP COLUMN "school_enrollment_class",
  DROP COLUMN "school_status",
  DROP COLUMN "school_mathematics_final_exam_score",
  DROP COLUMN "school_indonesian_final_exam_score",
  DROP COLUMN "school_english_final_exam_score",
  DROP COLUMN "school_mathematics_report_score",
  DROP COLUMN "school_indonesian_report_score",
  DROP COLUMN "school_english_report_score";

-- +goose StatementEnd
