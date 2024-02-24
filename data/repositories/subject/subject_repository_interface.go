package subject

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type SubjectRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, req objects.GetSubjectRequest) ([]models.GetSubject, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetSubjectDetail, *constants.ErrorResponse)
	GetDetailByIds(ctx context.Context, tx *sqlx.Tx, ids []string) ([]models.GetSubjectDetail, *constants.ErrorResponse)
	GetThesisByCurriculumId(ctx context.Context, tx *sqlx.Tx, curriculumId string) (models.GetSubjectDetail, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateSubject) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateSubject) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
}

func NewSubjectRepository(db *db.DB) SubjectRepositoryInterface {
	return &subjectRepository{
		db,
	}
}
