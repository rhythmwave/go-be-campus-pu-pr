package lesson_plan

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type LessonPlanRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, subjectId string) ([]models.GetLessonPlan, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetLessonPlan, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateLessonPlan) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateLessonPlan) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
}

func NewLessonPlanRepository(db *db.DB) LessonPlanRepositoryInterface {
	return &lessonPlanRepository{
		db,
	}
}
