package objects

import (
	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetThesisExaminerRole struct {
	Id      string
	Name    string
	Remarks string
	Sort    uint32
}

type ThesisExaminerRoleListWithPagination struct {
	Pagination common.Pagination
	Data       []GetThesisExaminerRole
}

type CreateThesisExaminerRole struct {
	Name    string
	Remarks string
	Sort    uint32
}

type UpdateThesisExaminerRole struct {
	Id      string
	Name    string
	Remarks string
	Sort    uint32
}
