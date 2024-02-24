package student_activity

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type StudentActivityRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, activityType, studyProgramId, semesterId string, isMbkm bool) ([]models.GetStudentActivity, common.Pagination, *constants.ErrorResponse)
	GetListParticipantByStudentActivityId(ctx context.Context, tx *sqlx.Tx, studentActivityId string) ([]models.GetStudentActivityParticipant, *constants.ErrorResponse)
	GetListLecturerByStudentActivityId(ctx context.Context, tx *sqlx.Tx, studentActivityId string) ([]models.GetStudentActivityLecturer, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetStudentActivityDetail, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateStudentActivity) (string, *constants.ErrorResponse)
	DeleteParticipantExcludingStudentIds(ctx context.Context, tx *sqlx.Tx, studentActivityId string, excludedStudentIds []string) *constants.ErrorResponse
	DeleteLecturerExcludingLecturerIds(ctx context.Context, tx *sqlx.Tx, studentActivityId, role string, excludedLecturerIds []string) *constants.ErrorResponse
	UpsertParticipant(ctx context.Context, tx *sqlx.Tx, data []models.UpsertStudentActivityParticipant) *constants.ErrorResponse
	UpsertLecturer(ctx context.Context, tx *sqlx.Tx, data []models.UpsertStudentActivityLecturer) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateStudentActivity) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
}

func NewStudentActivityRepository(db *db.DB) StudentActivityRepositoryInterface {
	return &studentActivityRepository{
		db,
	}
}
