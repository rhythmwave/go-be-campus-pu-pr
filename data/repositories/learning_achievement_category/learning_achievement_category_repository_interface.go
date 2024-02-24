package learning_achievement_category

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type LearningAchievementCategoryRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, curriculumId string) ([]models.GetLearningAchievementCategory, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetLearningAchievementCategory, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateLearningAchievementCategory) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateLearningAchievementCategory) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
}

func NewLearningAchievementCategoryRepository(db *db.DB) LearningAchievementCategoryRepositoryInterface {
	return &learningAchievementCategoryRepository{
		db,
	}
}
