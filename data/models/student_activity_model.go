package models

import (
	"database/sql"
	"time"
)

type GetStudentActivity struct {
	Id                string `db:"id"`
	StudyProgramId    string `db:"study_program_id"`
	StudyProgramName  string `db:"study_program_name"`
	SemesterId        string `db:"semester_id"`
	SemesterStartYear uint32 `db:"semester_start_year"`
	SemesterType      string `db:"semester_type"`
	ActivityType      string `db:"activity_type"`
	Title             string `db:"title"`
}

type GetStudentActivityParticipant struct {
	StudentId        string  `db:"student_id"`
	NimNumber        int64   `db:"nim_number"`
	Name             string  `db:"name"`
	StudyProgramId   *string `db:"study_program_id"`
	StudyProgramName *string `db:"study_program_name"`
	Role             string  `db:"role"`
}

type GetStudentActivityLecturer struct {
	LecturerId         string  `db:"lecturer_id"`
	IdNationalLecturer string  `db:"id_national_lecturer"`
	Name               string  `db:"name"`
	FrontTitle         *string `db:"front_title"`
	BackDegree         *string `db:"back_degree"`
	ActivityCategory   string  `db:"activity_category"`
	Role               string  `db:"role"`
	Sort               uint32  `db:"sort"`
}

type GetStudentActivityDetail struct {
	Id                string     `db:"id"`
	StudyProgramId    string     `db:"study_program_id"`
	StudyProgramName  string     `db:"study_program_name"`
	SemesterId        string     `db:"semester_id"`
	SemesterStartYear uint32     `db:"semester_start_year"`
	SemesterType      string     `db:"semester_type"`
	ActivityType      string     `db:"activity_type"`
	Title             string     `db:"title"`
	Location          *string    `db:"location"`
	DecisionNumber    *string    `db:"decision_number"`
	DecisionDate      *time.Time `db:"decision_date"`
	IsGroupActivity   bool       `db:"is_group_activity"`
	Remarks           *string    `db:"remarks"`
}

type UpsertStudentActivityParticipant struct {
	StudentActivityId string `db:"student_activity_id"`
	StudentId         string `db:"student_id"`
	Role              string `db:"role"`
	CreatedBy         string `db:"created_by"`
}

type UpsertStudentActivityLecturer struct {
	StudentActivityId string `db:"student_activity_id"`
	LecturerId        string `db:"lecturer_id"`
	ActivityCategory  string `db:"activity_category"`
	Role              string `db:"role"`
	Sort              uint32 `db:"sort"`
	CreatedBy         string `db:"created_by"`
}

type CreateStudentActivity struct {
	StudyProgramId  string         `db:"study_program_id"`
	SemesterId      string         `db:"semester_id"`
	ActivityType    string         `db:"activity_type"`
	Title           string         `db:"title"`
	Location        sql.NullString `db:"location"`
	DecisionNumber  sql.NullString `db:"decision_number"`
	DecisionDate    sql.NullString `db:"decision_date"`
	IsGroupActivity bool           `db:"is_group_activity"`
	Remarks         sql.NullString `db:"remarks"`
	IsMbkm          bool           `db:"is_mbkm"`
	CreatedBy       string         `db:"created_by"`
}

type UpdateStudentActivity struct {
	Id              string         `db:"id"`
	StudyProgramId  string         `db:"study_program_id"`
	SemesterId      string         `db:"semester_id"`
	ActivityType    string         `db:"activity_type"`
	Title           string         `db:"title"`
	Location        sql.NullString `db:"location"`
	DecisionNumber  sql.NullString `db:"decision_number"`
	DecisionDate    sql.NullString `db:"decision_date"`
	IsGroupActivity bool           `db:"is_group_activity"`
	Remarks         sql.NullString `db:"remarks"`
	IsMbkm          bool           `db:"is_mbkm"`
	UpdatedBy       string         `db:"updated_by"`
}
