package objects

import (
	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetBuilding struct {
	Id          string
	Level       string
	FacultyId   string
	FacultyName string
	MajorId     string
	MajorName   string
	Code        string
	Name        string
}

type BuildingListWithPagination struct {
	Pagination common.Pagination
	Data       []GetBuilding
}

type CreateBuilding struct {
	FacultyId string
	MajorId   string
	Code      string
	Name      string
}

type UpdateBuilding struct {
	Id        string
	FacultyId string
	MajorId   string
	Code      string
	Name      string
}
