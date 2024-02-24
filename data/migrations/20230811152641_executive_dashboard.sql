-- +goose Up
-- +goose StatementBegin

CREATE TABLE "stat_total_students" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "study_program_id" uuid NOT NULL REFERENCES "study_programs"("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "student_force" integer NOT NULL,
  "total" integer NOT NULL DEFAULT 0,
  UNIQUE("study_program_id", "student_force")
);

WITH d AS (
  SELECT
    study_program_id,
    student_force,
    COUNT(1) AS total
  FROM students
  WHERE study_program_id IS NOT NULL AND student_force IS NOT NULL
  GROUP BY
    study_program_id,
    student_force
)
INSERT INTO stat_total_students (
  study_program_id,
  student_force,
  total
) SELECT 
  study_program_id,
  student_force,
  total
FROM d;

CREATE FUNCTION "stat_total_students_func"()
RETURNS TRIGGER AS $$
BEGIN
  IF (num_nulls(NEW.study_program_id, NEW.student_force) = 0) THEN
    INSERT INTO stat_total_students AS x (
      study_program_id,
      student_force,
      total
    ) VALUES ( 
      NEW.study_program_id,
      NEW.student_force,
      1
    ) ON CONFLICT (study_program_id, student_force) DO UPDATE SET total = x.total + 1;
  END IF;

  UPDATE stat_total_students SET total = total - 1
  WHERE study_program_id = OLD.study_program_id AND student_force = OLD.student_force;

  IF NEW.id IS NULL THEN
    RETURN OLD;
  ELSE
    RETURN NEW;
  END IF;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER stat_total_students AFTER INSERT OR UPDATE OR DELETE ON students FOR EACH ROW EXECUTE PROCEDURE stat_total_students_func();

------------------------------------------------------------------------------------

CREATE TABLE "stat_student_sexes" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "study_program_id" uuid NOT NULL REFERENCES "study_programs"("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "student_force" integer NOT NULL,
  "sex" character varying NOT NULL,
  "total" integer NOT NULL DEFAULT 0,
  UNIQUE("study_program_id", "student_force", "sex")
);

WITH d AS (
  SELECT
    study_program_id,
    student_force,
    sex,
    COUNT(1) AS total
  FROM students
  WHERE study_program_id IS NOT NULL AND student_force IS NOT NULL AND sex IS NOT NULL
  GROUP BY
    study_program_id,
    student_force,
    sex
)
INSERT INTO stat_student_sexes (
  study_program_id,
  student_force,
  sex,
  total
) SELECT 
  study_program_id,
  student_force,
  sex,
  total
FROM d;

CREATE FUNCTION "stat_student_sexes_func"()
RETURNS TRIGGER AS $$
BEGIN
  IF (num_nulls(NEW.study_program_id, NEW.student_force, NEW.sex) = 0) THEN
    INSERT INTO stat_student_sexes AS x (
      study_program_id,
      student_force,
      sex,
      total
    ) VALUES ( 
      NEW.study_program_id,
      NEW.student_force,
      NEW.sex,
      1
    ) ON CONFLICT (study_program_id, student_force, sex) DO UPDATE SET total = x.total + 1;
  END IF;
  
  UPDATE stat_student_sexes SET total = total - 1
  WHERE study_program_id = OLD.study_program_id AND student_force = OLD.student_force AND sex = OLD.sex;

  IF NEW.id IS NULL THEN
    RETURN OLD;
  ELSE
    RETURN NEW;
  END IF;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER stat_student_sexes AFTER INSERT OR UPDATE OR DELETE ON students FOR EACH ROW EXECUTE PROCEDURE stat_student_sexes_func();

------------------------------------------------------------------------------------

CREATE TABLE "stat_student_provinces" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "study_program_id" uuid NOT NULL REFERENCES "study_programs"("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "student_force" integer NOT NULL,
  "province_id" integer NOT NULL REFERENCES "provinces"("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "total" integer NOT NULL DEFAULT 0,
  UNIQUE("study_program_id", "student_force", "province_id")
);

WITH d AS (
  SELECT
    s.study_program_id,
    s.student_force,
    r.province_id,
    COUNT(1) AS total
  FROM students s
  JOIN villages v ON v.id = s.village_id
  JOIN districts d ON d.id = v.district_id
  JOIN regencies r ON r.id = d.regency_id
  WHERE s.study_program_id IS NOT NULL AND s.student_force IS NOT NULL AND s.village_id IS NOT NULL
  GROUP BY
    s.study_program_id,
    s.student_force,
    r.province_id
)
INSERT INTO stat_student_provinces (
  study_program_id,
  student_force,
  province_id,
  total
) SELECT 
  study_program_id,
  student_force,
  province_id,
  total
FROM d;

CREATE FUNCTION "stat_student_provinces_func"()
RETURNS TRIGGER AS $$
BEGIN
  IF (num_nulls(NEW.study_program_id, NEW.student_force, NEW.village_id) = 0) THEN
    INSERT INTO stat_student_provinces AS x (
      study_program_id,
      student_force,
      province_id,
      total
    ) 
    SELECT
      NEW.study_program_id,
      NEW.student_force,
      r.province_id,
      1
    FROM villages v
    JOIN districts d ON d.id = v.district_id
    JOIN regencies r ON r.id = d.regency_id
    WHERE v.id = NEW.village_id
    ON CONFLICT (study_program_id, student_force, province_id) DO UPDATE SET total = x.total + 1;
  END IF;

  UPDATE stat_student_provinces ssp SET total = ssp.total - 1
  FROM villages v
  JOIN districts d ON d.id = v.district_id
  JOIN regencies r ON r.id = d.regency_id
  WHERE ssp.study_program_id = OLD.study_program_id AND ssp.student_force = OLD.student_force AND v.id = OLD.village_id AND r.province_id = ssp.province_id;

  IF NEW.id IS NULL THEN
    RETURN OLD;
  ELSE
    RETURN NEW;
  END IF;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER stat_student_provinces AFTER INSERT OR UPDATE OR DELETE ON students FOR EACH ROW EXECUTE PROCEDURE stat_student_provinces_func();

------------------------------------------------------------------------------------

CREATE TABLE "stat_student_schools" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "study_program_id" uuid NOT NULL REFERENCES "study_programs"("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "student_force" integer NOT NULL,
  "is_vocational" boolean NOT NULL,
  "total" integer NOT NULL DEFAULT 0,
  UNIQUE("study_program_id", "student_force", "is_vocational")
);

WITH d AS (
  SELECT
    study_program_id,
    student_force,
    (CASE WHEN school_type = 'SMK' THEN true ELSE false END) AS is_vocational,
    COUNT(1) AS total
  FROM students
  WHERE study_program_id IS NOT NULL AND student_force IS NOT NULL AND school_type IS NOT NULL
  GROUP BY
    study_program_id,
    student_force,
    is_vocational
)
INSERT INTO stat_student_schools (
  study_program_id,
  student_force,
  is_vocational,
  total
) SELECT 
  study_program_id,
  student_force,
  is_vocational,
  total
FROM d;

CREATE FUNCTION "stat_student_schools_func"()
RETURNS TRIGGER AS $$
BEGIN
  IF (num_nulls(NEW.study_program_id, NEW.student_force, NEW.school_type) = 0) THEN
    INSERT INTO stat_student_schools AS x (
      study_program_id,
      student_force,
      is_vocational,
      total
    ) VALUES ( 
      NEW.study_program_id,
      NEW.student_force,
      (CASE WHEN NEW.school_type = 'SMK' THEN true ELSE false END),
      1
    ) ON CONFLICT (study_program_id, student_force, is_vocational) DO UPDATE SET total = x.total + 1;
  END IF;
  
  UPDATE stat_student_schools SET total = total - 1
  WHERE study_program_id = OLD.study_program_id AND student_force = OLD.student_force AND is_vocational = (CASE WHEN OLD.school_type = 'SMK' THEN true ELSE false END);

  IF NEW.id IS NULL THEN
    RETURN OLD;
  ELSE
    RETURN NEW;
  END IF;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER stat_student_schools AFTER INSERT OR UPDATE OR DELETE ON students FOR EACH ROW EXECUTE PROCEDURE stat_student_schools_func();

------------------------------------------------------------------------------------

CREATE TABLE "stat_student_graduations" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "study_program_id" uuid NOT NULL REFERENCES "study_programs"("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "student_force" integer NOT NULL,
  "is_graduated" boolean NOT NULL,
  "max_gpa" numeric(3,2) NOT NULL,
  "total" integer NOT NULL DEFAULT 0,
  UNIQUE("study_program_id", "student_force", "is_graduated")
);

WITH d AS (
  SELECT
    study_program_id,
    student_force,
    (CASE WHEN status = 'LULUS' THEN true ELSE false END) AS is_graduated,
    MAX(gpa) AS max_gpa,
    COUNT(1) AS total
  FROM students
  WHERE study_program_id IS NOT NULL AND student_force IS NOT NULL AND status IS NOT NULL AND gpa IS NOT NULL
  GROUP BY
    study_program_id,
    student_force,
    is_graduated
)
INSERT INTO stat_student_graduations (
  study_program_id,
  student_force,
  is_graduated,
  max_gpa,
  total
) SELECT 
  study_program_id,
  student_force,
  is_graduated,
  max_gpa,
  total
FROM d;

CREATE FUNCTION "stat_student_graduations_func"()
RETURNS TRIGGER AS $$
BEGIN
  IF (num_nulls(NEW.study_program_id, NEW.student_force, NEW.status, NEW.gpa) = 0) THEN
    INSERT INTO stat_student_graduations AS x (
      study_program_id,
      student_force,
      is_graduated,
      max_gpa,
      total
    ) VALUES ( 
      NEW.study_program_id,
      NEW.student_force,
      (CASE WHEN NEW.status = 'LULUS' THEN true ELSE false END),
      NEW.gpa,
      1
    ) ON CONFLICT (study_program_id, student_force, is_graduated) DO UPDATE SET 
      total = x.total + 1,
      max_gpa = (CASE WHEN x.max_gpa < EXCLUDED.max_gpa THEN EXCLUDED.max_gpa ELSE x.max_gpa END);
  END IF;
  
  UPDATE stat_student_graduations SET total = total - 1
  WHERE study_program_id = OLD.study_program_id AND student_force = OLD.student_force AND is_graduated = (CASE WHEN OLD.status = 'LULUS' THEN true ELSE false END);

  IF NEW.id IS NULL THEN
    RETURN OLD;
  ELSE
    RETURN NEW;
  END IF;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER stat_student_graduations AFTER INSERT OR UPDATE OR DELETE ON students FOR EACH ROW EXECUTE PROCEDURE stat_student_graduations_func();

------------------------------------------------------------------------------------

CREATE TABLE "stat_student_semester_grade_point" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "study_program_id" uuid NOT NULL REFERENCES "study_programs"("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "student_force" integer NOT NULL,
  "average_grade_point" numeric(3,2) NOT NULL,
  UNIQUE("study_program_id", "student_force")
);

WITH d AS (
  SELECT
    s.study_program_id,
    s.student_force,
    SUM(sp.grade_point) AS value_sum,
    COUNT(1) AS total_value
  FROM students s
  JOIN study_plans sp ON sp.student_id = s.id
  WHERE s.study_program_id IS NOT NULL AND s.student_force IS NOT NULL
  GROUP BY
    s.study_program_id,
    s.student_force
)
INSERT INTO stat_student_semester_grade_point (
  study_program_id,
  student_force,
  average_grade_point
) SELECT 
  study_program_id,
  student_force,
  value_sum/total_value
FROM d;

CREATE FUNCTION "stat_student_semester_grade_point_func"()
RETURNS TRIGGER AS $$
DECLARE 
  studyProgramId uuid;
  studentForce integer;
BEGIN
  SELECT study_program_id, student_force INTO studyProgramId, studentForce FROM students WHERE id = NEW.student_id;

  IF studyProgramId IS NOT NULL AND studentForce IS NOT NULL THEN
    WITH d AS (
      SELECT
        s.study_program_id,
        s.student_force,
        SUM(sp.grade_point) AS value_sum,
        COUNT(1) AS total_value
      FROM students s
      JOIN study_plans sp ON sp.student_id = s.id
      WHERE s.study_program_id = studyProgramId AND s.student_force = studentForce
      GROUP BY
        s.study_program_id,
        s.student_force
    )
    INSERT INTO stat_student_semester_grade_point (
      study_program_id,
      student_force,
      average_grade_point
    ) SELECT 
      study_program_id,
      student_force,
      value_sum/total_value
    FROM d ON CONFLICT (study_program_id, student_force) DO UPDATE SET
      average_grade_point = EXCLUDED.average_grade_point;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER stat_student_semester_grade_point AFTER UPDATE OF grade_point ON study_plans FOR EACH ROW EXECUTE PROCEDURE stat_student_semester_grade_point_func();

------------------------------------------------------------------------------------

CREATE TABLE "stat_student_graduation_predicates" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "study_program_id" uuid NOT NULL REFERENCES "study_programs"("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "student_force" integer NOT NULL,
  "graduation_predicate_id" uuid NOT NULL REFERENCES "graduation_predicates"("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "total" integer NOT NULL DEFAULT 0,
  UNIQUE("study_program_id", "student_force", "graduation_predicate_id")
);

WITH d AS (
  SELECT
    study_program_id,
    student_force,
    graduation_predicate_id,
    COUNT(1) AS total
  FROM students
  WHERE study_program_id IS NOT NULL AND student_force IS NOT NULL AND graduation_predicate_id IS NOT NULL
  GROUP BY
    study_program_id,
    student_force,
    graduation_predicate_id
)
INSERT INTO stat_student_graduation_predicates (
  study_program_id,
  student_force,
  graduation_predicate_id,
  total
) SELECT 
  study_program_id,
  student_force,
  graduation_predicate_id,
  total
FROM d;

CREATE FUNCTION "stat_student_graduation_predicates_func"()
RETURNS TRIGGER AS $$
BEGIN
  IF (num_nulls(NEW.study_program_id, NEW.student_force, NEW.graduation_predicate_id) = 0) THEN
    INSERT INTO stat_student_graduation_predicates AS x (
      study_program_id,
      student_force,
      graduation_predicate_id,
      total
    ) VALUES ( 
      NEW.study_program_id,
      NEW.student_force,
      NEW.graduation_predicate_id,
      1
    ) ON CONFLICT (study_program_id, student_force, graduation_predicate_id) DO UPDATE SET total = x.total + 1;
  END IF;
  
  UPDATE stat_student_graduation_predicates SET total = total - 1
  WHERE study_program_id = OLD.study_program_id AND student_force = OLD.student_force AND graduation_predicate_id = OLD.graduation_predicate_id;

  IF NEW.id IS NULL THEN
    RETURN OLD;
  ELSE
    RETURN NEW;
  END IF;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER stat_student_graduation_predicates AFTER INSERT OR UPDATE OR DELETE ON students FOR EACH ROW EXECUTE PROCEDURE stat_student_graduation_predicates_func();

------------------------------------------------------------------------------------

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

------------------------------------------------------------------------------------

CREATE TABLE "stat_lecturer_highest_degrees" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "highest_degree" character varying NOT NULL UNIQUE,
  "total" integer NOT NULL DEFAULT 0
);

WITH d AS (
  SELECT
    highest_degree,
    COUNT(1) AS total
  FROM lecturers
  WHERE highest_degree IS NOT NULL
  GROUP BY
    highest_degree
)
INSERT INTO stat_lecturer_highest_degrees (
  highest_degree,
  total
) SELECT 
  highest_degree,
  total
FROM d;

CREATE FUNCTION "stat_lecturer_highest_degrees_func"()
RETURNS TRIGGER AS $$
BEGIN
  IF NEW.highest_degree IS NOT NULL THEN
    INSERT INTO stat_lecturer_highest_degrees AS x (
      highest_degree,
      total
    ) VALUES ( 
      NEW.highest_degree,
      1
    ) ON CONFLICT (highest_degree) DO UPDATE SET total = x.total + 1;
  END IF;
  
  UPDATE stat_lecturer_highest_degrees SET total = total - 1
  WHERE highest_degree = OLD.highest_degree;

  IF NEW.id IS NULL THEN
    RETURN OLD;
  ELSE
    RETURN NEW;
  END IF;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER stat_lecturer_highest_degrees AFTER INSERT OR UPDATE OR DELETE ON lecturers FOR EACH ROW EXECUTE PROCEDURE stat_lecturer_highest_degrees_func();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER stat_total_students AFTER ON students;
DROP FUNCTION "stat_total_students_func"();
DROP TABLE "stat_total_students";

DROP TRIGGER stat_student_sexes AFTER ON students;
DROP FUNCTION "stat_student_sexes_func"();
DROP TABLE "stat_student_sexes";

DROP TRIGGER stat_student_provinces AFTER ON students;
DROP FUNCTION "stat_student_provinces_func"();
DROP TABLE "stat_student_provinces";

DROP TRIGGER stat_student_schools AFTER ON students;
DROP FUNCTION "stat_student_schools_func"();
DROP TABLE "stat_student_schools";

DROP TRIGGER stat_student_graduations AFTER ON students;
DROP FUNCTION "stat_student_graduations_func"();
DROP TABLE "stat_student_graduations";

DROP TRIGGER stat_student_semester_grade_point ON study_plans;
DROP FUNCTION "stat_student_semester_grade_point_func"();
DROP TABLE "stat_student_semester_grade_point";

DROP TRIGGER stat_student_graduation_predicates AFTER ON students;
DROP FUNCTION "stat_student_graduation_predicates_func"();
DROP TABLE "stat_student_graduation_predicates";

DROP TRIGGER stat_lecturer_sexes ON lecturers;
DROP FUNCTION "stat_lecturer_sexes_func"();
DROP TABLE "stat_lecturer_sexes";

-- +goose StatementEnd
