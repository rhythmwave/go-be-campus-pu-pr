package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type ApplyYudicium struct {
	StudentId       string
	ApplicationDate time.Time
	WithThesis      bool
}

type GetListStudentYudiciumRequest struct {
	StudyProgramId   string
	NimNumberFrom    int64
	NimNumberTo      int64
	SemesterId       string
	StudentForceFrom uint32
	StudentForceTo   uint32
}

type GetListStudentYudicium struct {
	Id                    string
	NimNumber             int64
	Name                  string
	DiktiStudyProgramCode *string
	StudyProgramName      *string
	StudyLevelShortName   *string
	DiktiStudyProgramType *string
	TotalCredit           *uint32
	Gpa                   *float64
	Status                *string
	ApplicationDate       time.Time
	DoneYudicium          bool
}

type GetListStudentYudiciumWithPagination struct {
	Pagination common.Pagination
	Data       []GetListStudentYudicium
}

type DoYudicium struct {
	YudiciumSessionId string
	YudiciumNumber    string
	YudiciumDate      time.Time
	StudentIds        []string
}
