package yudicium_session

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type YudiciumSessionRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest) ([]models.GetYudiciumSession, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetYudiciumSession, *constants.ErrorResponse)
	GetUpcoming(ctx context.Context, tx *sqlx.Tx) (models.GetYudiciumSession, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateYudiciumSession) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateYudiciumSession) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
	Do(ctx context.Context, tx *sqlx.Tx, data models.DoYudicium) *constants.ErrorResponse
}

func NewYudiciumSessionRepository(db *db.DB) YudiciumSessionRepositoryInterface {
	return &yudiciumSessionRepository{
		db,
	}
}
