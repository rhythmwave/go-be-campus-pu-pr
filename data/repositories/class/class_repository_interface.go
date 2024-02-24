package class

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type ClassRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, in objects.GetClassListRequest) ([]models.GetClass, common.Pagination, *constants.ErrorResponse)
	GetBySubjectIdsSemesterId(ctx context.Context, tx *sqlx.Tx, subjectIds []string, semesterId string) ([]models.GetClass, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id, lecturerId string) (models.GetClassDetail, *constants.ErrorResponse)
	GetDetailByIds(ctx context.Context, tx *sqlx.Tx, ids []string, lecturerId string) ([]models.GetClassDetail, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateClass) (string, *constants.ErrorResponse)
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateClass) *constants.ErrorResponse
	UpsertMaximumParticipant(ctx context.Context, tx *sqlx.Tx, data []models.CreateClass) *constants.ErrorResponse
	UpdateActivation(ctx context.Context, tx *sqlx.Tx, id string, isActive bool) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
	GetClassLecturerByClassIds(ctx context.Context, tx *sqlx.Tx, classIds []string) ([]models.GetClassLecturer, *constants.ErrorResponse)
	GetClassLecturersBySemesterIdLecturerId(ctx context.Context, tx *sqlx.Tx, semesterId, lecturerId string) ([]models.GetClassLecturer, *constants.ErrorResponse)
	DeleteClassLecturerExcludingLecturerIds(ctx context.Context, tx *sqlx.Tx, classId string, excludedLecturerIds []string) *constants.ErrorResponse
	UpsertClassLecturer(ctx context.Context, tx *sqlx.Tx, data []models.UpsertClassLecturer) *constants.ErrorResponse
	Duplicate(ctx context.Context, tx *sqlx.Tx, fromSemesterId, toSemesterId, adminId string) *constants.ErrorResponse
	DuplicateLecturer(ctx context.Context, tx *sqlx.Tx, fromSemesterId, toSemesterId, adminId string) *constants.ErrorResponse
	GetActiveBySemesterId(ctx context.Context, tx *sqlx.Tx, semesterId string) ([]models.GetClass, *constants.ErrorResponse)
	InactivateClasses(ctx context.Context, tx *sqlx.Tx, classIds []string) *constants.ErrorResponse
	GetThesisClass(ctx context.Context, tx *sqlx.Tx, studentId string, semesterId string) (models.GetClass, *constants.ErrorResponse)
	UpsertThesisClass(ctx context.Context, tx *sqlx.Tx, subjectId, lecturerId, semesterId, adminId string) (string, *constants.ErrorResponse)
}

func NewClassRepository(db *db.DB) ClassRepositoryInterface {
	return &classRepository{
		db,
	}
}
