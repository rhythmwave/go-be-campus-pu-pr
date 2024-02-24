-- +goose Up
-- +goose StatementBegin

ALTER TABLE "lecture_participants" DROP CONSTRAINT "lecture_participants_lecture_id_fkey";
ALTER TABLE "lecture_participants" ADD CONSTRAINT "lecture_participants_lecture_id_fkey" FOREIGN KEY ("lecture_id") REFERENCES "lectures" ("id") ON DELETE CASCADE ON UPDATE CASCADE; 

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "lecture_participants" DROP CONSTRAINT "lecture_participants_lecture_id_fkey";
ALTER TABLE "lecture_participants" ADD CONSTRAINT "lecture_participants_lecture_id_fkey" FOREIGN KEY ("lecture_id") REFERENCES "lectures" ("id") ON DELETE RESTRICT ON UPDATE CASCADE; 

-- +goose StatementEnd
