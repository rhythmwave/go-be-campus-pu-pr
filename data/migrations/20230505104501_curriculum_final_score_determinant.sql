-- +goose Up
-- +goose StatementBegin

CREATE TYPE "curriculums_final_score_determinant" AS ENUM('best', 'last');

ALTER TABLE "curriculums" ADD COLUMN "final_score_determinant" curriculums_final_score_determinant NOT NULL DEFAULT 'best'::curriculums_final_score_determinant;
ALTER TABLE "deleted_curriculums" ADD COLUMN "final_score_determinant" curriculums_final_score_determinant NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "curriculums" DROP COLUMN "final_score_determinant";
ALTER TABLE "deleted_curriculums" DROP COLUMN "final_score_determinant";

DROP TYPE "curriculums_final_score_determinant";

-- +goose StatementEnd
