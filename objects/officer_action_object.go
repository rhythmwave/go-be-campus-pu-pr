package objects

import (
	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetOfficerAction struct {
	Id                          string
	DocumentTypeId              string
	DocumentTypeName            string
	DocumentActionId            string
	DocumentActionAction        string
	DocumentActionEnglishAction string
	OfficerId                   string
	OfficerName                 string
	OfficerTitle                *string
	OfficerEnglishTitle         *string
	OfficerStudyProgramId       *string
	OfficerStudyProgramName     *string
}

type OfficerActionListWithPagination struct {
	Pagination common.Pagination
	Data       []GetOfficerAction
}

type CreateOfficerAction struct {
	DocumentTypeId   string
	DocumentActionId string
	OfficerId        string
}

type UpdateOfficerAction struct {
	Id               string
	DocumentTypeId   string
	DocumentActionId string
	OfficerId        string
}
