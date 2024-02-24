package models

import (
	"database/sql"
	"time"
)

type GetStudentClass struct {
	Id                          string  `db:"id"`
	ClassId                     string  `db:"class_id"`
	ClassName                   string  `db:"class_name"`
	SubjectId                   string  `db:"subject_id"`
	SubjectCode                 string  `db:"subject_code"`
	SubjectName                 string  `db:"subject_name"`
	SemesterId                  string  `db:"semester_id"`
	SemesterStartYear           uint32  `db:"semester_start_year"`
	SemesterType                string  `db:"semester_type"`
	SubjectTheoryCredit         uint32  `db:"subject_theory_credit"`
	SubjectPracticumCredit      uint32  `db:"subject_practicum_credit"`
	SubjectFieldPracticumCredit uint32  `db:"subject_field_practicum_credit"`
	SubjectRepetition           uint32  `db:"subject_repetition"`
	StudyPlanId                 string  `db:"study_plan_id"`
	StudentId                   string  `db:"student_id"`
	SubjectIsMandatory          bool    `db:"subject_is_mandatory"`
	CurriculumId                string  `db:"curriculum_id"`
	StudentCurriculumId         string  `db:"student_curriculum_id"`
	TotalCredit                 uint32  `db:"total_credit"`
	TotalAttendance             uint32  `db:"total_attendance"`
	TotalSick                   uint32  `db:"total_sick"`
	TotalLeave                  uint32  `db:"total_leave"`
	TotalAwol                   uint32  `db:"total_awol"`
	GradePoint                  float64 `db:"grade_point"`
	GradeCode                   *string `db:"grade_code"`
	GradedByAdminId             *string `db:"graded_by_admin_id"`
	GradedByAdminName           *string `db:"graded_by_admin_name"`
	GradedByLecturerId          *string `db:"graded_by_lecturer_id"`
	GradedByLecturerName        *string `db:"graded_by_lecturer_name"`
	GradedAt                    *string `db:"graded_at"`
	TotalLectureDone            uint32  `db:"total_lecture_done"`
	ActiveLectureId             *string `db:"active_lecture_id"`
	ActiveLectureHasAttend      *bool   `db:"active_lecture_has_attend"`
	SubjectIsMbkm               bool    `db:"subject_is_mbkm"`
	MbkmUsedCredit              uint32  `db:"mbkm_used_credit"`
}

type GetClassParticipant struct {
	StudentId             string     `db:"student_id"`
	StudentNimNumber      int64      `db:"student_nim_number"`
	StudentName           string     `db:"student_name"`
	StudyProgramId        string     `db:"study_program_id"`
	StudyProgramName      string     `db:"study_program_name"`
	DiktiStudyProgramCode string     `db:"dikti_study_program_code"`
	DiktiStudyProgramType string     `db:"dikti_study_program_type"`
	StudyLevelShortName   string     `db:"study_level_short_name"`
	StudentClassId        string     `db:"student_class_id"`
	StudyPlanId           string     `db:"study_plan_id"`
	SubjectId             string     `db:"subject_id"`
	SubjectName           string     `db:"subject_name"`
	CurriculumId          string     `db:"curriculum_id"`
	StudentCurriculumId   string     `db:"student_curriculum_id"`
	ClassId               string     `db:"class_id"`
	TotalAttendance       uint32     `db:"total_attendance"`
	TotalSick             uint32     `db:"total_sick"`
	TotalLeave            uint32     `db:"total_leave"`
	TotalAwol             uint32     `db:"total_awol"`
	IsAttend              *bool      `db:"is_attend"`
	IsSick                *bool      `db:"is_sick"`
	IsLeave               *bool      `db:"is_leave"`
	IsAwol                *bool      `db:"is_awol"`
	GradePoint            float64    `db:"grade_point"`
	GradeCode             *string    `db:"grade_code"`
	GradedByAdminId       *string    `db:"graded_by_admin_id"`
	GradedByAdminName     *string    `db:"graded_by_admin_name"`
	GradedByLecturerId    *string    `db:"graded_by_lecturer_id"`
	GradedByLecturerName  *string    `db:"graded_by_lecturer_name"`
	GradedAt              *time.Time `db:"graded_at"`
	SubjectRepetition     uint32     `db:"subject_repetition"`
}

type CreateStudentClass struct {
	Id                  string `db:"id"`
	StudyPlanId         string `db:"study_plan_id"`
	CurriculumId        string `db:"curriculum_id"`
	StudentCurriculumId string `db:"student_curriculum_id"`
	ClassId             string `db:"class_id"`
}

type StudentIdClassId struct {
	StudentId string `db:"student_id"`
	ClassId   string `db:"class_id"`
}

type GradeStudentClass struct {
	ClassId               string         `db:"class_id"`
	StudentId             string         `db:"student_id"`
	ClassGradeComponentId string         `db:"class_grade_component_id"`
	InitialGrade          float64        `db:"initial_grade"`
	GradedByAdminId       sql.NullString `db:"graded_by_admin_id"`
	GradedByLecturerId    sql.NullString `db:"graded_by_lecturer_id"`
}

type GetStudentClassGrade struct {
	ClassId                 string  `db:"class_id"`
	StudentId               string  `db:"student_id"`
	ClassGradeComponentId   string  `db:"class_grade_component_id"`
	ClassGradeComponentName string  `db:"class_grade_component_name"`
	InitialGrade            float64 `db:"initial_grade"`
	FinalGrade              float64 `db:"final_grade"`
}

type DeleteStudentClassExcludingClassIds struct {
	StudyPlanId      string
	ExcludedClassIds []string
}
