package subject_grade_component

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
)

type SubjectGradeComponentRepositoryInterface interface {
	GetBySubjectId(ctx context.Context, tx *sqlx.Tx, subjectId string) ([]models.GetSubjectGradeComponent, *constants.ErrorResponse)
	Upsert(ctx context.Context, tx *sqlx.Tx, data []models.CreateSubjectGradeComponent) *constants.ErrorResponse
	DeleteAllBySubjectIdExcludingNames(ctx context.Context, tx *sqlx.Tx, subjectId string, names []string) *constants.ErrorResponse
	DeleteAllBySubjectId(ctx context.Context, tx *sqlx.Tx, subjectId string) *constants.ErrorResponse
}

func NewSubjectGradeComponentRepository(db *db.DB) SubjectGradeComponentRepositoryInterface {
	return &subjectGradeComponentRepository{
		db,
	}
}
