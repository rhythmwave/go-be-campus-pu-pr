-- +goose Up
-- +goose StatementBegin

CREATE TABLE "credit_quotas" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "minimum_grade_point" numeric(3,2) NOT NULL,
  "maximum_grade_point" numeric(3,2) NOT NULL,
  "maximum_credit" integer NOT NULL,
  "created_by" uuid NOT NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "admins" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_credit_quotas" BEFORE UPDATE ON "credit_quotas" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_credit_quotas" AS TABLE "credit_quotas" WITH NO DATA;
CREATE TRIGGER "soft_delete_credit_quotas" BEFORE DELETE ON "credit_quotas" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "deleted_credit_quotas";
DROP TABLE "credit_quotas";

-- +goose StatementEnd
