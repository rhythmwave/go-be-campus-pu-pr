package objects

import (
	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetDocumentType struct {
	Id   string
	Name string
}

type DocumentTypeListWithPagination struct {
	Pagination common.Pagination
	Data       []GetDocumentType
}

type CreateDocumentType struct {
	Name string
}

type UpdateDocumentType struct {
	Id   string
	Name string
}
