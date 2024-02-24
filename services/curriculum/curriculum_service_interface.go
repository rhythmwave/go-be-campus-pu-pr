package curriculum

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type CurriculumServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest, studyProgramId string) (objects.CurriculumListWithPagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, id string) (objects.GetCurriculumDetail, *constants.ErrorResponse)
	GetActiveByStudyProgramId(ctx context.Context, studyProgramId string) (objects.GetCurriculumDetail, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateCurriculum) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateCurriculum) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewCurriculumService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) CurriculumServiceInterface {
	return &curriculumService{
		repoCtx,
		infraCtx,
	}
}
