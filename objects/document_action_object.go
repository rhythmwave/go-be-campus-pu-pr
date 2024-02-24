package objects

import (
	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetDocumentAction struct {
	Id            string
	Action        string
	EnglishAction string
}

type DocumentActionListWithPagination struct {
	Pagination common.Pagination
	Data       []GetDocumentAction
}

type CreateDocumentAction struct {
	Action        string
	EnglishAction string
}

type UpdateDocumentAction struct {
	Id            string
	Action        string
	EnglishAction string
}
