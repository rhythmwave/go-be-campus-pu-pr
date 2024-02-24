-- +goose Up
-- +goose StatementBegin

CREATE TABLE "countries"
(
  "id" serial PRIMARY KEY,
  "name" varchar(100) NOT NULL UNIQUE
);

----------------------------------------------------------------------

CREATE TABLE "provinces"
(
  "id" serial PRIMARY KEY,
  "country_id" integer NOT NULL REFERENCES "countries" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "name" varchar(100) NOT NULL,
  UNIQUE("country_id", "name")
);

----------------------------------------------------------------------

CREATE TABLE "regencies"
(
  "id" serial PRIMARY KEY,
  "province_id" integer NOT NULL REFERENCES "provinces" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "name" varchar(100) NOT NULL,
  UNIQUE("province_id", "name")
);

----------------------------------------------------------------------

CREATE TABLE "districts"
(
  "id" serial PRIMARY KEY,
  "regency_id" integer NOT NULL REFERENCES "regencies" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "name" varchar(100) NOT NULL,
  UNIQUE("regency_id", "name")
);

----------------------------------------------------------------------

CREATE TABLE "villages"
(
  "id" bigint PRIMARY KEY,
  "district_id" integer NOT NULL REFERENCES "districts" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "name" varchar(100) NOT NULL,
  UNIQUE("district_id", "name")
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "villages";
DROP TABLE "districts";
DROP TABLE "regencies";
DROP TABLE "provinces";
DROP TABLE "countries";

-- +goose StatementEnd
