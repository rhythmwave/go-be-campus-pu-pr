package lesson_plan

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type LessonPlanServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest, subjectId string) (objects.LessonPlanListWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateLessonPlan) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateLessonPlan) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewLessonPlanService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) LessonPlanServiceInterface {
	return &lessonPlanService{
		repoCtx,
		infraCtx,
	}
}
