package location

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type LocationServiceInterface interface {
	GetListCountry(ctx context.Context, paginationData common.PaginationRequest) (objects.LocationListWithPagination, *constants.ErrorResponse)
	GetListProvince(ctx context.Context, paginationData common.PaginationRequest, countryId uint32) (objects.LocationListWithPagination, *constants.ErrorResponse)
	GetListRegency(ctx context.Context, paginationData common.PaginationRequest, provinceId uint32) (objects.LocationListWithPagination, *constants.ErrorResponse)
	GetListDistrict(ctx context.Context, paginationData common.PaginationRequest, regencyId uint32) (objects.LocationListWithPagination, *constants.ErrorResponse)
	GetListVillage(ctx context.Context, paginationData common.PaginationRequest, districtId uint32) (objects.LocationListWithPagination, *constants.ErrorResponse)
	TempCreateData(ctx context.Context, data objects.TempCreateData) *constants.ErrorResponse
	TempGetData(ctx context.Context, paginationData common.PaginationRequest) (objects.TempGetDataWithPagination, *constants.ErrorResponse)
}

func NewLocationService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) LocationServiceInterface {
	return &locationService{
		repoCtx,
		infraCtx,
	}
}
