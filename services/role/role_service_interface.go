package role

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type RoleServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.RoleListWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateRole) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateRole) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewRoleService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) RoleServiceInterface {
	return &roleService{
		repoCtx,
		infraCtx,
	}
}
