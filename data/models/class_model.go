package models

import (
	"database/sql"
	"time"
)

type GetClassLecturer struct {
	Id                   string  `db:"id"`
	ClassId              string  `db:"class_id"`
	LecturerId           string  `db:"lecturer_id"`
	LecturerName         string  `db:"lecturer_name"`
	LecturerFrontTitle   *string `db:"lecturer_front_title"`
	LecturerBackDegree   *string `db:"lecturer_back_degree"`
	IsGradingResponsible bool    `db:"is_grading_responsible"`
}

type GetClass struct {
	Id                          string     `db:"id"`
	Name                        string     `db:"name"`
	StudyProgramId              string     `db:"study_program_id"`
	CurriculumId                string     `db:"curriculum_id"`
	CurriculumName              string     `db:"curriculum_name"`
	SemesterId                  string     `db:"semester_id"`
	SubjectId                   string     `db:"subject_id"`
	SubjectCode                 string     `db:"subject_code"`
	SubjectName                 string     `db:"subject_name"`
	SubjectIsMandatory          bool       `db:"subject_is_mandatory"`
	SubjectSemesterPackage      uint32     `db:"subject_semester_package"`
	SubjectTotalLessonPlan      uint32     `db:"subject_total_lesson_plan"`
	MaximumParticipant          *uint32    `db:"maximum_participant"`
	TotalParticipant            uint32     `db:"total_participant"`
	SubjectTheoryCredit         uint32     `db:"subject_theory_credit"`
	SubjectPracticumCredit      uint32     `db:"subject_practicum_credit"`
	SubjectFieldPracticumCredit uint32     `db:"subject_field_practicum_credit"`
	UnapprovedStudyPlan         uint32     `db:"unapproved_study_plan"`
	TotalMaterial               uint32     `db:"total_material"`
	TotalWork                   uint32     `db:"total_work"`
	TotalDiscussion             uint32     `db:"total_discussion"`
	TotalExam                   uint32     `db:"total_exam"`
	TotalEvent                  uint32     `db:"total_event"`
	TotalLecturePlan            uint32     `db:"total_lecture_plan"`
	TotalLectureDone            uint32     `db:"total_lecture_done"`
	TotalGradedParticipant      uint32     `db:"total_graded_participant"`
	IsActive                    bool       `db:"is_active"`
	StudyLevelId                string     `db:"study_level_id"`
	ApplicationDeadline         *time.Time `db:"application_deadline"`
	StudyProgramName            string     `db:"study_program_name"`
	SemesterStartYear           uint32     `db:"semester_start_year"`
	SemesterType                string     `db:"semester_type"`
}

type GetClassDetail struct {
	Id                              string     `db:"id"`
	Name                            string     `db:"name"`
	StudyProgramId                  string     `db:"study_program_id"`
	StudyProgramName                string     `db:"study_program_name"`
	DiktiStudyProgramType           string     `db:"dikti_study_program_type"`
	StudyLevelId                    string     `db:"study_level_id"`
	StudyLevelShortName             string     `db:"study_level_short_name"`
	CurriculumId                    string     `db:"curriculum_id"`
	CurriculumName                  string     `db:"curriculum_name"`
	CurriculumYear                  string     `db:"curriculum_year"`
	SemesterId                      string     `db:"semester_id"`
	SemesterStartYear               uint32     `db:"semester_start_year"`
	SemesterType                    string     `db:"semester_type"`
	SemesterIsActive                bool       `db:"semester_is_active"`
	GradingStartDate                *time.Time `db:"grading_start_date"`
	GradingEndDate                  *time.Time `db:"grading_end_date"`
	SubjectId                       string     `db:"subject_id"`
	SubjectCode                     string     `db:"subject_code"`
	SubjectName                     string     `db:"subject_name"`
	SubjectTheoryCredit             uint32     `db:"subject_theory_credit"`
	SubjectPracticumCredit          uint32     `db:"subject_practicum_credit"`
	SubjectFieldPracticumCredit     uint32     `db:"subject_field_practicum_credit"`
	Scope                           *string    `db:"scope"`
	IsOnline                        *bool      `db:"is_online"`
	IsOffline                       *bool      `db:"is_offline"`
	MinimumParticipant              *uint32    `db:"minimum_participant"`
	MaximumParticipant              *uint32    `db:"maximum_participant"`
	TotalParticipant                uint32     `db:"total_participant"`
	Remarks                         *string    `db:"remarks"`
	IsActive                        bool       `db:"is_active"`
	SemesterPackage                 uint32     `db:"semester_package"`
	TheoryCredit                    uint32     `db:"theory_credit"`
	PracticumCredit                 uint32     `db:"practicum_credit"`
	FieldPracticumCredit            uint32     `db:"field_practicum_credit"`
	SubjectMinimumPassingGradePoint float64    `db:"subject_minimum_passing_grade_point"`
	SubjectIsMandatory              bool       `db:"subject_is_mandatory"`
	TotalMaterial                   uint32     `db:"total_material"`
	TotalWork                       uint32     `db:"total_work"`
	TotalDiscussion                 uint32     `db:"total_discussion"`
	TotalExam                       uint32     `db:"total_exam"`
	TotalEvent                      uint32     `db:"total_event"`
	TotalLecturePlan                uint32     `db:"total_lecture_plan"`
	TotalLectureDone                uint32     `db:"total_lecture_done"`
	TotalGradedParticipant          uint32     `db:"total_graded_participant"`
	ApplicationDeadline             *time.Time `db:"application_deadline"`
	IsGradingResponsible            *bool      `db:"is_grading_responsible"`
	SubjectIsMbkm                   bool       `db:"subject_is_mbkm"`
}

type CreateClass struct {
	Id                  string         `db:"id"`
	SubjectId           string         `db:"subject_id"`
	SemesterId          string         `db:"semester_id"`
	Name                string         `db:"name"`
	Scope               sql.NullString `db:"scope"`
	IsOnline            sql.NullBool   `db:"is_online"`
	IsOffline           sql.NullBool   `db:"is_offline"`
	MinimumParticipant  sql.NullInt32  `db:"minimum_participant"`
	MaximumParticipant  sql.NullInt32  `db:"maximum_participant"`
	Remarks             sql.NullString `db:"remarks"`
	ApplicationDeadline sql.NullTime   `db:"application_deadline"`
	CreatedBy           string         `db:"created_by"`
}

type UpdateClass struct {
	Id                  string         `db:"id"`
	SubjectId           string         `db:"subject_id"`
	Name                string         `db:"name"`
	Scope               string         `db:"scope"`
	IsOnline            bool           `db:"is_online"`
	IsOffline           bool           `db:"is_offline"`
	MinimumParticipant  uint32         `db:"minimum_participant"`
	MaximumParticipant  sql.NullInt32  `db:"maximum_participant"`
	Remarks             sql.NullString `db:"remarks"`
	ApplicationDeadline sql.NullTime   `db:"application_deadline"`
	UpdatedBy           string         `db:"updated_by"`
}

type UpsertClassLecturer struct {
	ClassId              string `db:"class_id"`
	LecturerId           string `db:"lecturer_id"`
	IsGradingResponsible bool   `db:"is_grading_responsible"`
	CreatedBy            string `db:"created_by"`
}
