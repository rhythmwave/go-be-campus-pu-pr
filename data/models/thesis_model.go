package models

import (
	"database/sql"
	"time"
)

type GetListThesis struct {
	Id                        string  `db:"id"`
	Topic                     string  `db:"topic"`
	Title                     string  `db:"title"`
	Status                    string  `db:"status"`
	StudentId                 string  `db:"student_id"`
	StudentName               string  `db:"student_name"`
	StudentNimNumber          int64   `db:"student_nim_number"`
	StudentStatus             *string `db:"student_status"`
	StudyProgramId            string  `db:"study_program_id"`
	StudyProgramName          string  `db:"study_program_name"`
	DiktiStudyProgramCode     string  `db:"dikti_study_program_code"`
	DiktiStudyProgramType     string  `db:"dikti_study_program_type"`
	StudyLevelShortName       string  `db:"study_level_short_name"`
	StudentHasThesisStudyPlan bool    `db:"student_has_thesis_study_plan"`
	StartSemesterId           string  `db:"start_semester_id"`
	StartSemesterType         string  `db:"start_semester_type"`
	StartSemesterStartYear    uint32  `db:"start_semester_start_year"`
}

type GetThesisFile struct {
	Id              string  `db:"id"`
	ThesisId        string  `db:"thesis_id"`
	FilePath        string  `db:"file_path"`
	FilePathType    string  `db:"file_path_type"`
	FileDescription *string `db:"description"`
}

type GetThesisSupervisor struct {
	Id                       string  `db:"id"`
	ThesisId                 string  `db:"thesis_id"`
	LecturerId               string  `db:"lecturer_id"`
	LecturerName             string  `db:"lecturer_name"`
	LecturerFrontTitle       *string `db:"lecturer_front_title"`
	LecturerBackDegree       *string `db:"lecturer_back_degree"`
	ThesisSupervisorRoleId   string  `db:"thesis_supervisor_role_id"`
	ThesisSupervisorRoleName string  `db:"thesis_supervisor_role_name"`
	ThesisSupervisorRoleSort uint32  `db:"thesis_supervisor_role_sort"`
}

type GetDetailThesis struct {
	Id                        string     `db:"id"`
	StudyProgramId            *string    `db:"study_program_id"`
	StudentId                 string     `db:"student_id"`
	StudentName               string     `db:"student_name"`
	StudentNimNumber          int64      `db:"student_nim_number"`
	StartSemesterId           string     `db:"start_semester_id"`
	StartSemesterType         string     `db:"start_semester_type"`
	StartSemesterStartYear    uint32     `db:"start_semester_start_year"`
	FinishSemesterId          *string    `db:"finish_semester_id"`
	FinishSemesterType        *string    `db:"finish_semester_type"`
	FinishSemesterStartYear   *uint32    `db:"finish_semester_start_year"`
	Topic                     string     `db:"topic"`
	Title                     string     `db:"title"`
	EnglishTitle              *string    `db:"english_title"`
	StartDate                 time.Time  `db:"start_date"`
	FinishDate                *time.Time `db:"finish_date"`
	Remarks                   *string    `db:"remarks"`
	IsJointThesis             bool       `db:"is_joint_thesis"`
	Status                    string     `db:"status"`
	ProposalSeminarDate       *time.Time `db:"proposal_seminar_date"`
	ProposalCertificateNumber *string    `db:"proposal_certificate_number"`
	ProposalCertificateDate   *time.Time `db:"proposal_certificate_date"`
	ThesisDefenseCount        uint32     `db:"thesis_defense_count"`
	GradePoint                float64    `db:"grade_point"`
	GradeCode                 *string    `db:"grade_code"`
}

type CreateThesis struct {
	Id                        string         `db:"id"`
	StudentId                 string         `db:"student_id"`
	Topic                     string         `db:"topic"`
	Status                    string         `db:"status"`
	Title                     string         `db:"title"`
	EnglishTitle              sql.NullString `db:"english_title"`
	StartSemesterId           string         `db:"start_semester_id"`
	StartDate                 time.Time      `db:"start_date"`
	Remarks                   sql.NullString `db:"remarks"`
	IsJointThesis             bool           `db:"is_joint_thesis"`
	ProposalSeminarDate       sql.NullTime   `db:"proposal_seminar_date"`
	ProposalCertificateNumber sql.NullString `db:"proposal_certificate_number"`
	ProposalCertificateDate   sql.NullTime   `db:"proposal_certificate_date"`
}

type UpsertThesisFile struct {
	ThesisId     string         `db:"thesis_id"`
	FilePath     string         `db:"file_path"`
	FilePathType string         `db:"file_path_type"`
	Description  sql.NullString `db:"description"`
}

type UpsertThesisSupervisor struct {
	ThesisId               string `db:"thesis_id"`
	LecturerId             string `db:"lecturer_id"`
	ThesisSupervisorRoleId string `db:"thesis_supervisor_role_id"`
}

type UpdateThesis struct {
	Id                        string         `db:"id"`
	StudentId                 string         `db:"student_id"`
	Topic                     string         `db:"topic"`
	Status                    string         `db:"status"`
	Title                     string         `db:"title"`
	EnglishTitle              sql.NullString `db:"english_title"`
	StartSemesterId           string         `db:"start_semester_id"`
	StartDate                 time.Time      `db:"start_date"`
	Remarks                   sql.NullString `db:"remarks"`
	IsJointThesis             bool           `db:"is_joint_thesis"`
	ProposalSeminarDate       sql.NullTime   `db:"proposal_seminar_date"`
	ProposalCertificateNumber sql.NullString `db:"proposal_certificate_number"`
	ProposalCertificateDate   sql.NullTime   `db:"proposal_certificate_date"`
}

type GetThesisDefenseExaminer struct {
	Id                     string  `db:"id"`
	ThesisDefenseId        string  `db:"thesis_defense_id"`
	LecturerId             string  `db:"lecturer_id"`
	LecturerName           string  `db:"lecturer_name"`
	LecturerFrontTitle     *string `db:"lecturer_front_title"`
	LecturerBackDegree     *string `db:"lecturer_back_degree"`
	ThesisExaminerRoleId   string  `db:"thesis_examiner_role_id"`
	ThesisExaminerRoleName string  `db:"thesis_examiner_role_name"`
}

type GetThesisDefenseRequest struct {
	Id                           string     `db:"id"`
	StudentId                    string     `db:"student_id"`
	StudentName                  string     `db:"student_name"`
	StudentNimNumber             int64      `db:"student_nim_number"`
	StudentStatus                string     `db:"student_status"`
	StudyProgramId               string     `db:"study_program_id"`
	StudyProgramName             string     `db:"study_program_name"`
	DiktiStudyProgramCode        string     `db:"dikti_study_program_code"`
	DiktiStudyProgramType        string     `db:"dikti_study_program_type"`
	StudyLevelId                 string     `db:"study_level_id"`
	StudyLevelShortName          string     `db:"study_level_short_name"`
	ThesisId                     string     `db:"thesis_id"`
	ThesisTitle                  string     `db:"thesis_title"`
	ThesisStatus                 string     `db:"thesis_status"`
	ThesisDefenseCount           uint32     `db:"thesis_defense_count"`
	ThesisDefenseId              *string    `db:"thesis_defense_id"`
	ThesisDefensePlanDate        *time.Time `db:"thesis_defense_plan_date"`
	ThesisDefensePlanStartTime   *uint32    `db:"thesis_defense_plan_start_time"`
	ThesisDefensePlanEndTime     *uint32    `db:"thesis_defense_plan_end_time"`
	ThesisDefenseActualDate      *time.Time `db:"thesis_defense_actual_date"`
	ThesisDefenseActualStartTime *uint32    `db:"thesis_defense_actual_start_time"`
	ThesisDefenseActualEndTime   *uint32    `db:"thesis_defense_actual_end_time"`
	ThesisDefenseRoomId          *string    `db:"thesis_defense_room_id"`
	ThesisDefenseRoomName        *string    `db:"thesis_defense_room_name"`
	ThesisDefenseIsPassed        *bool      `db:"thesis_defense_is_passed"`
	ThesisDefenseRevision        *string    `db:"thesis_defense_revision"`
	ThesisGradeCode              *string    `db:"thesis_grade_code"`
	CreatedAt                    time.Time  `db:"created_at"`
}

type GetThesisDefense struct {
	Id                    string     `db:"id"`
	PlanDate              time.Time  `db:"plan_date"`
	PlanStartTime         uint32     `db:"plan_start_time"`
	PlanEndTime           uint32     `db:"plan_end_time"`
	ActualDate            *time.Time `db:"actual_date"`
	ActualStartTime       *uint32    `db:"actual_start_time"`
	ActualEndTime         *uint32    `db:"actual_end_time"`
	RoomId                string     `db:"room_id"`
	RoomName              string     `db:"room_name"`
	IsPassed              bool       `db:"is_passed"`
	StudentId             string     `db:"student_id"`
	StudentName           string     `db:"student_name"`
	StudentNimNumber      int64      `db:"student_nim_number"`
	StudentStatus         string     `db:"student_status"`
	StudyProgramId        string     `db:"study_program_id"`
	StudyProgramName      string     `db:"study_program_name"`
	DiktiStudyProgramCode string     `db:"dikti_study_program_code"`
	DiktiStudyProgramType string     `db:"dikti_study_program_type"`
	StudyLevelId          string     `db:"study_level_id"`
	StudyLevelShortName   string     `db:"study_level_short_name"`
	ThesisId              string     `db:"thesis_id"`
	ThesisTitle           string     `db:"thesis_title"`
	ThesisStatus          string     `db:"thesis_status"`
	ThesisDefenseCount    uint32     `db:"thesis_defense_count"`
	CreatedAt             time.Time  `db:"created_at"`
}

type CreateThesisDefense struct {
	Id            string    `db:"id"`
	ThesisId      string    `db:"thesis_id"`
	PlanDate      time.Time `db:"plan_date"`
	PlanStartTime uint32    `db:"plan_start_time"`
	PlanEndTime   uint32    `db:"plan_end_time"`
	RoomId        string    `db:"room_id"`
}

type UpsertThesisDefenseExaminer struct {
	ThesisDefenseId      string `db:"thesis_defense_id"`
	LecturerId           string `db:"lecturer_id"`
	ThesisExaminerRoleId string `db:"thesis_examiner_role_id"`
}

type UpdateThesisDefense struct {
	Id              string         `db:"id"`
	PlanDate        time.Time      `db:"plan_date"`
	PlanStartTime   uint32         `db:"plan_start_time"`
	PlanEndTime     uint32         `db:"plan_end_time"`
	RoomId          string         `db:"room_id"`
	ActualDate      sql.NullTime   `db:"actual_date"`
	ActualStartTime sql.NullInt32  `db:"actual_start_time"`
	ActualEndTime   sql.NullInt32  `db:"actual_end_time"`
	IsPassed        bool           `db:"is_passed"`
	Revision        sql.NullString `db:"revision"`
}

type FinishThesisDefense struct {
	Id               string  `db:"id"`
	FinishSemesterId string  `db:"finish_semester_id"`
	GradePoint       float64 `db:"grade_point"`
	GradeCode        string  `db:"grade_code"`
	Status           string  `db:"status"`
}

type GetThesisSupervisorLog struct {
	LecturerId               string `db:"lecturer_id"`
	ThesisSupervisorRoleId   string `db:"thesis_supervisor_role_id"`
	ThesisSupervisorRoleName string `db:"thesis_supervisor_role_name"`
	Total                    uint32 `db:"total"`
}
