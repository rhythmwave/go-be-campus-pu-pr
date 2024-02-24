package exam_supervisor

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type ExamSupervisorRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studyProgramId string) ([]models.GetExamSupervisorList, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetExamSupervisorDetail, *constants.ErrorResponse)
	GetDetailByIds(ctx context.Context, tx *sqlx.Tx, ids []string) ([]models.GetExamSupervisorDetail, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateExamSupervisor) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateExamSupervisor) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
	GetExamLectureSupervisorByLectureIds(ctx context.Context, tx *sqlx.Tx, lectureIds []string) ([]models.GetExamLectureSupervisor, *constants.ErrorResponse)
	DeleteExamLectureSupervisorExcludingExamLectureIds(ctx context.Context, tx *sqlx.Tx, lectureId string, excludedExamSupervisorIds []string) *constants.ErrorResponse
	UpsertExamLectureSupervisor(ctx context.Context, tx *sqlx.Tx, data []models.UpsertExamLectureSupervisor) *constants.ErrorResponse
}

func NewExamSupervisorRepository(db *db.DB) ExamSupervisorRepositoryInterface {
	return &examSupervisorRepository{
		db,
	}
}
