package shared_file

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type SharedFileRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, appType string, lecturerId string, isApproved *bool) ([]models.GetSharedFile, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetSharedFile, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateSharedFile) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateSharedFile) *constants.ErrorResponse
	Approve(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
}

func NewSharedFileRepository(db *db.DB) SharedFileRepositoryInterface {
	return &sharedFileRepository{
		db,
	}
}
