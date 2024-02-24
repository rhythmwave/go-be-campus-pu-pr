package lecturer

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type LecturerRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, req objects.GetLecturerRequest) ([]models.GetLecturerList, common.Pagination, *constants.ErrorResponse)
	GetSchedule(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studyProgramId, idNationalLecturer, semesterId string) ([]models.GetLecturerSchedule, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetLecturerDetail, *constants.ErrorResponse)
	GetDetailByIds(ctx context.Context, tx *sqlx.Tx, ids []string) ([]models.GetLecturerDetail, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateLecturer) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateLecturer) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
	UpdateStatus(ctx context.Context, tx *sqlx.Tx, ids []string, status string) *constants.ErrorResponse
	GetAssignedClass(ctx context.Context, tx *sqlx.Tx, lecturerId, semesterId, classId string, classIsActive *bool) ([]models.GetLecturerAssignedClass, *constants.ErrorResponse)
}

func NewLecturerRepository(db *db.DB) LecturerRepositoryInterface {
	return &lecturerRepository{
		db,
	}
}
