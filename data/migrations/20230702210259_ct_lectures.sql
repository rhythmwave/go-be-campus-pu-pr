-- +goose Up
-- +goose StatementBegin

ALTER TABLE "classes"
  ADD COLUMN "total_lecture_plan" integer NOT NULL DEFAULT 0,
  ADD COLUMN "total_lecture_done" integer NOT NULL DEFAULT 0;

ALTER TABLE "deleted_classes"
  ADD COLUMN "total_lecture_plan" integer NULL,
  ADD COLUMN "total_lecture_done" integer NULL;

ALTER TABLE "student_classes"
  ADD COLUMN "total_attendance" integer NOT NULL DEFAULT 0,
  ADD COLUMN "total_sick" integer NOT NULL DEFAULT 0,
  ADD COLUMN "total_leave" integer NOT NULL DEFAULT 0,
  ADD COLUMN "total_awol" integer NOT NULL DEFAULT 0;

ALTER TABLE "deleted_student_classes"
  ADD COLUMN "total_attendance" integer NULL,
  ADD COLUMN "total_sick" integer NULL,
  ADD COLUMN "total_leave" integer NULL,
  ADD COLUMN "total_awol" integer NULL;

CREATE TABLE "lectures" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "class_id" uuid NOT NULL REFERENCES "classes" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "schedule_id" uuid NULL REFERENCES "schedules" ("id") ON DELETE SET NULL ON UPDATE CASCADE,
  "room_id" uuid NULL REFERENCES "rooms" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "lecture_plan_date" date NOT NULL,
  "lecture_plan_day_of_week" integer NOT NULL,
  "lecture_plan_start_time" integer NULL,
  "lecture_plan_end_time" integer NULL,
  "lecturer_id" uuid NULL REFERENCES "lecturers" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "foreign_lecturer_name" character varying NULL,
  "foreign_lecturer_source_instance" character varying NULL,
  "is_original_lecturer" boolean NULL,
  "lecture_actual_date" date NULL,
  "lecture_actual_day_of_week" integer NULL,
  "lecture_actual_start_time" integer NULL,
  "lecture_actual_end_time" integer NULL,
  "lecture_theme" character varying NULL,
  "lecture_subject" text NULL,
  "remarks" text NULL,
  "total_participant" integer NOT NULL DEFAULT 0,
  "attending_participant" integer NOT NULL DEFAULT 0,
  "is_manual_participation" boolean NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  CHECK(num_nonnulls("lecturer_id", "foreign_lecturer_name") IN (0,1)),
  CHECK(num_nonnulls("foreign_lecturer_name", "foreign_lecturer_source_instance") IN (0,2))
);
CREATE TRIGGER "updated_at_lectures" BEFORE UPDATE ON "lectures" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_lectures" AS TABLE "lectures" WITH NO DATA;
CREATE TRIGGER "soft_delete_lectures" BEFORE DELETE ON "lectures" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

------------------------------------------------------------------------

CREATE TABLE "lecture_participants" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "lecture_id" uuid NOT NULL REFERENCES "lectures" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "student_id" uuid NOT NULL REFERENCES "students" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "is_attend" boolean NULL,
  "is_sick" boolean NULL,
  "is_leave" boolean NULL,
  "is_awol" boolean NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE("lecture_id", "student_id"),
  CHECK(num_nonnulls(is_attend, is_sick, is_leave, is_awol) IN (0,1))
);
CREATE TRIGGER "updated_at_lecture_participants" BEFORE UPDATE ON "lecture_participants" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_lecture_participants" AS TABLE "lecture_participants" WITH NO DATA;
CREATE TRIGGER "soft_delete_lecture_participants" BEFORE DELETE ON "lecture_participants" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

------------------------------------------------------------------------

CREATE FUNCTION lecture_total_plan_func()
RETURNS TRIGGER AS $$
DECLARE totalParticipant integer;
BEGIN
  UPDATE classes SET total_lecture_plan = total_lecture_plan + 1
  WHERE id = NEW.class_id;

  NEW.lecture_plan_day_of_week = EXTRACT('isodow' FROM NEW.lecture_plan_date);

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER lecture_total_plan BEFORE INSERT ON lectures FOR EACH ROW EXECUTE PROCEDURE lecture_total_plan_func();

------------------------------------------------------------------------

CREATE FUNCTION lecture_participant_func()
RETURNS TRIGGER AS $$
DECLARE totalParticipant integer;
BEGIN
  INSERT INTO lecture_participants (
    "lecture_id",
    "student_id"
  ) SELECT
    NEW.id,
    sc.student_id
  FROM student_classes sc WHERE sc.class_id = NEW.class_id
  ON CONFLICT (lecture_id, student_id) DO NOTHING;

  NEW.lecture_actual_day_of_week = EXTRACT('isodow' FROM NEW.lecture_actual_date);

  UPDATE classes SET total_lecture_done = total_lecture_done + 1
  WHERE id = NEW.class_id;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER lecture_participant BEFORE UPDATE OF lecture_actual_date ON lectures FOR EACH ROW EXECUTE PROCEDURE lecture_participant_func();

------------------------------------------------------------------------

CREATE FUNCTION lecture_participant_attendance_func()
RETURNS TRIGGER AS $$
  DECLARE
    isAttend integer := 0;
    isSick integer := 0;
    isLeave integer := 0;
    isAwol integer := 0;
    newAttend boolean := NEW.is_attend;
    newSick boolean := NEW.is_sick;
    newLeave boolean := NEW.is_leave;
    newAwol boolean := NEW.is_awol;
BEGIN
  IF newAttend IS true THEN
    NEW.is_sick = NULL;
    NEW.is_leave = NULL;
    NEW.is_awol = NULL;
    IF OLD.is_attend IS NULL THEN
      isAttend = 1;
    END IF;
  END IF;
  IF newSick IS true THEN
    NEW.is_attend = NULL;
    NEW.is_leave = NULL;
    NEW.is_awol = NULL;
    IF OLD.is_sick IS NULL THEN
      isSick = 1;
    END IF;
  END IF;
  IF newLeave IS true THEN
    NEW.is_attend = NULL;
    NEW.is_sick = NULL;
    NEW.is_awol = NULL;
    IF OLD.is_leave IS NULL THEN
      isLeave = 1;
    END IF;
  END IF;
  IF newAwol IS true THEN
    NEW.is_attend = NULL;
    NEW.is_sick = NULL;
    NEW.is_leave = NULL;
    IF OLD.is_awol IS NULL THEN
      isAwol = 1;
    END IF;
  END IF;

  IF OLD.is_attend IS true THEN
    IF newAttend IS NULL THEN
      isAttend = -1;
    END IF;
  END IF;
  IF OLD.is_sick IS true THEN
    IF newSick IS NULL THEN
      isSick = -1;
    END IF;
  END IF;
  IF OLD.is_leave IS true THEN
    IF newLeave IS NULL THEN
      isLeave = -1;
    END IF;
  END IF;
  IF OLD.is_awol IS true THEN
    IF newAwol IS NULL THEN
      isAwol = -1;
    END IF;
  END IF;
  
  UPDATE student_classes sc SET 
    total_attendance = sc.total_attendance + (1 * isAttend),
    total_sick = sc.total_sick + (1 * isSick),
    total_leave = sc.total_leave + (1 * isLeave),
    total_awol = sc.total_awol + (1 * isAwol)
  FROM lectures l
  WHERE l.class_id = sc.class_id AND sc.student_id = NEW.student_id AND l.id = NEW.lecture_id;

  UPDATE lectures SET attending_participant = attending_participant + (1 * isAttend)
  WHERE id = NEW.lecture_id;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER lecture_participant_attendance BEFORE UPDATE ON lecture_participants FOR EACH ROW EXECUTE PROCEDURE lecture_participant_attendance_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER lecture_total_plan ON lectures;
DROP FUNCTION lecture_total_plan_func();

DROP TRIGGER lecture_participant ON lectures;
DROP FUNCTION lecture_participant_func();

DROP TRIGGER lecture_participant_attendance ON lecture_participants;
DROP FUNCTION lecture_participant_attendance_func();

DROP TABLE "deleted_lecture_participants";
DROP TABLE "lecture_participants";

DROP TABLE "deleted_lectures";
DROP TABLE "lectures";

ALTER TABLE "student_classes"
  DROP COLUMN "total_attendance",
  DROP COLUMN "total_sick",
  DROP COLUMN "total_leave",
  DROP COLUMN "total_awol";

ALTER TABLE "deleted_student_classes"
  DROP COLUMN "total_attendance",
  DROP COLUMN "total_sick",
  DROP COLUMN "total_leave",
  DROP COLUMN "total_awol";

ALTER TABLE "deleted_classes"
  DROP COLUMN "total_lecture_plan",
  DROP COLUMN "total_lecture_done";

ALTER TABLE "classes"
  DROP COLUMN "total_lecture_plan",
  DROP COLUMN "total_lecture_done";

-- +goose StatementEnd
