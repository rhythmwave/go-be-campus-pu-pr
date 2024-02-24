package class_grade_component

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
)

type classGradeComponentRepository struct {
	*db.DB
}

func (s classGradeComponentRepository) GetByClassId(ctx context.Context, tx *sqlx.Tx, classId string) ([]models.GetClassGradeComponent, *constants.ErrorResponse) {
	results := []models.GetClassGradeComponent{}

	err := tx.SelectContext(
		ctx,
		&results,
		getByClassIdQuery,
		classId,
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (s classGradeComponentRepository) Upsert(ctx context.Context, tx *sqlx.Tx, data []models.CreateClassGradeComponent) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		upsertQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (s classGradeComponentRepository) DeleteAllByClassIdExcludingNames(ctx context.Context, tx *sqlx.Tx, classId string, names []string) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		deleteAllByClassIdExcludingNamesQuery,
		classId,
		pq.Array(names),
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (s classGradeComponentRepository) DeleteAllByClassId(ctx context.Context, tx *sqlx.Tx, classId string) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		deleteAllByClassIdQuery,
		classId,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}
