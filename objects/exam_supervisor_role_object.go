package objects

import (
	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetExamSupervisorRole struct {
	Id   string
	Name string
	Sort uint32
}

type ExamSupervisorRoleListWithPagination struct {
	Pagination common.Pagination
	Data       []GetExamSupervisorRole
}

type CreateExamSupervisorRole struct {
	Name string
	Sort uint32
}

type UpdateExamSupervisorRole struct {
	Id   string
	Name string
	Sort uint32
}
