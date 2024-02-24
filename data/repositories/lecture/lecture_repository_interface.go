package lecture

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type LectureRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, classId, semesterId string, hasActualLecture, isExam *bool, examType string) ([]models.GetLectureList, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetLectureDetail, *constants.ErrorResponse)
	GetByClassIds(ctx context.Context, tx *sqlx.Tx, classIds []string) ([]models.GetLectureDetail, *constants.ErrorResponse)
	BulkCreate(ctx context.Context, tx *sqlx.Tx, data []models.CreateLecture) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateLecture) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
	ResetParticipation(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
	GetParticipantByLectureId(ctx context.Context, tx *sqlx.Tx, lectureId string) ([]models.GetLectureParticipant, *constants.ErrorResponse)
	BulkUpdateParticipant(ctx context.Context, tx *sqlx.Tx, data []models.UpdateLectureParticipant) *constants.ErrorResponse
	GetStudentParticipation(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, classId, studentId string) ([]models.GetLectureParticipation, common.Pagination, *constants.ErrorResponse)
	AttendLecture(ctx context.Context, tx *sqlx.Tx, lectureId, studentId string) *constants.ErrorResponse
	GetLectureCalendar(ctx context.Context, tx *sqlx.Tx, req objects.GetLectureCalendarRequest) ([]models.GetLectureCalendar, *constants.ErrorResponse)
	GetHistory(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studentId string, startDate, endDate time.Time) ([]models.GetLectureHistory, common.Pagination, *constants.ErrorResponse)
}

func NewLectureRepository(db *db.DB) LectureRepositoryInterface {
	return &lectureRepository{
		db,
	}
}
