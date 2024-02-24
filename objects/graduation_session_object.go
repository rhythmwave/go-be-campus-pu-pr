package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetGraduationSession struct {
	Id                string
	SessionYear       uint32
	SessionSchoolYear string
	SessionNumber     uint32
	SessionDate       time.Time
}

type GraduationSessionListWithPagination struct {
	Pagination common.Pagination
	Data       []GetGraduationSession
}

type CreateGraduationSession struct {
	SessionYear   uint32
	SessionNumber uint32
	SessionDate   time.Time
}

type UpdateGraduationSession struct {
	Id            string
	SessionYear   uint32
	SessionNumber uint32
	SessionDate   time.Time
}
