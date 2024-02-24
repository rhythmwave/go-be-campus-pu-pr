package academic_guidance

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type AcademicGuidanceRepositoryInterface interface {
	GetListStudent(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, lecturerId, semesterId string) ([]models.GetAcademicGuidanceStudentList, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetAcademicGuidanceDetail, *constants.ErrorResponse)
	GetDetailBySemesterIdLecturerId(ctx context.Context, tx *sqlx.Tx, semesterId, lecturerId string) (models.GetAcademicGuidanceDetail, *constants.ErrorResponse)
	GetDetailBySemesterIdStudentId(ctx context.Context, tx *sqlx.Tx, semesterId, studentId string) (models.GetAcademicGuidanceDetail, *constants.ErrorResponse)
	Upsert(ctx context.Context, tx *sqlx.Tx, data models.UpsertAcademicGuidance) (string, *constants.ErrorResponse)
	UpsertStudent(ctx context.Context, tx *sqlx.Tx, data []models.UpsertAcademicGuidanceStudent) *constants.ErrorResponse
	UpsertDecision(ctx context.Context, tx *sqlx.Tx, data models.UpsertDecisionAcademicGuidance) *constants.ErrorResponse
	GetSession(ctx context.Context, tx *sqlx.Tx, academicGuidanceId, semesterId, lecturerId string) ([]models.GetAcademicGuidanceSession, *constants.ErrorResponse)
	GetSessionById(ctx context.Context, tx *sqlx.Tx, id string) (models.GetAcademicGuidanceSession, *constants.ErrorResponse)
	GetSessionStudent(ctx context.Context, tx *sqlx.Tx, academicGuidanceSessionIds []string) ([]models.GetAcademicGuidanceSessionStudent, *constants.ErrorResponse)
	GetSessionFile(ctx context.Context, tx *sqlx.Tx, academicGuidanceSessionIds []string) ([]models.GetAcademicGuidanceSessionFile, *constants.ErrorResponse)
	UpsertSession(ctx context.Context, tx *sqlx.Tx, data models.UpsertAcademicGuidanceSession) *constants.ErrorResponse
	DeleteSessionStudentExcludingStudentIds(ctx context.Context, tx *sqlx.Tx, academicGuidanceSessionId string, studentIds []string) *constants.ErrorResponse
	DeleteSessionFileExcludingFilePaths(ctx context.Context, tx *sqlx.Tx, academicGuidanceSessionId string, filePaths []string) *constants.ErrorResponse
	UpsertSessionStudent(ctx context.Context, tx *sqlx.Tx, data []models.UpsertAcademicGuidanceSessionStudent) *constants.ErrorResponse
	UpsertSessionFile(ctx context.Context, tx *sqlx.Tx, data []models.UpsertAcademicGuidanceSessionFile) *constants.ErrorResponse
	DeleteSession(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
}

func NewAcademicGuidanceRepository(db *db.DB) AcademicGuidanceRepositoryInterface {
	return &academicGuidanceRepository{
		db,
	}
}
