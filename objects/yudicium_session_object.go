package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetYudiciumSession struct {
	Id                 string
	SemesterId         string
	SemesterSchoolYear string
	SemesterType       string
	Name               string
	SessionDate        time.Time
}

type YudiciumSessionListWithPagination struct {
	Pagination common.Pagination
	Data       []GetYudiciumSession
}

type CreateYudiciumSession struct {
	SemesterId  string
	Name        string
	SessionDate time.Time
}

type UpdateYudiciumSession struct {
	Id          string
	SemesterId  string
	Name        string
	SessionDate time.Time
}
