package learning_achievement

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type LearningAchievementServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest, learningAchievementCategoryId string) (objects.LearningAchievementListWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateLearningAchievement) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateLearningAchievement) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewLearningAchievementService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) LearningAchievementServiceInterface {
	return &learningAchievementService{
		repoCtx,
		infraCtx,
	}
}
