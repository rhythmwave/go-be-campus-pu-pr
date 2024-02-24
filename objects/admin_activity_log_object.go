package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetAdminActivityLog struct {
	Id            string
	AdminId       string
	AdminName     string
	AdminUsername string
	Module        string
	Action        string
	IpAddress     string
	UserAgent     string
	ExecutionTime float64
	MemoryUsage   float64
	CreatedAt     time.Time
}

type AdminActivityLogWithPagination struct {
	Pagination common.Pagination
	Data       []GetAdminActivityLog
}
