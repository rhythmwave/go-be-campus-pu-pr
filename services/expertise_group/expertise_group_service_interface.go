package expertise_group

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type ExpertiseGroupServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.ExpertiseGroupListWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateExpertiseGroup) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateExpertiseGroup) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewExpertiseGroupService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) ExpertiseGroupServiceInterface {
	return &expertiseGroupService{
		repoCtx,
		infraCtx,
	}
}
