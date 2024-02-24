-- +goose Up
-- +goose StatementBegin

ALTER TABLE "student_skpi" DROP COLUMN "created_by", DROP COLUMN "updated_by";
ALTER TABLE "deleted_student_skpi" DROP COLUMN "created_by", DROP COLUMN "updated_by";
ALTER TABLE "student_skpi_achievements" DROP COLUMN "created_by", DROP COLUMN "updated_by";
ALTER TABLE "deleted_student_skpi_achievements" DROP COLUMN "created_by", DROP COLUMN "updated_by";
ALTER TABLE "student_skpi_certificates" DROP COLUMN "created_by", DROP COLUMN "updated_by";
ALTER TABLE "deleted_student_skpi_certificates" DROP COLUMN "created_by", DROP COLUMN "updated_by";
ALTER TABLE "student_skpi_character_buildings" DROP COLUMN "created_by", DROP COLUMN "updated_by";
ALTER TABLE "deleted_student_skpi_character_buildings" DROP COLUMN "created_by", DROP COLUMN "updated_by";
ALTER TABLE "student_skpi_internships" DROP COLUMN "created_by", DROP COLUMN "updated_by";
ALTER TABLE "deleted_student_skpi_internships" DROP COLUMN "created_by", DROP COLUMN "updated_by";
ALTER TABLE "student_skpi_languages" DROP COLUMN "created_by", DROP COLUMN "updated_by";
ALTER TABLE "deleted_student_skpi_languages" DROP COLUMN "created_by", DROP COLUMN "updated_by";
ALTER TABLE "student_skpi_organizations" DROP COLUMN "created_by", DROP COLUMN "updated_by";
ALTER TABLE "deleted_student_skpi_organizations" DROP COLUMN "created_by", DROP COLUMN "updated_by";

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "student_skpi" ADD COLUMN "created_by" uuid NULL, ADD COLUMN "updated_by" uuid NULL;
ALTER TABLE "deleted_student_skpi" ADD COLUMN "created_by" uuid NULL, ADD COLUMN "updated_by" uuid NULL;
ALTER TABLE "student_skpi_achievements" ADD COLUMN "created_by" uuid NULL, ADD COLUMN "updated_by" uuid NULL;
ALTER TABLE "deleted_student_skpi_achievements" ADD COLUMN "created_by" uuid NULL, ADD COLUMN "updated_by" uuid NULL;
ALTER TABLE "student_skpi_certificates" ADD COLUMN "created_by" uuid NULL, ADD COLUMN "updated_by" uuid NULL;
ALTER TABLE "deleted_student_skpi_certificates" ADD COLUMN "created_by" uuid NULL, ADD COLUMN "updated_by" uuid NULL;
ALTER TABLE "student_skpi_character_buildings" ADD COLUMN "created_by" uuid NULL, ADD COLUMN "updated_by" uuid NULL;
ALTER TABLE "deleted_student_skpi_character_buildings" ADD COLUMN "created_by" uuid NULL, ADD COLUMN "updated_by" uuid NULL;
ALTER TABLE "student_skpi_internships" ADD COLUMN "created_by" uuid NULL, ADD COLUMN "updated_by" uuid NULL;
ALTER TABLE "deleted_student_skpi_internships" ADD COLUMN "created_by" uuid NULL, ADD COLUMN "updated_by" uuid NULL;
ALTER TABLE "student_skpi_languages" ADD COLUMN "created_by" uuid NULL, ADD COLUMN "updated_by" uuid NULL;
ALTER TABLE "deleted_student_skpi_languages" ADD COLUMN "created_by" uuid NULL, ADD COLUMN "updated_by" uuid NULL;
ALTER TABLE "student_skpi_organizations" ADD COLUMN "created_by" uuid NULL, ADD COLUMN "updated_by" uuid NULL;
ALTER TABLE "deleted_student_skpi_organizations" ADD COLUMN "created_by" uuid NULL, ADD COLUMN "updated_by" uuid NULL;

-- +goose StatementEnd
