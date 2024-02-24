package exam_supervisor_role

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type ExamSupervisorRoleRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest) ([]models.GetExamSupervisorRole, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetExamSupervisorRole, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateExamSupervisorRole) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateExamSupervisorRole) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
}

func NewExamSupervisorRoleRepository(db *db.DB) ExamSupervisorRoleRepositoryInterface {
	return &examSupervisorRoleRepository{
		db,
	}
}
