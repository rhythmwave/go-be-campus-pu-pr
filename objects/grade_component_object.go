package objects

import "github.com/sccicitb/pupr-backend/objects/common"

type GetGradeComponent struct {
	Id                  string
	StudyProgramId      string
	StudyProgramName    string
	SubjectCategoryId   string
	SubjectCategoryName string
	Name                string
	IsActive            bool
	DefaultPercentage   float64
}

type GradeComponentListWithPagination struct {
	Pagination common.Pagination
	Data       []GetGradeComponent
}

type CreateGradeComponent struct {
	StudyProgramId    string
	SubjectCategoryId string
	Name              string
	IsActive          bool
}

type UpdateGradeComponent struct {
	Id                string
	SubjectCategoryId string
	Name              string
	IsActive          bool
}

type GetGradeComponentBySubjectCategoryGradeComponent struct {
	Id                string
	Name              string
	DefaultPercentage float64
	IsActive          bool
}

type GetGradeComponentBySubjectCategory struct {
	StudyProgramId      string
	StudyProgramName    string
	SubjectCategoryId   string
	SubjectCategoryName string
	GradeComponents     []GetGradeComponentBySubjectCategoryGradeComponent
}

type GradeComponentBySubjectCategoryListWithPagination struct {
	Pagination common.Pagination
	Data       []GetGradeComponentBySubjectCategory
}

type BulkUpdatePercentageGradeComponent struct {
	StudyProgramId    string
	SubjectCategoryId string
}

type BulkUpdatePercentageGradeComponentData struct {
	Id                string
	DefaultPercentage float64
	IsActive          bool
}
