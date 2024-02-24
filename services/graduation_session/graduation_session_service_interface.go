package graduation_session

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type GraduationSessionServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.GraduationSessionListWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateGraduationSession) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateGraduationSession) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewGraduationSessionService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) GraduationSessionServiceInterface {
	return &graduationSessionService{
		repoCtx,
		infraCtx,
	}
}
