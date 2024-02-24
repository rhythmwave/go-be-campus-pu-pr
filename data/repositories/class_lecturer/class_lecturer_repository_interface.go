package class_lecturer

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
)

type ClassLecturerRepositoryInterface interface {
	GetByClassIdLecturerId(ctx context.Context, tx *sqlx.Tx, classId string, lecturerId string) (models.GetClassLecturer, *constants.ErrorResponse)
}

func NewClassLecturerRepository(db *db.DB) ClassLecturerRepositoryInterface {
	return &classLecturerRepository{
		db,
	}
}
