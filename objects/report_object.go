package objects

import "github.com/sccicitb/pupr-backend/objects/common"

type ReportStudentStatusStatus struct {
	Status string
	Total  uint32
}

type ReportStudentStatus struct {
	StudyProgramId        string
	StudyProgramName      string
	DiktiStudyProgramCode string
	DiktiStudyProgramType string
	StudyLevelShortName   string
	Statuses              []ReportStudentStatusStatus
}

type ReportStudentClassGradeGrade struct {
	GradeCode string
	Total     uint32
}

type ReportStudentClassGrade struct {
	SubjectId   string
	SubjectName string
	Grades      []ReportStudentClassGradeGrade
}

type ReportStudentClassGradeWithPagination struct {
	Pagination common.Pagination
	Data       []ReportStudentClassGrade
}

type ReportStudentProvinceStudentForce struct {
	StudentForce uint32
	Total        uint32
}

type ReportStudentProvince struct {
	ProvinceId    uint32
	ProvinceName  string
	StudentForces []ReportStudentProvinceStudentForce
}

type ReportStudentSchoolProvinceStudentForce struct {
	StudentForce uint32
	Total        uint32
}

type ReportStudentSchoolProvince struct {
	ProvinceId    uint32
	ProvinceName  string
	StudentForces []ReportStudentSchoolProvinceStudentForce
}
