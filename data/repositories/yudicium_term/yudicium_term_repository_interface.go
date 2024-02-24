package yudicium_term

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type YudiciumTermRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, curriculumId string) ([]models.GetYudiciumTerm, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetYudiciumTerm, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateYudiciumTerm) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateYudiciumTerm) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
}

func NewYudiciumTermRepository(db *db.DB) YudiciumTermRepositoryInterface {
	return &yudiciumTermRepository{
		db,
	}
}
