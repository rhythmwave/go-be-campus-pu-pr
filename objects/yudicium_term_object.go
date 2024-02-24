package objects

import (
	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetYudiciumTerm struct {
	Id               string
	CurriculumId     string
	CurriculumName   string
	StudyProgramId   string
	StudyProgramName string
	Term             string
	Remarks          string
}

type YudiciumTermListWithPagination struct {
	Pagination common.Pagination
	Data       []GetYudiciumTerm
}

type CreateYudiciumTerm struct {
	CurriculumId string
	Term         string
	Remarks      string
}

type UpdateYudiciumTerm struct {
	Id      string
	Term    string
	Remarks string
}
