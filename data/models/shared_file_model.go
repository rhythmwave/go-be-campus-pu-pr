package models

import (
	"database/sql"
	"time"
)

type GetSharedFile struct {
	Id                 string    `db:"id"`
	LecturerId         string    `db:"lecturer_id"`
	LecturerName       string    `db:"lecturer_name"`
	LecturerFrontTitle *string   `db:"lecturer_front_title"`
	LecturerBackDegree *string   `db:"lecturer_back_degree"`
	Title              string    `db:"title"`
	FilePath           string    `db:"file_path"`
	FilePathType       string    `db:"file_path_type"`
	Remarks            *string   `db:"remarks"`
	IsApproved         bool      `db:"is_approved"`
	CreatedAt          time.Time `db:"created_at"`
}

type CreateSharedFile struct {
	LecturerId   string         `db:"lecturer_id"`
	Title        string         `db:"title"`
	FilePath     string         `db:"file_path"`
	FilePathType string         `db:"file_path_type"`
	Remarks      sql.NullString `db:"remarks"`
}

type UpdateSharedFile struct {
	Id      string         `db:"id"`
	Title   string         `db:"title"`
	Remarks sql.NullString `db:"remarks"`
}
