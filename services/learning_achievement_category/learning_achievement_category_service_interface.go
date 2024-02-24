package learning_achievement_category

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type LearningAchievementCategoryServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest, curriculumId string) (objects.LearningAchievementCategoryListWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateLearningAchievementCategory) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateLearningAchievementCategory) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewLearningAchievementCategoryService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) LearningAchievementCategoryServiceInterface {
	return &learningAchievementCategoryService{
		repoCtx,
		infraCtx,
	}
}
