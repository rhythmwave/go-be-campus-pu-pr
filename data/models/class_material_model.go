package models

import "database/sql"

type GetClassMaterial struct {
	Id                 string  `db:"id"`
	Title              string  `db:"title"`
	Abstraction        *string `db:"abstraction"`
	FilePath           *string `db:"file_path"`
	FilePathType       *string `db:"file_path_type"`
	LecturerId         string  `db:"lecturer_id"`
	LecturerName       string  `db:"lecturer_name"`
	LecturerFrontTitle *string `db:"lecturer_front_title"`
	LecturerBackDegree *string `db:"lecturer_back_degree"`
	IsActive           bool    `db:"is_active"`
	CreatedAt          string  `db:"created_at"`
}

type CreateClassMaterial struct {
	LecturerId   string         `db:"lecturer_id"`
	ClassId      string         `db:"class_id"`
	Title        string         `db:"title"`
	Abstraction  sql.NullString `db:"abstraction"`
	FilePath     sql.NullString `db:"file_path"`
	FilePathType sql.NullString `db:"file_path_type"`
	IsActive     bool           `db:"is_active"`
}

type UpdateClassMaterial struct {
	Id           string         `db:"id"`
	Title        string         `db:"title"`
	Abstraction  sql.NullString `db:"abstraction"`
	FilePath     sql.NullString `db:"file_path"`
	FilePathType sql.NullString `db:"file_path_type"`
	IsActive     bool           `db:"is_active"`
}
