package yudicium_session

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type YudiciumSessionServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.YudiciumSessionListWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateYudiciumSession) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateYudiciumSession) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewYudiciumSessionService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) YudiciumSessionServiceInterface {
	return &yudiciumSessionService{
		repoCtx,
		infraCtx,
	}
}
