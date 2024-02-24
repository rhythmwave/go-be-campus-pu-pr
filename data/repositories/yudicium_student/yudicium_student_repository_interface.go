package yudicium_student

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type YudiciumStudentRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, req objects.GetListStudentYudiciumRequest) ([]models.GetListStudentYudicium, common.Pagination, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateYudiciumStudent) *constants.ErrorResponse
	Do(ctx context.Context, tx *sqlx.Tx, yudiciumSessionId string, studentIds []string) *constants.ErrorResponse
}

func NewYudiciumStudentRepository(db *db.DB) YudiciumStudentRepositoryInterface {
	return &yudiciumStudentRepository{
		db,
	}
}
