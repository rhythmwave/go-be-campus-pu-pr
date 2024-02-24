package lecturer_student_activity_log

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type LecturerStudentActivityLogInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, year, month uint32) ([]models.GetLecturerStudentActivityLog, common.Pagination, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateLecturerStudentActivityLog) *constants.ErrorResponse
}

func NewLecturerStudentActivityLogRepository(db *db.DB) LecturerStudentActivityLogInterface {
	return &lecturerStudentActivityLog{
		db,
	}
}
