package models

type GetReportStudentStatus struct {
	SemesterId            string `db:"semester_id"`
	SemesterType          string `db:"semester_type"`
	SemesterStartYear     string `db:"semester_start_year"`
	StudyProgramId        string `db:"study_program_id"`
	StudyProgramName      string `db:"study_program_name"`
	DiktiStudyProgramType string `db:"dikti_study_program_type"`
	DiktiStudyProgramCode string `db:"dikti_study_program_code"`
	StudyLevelShortName   string `db:"study_level_short_name"`
	Status                string `db:"status"`
	Total                 uint32 `db:"total"`
}

type GetReportStudentClassGrade struct {
	SubjectId string `db:"subject_id"`
	GradeCode string `db:"grade_code"`
	Total     uint32 `db:"total"`
}

type GetReportStudentProvince struct {
	ProvinceId   uint32 `db:"province_id"`
	ProvinceName string `db:"province_name"`
	StudentForce uint32 `db:"student_force"`
	Total        uint32 `db:"total"`
}

type GetReportStudentSchoolProvince struct {
	ProvinceId   uint32 `db:"province_id"`
	ProvinceName string `db:"province_name"`
	StudentForce uint32 `db:"student_force"`
	Total        uint32 `db:"total"`
}
