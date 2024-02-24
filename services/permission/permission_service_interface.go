package permission

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type PermissionServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.PermissionListWithPagination, *constants.ErrorResponse)
}

func NewPermissionService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) PermissionServiceInterface {
	return &permissionService{
		repoCtx,
		infraCtx,
	}
}
