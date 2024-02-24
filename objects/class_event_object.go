package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetClassEvent struct {
	Id                 string
	Title              string
	Frequency          string
	EventTime          time.Time
	LecturerId         string
	LecturerName       string
	LecturerFrontTitle *string
	LecturerBackDegree *string
	Remarks            *string
	IsActive           bool
	CreatedAt          string
}

type ClassEventListWithPagination struct {
	Pagination common.Pagination
	Data       []GetClassEvent
}

type CreateClassEvent struct {
	ClassId   string
	Title     string
	Frequency string
	EventTime time.Time
	Remarks   string
	IsActive  bool
}

type UpdateClassEvent struct {
	Id        string
	Title     string
	Frequency string
	EventTime time.Time
	Remarks   string
	IsActive  bool
}
