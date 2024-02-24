package models

import (
	"database/sql"
	"time"
)

type GetLecturerLeave struct {
	Id                    string     `db:"id"`
	Name                  string     `db:"name"`
	IdNationalLecturer    string     `db:"id_national_lecturer"`
	FrontTitle            *string    `db:"front_title"`
	BackDegree            *string    `db:"back_degree"`
	SemesterStartYear     uint32     `db:"semester_start_year"`
	SemesterType          string     `db:"semester_type"`
	DiktiStudyProgramCode *string    `db:"dikti_study_program_code"`
	StudyProgramName      *string    `db:"study_program_name"`
	StudyLevelShortName   *string    `db:"study_level_short_name"`
	DiktiStudyProgramType *string    `db:"dikti_study_program_type"`
	StartDate             time.Time  `db:"start_date"`
	EndDate               *time.Time `db:"end_date"`
	PermitNumber          string     `db:"permit_number"`
	Purpose               string     `db:"purpose"`
	Remarks               string     `db:"remarks"`
	FilePath              *string    `db:"file_path"`
	FilePathType          *string    `db:"file_path_type"`
	IsActive              bool       `db:"is_active"`
}

type CreateLecturerLeave struct {
	LecturerId   string         `db:"lecturer_id"`
	SemesterId   string         `db:"semester_id"`
	StartDate    time.Time      `db:"start_date"`
	PermitNumber string         `db:"permit_number"`
	Purpose      string         `db:"purpose"`
	Remarks      string         `db:"remarks"`
	FilePath     sql.NullString `db:"file_path"`
	FilePathType sql.NullString `db:"file_path_type"`
	CreatedBy    string         `db:"created_by"`
}

type UpdateLecturerLeave struct {
	Id           string         `db:"id"`
	PermitNumber string         `db:"permit_number"`
	Purpose      string         `db:"purpose"`
	Remarks      string         `db:"remarks"`
	FilePath     sql.NullString `db:"file_path"`
	FilePathType sql.NullString `db:"file_path_type"`
	UpdatedBy    string         `db:"updated_by"`
}
