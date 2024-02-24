package thesis_supervisor_role

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type ThesisSupervisorRoleRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest) ([]models.GetThesisSupervisorRole, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetThesisSupervisorRole, *constants.ErrorResponse)
	GetFirstOrder(ctx context.Context, tx *sqlx.Tx) (models.GetThesisSupervisorRole, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateThesisSupervisorRole) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateThesisSupervisorRole) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
}

func NewThesisSupervisorRoleRepository(db *db.DB) ThesisSupervisorRoleRepositoryInterface {
	return &thesisSupervisorRoleRepository{
		db,
	}
}
