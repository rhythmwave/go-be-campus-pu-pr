package class_material

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type ClassMaterialRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, classId string, isActive *bool) ([]models.GetClassMaterial, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetClassMaterial, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateClassMaterial) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateClassMaterial) *constants.ErrorResponse
	BulkUpdateActivation(ctx context.Context, tx *sqlx.Tx, ids []string, isActive bool) *constants.ErrorResponse
	BulkDelete(ctx context.Context, tx *sqlx.Tx, ids []string) *constants.ErrorResponse
}

func NewClassMaterialRepository(db *db.DB) ClassMaterialRepositoryInterface {
	return &classMaterialRepository{
		db,
	}
}
