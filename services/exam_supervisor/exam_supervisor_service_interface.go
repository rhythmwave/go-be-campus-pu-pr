package exam_supervisor

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type ExamSupervisorServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest, studyProgramId string) (objects.ExamSupervisorListWithPagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, id string) (objects.GetExamSupervisorDetail, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateExamSupervisor) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateExamSupervisor) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewExamSupervisorService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) ExamSupervisorServiceInterface {
	return &examSupervisorService{
		repoCtx,
		infraCtx,
	}
}
