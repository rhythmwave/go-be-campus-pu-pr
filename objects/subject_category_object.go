package objects

import (
	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetSubjectCategory struct {
	Id   string
	Code string
	Name string
}

type SubjectCategoryListWithPagination struct {
	Pagination common.Pagination
	Data       []GetSubjectCategory
}

type CreateSubjectCategory struct {
	Code string
	Name string
}

type UpdateSubjectCategory struct {
	Id   string
	Code string
	Name string
}
