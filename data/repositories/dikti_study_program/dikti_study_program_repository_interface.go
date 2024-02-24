package dikti_study_program

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type DiktiStudyProgramRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest) ([]models.GetDiktiStudyProgramList, common.Pagination, *constants.ErrorResponse)
}

func NewDiktiStudyProgramRepository(db *db.DB) DiktiStudyProgramRepositoryInterface {
	return &diktiStudyProgramRepository{
		db,
	}
}
