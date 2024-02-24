package graduation_predicate

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type GraduationPredicateRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest) ([]models.GetGraduationPredicate, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetGraduationPredicate, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateGraduationPredicate) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateGraduationPredicate) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
}

func NewGraduationPredicateRepository(db *db.DB) GraduationPredicateRepositoryInterface {
	return &graduationPredicateRepository{
		db,
	}
}
