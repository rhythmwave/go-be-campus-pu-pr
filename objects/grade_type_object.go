package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetGradeType struct {
	Id                  string
	StudyLevelId        string
	StudyLevelShortName string
	Code                string
	GradePoint          float64
	MinimumGrade        float64
	MaximumGrade        float64
	GradeCategory       string
	GradePointCategory  float64
	Label               *string
	EnglishLabel        *string
	StartDate           time.Time
	EndDate             time.Time
}

type GradeTypeListWithPagination struct {
	Pagination common.Pagination
	Data       []GetGradeType
}

type CreateGradeType struct {
	StudyLevelId       string
	Code               string
	GradePoint         float64
	MinimumGrade       float64
	MaximumGrade       float64
	GradeCategory      string
	GradePointCategory float64
	Label              string
	EnglishLabel       string
	StartDate          time.Time
	EndDate            time.Time
}

type UpdateGradeType struct {
	Id                 string
	Code               string
	GradePoint         float64
	MinimumGrade       float64
	MaximumGrade       float64
	GradeCategory      string
	GradePointCategory float64
	Label              string
	EnglishLabel       string
	StartDate          time.Time
	EndDate            time.Time
}
