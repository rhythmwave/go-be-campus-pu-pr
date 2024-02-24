package faculty

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type FacultyRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest) ([]models.GetFacultyList, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetFacultyDetail, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateFaculty) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateFaculty) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
}

func NewFacultyRepository(db *db.DB) FacultyRepositoryInterface {
	return &facultyRepository{
		db,
	}
}
