package study_level

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type StudyLevelRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest) ([]models.GetStudyLevel, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetStudyLevel, *constants.ErrorResponse)
	UpdateSkpi(ctx context.Context, tx *sqlx.Tx, data models.UpdateStudyLevelSkpi) *constants.ErrorResponse
}

func NewStudyLevelRepository(db *db.DB) StudyLevelRepositoryInterface {
	return &studyLevelRepository{
		db,
	}
}
