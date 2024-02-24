package accreditation

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type AccreditationServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest, studyProgramId string) (objects.AccreditationListWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateAccreditation) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateAccreditation) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewAccreditationService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) AccreditationServiceInterface {
	return &accreditationService{
		repoCtx,
		infraCtx,
	}
}
