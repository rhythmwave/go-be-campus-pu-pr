-- +goose Up
-- +goose StatementBegin

CREATE TABLE "stat_lecturer_academic_positions" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "study_program_id" uuid NOT NULL REFERENCES "study_programs"("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "academic_position" character varying NOT NULL,
  "total" integer NOT NULL DEFAULT 0,
  UNIQUE("study_program_id", "academic_position")
);

WITH d AS (
  SELECT
    study_program_id,
    academic_position,
    COUNT(1) AS total
  FROM lecturers
  WHERE study_program_id IS NOT NULL AND academic_position IS NOT NULL
  GROUP BY
    study_program_id,
    academic_position
)
INSERT INTO stat_lecturer_academic_positions (
  study_program_id,
  academic_position,
  total
) SELECT 
  study_program_id,
  academic_position,
  total
FROM d;

CREATE FUNCTION "stat_lecturer_academic_positions_func"()
RETURNS TRIGGER AS $$
BEGIN
  IF (num_nulls(NEW.study_program_id, NEW.academic_position) = 0) THEN
    INSERT INTO stat_lecturer_academic_positions AS x (
      study_program_id,
      academic_position,
      total
    ) VALUES ( 
      NEW.study_program_id,
      NEW.academic_position,
      1
    ) ON CONFLICT (study_program_id, academic_position) DO UPDATE SET total = x.total + 1;
  END IF;
  
  UPDATE stat_lecturer_academic_positions SET total = total - 1
  WHERE study_program_id = OLD.study_program_id AND academic_position = OLD.academic_position;

  IF NEW.id IS NULL THEN
    RETURN OLD;
  ELSE
    RETURN NEW;
  END IF;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER stat_lecturer_academic_positions AFTER INSERT OR UPDATE OR DELETE ON lecturers FOR EACH ROW EXECUTE PROCEDURE stat_lecturer_academic_positions_func();

----------------------------------------------------------------

CREATE TABLE "stat_lecturer_number_type" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "study_program_id" uuid NOT NULL REFERENCES "study_programs"("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "number_type" character varying NOT NULL,
  "total" integer NOT NULL DEFAULT 0,
  UNIQUE("study_program_id", "number_type")
);

WITH d AS (
  SELECT
    study_program_id,
    COUNT(1) AS total
  FROM lecturers
  WHERE study_program_id IS NOT NULL
  GROUP BY
    study_program_id
)
INSERT INTO stat_lecturer_number_type (
  study_program_id,
  number_type,
  total
) SELECT 
  study_program_id,
  'NIDN',
  total
FROM d;

CREATE FUNCTION "stat_lecturer_number_type_func"()
RETURNS TRIGGER AS $$
BEGIN
  IF (num_nulls(NEW.study_program_id) = 0) THEN
    INSERT INTO stat_lecturer_number_type AS x (
      study_program_id,
      number_type,
      total
    ) VALUES ( 
      NEW.study_program_id,
      'NIDN',
      1
    ) ON CONFLICT (study_program_id, number_type) DO UPDATE SET total = x.total + 1;
  END IF;
  
  UPDATE stat_lecturer_number_type SET total = total - 1
  WHERE study_program_id = OLD.study_program_id AND number_type = 'NIDN';

  IF NEW.id IS NULL THEN
    RETURN OLD;
  ELSE
    RETURN NEW;
  END IF;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER stat_lecturer_number_type AFTER INSERT OR UPDATE OR DELETE ON lecturers FOR EACH ROW EXECUTE PROCEDURE stat_lecturer_number_type_func();

----------------------------------------------------------------

CREATE TABLE "stat_lecturer_statuses" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "study_program_id" uuid NOT NULL REFERENCES "study_programs"("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "status" character varying NOT NULL,
  "total" integer NOT NULL DEFAULT 0,
  UNIQUE("study_program_id", "status")
);

WITH a AS (
  SELECT 
    study_program_id,
    (
      CASE 
        WHEN employee_status LIKE '%PNS%' AND employee_status != 'NON PNS' THEN 'PNS'
        ELSE 'NON PNS'
      END
    ) AS status
  FROM lecturers
  WHERE study_program_id IS NOT NULL AND employee_status IS NOT NULL
), d AS (
  SELECT
    study_program_id,
    status,
    COUNT(1) AS total
  FROM a
  GROUP BY
    study_program_id,
    status
)
INSERT INTO stat_lecturer_statuses (
  study_program_id,
  status,
  total
) SELECT 
  study_program_id,
  status,
  total
FROM d;

CREATE FUNCTION "stat_lecturer_statuses_func"()
RETURNS TRIGGER AS $$
BEGIN
  IF (num_nulls(NEW.study_program_id, NEW.employee_status) = 0) THEN
    INSERT INTO stat_lecturer_statuses AS x (
      study_program_id,
      status,
      total
    ) VALUES ( 
      NEW.study_program_id,
      (
        CASE 
          WHEN NEW.employee_status LIKE '%PNS%' AND NEW.employee_status != 'NON PNS' THEN 'PNS'
          ELSE 'NON PNS'
        END
      ),
      1
    ) ON CONFLICT (study_program_id, status) DO UPDATE SET total = x.total + 1;
  END IF;
  
  IF OLD.employee_status IS NOT NULL THEN
    UPDATE stat_lecturer_statuses SET total = total - 1
    WHERE study_program_id = OLD.study_program_id AND 
      status = (
        CASE 
          WHEN OLD.employee_status LIKE '%PNS%' AND OLD.employee_status != 'NON PNS' THEN 'PNS'
          ELSE 'NON PNS'
        END
      );
  END IF;

  IF NEW.id IS NULL THEN
    RETURN OLD;
  ELSE
    RETURN NEW;
  END IF;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER stat_lecturer_statuses AFTER INSERT OR UPDATE OR DELETE ON lecturers FOR EACH ROW EXECUTE PROCEDURE stat_lecturer_statuses_func();

----------------------------------------------------------------

DROP TRIGGER stat_lecturer_sexes ON lecturers;
DROP FUNCTION "stat_lecturer_sexes_func"();
DROP TABLE "stat_lecturer_sexes";
CREATE TABLE "stat_lecturer_sexes" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "study_program_id" uuid NOT NULL REFERENCES "study_programs"("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "sex" character varying NOT NULL,
  "total" integer NOT NULL DEFAULT 0,
  UNIQUE("study_program_id", "sex")
);

WITH d AS (
  SELECT
    study_program_id,
    sex,
    COUNT(1) AS total
  FROM lecturers
  WHERE study_program_id IS NOT NULL AND sex IS NOT NULL
  GROUP BY
    study_program_id,
    sex
)
INSERT INTO stat_lecturer_sexes (
  study_program_id,
  sex,
  total
) SELECT 
  study_program_id,
  sex,
  total
FROM d;

CREATE FUNCTION "stat_lecturer_sexes_func"()
RETURNS TRIGGER AS $$
BEGIN
  IF num_nulls(NEW.study_program_id, NEW.sex) = 0 THEN
    INSERT INTO stat_lecturer_sexes AS x (
      study_program_id,
      sex,
      total
    ) VALUES ( 
      NEW.study_program_id,
      NEW.sex,
      1
    ) ON CONFLICT (study_program_id, sex) DO UPDATE SET total = x.total + 1;
  END IF;
  
  UPDATE stat_lecturer_sexes SET total = total - 1
  WHERE study_program_id = OLD.study_program_id AND sex = OLD.sex;

  IF NEW.id IS NULL THEN
    RETURN OLD;
  ELSE
    RETURN NEW;
  END IF;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER stat_lecturer_sexes AFTER INSERT OR UPDATE OR DELETE ON lecturers FOR EACH ROW EXECUTE PROCEDURE stat_lecturer_sexes_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER stat_lecturer_statuses ON lecturers;
DROP FUNCTION "stat_lecturer_statuses_func"();
DROP TABLE "stat_lecturer_statuses";

DROP TRIGGER stat_lecturer_number_type ON lecturers;
DROP FUNCTION "stat_lecturer_number_type_func"();
DROP TABLE "stat_lecturer_number_type";

DROP TRIGGER stat_lecturer_sexes ON lecturers;
DROP FUNCTION "stat_lecturer_sexes_func"();
DROP TABLE "stat_lecturer_sexes";
CREATE TABLE "stat_lecturer_sexes" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "sex" character varying NOT NULL UNIQUE,
  "total" integer NOT NULL DEFAULT 0
);

WITH d AS (
  SELECT
    sex,
    COUNT(1) AS total
  FROM lecturers
  WHERE sex IS NOT NULL
  GROUP BY
    sex
)
INSERT INTO stat_lecturer_sexes (
  sex,
  total
) SELECT 
  sex,
  total
FROM d;

CREATE FUNCTION "stat_lecturer_sexes_func"()
RETURNS TRIGGER AS $$
BEGIN
  IF NEW.sex IS NOT NULL THEN
    INSERT INTO stat_lecturer_sexes AS x (
      sex,
      total
    ) VALUES ( 
      NEW.sex,
      1
    ) ON CONFLICT (sex) DO UPDATE SET total = x.total + 1;
  END IF;
  
  UPDATE stat_lecturer_sexes SET total = total - 1
  WHERE sex = OLD.sex;

  IF NEW.id IS NULL THEN
    RETURN OLD;
  ELSE
    RETURN NEW;
  END IF;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER stat_lecturer_sexes AFTER INSERT OR UPDATE OR DELETE ON lecturers FOR EACH ROW EXECUTE PROCEDURE stat_lecturer_sexes_func();

----------------------------------------------------------------

DROP TRIGGER stat_lecturer_academic_positions ON lecturers;
DROP FUNCTION "stat_lecturer_academic_positions_func"();
DROP TABLE "stat_lecturer_academic_positions";


-- +goose StatementEnd
