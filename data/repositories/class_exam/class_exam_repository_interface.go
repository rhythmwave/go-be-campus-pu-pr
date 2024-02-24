package class_exam

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type ClassExamRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, classIds []string, studentId string, ids []string) ([]models.GetClassExam, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetClassExam, *constants.ErrorResponse)
	GetSubmission(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, classId, classExamId string, ids []string, isSubmitted *bool) ([]models.GetClassExamSubmission, common.Pagination, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateClassExam) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateClassExam) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, ids []string) *constants.ErrorResponse
	GradeSubmission(ctx context.Context, tx *sqlx.Tx, data []models.GradeClassExamSubmission) *constants.ErrorResponse
	Submit(ctx context.Context, tx *sqlx.Tx, classExamId, studentId, filePath, filePathType string) *constants.ErrorResponse
}

func NewClassExamRepository(db *db.DB) ClassExamRepositoryInterface {
	return &classExamRepository{
		db,
	}
}
