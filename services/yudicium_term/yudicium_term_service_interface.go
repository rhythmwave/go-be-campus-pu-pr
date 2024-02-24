package yudicium_term

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type YudiciumTermServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest, curriculumId string) (objects.YudiciumTermListWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateYudiciumTerm) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateYudiciumTerm) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewYudiciumTermService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) YudiciumTermServiceInterface {
	return &yudiciumTermService{
		repoCtx,
		infraCtx,
	}
}
