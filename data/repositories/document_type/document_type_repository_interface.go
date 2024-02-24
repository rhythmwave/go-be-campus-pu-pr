package document_type

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type DocumentTypeRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest) ([]models.GetDocumentType, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetDocumentType, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateDocumentType) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateDocumentType) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
}

func NewDocumentTypeRepository(db *db.DB) DocumentTypeRepositoryInterface {
	return &documentTypeRepository{
		db,
	}
}
