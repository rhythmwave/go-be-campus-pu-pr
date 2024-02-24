package graduation_student

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type GraduationStudentRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studyProgramId, graduationSessionId string) ([]models.GetListStudentGraduation, common.Pagination, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateGraduationStudent) *constants.ErrorResponse
}

func NewGraduationStudentRepository(db *db.DB) GraduationStudentRepositoryInterface {
	return &graduationStudentRepository{
		db,
	}
}
