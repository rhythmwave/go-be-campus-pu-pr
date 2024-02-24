package student_leave

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type StudentLeaveRepositoryInterface interface {
	GetListRequest(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studyProgramId, studentId string, isApproved *bool) ([]models.GetStudentLeaveRequest, common.Pagination, *constants.ErrorResponse)
	GetDetailRequest(ctx context.Context, tx *sqlx.Tx, id string) (models.GetStudentLeaveRequest, *constants.ErrorResponse)
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studyProgramId, semesterId string) ([]models.GetStudentLeave, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetStudentLeave, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateStudentLeave) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateStudentLeave) *constants.ErrorResponse
	Approve(ctx context.Context, tx *sqlx.Tx, id string, isApproved bool) *constants.ErrorResponse
	End(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
}

func NewStudentLeaveRepository(db *db.DB) StudentLeaveRepositoryInterface {
	return &studentLeaveRepository{
		db,
	}
}
