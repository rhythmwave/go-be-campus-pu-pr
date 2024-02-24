-- +goose Up
-- +goose StatementBegin

ALTER TABLE "student_classes" ADD CONSTRAINT "student_classes_student_id_class_id" UNIQUE("student_id", "class_id");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "student_classes" DROP CONSTRAINT "student_classes_student_id_class_id";

-- +goose StatementEnd
