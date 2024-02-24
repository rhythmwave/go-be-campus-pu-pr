package officer

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type OfficerRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest) ([]models.GetOfficer, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetOfficer, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateOfficer) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateOfficer) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
}

func NewOfficerRepository(db *db.DB) OfficerRepositoryInterface {
	return &officerRepository{
		db,
	}
}
