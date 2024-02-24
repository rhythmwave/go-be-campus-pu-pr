package expertise_group

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type ExpertiseGroupRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest) ([]models.GetExpertiseGroup, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetExpertiseGroup, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateExpertiseGroup) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateExpertiseGroup) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
}

func NewExpertiseGroupRepository(db *db.DB) ExpertiseGroupRepositoryInterface {
	return &expertiseGroupRepository{
		db,
	}
}
