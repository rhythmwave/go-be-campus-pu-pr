-- +goose Up
-- +goose StatementBegin

CREATE TABLE "role_study_program" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "role_id" uuid NOT NULL REFERENCES "roles" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "study_program_id" uuid NOT NULL REFERENCES "study_programs" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  UNIQUE("role_id", "study_program_id")
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "role_study_program";

-- +goose StatementEnd
