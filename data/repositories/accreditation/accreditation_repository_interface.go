package accreditation

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type AccreditationRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studyProgramId string) ([]models.GetAccreditation, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetAccreditation, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateAccreditation) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateAccreditation) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
}

func NewAccreditationRepository(db *db.DB) AccreditationRepositoryInterface {
	return &accreditationRepository{
		db,
	}
}
