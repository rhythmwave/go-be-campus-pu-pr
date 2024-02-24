package class_lecturer

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
)

type classLecturerRepository struct {
	*db.DB
}

func (c classLecturerRepository) GetByClassIdLecturerId(ctx context.Context, tx *sqlx.Tx, classId string, lecturerId string) (models.GetClassLecturer, *constants.ErrorResponse) {
	results := []models.GetClassLecturer{}

	err := tx.SelectContext(
		ctx,
		&results,
		getByClassIdLecturerIdQuery,
		classId,
		lecturerId,
	)
	if err != nil {
		return models.GetClassLecturer{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetClassLecturer{}, nil
	}

	return results[0], nil
}
