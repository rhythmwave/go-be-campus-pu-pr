package building

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type BuildingRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest) ([]models.GetBuilding, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetBuilding, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateBuilding) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateBuilding) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
}

func NewBuildingRepository(db *db.DB) BuildingRepositoryInterface {
	return &buildingRepository{
		db,
	}
}
