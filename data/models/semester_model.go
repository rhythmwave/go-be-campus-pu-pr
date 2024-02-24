package models

import (
	"database/sql"
	"time"
)

type GetSemester struct {
	Id                string     `db:"id"`
	SemesterStartYear uint32     `db:"semester_start_year"`
	SemesterType      string     `db:"semester_type"`
	IsActive          bool       `db:"is_active"`
	StartDate         time.Time  `db:"start_date"`
	EndDate           time.Time  `db:"end_date"`
	MidtermStartDate  *time.Time `db:"midterm_start_date"`
	MidtermEndDate    *time.Time `db:"midterm_end_date"`
	EndtermStartDate  *time.Time `db:"endterm_start_date"`
	EndtermEndDate    *time.Time `db:"endterm_end_date"`
}

type GetSemesterDetail struct {
	Id                         string     `db:"id"`
	SemesterStartYear          uint32     `db:"semester_start_year"`
	SemesterType               string     `db:"semester_type"`
	IsActive                   bool       `db:"is_active"`
	StartDate                  time.Time  `db:"start_date"`
	EndDate                    time.Time  `db:"end_date"`
	MidtermStartDate           *time.Time `db:"midterm_start_date"`
	MidtermEndDate             *time.Time `db:"midterm_end_date"`
	EndtermStartDate           *time.Time `db:"endterm_start_date"`
	EndtermEndDate             *time.Time `db:"endterm_end_date"`
	StudyPlanInputStartDate    time.Time  `db:"study_plan_input_start_date"`
	StudyPlanInputEndDate      time.Time  `db:"study_plan_input_end_date"`
	StudyPlanApprovalStartDate time.Time  `db:"study_plan_approval_start_date"`
	StudyPlanApprovalEndDate   time.Time  `db:"study_plan_approval_end_date"`
	GradingStartDate           *time.Time `db:"grading_start_date"`
	GradingEndDate             *time.Time `db:"grading_end_date"`
	ReferenceSemesterId        *string    `db:"reference_semester_id"`
	ReferenceSemesterStartYear *uint32    `db:"reference_semester_start_year"`
	ReferenceSemesterType      *string    `db:"reference_semester_type"`
	CheckMinimumGpa            bool       `db:"check_minimum_gpa"`
	CheckPassedCredit          bool       `db:"check_passed_credit"`
	DefaultCredit              uint32     `db:"default_credit"`
}

type GetSemesterCurriculum struct {
	SemesterId       string `db:"semester_id"`
	StudyProgramId   string `db:"study_program_id"`
	StudyProgramName string `db:"study_program_name"`
	CurriculumId     string `db:"curriculum_id"`
	CurriculumName   string `db:"curriculum_name"`
}

type CreateSemester struct {
	SemesterStartYear          uint32         `db:"semester_start_year"`
	SemesterType               string         `db:"semester_type"`
	StartDate                  time.Time      `db:"start_date"`
	EndDate                    time.Time      `db:"end_date"`
	MidtermStartDate           sql.NullTime   `db:"midterm_start_date"`
	MidtermEndDate             sql.NullTime   `db:"midterm_end_date"`
	EndtermStartDate           sql.NullTime   `db:"endterm_start_date"`
	EndtermEndDate             sql.NullTime   `db:"endterm_end_date"`
	StudyPlanInputStartDate    time.Time      `db:"study_plan_input_start_date"`
	StudyPlanInputEndDate      time.Time      `db:"study_plan_input_end_date"`
	StudyPlanApprovalStartDate time.Time      `db:"study_plan_approval_start_date"`
	StudyPlanApprovalEndDate   time.Time      `db:"study_plan_approval_end_date"`
	GradingStartDate           sql.NullTime   `db:"grading_start_date"`
	GradingEndDate             sql.NullTime   `db:"grading_end_date"`
	ReferenceSemesterId        sql.NullString `db:"reference_semester_id"`
	CheckMinimumGpa            bool           `db:"check_minimum_gpa"`
	CheckPassedCredit          bool           `db:"check_passed_credit"`
	DefaultCredit              uint32         `db:"default_credit"`
	CreatedBy                  string         `db:"created_by"`
}

type UpdateSemester struct {
	Id                         string         `db:"id"`
	SemesterStartYear          uint32         `db:"semester_start_year"`
	SemesterType               string         `db:"semester_type"`
	StartDate                  time.Time      `db:"start_date"`
	EndDate                    time.Time      `db:"end_date"`
	MidtermStartDate           sql.NullTime   `db:"midterm_start_date"`
	MidtermEndDate             sql.NullTime   `db:"midterm_end_date"`
	EndtermStartDate           sql.NullTime   `db:"endterm_start_date"`
	EndtermEndDate             sql.NullTime   `db:"endterm_end_date"`
	StudyPlanInputStartDate    time.Time      `db:"study_plan_input_start_date"`
	StudyPlanInputEndDate      time.Time      `db:"study_plan_input_end_date"`
	StudyPlanApprovalStartDate time.Time      `db:"study_plan_approval_start_date"`
	StudyPlanApprovalEndDate   time.Time      `db:"study_plan_approval_end_date"`
	GradingStartDate           sql.NullTime   `db:"grading_start_date"`
	GradingEndDate             sql.NullTime   `db:"grading_end_date"`
	ReferenceSemesterId        sql.NullString `db:"reference_semester_id"`
	CheckMinimumGpa            bool           `db:"check_minimum_gpa"`
	CheckPassedCredit          bool           `db:"check_passed_credit"`
	DefaultCredit              uint32         `db:"default_credit"`
	UpdatedBy                  string         `db:"updated_by"`
}

type UpsertSemesterCurriculum struct {
	SemesterId   string `db:"semester_id"`
	CurriculumId string `db:"curriculum_id"`
	CreatedBy    string `db:"created_by"`
}
