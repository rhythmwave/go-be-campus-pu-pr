-- +goose Up
-- +goose StatementBegin

CREATE TABLE "study_levels" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" character varying NOT NULL UNIQUE,
  "short_name" character varying(10) NOT NULL UNIQUE
);

CREATE TABLE "dikti_study_programs" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "code" character varying NOT NULL,
  "name" character varying NOT NULL,
  "type" character varying NOT NULL DEFAULT 'regular',
  "study_level_id" uuid NOT NULL REFERENCES "study_levels" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "dikti_study_programs";
DROP TABLE "study_levels";

-- +goose StatementEnd
