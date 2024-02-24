package thesis_supervisor_role

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type ThesisSupervisorRoleServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.ThesisSupervisorRoleListWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateThesisSupervisorRole) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateThesisSupervisorRole) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewThesisSupervisorRoleService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) ThesisSupervisorRoleServiceInterface {
	return &thesisSupervisorRoleService{
		repoCtx,
		infraCtx,
	}
}
