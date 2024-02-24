package grade_component

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type GradeComponentRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studyProgramId, subjectCategoryId string) ([]models.GetGradeComponent, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetGradeComponent, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateGradeComponent) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateGradeComponent) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
	GetDistinctSubjectCategoryList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studyProgramId string) ([]models.GetGradeComponentDistinctSubjectCategory, common.Pagination, *constants.ErrorResponse)
	GetPercentageBySubjectCategories(ctx context.Context, tx *sqlx.Tx, studyProgramId string, subjectCategoryIds []string) ([]models.GetPercentageBySubjectCategories, *constants.ErrorResponse)
	BulkUpdatePercentage(ctx context.Context, tx *sqlx.Tx, data []models.BulkUpdateGradeComponentPercentage) *constants.ErrorResponse
}

func NewGradeComponentRepository(db *db.DB) GradeComponentRepositoryInterface {
	return &gradeComponentRepository{
		db,
	}
}
