package objects

import (
	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetExpertiseGroup struct {
	Id   string
	Name string
}

type ExpertiseGroupListWithPagination struct {
	Pagination common.Pagination
	Data       []GetExpertiseGroup
}

type CreateExpertiseGroup struct {
	Name string
}

type UpdateExpertiseGroup struct {
	Id   string
	Name string
}
