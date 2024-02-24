package curriculum

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type CurriculumRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studyProgramIds []string) ([]models.GetCurriculum, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetCurriculumDetail, *constants.ErrorResponse)
	GetActiveByStudyProgramId(ctx context.Context, tx *sqlx.Tx, studyProgramId string) (models.GetCurriculumDetail, *constants.ErrorResponse)
	GetActive(ctx context.Context, tx *sqlx.Tx) ([]models.GetCurriculumDetail, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateCurriculum) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateCurriculum) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
}

func NewCurriculumRepository(db *db.DB) CurriculumRepositoryInterface {
	return &curriculumRepository{
		db,
	}
}
