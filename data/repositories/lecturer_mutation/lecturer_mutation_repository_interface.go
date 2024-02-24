package lecturer_mutation

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type LecturerMutationRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studyProgramId, idNationalLecturer, semesterId string) ([]models.GetLecturerMutation, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetLecturerMutation, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateLecturerMutation) *constants.ErrorResponse
}

func NewLecturerMutationRepository(db *db.DB) LecturerMutationRepositoryInterface {
	return &lecturerMutationRepository{
		db,
	}
}
