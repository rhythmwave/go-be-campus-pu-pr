package document_action

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type DocumentActionRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest) ([]models.GetDocumentAction, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetDocumentAction, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateDocumentAction) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateDocumentAction) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
}

func NewDocumentActionRepository(db *db.DB) DocumentActionRepositoryInterface {
	return &documentActionRepository{
		db,
	}
}
