-- +goose Up
-- +goose StatementBegin

ALTER TABLE "authentications"
  ADD COLUMN "is_active" boolean NOT NULL DEFAULT true,
  ADD COLUMN "suspension_remarks" character varying NULL;

----------------------------------------

CREATE FUNCTION authentications_suspension_remarks_func()
RETURNS TRIGGER AS $$
BEGIN
  IF OLD.is_active IS false AND NEW.is_active IS true THEN
    NEW.suspension_remarks = NULL;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER authentications_suspension_remarks BEFORE UPDATE OF is_active ON authentications FOR EACH ROW EXECUTE PROCEDURE authentications_suspension_remarks_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER authentications_suspension_remarks ON authentications;
DROP FUNCTION authentications_suspension_remarks_func();

ALTER TABLE "authentications"
  DROP COLUMN "is_active",
  DROP COLUMN "suspension_remarks";

-- +goose StatementEnd
