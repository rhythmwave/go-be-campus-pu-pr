package student

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type StudentRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, requestData objects.GetStudentRequest) ([]models.GetStudent, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string, previousSemesterId string) (models.GetStudent, *constants.ErrorResponse)
	GetDetailByIds(ctx context.Context, tx *sqlx.Tx, ids []string) ([]models.GetStudent, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateStudent) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateStudent) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
	GetActive(ctx context.Context, tx *sqlx.Tx, previousSemesterId string, studentIds []string) ([]models.GetActiveStudent, *constants.ErrorResponse)
	UpdateActiveSemesterPackage(ctx context.Context, tx *sqlx.Tx, studentIds []string) *constants.ErrorResponse
	BulkUpdateStatus(ctx context.Context, tx *sqlx.Tx, data models.BulkUpdateStatusStudent) *constants.ErrorResponse
	GetStatusSummary(ctx context.Context, tx *sqlx.Tx, studyProgramIds []string, semesterId string) ([]models.StudentStatusSummary, *constants.ErrorResponse)
	UpdateProfile(ctx context.Context, tx *sqlx.Tx, data models.UpdateStudentProfile) *constants.ErrorResponse
	UpdateParentProfile(ctx context.Context, tx *sqlx.Tx, data models.UpdateStudentParentProfile) *constants.ErrorResponse
	GetPreHighSchoolHistoryByStudentIds(ctx context.Context, tx *sqlx.Tx, studentIds []string) ([]models.GetStudentPreHighSchoolHistory, *constants.ErrorResponse)
	UpdateSchoolProfile(ctx context.Context, tx *sqlx.Tx, data models.UpdateStudentSchoolProfile) *constants.ErrorResponse
	DeletePreHighSchoolHistoryExcludingLevel(ctx context.Context, tx *sqlx.Tx, studentId string, excludedLevel []string) *constants.ErrorResponse
	UpsertPreHighSchoolHistory(ctx context.Context, tx *sqlx.Tx, data []models.UpsertStudentPreHighSchoolHistory) *constants.ErrorResponse
	GetStudentSubject(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studentId string) ([]models.GetStudentSubject, common.Pagination, *constants.ErrorResponse)
	UpdatePayment(ctx context.Context, tx *sqlx.Tx, studentIds []string, adminId string) *constants.ErrorResponse
	GetPaymentLog(ctx context.Context, tx *sqlx.Tx, studentId string) ([]models.GetStudentPaymentLog, *constants.ErrorResponse)
	BulkCreate(ctx context.Context, tx *sqlx.Tx, data []models.BulkCreateStudent) *constants.ErrorResponse
	ConvertGrade(ctx context.Context, tx *sqlx.Tx, studentId string, subjectIds []string, data models.ConvertStudentGrade) *constants.ErrorResponse
}

func NewStudentRepository(db *db.DB) StudentRepositoryInterface {
	return &studentRepository{
		db,
	}
}
