package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetAccreditation struct {
	Id             string
	StudyProgramId string
	DecreeNumber   string
	DecreeDate     time.Time
	DecreeDueDate  time.Time
	Accreditation  string
	IsActive       bool
}

type AccreditationListWithPagination struct {
	Pagination common.Pagination
	Data       []GetAccreditation
}

type CreateAccreditation struct {
	StudyProgramId string
	DecreeNumber   string
	DecreeDate     time.Time
	DecreeDueDate  time.Time
	Accreditation  string
	IsActive       bool
}

type UpdateAccreditation struct {
	Id            string
	DecreeNumber  string
	DecreeDate    time.Time
	DecreeDueDate time.Time
	Accreditation string
	IsActive      bool
}
