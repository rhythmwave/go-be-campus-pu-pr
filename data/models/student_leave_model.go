package models

import (
	"database/sql"
	"time"
)

type GetStudentLeaveRequest struct {
	Id                         string    `db:"id"`
	NimNumber                  int64     `db:"nim_number"`
	Name                       string    `db:"name"`
	DiktiStudyProgramCode      string    `db:"dikti_study_program_code"`
	StudyProgramName           string    `db:"study_program_name"`
	StudyLevelShortName        string    `db:"study_level_short_name"`
	DiktiStudyProgramType      string    `db:"dikti_study_program_type"`
	StartDate                  time.Time `db:"start_date"`
	TotalLeaveDurationSemester uint32    `db:"total_leave_duration_semester"`
	PermitNumber               *string   `db:"permit_number"`
	Purpose                    string    `db:"purpose"`
	Remarks                    string    `db:"remarks"`
	IsApproved                 *bool     `db:"is_approved"`
	SemesterType               string    `db:"semester_type"`
	SemesterStartYear          uint32    `db:"semester_start_year"`
}

type GetStudentLeave struct {
	Id                    string  `db:"id"`
	NimNumber             int64   `db:"nim_number"`
	Name                  string  `db:"name"`
	DiktiStudyProgramCode string  `db:"dikti_study_program_code"`
	StudyProgramName      string  `db:"study_program_name"`
	StudyLevelShortName   string  `db:"study_level_short_name"`
	DiktiStudyProgramType string  `db:"dikti_study_program_type"`
	SemesterStartYear     uint32  `db:"semester_start_year"`
	SemesterType          string  `db:"semester_type"`
	PermitNumber          *string `db:"permit_number"`
	Purpose               string  `db:"purpose"`
	Remarks               string  `db:"remarks"`
}

type CreateStudentLeave struct {
	StudentId                  string         `db:"student_id"`
	TotalLeaveDurationSemester uint32         `db:"total_leave_duration_semester"`
	StartDate                  time.Time      `db:"start_date"`
	PermitNumber               sql.NullString `db:"permit_number"`
	Purpose                    string         `db:"purpose"`
	Remarks                    string         `db:"remarks"`
}

type UpdateStudentLeave struct {
	Id           string         `db:"id"`
	PermitNumber sql.NullString `db:"permit_number"`
	Purpose      string         `db:"purpose"`
	Remarks      string         `db:"remarks"`
}
