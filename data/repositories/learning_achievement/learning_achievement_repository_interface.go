package learning_achievement

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type LearningAchievementRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, learningAchievementCategoryId string) ([]models.GetLearningAchievement, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetLearningAchievement, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateLearningAchievement) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateLearningAchievement) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
}

func NewLearningAchievementRepository(db *db.DB) LearningAchievementRepositoryInterface {
	return &learningAchievementRepository{
		db,
	}
}
