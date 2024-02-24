package officer

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type OfficerServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.OfficerListWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateOfficer) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateOfficer) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewOfficerService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) OfficerServiceInterface {
	return &officerService{
		repoCtx,
		infraCtx,
	}
}
