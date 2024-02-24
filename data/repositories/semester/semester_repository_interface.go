package semester

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type SemesterRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studyProgramId, excludedId string) ([]models.GetSemester, common.Pagination, *constants.ErrorResponse)
	GetById(ctx context.Context, tx *sqlx.Tx, id string) (models.GetSemesterDetail, *constants.ErrorResponse)
	GetActive(ctx context.Context, tx *sqlx.Tx) (models.GetSemesterDetail, *constants.ErrorResponse)
	GetCurriculumBySemesterIds(ctx context.Context, tx *sqlx.Tx, semesterIds []string) ([]models.GetSemesterCurriculum, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateSemester) (string, *constants.ErrorResponse)
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateSemester) *constants.ErrorResponse
	DeleteCurriculumSemesterExcludingCurriculumId(ctx context.Context, tx *sqlx.Tx, semesterId string, excludedCurriculumIds []string) *constants.ErrorResponse
	UpsertCurriculum(ctx context.Context, tx *sqlx.Tx, data []models.UpsertSemesterCurriculum) *constants.ErrorResponse
	UpdateActivation(ctx context.Context, tx *sqlx.Tx, id string, isActive bool) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
	AutoSetActive(ctx context.Context, tx *sqlx.Tx) (string, *constants.ErrorResponse)
	GetPreviousSemester(ctx context.Context, tx *sqlx.Tx, semesterId string) (models.GetSemesterDetail, *constants.ErrorResponse)
}

func NewSemesterRepository(db *db.DB) SemesterRepositoryInterface {
	return &semesterRepository{
		db,
	}
}
