package student_class

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type StudentClassRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studyPlanId, studentId, semesterId string, isMbkm *bool) ([]models.GetStudentClass, common.Pagination, *constants.ErrorResponse)
	GetClassParticipant(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, classIds []string, lectureId string, isGraded *bool, studentId string) ([]models.GetClassParticipant, common.Pagination, *constants.ErrorResponse)
	BulkCreate(ctx context.Context, tx *sqlx.Tx, data []models.CreateStudentClass) *constants.ErrorResponse
	GetStudentClassByStudentIdClassId(ctx context.Context, tx *sqlx.Tx, data []models.StudentIdClassId) ([]models.GetStudentClass, *constants.ErrorResponse)
	BulkUpdateClass(ctx context.Context, tx *sqlx.Tx, data []models.CreateStudentClass) *constants.ErrorResponse
	BulkGradeStudentClass(ctx context.Context, tx *sqlx.Tx, data []models.GradeStudentClass) *constants.ErrorResponse
	GetStudentClassGradeByClassIdAndStudentIds(ctx context.Context, tx *sqlx.Tx, classId string, studentIds []string) ([]models.GetStudentClassGrade, *constants.ErrorResponse)
	BulkDeleteExcludingClassIds(ctx context.Context, tx *sqlx.Tx, data []models.DeleteStudentClassExcludingClassIds) *constants.ErrorResponse
	UpdateMbkmConvertedCredit(ctx context.Context, tx *sqlx.Tx, id string, convertedCredit uint32) *constants.ErrorResponse
}

func NewStudentClassRepository(db *db.DB) StudentClassRepositoryInterface {
	return &studentClassRepository{
		db,
	}
}
