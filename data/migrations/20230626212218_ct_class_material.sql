-- +goose Up
-- +goose StatementBegin

ALTER TABLE "classes" ADD COLUMN "total_material" integer NOT NULL DEFAULT 0;
ALTER TABLE "deleted_classes" ADD COLUMN "total_material" integer NULL;

CREATE TABLE "class_materials" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "lecturer_id" uuid NOT NULL REFERENCES "lecturers" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "class_id" uuid NOT NULL REFERENCES "classes" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "title" character varying NOT NULL,
  "abstraction" text NULL, 
  "file_path" character varying NULL,
  "file_path_type" character varying(20) NULL,
  "is_active" boolean NOT NULL DEFAULT true,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("class_id", "title")
);
CREATE TRIGGER "updated_at_class_materials" BEFORE UPDATE ON "class_materials" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_class_materials" AS TABLE "class_materials" WITH NO DATA;
CREATE TRIGGER "soft_delete_class_materials" BEFORE DELETE ON "class_materials" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

--------------------------------------------

CREATE FUNCTION classes_total_material_func()
RETURNS TRIGGER AS $$
BEGIN
  IF OLD.class_id IS NOT NULL THEN
    UPDATE classes SET total_material = total_material - 1 WHERE id = OLD.class_id;
  END IF;
  IF NEW.class_id IS NOT NULL THEN
    UPDATE classes SET total_material = total_material + 1 WHERE id = NEW.class_id;
  END IF;

  IF NEW.class_id IS NOT NULL THEN
    RETURN NEW;
  ELSE
    RETURN OLD;
  END IF;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER classes_total_material AFTER INSERT OR DELETE ON class_materials FOR EACH ROW EXECUTE PROCEDURE classes_total_material_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER classes_total_material ON class_materials;
DROP FUNCTION classes_total_material_func();

DROP TABLE "class_materials";
DROP TABLE "deleted_class_materials";

ALTER TABLE "classes" DROP COLUMN "total_material";
ALTER TABLE "deleted_classes" DROP COLUMN "total_material";

-- +goose StatementEnd
