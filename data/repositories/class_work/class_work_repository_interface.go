package class_work

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type ClassWorkRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, classIds []string, studentId string) ([]models.GetClassWork, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetClassWork, *constants.ErrorResponse)
	GetSubmission(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, classId, classWorkId string, ids []string, isSubmitted *bool) ([]models.GetClassWorkSubmission, common.Pagination, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateClassWork) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateClassWork) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, ids []string) *constants.ErrorResponse
	GradeSubmission(ctx context.Context, tx *sqlx.Tx, data []models.GradeClassWorkSubmission) *constants.ErrorResponse
	Submit(ctx context.Context, tx *sqlx.Tx, classWorkId, studentId, filePath, filePathType string) *constants.ErrorResponse
}

func NewClassWorkRepository(db *db.DB) ClassWorkRepositoryInterface {
	return &classWorkRepository{
		db,
	}
}
