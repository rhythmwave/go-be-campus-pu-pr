-- +goose Up
-- +goose StatementBegin

ALTER TABLE "classes" ADD COLUMN "total_event" integer NOT NULL DEFAULT 0;
ALTER TABLE "deleted_classes" ADD COLUMN "total_event" integer NULL;

CREATE TABLE "class_events" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "lecturer_id" uuid NOT NULL REFERENCES "lecturers" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "class_id" uuid NOT NULL REFERENCES "classes" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "title" character varying NOT NULL,
  "frequency" character varying NOT NULL, 
  "event_time" timestamp NOT NULL,
  "remarks" character varying NULL,
  "is_active" boolean NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("class_id", "title")
);
CREATE TRIGGER "updated_at_class_events" BEFORE UPDATE ON "class_events" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_class_events" AS TABLE "class_events" WITH NO DATA;
CREATE TRIGGER "soft_delete_class_events" BEFORE DELETE ON "class_events" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

--------------------------------------------

CREATE FUNCTION classes_total_event_func()
RETURNS TRIGGER AS $$
BEGIN
  IF OLD.class_id IS NOT NULL THEN
    UPDATE classes SET total_event = total_event - 1 WHERE id = OLD.class_id;
  END IF;
  IF NEW.class_id IS NOT NULL THEN
    UPDATE classes SET total_event = total_event + 1 WHERE id = NEW.class_id;
  END IF;

  IF NEW.class_id IS NOT NULL THEN
    RETURN NEW;
  ELSE
    RETURN OLD;
  END IF;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER classes_total_event AFTER INSERT OR DELETE ON class_events FOR EACH ROW EXECUTE PROCEDURE classes_total_event_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER classes_total_event ON class_events;
DROP FUNCTION classes_total_event_func();

DROP TABLE "class_events";
DROP TABLE "deleted_class_events";

ALTER TABLE "classes" DROP COLUMN "total_event";
ALTER TABLE "deleted_classes" DROP COLUMN "total_event";

-- +goose StatementEnd
