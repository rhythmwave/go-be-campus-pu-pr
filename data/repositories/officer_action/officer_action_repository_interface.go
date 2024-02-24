package officer_action

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type OfficerActionRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest) ([]models.GetOfficerAction, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetOfficerAction, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateOfficerAction) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateOfficerAction) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
}

func NewOfficerActionRepository(db *db.DB) OfficerActionRepositoryInterface {
	return &officerActionRepository{
		db,
	}
}
