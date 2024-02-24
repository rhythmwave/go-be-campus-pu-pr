package objects

import (
	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetThesisSupervisorRole struct {
	Id   string
	Name string
	Sort uint32
}

type ThesisSupervisorRoleListWithPagination struct {
	Pagination common.Pagination
	Data       []GetThesisSupervisorRole
}

type CreateThesisSupervisorRole struct {
	Name string
	Sort uint32
}

type UpdateThesisSupervisorRole struct {
	Id   string
	Name string
	Sort uint32
}
