package semester

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type SemesterServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest, studyProgramId, excludedId string) (objects.SemesterListWithPagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, id string) (objects.GetSemesterDetail, *constants.ErrorResponse)
	GetActive(ctx context.Context) (objects.GetSemesterDetail, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateSemester) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateSemester) *constants.ErrorResponse
	UpdateActivation(ctx context.Context, id string, isActive bool) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewSemesterService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) SemesterServiceInterface {
	return &semesterService{
		repoCtx,
		infraCtx,
	}
}
