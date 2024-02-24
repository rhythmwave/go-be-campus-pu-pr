package class_event

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type ClassEventRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, classId, frequency string, futureEventOnly bool, isActive *bool) ([]models.GetClassEvent, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetClassEvent, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateClassEvent) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateClassEvent) *constants.ErrorResponse
	BulkUpdateActivation(ctx context.Context, tx *sqlx.Tx, ids []string, isActive bool) *constants.ErrorResponse
	BulkDelete(ctx context.Context, tx *sqlx.Tx, ids []string) *constants.ErrorResponse
}

func NewClassEventRepository(db *db.DB) ClassEventRepositoryInterface {
	return &classEventRepository{
		db,
	}
}
