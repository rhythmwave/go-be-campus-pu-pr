package models

import (
	"database/sql"
	"time"
)

type GetClassEvent struct {
	Id                 string    `db:"id"`
	Title              string    `db:"title"`
	Frequency          string    `db:"frequency"`
	EventTime          time.Time `db:"event_time"`
	LecturerId         string    `db:"lecturer_id"`
	LecturerName       string    `db:"lecturer_name"`
	LecturerFrontTitle *string   `db:"lecturer_front_title"`
	LecturerBackDegree *string   `db:"lecturer_back_degree"`
	Remarks            *string   `db:"remarks"`
	IsActive           bool      `db:"is_active"`
	CreatedAt          string    `db:"created_at"`
}

type CreateClassEvent struct {
	LecturerId string         `db:"lecturer_id"`
	ClassId    string         `db:"class_id"`
	Title      string         `db:"title"`
	Frequency  string         `db:"frequency"`
	EventTime  time.Time      `db:"event_time"`
	Remarks    sql.NullString `db:"remarks"`
	IsActive   bool           `db:"is_active"`
}

type UpdateClassEvent struct {
	Id        string         `db:"id"`
	Title     string         `db:"title"`
	Frequency string         `db:"frequency"`
	EventTime time.Time      `db:"event_time"`
	Remarks   sql.NullString `db:"remarks"`
	IsActive  bool           `db:"is_active"`
}
