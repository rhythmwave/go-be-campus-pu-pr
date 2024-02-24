-- +goose Up
-- +goose StatementBegin

CREATE OR REPLACE FUNCTION semester_conflict_date_func()
RETURNS TRIGGER AS $$
DECLARE conflictId uuid;
BEGIN
  SELECT s.id INTO conflictId
  FROM semesters s WHERE NEW.start_date <= s.end_date AND NEW.end_date >= s.start_date AND s.id != NEW.id;

  IF conflictId IS NOT NULL THEN
    RAISE EXCEPTION 'conflicting semester date is exist.';
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

CREATE OR REPLACE FUNCTION semester_conflict_date_func()
RETURNS TRIGGER AS $$
DECLARE conflictId uuid;
BEGIN
  SELECT s.id INTO conflictId
  FROM semesters s WHERE NEW.start_date <= s.end_date AND NEW.end_date >= s.start_date;

  IF conflictId IS NOT NULL THEN
    RAISE EXCEPTION 'conflicting semester date is exist.';
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- +goose StatementEnd
