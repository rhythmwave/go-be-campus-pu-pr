package thesis_examiner_role

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type ThesisExaminerRoleRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest) ([]models.GetThesisExaminerRole, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetThesisExaminerRole, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateThesisExaminerRole) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateThesisExaminerRole) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
}

func NewThesisExaminerRoleRepository(db *db.DB) ThesisExaminerRoleRepositoryInterface {
	return &thesisExaminerRoleRepository{
		db,
	}
}
