package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type ApplyGraduation struct {
	StudentId           string
	ApplicationDate     time.Time
	GraduationSessionId string
}

type GetListStudentGraduation struct {
	Id                    string
	NimNumber             int64
	Name                  string
	DiktiStudyProgramCode *string
	StudyProgramName      *string
	StudyLevelShortName   *string
	DiktiStudyProgramType *string
	ApplicationDate       time.Time
}

type GetListStudentGraduationWithPagination struct {
	Pagination common.Pagination
	Data       []GetListStudentGraduation
}
