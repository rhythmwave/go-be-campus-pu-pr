package role

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type RoleRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest) ([]models.GetRoleList, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetRoleList, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateRole) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateRole) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
}

func NewRoleRepository(db *db.DB) RoleRepositoryInterface {
	return &roleRepository{
		db,
	}
}
