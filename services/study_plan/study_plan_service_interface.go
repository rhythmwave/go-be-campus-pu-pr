package study_plan

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type StudyPlanServiceInterface interface {
	BulkCreate(ctx context.Context, data objects.BulkCreateStudyPlan) *constants.ErrorResponse
	BulkApprove(ctx context.Context, studyPlanIds []string, isApproved bool) *constants.ErrorResponse
	GetList(ctx context.Context, paginationData common.PaginationRequest, studentId, semesterId string) (objects.GetStudyPlanWithPagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, studentId, semesterId string) (objects.GetStudentStudyPlanDetail, *constants.ErrorResponse)
}

func NewStudyPlanService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) StudyPlanServiceInterface {
	return &studyPlanService{
		repoCtx,
		infraCtx,
	}
}
