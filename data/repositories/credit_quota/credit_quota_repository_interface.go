package credit_quota

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type CreditQuotaRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest) ([]models.GetCreditQuota, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetCreditQuota, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateCreditQuota) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateCreditQuota) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
}

func NewCreditQuotaRepository(db *db.DB) CreditQuotaRepositoryInterface {
	return &creditQuotaRepository{
		db,
	}
}
