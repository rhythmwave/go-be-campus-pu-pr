package grade_type

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type GradeTypeRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studyLevelId string) ([]models.GetGradeType, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetGradeType, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateGradeType) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateGradeType) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
	GetByGradeCode(ctx context.Context, tx *sqlx.Tx, studyLevelId string, gradeCode string) (models.GetGradeType, *constants.ErrorResponse)
}

func NewGradeTypeRepository(db *db.DB) GradeTypeRepositoryInterface {
	return &gradeTypeRepository{
		db,
	}
}
