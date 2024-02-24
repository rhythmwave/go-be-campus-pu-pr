package models

import (
	"database/sql"
	"time"
)

type GetClassWork struct {
	Id                     string    `db:"id"`
	ClassId                string    `db:"class_id"`
	Title                  string    `db:"title"`
	Abstraction            *string   `db:"abstraction"`
	FilePath               *string   `db:"file_path"`
	FilePathType           *string   `db:"file_path_type"`
	LecturerId             string    `db:"lecturer_id"`
	LecturerName           string    `db:"lecturer_name"`
	LecturerFrontTitle     *string   `db:"lecturer_front_title"`
	LecturerBackDegree     *string   `db:"lecturer_back_degree"`
	StartTime              time.Time `db:"start_time"`
	EndTime                time.Time `db:"end_time"`
	TotalSubmission        uint32    `db:"total_submission"`
	SubmissionFilePath     *string   `db:"submission_file_path"`
	SubmissionFilePathType *string   `db:"submission_file_path_type"`
	SubmissionPoint        *float64  `db:"submission_point"`
}

type GetClassWorkSubmission struct {
	Id               *string  `db:"id"`
	StudentId        string   `db:"student_id"`
	NimNumber        int64    `db:"nim_number"`
	Name             string   `db:"name"`
	StudyProgramName string   `db:"study_program_name"`
	FilePath         *string  `db:"file_path"`
	FilePathType     *string  `db:"file_path_type"`
	Point            *float64 `db:"point"`
}

type CreateClassWork struct {
	LecturerId   string         `db:"lecturer_id"`
	ClassId      string         `db:"class_id"`
	Title        string         `db:"title"`
	Abstraction  sql.NullString `db:"abstraction"`
	FilePath     sql.NullString `db:"file_path"`
	FilePathType sql.NullString `db:"file_path_type"`
	StartTime    time.Time      `db:"start_time"`
	EndTime      time.Time      `db:"end_time"`
}

type UpdateClassWork struct {
	Id           string         `db:"id"`
	Title        string         `db:"title"`
	Abstraction  sql.NullString `db:"abstraction"`
	FilePath     sql.NullString `db:"file_path"`
	FilePathType sql.NullString `db:"file_path_type"`
	StartTime    time.Time      `db:"start_time"`
	EndTime      time.Time      `db:"end_time"`
}

type GradeClassWorkSubmission struct {
	ClassWorkId  string  `db:"class_work_id"`
	StudentId    string  `db:"student_id"`
	FilePath     string  `db:"file_path"`
	FilePathType string  `db:"file_path_type"`
	Point        float64 `db:"point"`
}
