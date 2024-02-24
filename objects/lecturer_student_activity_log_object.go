package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetLecturerStudentActivityLog struct {
	Id            string
	UserType      string
	UserId        string
	UserName      string
	UserUsername  string
	Module        string
	Action        string
	IpAddress     string
	UserAgent     string
	ExecutionTime float64
	MemoryUsage   float64
	CreatedAt     time.Time
}

type LecturerStudentActivityLogWithPagination struct {
	Pagination common.Pagination
	Data       []GetLecturerStudentActivityLog
}
