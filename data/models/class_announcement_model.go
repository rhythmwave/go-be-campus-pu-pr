package models

import (
	"database/sql"
	"time"
)

type GetClassAnnouncement struct {
	Id                 string     `db:"id"`
	ClassId            string     `db:"class_id"`
	Title              string     `db:"title"`
	Content            string     `db:"content"`
	LecturerId         string     `db:"lecturer_id"`
	LecturerName       string     `db:"lecturer_name"`
	LecturerFrontTitle *string    `db:"lecturer_front_title"`
	LecturerBackDegree *string    `db:"lecturer_back_degree"`
	FilePath           *string    `db:"file_path"`
	FilePathType       *string    `db:"file_path_type"`
	StartTime          *time.Time `db:"start_time"`
	EndTime            *time.Time `db:"end_time"`
}

type CreateClassAnnouncement struct {
	LecturerId   string         `db:"lecturer_id"`
	ClassId      string         `db:"class_id"`
	Title        string         `db:"title"`
	Content      string         `db:"content"`
	FilePath     sql.NullString `db:"file_path"`
	FilePathType sql.NullString `db:"file_path_type"`
	StartTime    sql.NullTime   `db:"start_time"`
	EndTime      sql.NullTime   `db:"end_time"`
}

type UpdateClassAnnouncement struct {
	Id           string         `db:"id"`
	Title        string         `db:"title"`
	Content      string         `db:"content"`
	FilePath     sql.NullString `db:"file_path"`
	FilePathType sql.NullString `db:"file_path_type"`
	StartTime    sql.NullTime   `db:"start_time"`
	EndTime      sql.NullTime   `db:"end_time"`
}
