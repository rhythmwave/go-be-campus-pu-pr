package lecturer_leave

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type LecturerLeaveRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studyProgramId, idNationalLecturer, semesterId string, isActive bool) ([]models.GetLecturerLeave, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetLecturerLeave, *constants.ErrorResponse)
	GetDetailByLecturerIdAndStartDate(ctx context.Context, tx *sqlx.Tx, lecturerId string, startDate time.Time) (models.GetLecturerLeave, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateLecturerLeave) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateLecturerLeave) *constants.ErrorResponse
	End(ctx context.Context, tx *sqlx.Tx, id, adminId string) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
	AutoSetActive(ctx context.Context, tx *sqlx.Tx) ([]models.LecturerId, *constants.ErrorResponse)
}

func NewLecturerLeaveRepository(db *db.DB) LecturerLeaveRepositoryInterface {
	return &lecturerLeaveRepository{
		db,
	}
}
