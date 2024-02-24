-- +goose Up
-- +goose StatementBegin

CREATE TABLE "authentications" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "username" character varying NOT NULL UNIQUE,
  "password" character varying NOT NULL,
  "admin_id" uuid NULL REFERENCES "admins" ("id") ON DELETE CASCADE ON UPDATE CASCADE UNIQUE,
  "lecturer_id" uuid NULL REFERENCES "lecturers" ("id") ON DELETE CASCADE ON UPDATE CASCADE UNIQUE,
  "student_id" uuid NULL REFERENCES "students" ("id") ON DELETE CASCADE ON UPDATE CASCADE UNIQUE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  CHECK(num_nonnulls("admin_id", "lecturer_id", "student_id") = 1)
);
CREATE TRIGGER "updated_at_authentications" BEFORE UPDATE ON "authentications" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

INSERT INTO "authentications" (
  "username",
  "password",
  "admin_id"
) SELECT
  username,
  password,
  id
FROM admins;

------------------------------------------------------------

CREATE FUNCTION admin_authentication_func()
RETURNS TRIGGER AS $$
  BEGIN
    INSERT INTO "authentications" (
      "username",
      "password",
      "admin_id"
    ) VALUES (
      NEW.username,
      NEW.password,
      NEW.id
    ) ON CONFLICT (admin_id) DO UPDATE SET
      username = EXCLUDED.username,
      password = EXCLUDED.password;

    RETURN NEW;
  END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER admin_authentication AFTER INSERT OR UPDATE ON admins FOR EACH ROW EXECUTE PROCEDURE admin_authentication_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER admin_authentication ON admins;
DROP FUNCTION admin_authentication_func();

DROP TABLE "authentications";

-- +goose StatementEnd
