package models

import "time"

type GetAcademicGuidanceStudentList struct {
	Id                      string  `db:"id"`
	NimNumber               int64   `db:"nim_number"`
	StudentForce            *string `db:"student_force"`
	Name                    string  `db:"name"`
	Status                  *string `db:"status"`
	StudyPlanFormIsApproved *bool   `db:"study_plan_form_is_approved"`
}

type GetAcademicGuidanceDetail struct {
	Id                 string     `db:"id"`
	SemesterId         string     `db:"semester_id"`
	LecturerId         string     `db:"lecturer_id"`
	LecturerName       string     `db:"lecturer_name"`
	LecturerFrontTitle *string    `db:"lecturer_front_title"`
	LecturerBackDegree *string    `db:"lecturer_back_degree"`
	DecisionNumber     *string    `db:"decision_number"`
	DecisionDate       *time.Time `db:"decision_date"`
	TotalStudent       uint32     `db:"total_student"`
}

type UpsertAcademicGuidance struct {
	SemesterId string `db:"semester_id"`
	LecturerId string `db:"lecturer_id"`
	CreatedBy  string `db:"created_by"`
}

type UpsertAcademicGuidanceStudent struct {
	AcademicGuidanceId string `db:"academic_guidance_id"`
	StudentId          string `db:"student_id"`
	CreatedBy          string `db:"created_by"`
}

type UpsertDecisionAcademicGuidance struct {
	LecturerId     string    `db:"lecturer_id"`
	SemesterId     string    `db:"semester_id"`
	DecisionNumber string    `db:"decision_number"`
	DecisionDate   time.Time `db:"decision_date"`
	CreatedBy      string    `db:"created_by"`
}

type GetAcademicGuidanceSessionFile struct {
	Id                        string `db:"id"`
	AcademicGuidanceSessionId string `db:"academic_guidance_session_id"`
	Title                     string `db:"title"`
	FilePath                  string `db:"file_path"`
	FilePathType              string `db:"file_path_type"`
}

type GetAcademicGuidanceSessionStudent struct {
	Id                        string `db:"id"`
	AcademicGuidanceSessionId string `db:"academic_guidance_session_id"`
	Name                      string `db:"name"`
	NimNumber                 int64  `db:"nim_number"`
}

type GetAcademicGuidanceSession struct {
	Id                 string    `db:"id"`
	AcademicGuidanceId string    `db:"academic_guidance_id"`
	SemesterId         string    `db:"semester_id"`
	LecturerId         string    `db:"lecturer_id"`
	Subject            string    `db:"subject"`
	SessionDate        time.Time `db:"session_date"`
	Summary            string    `db:"summary"`
}

type UpsertAcademicGuidanceSessionFile struct {
	AcademicGuidanceSessionId string `db:"academic_guidance_session_id"`
	Title                     string `db:"title"`
	FilePath                  string `db:"file_path"`
	FilePathType              string `db:"file_path_type"`
}

type UpsertAcademicGuidanceSessionStudent struct {
	AcademicGuidanceSessionId string `db:"academic_guidance_session_id"`
	StudentId                 string `db:"student_id"`
}

type UpsertAcademicGuidanceSession struct {
	Id                 string    `db:"id"`
	AcademicGuidanceId string    `db:"academic_guidance_id"`
	Subject            string    `db:"subject"`
	SessionDate        time.Time `db:"session_date"`
	Summary            string    `db:"summary"`
}
