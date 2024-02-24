package officer_action

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type OfficerActionServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.OfficerActionListWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateOfficerAction) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateOfficerAction) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewOfficerActionService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) OfficerActionServiceInterface {
	return &officerActionService{
		repoCtx,
		infraCtx,
	}
}
