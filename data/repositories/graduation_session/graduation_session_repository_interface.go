package graduation_session

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type GraduationSessionRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest) ([]models.GetGraduationSession, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetGraduationSession, *constants.ErrorResponse)
	GetUpcoming(ctx context.Context, tx *sqlx.Tx) (models.GetGraduationSession, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateGraduationSession) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateGraduationSession) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
}

func NewGraduationSessionRepository(db *db.DB) GraduationSessionRepositoryInterface {
	return &graduationSessionRepository{
		db,
	}
}
