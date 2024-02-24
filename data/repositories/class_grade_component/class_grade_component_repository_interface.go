package class_grade_component

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
)

type ClassGradeComponentRepositoryInterface interface {
	GetByClassId(ctx context.Context, tx *sqlx.Tx, classId string) ([]models.GetClassGradeComponent, *constants.ErrorResponse)
	Upsert(ctx context.Context, tx *sqlx.Tx, data []models.CreateClassGradeComponent) *constants.ErrorResponse
	DeleteAllByClassIdExcludingNames(ctx context.Context, tx *sqlx.Tx, classId string, names []string) *constants.ErrorResponse
	DeleteAllByClassId(ctx context.Context, tx *sqlx.Tx, classId string) *constants.ErrorResponse
}

func NewClassGradeComponentRepository(db *db.DB) ClassGradeComponentRepositoryInterface {
	return &classGradeComponentRepository{
		db,
	}
}
