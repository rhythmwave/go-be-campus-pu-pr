package building

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type BuildingServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.BuildingListWithPagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, id string) (objects.GetBuilding, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateBuilding) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateBuilding) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewBuildingService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) BuildingServiceInterface {
	return &buildingService{
		repoCtx,
		infraCtx,
	}
}
