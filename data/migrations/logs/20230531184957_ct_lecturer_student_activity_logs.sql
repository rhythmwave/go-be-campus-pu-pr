-- +goose Up
-- +goose StatementBegin

CREATE TABLE "lecturer_student_activity_logs" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "lecturer_id" uuid NULL,
  "lecturer_name" character varying NULL,
  "lecturer_username" character varying(100) NULL,
  "student_id" uuid NULL,
  "student_name" character varying NULL,
  "student_username" character varying(100) NULL,
  "module" character varying NOT NULL,
  "action" text NOT NULL,
  "ip_address" character varying NOT NULL,
  "user_agent" character varying NOT NULL,
  "execution_time" numeric(12,9) NOT NULL,
  "memory_usage" numeric(15,12) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CHECK(num_nonnulls("lecturer_id", "student_id") = 1),
  CHECK(num_nonnulls("lecturer_id", "lecturer_name", "lecturer_username") IN (0, 3)),
  CHECK(num_nonnulls("student_id", "student_name", "student_username") IN (0, 3))
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "lecturer_student_activity_logs";

-- +goose StatementEnd
