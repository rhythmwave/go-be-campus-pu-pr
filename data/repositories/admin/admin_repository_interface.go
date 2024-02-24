package admin

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type AdminRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest) ([]models.GetAdmin, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetAdmin, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateAdmin) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateAdmin) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
	GetSingleSuperAdmin(ctx context.Context, tx *sqlx.Tx) (models.GetAdmin, *constants.ErrorResponse)
}

func NewAdminRepository(db *db.DB) AdminRepositoryInterface {
	return &adminRepository{
		db,
	}
}
