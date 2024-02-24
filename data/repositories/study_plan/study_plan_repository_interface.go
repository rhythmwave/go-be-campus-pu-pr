package study_plan

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type StudyPlanRepositoryInterface interface {
	BulkCreate(ctx context.Context, tx *sqlx.Tx, data []models.CreateStudyPlan) *constants.ErrorResponse
	BulkApprove(ctx context.Context, tx *sqlx.Tx, studyPlanIds []string, isApproved bool) *constants.ErrorResponse
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studentId, semesterId string) ([]models.GetStudyPlan, common.Pagination, *constants.ErrorResponse)
	GetApprovedByStudentId(ctx context.Context, tx *sqlx.Tx, studentId string) ([]models.GetStudyPlan, *constants.ErrorResponse)
	GetByStudentIdAndSemesterId(ctx context.Context, tx *sqlx.Tx, studentId, semesterId string) (models.GetStudyPlan, *constants.ErrorResponse)
	GetByStudentIdsAndSemesterId(ctx context.Context, tx *sqlx.Tx, studentIds []string, semesterId string) ([]models.GetStudyPlan, *constants.ErrorResponse)
}

func NewStudyPlanRepository(db *db.DB) StudyPlanRepositoryInterface {
	return &studyPlanRepository{
		db,
	}
}
