package admin_activity_log

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type AdminActivityLogInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, year, month uint32) ([]models.GetAdminActivityLog, common.Pagination, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateAdminActivityLog) *constants.ErrorResponse
}

func NewAdminActivityLogRepository(db *db.DB) AdminActivityLogInterface {
	return &adminActivityLog{
		db,
	}
}
