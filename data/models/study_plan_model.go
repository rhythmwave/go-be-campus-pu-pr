package models

type CreateStudyPlan struct {
	Id              string `db:"id"`
	StudentId       string `db:"student_id"`
	SemesterId      string `db:"semester_id"`
	SemesterPackage uint32 `db:"semester_package"`
	MaximumCredit   uint32 `db:"maximum_credit"`
	IsSubmitted     bool   `db:"is_submitted"`
	IsThesis        bool   `db:"is_thesis"`
}

type GetStudyPlan struct {
	Id                    string  `db:"id"`
	IsSubmitted           bool    `db:"is_submitted"`
	IsApproved            bool    `db:"is_approved"`
	SemesterId            string  `db:"semester_id"`
	SemesterStartYear     uint32  `db:"semester_start_year"`
	SemesterType          string  `db:"semester_type"`
	TotalMandatoryCredit  uint32  `db:"total_mandatory_credit"`
	TotalOptionalCredit   uint32  `db:"total_optional_credit"`
	GradePoint            float64 `db:"grade_point"`
	StudentId             string  `db:"student_id"`
	StudentNimNumber      int64   `db:"student_nim_number"`
	StudentName           string  `db:"student_name"`
	StudyProgramId        string  `db:"study_program_id"`
	StudyProgramName      string  `db:"study_program_name"`
	DiktiStudyProgramCode string  `db:"dikti_study_program_code"`
	DiktiStudyProgramType string  `db:"dikti_study_program_type"`
	StudyLevelShortName   string  `db:"study_level_short_name"`
	IsThesis              bool    `db:"is_thesis"`
}
