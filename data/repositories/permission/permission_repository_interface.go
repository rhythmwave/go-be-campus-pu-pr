package permission

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type PermissionRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest) ([]models.GetPermissionList, common.Pagination, *constants.ErrorResponse)
	GetByRoleIds(ctx context.Context, tx *sqlx.Tx, roleIds []string) ([]models.GetPermissionByRoleIds, *constants.ErrorResponse)
}

func NewPermissionRepository(db *db.DB) PermissionRepositoryInterface {
	return &permissionRepository{
		db,
	}
}
