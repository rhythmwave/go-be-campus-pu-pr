package subject_grade_component

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
)

type subjectGradeComponentRepository struct {
	*db.DB
}

func (s subjectGradeComponentRepository) GetBySubjectId(ctx context.Context, tx *sqlx.Tx, subjectId string) ([]models.GetSubjectGradeComponent, *constants.ErrorResponse) {
	results := []models.GetSubjectGradeComponent{}

	err := tx.SelectContext(
		ctx,
		&results,
		getBySubjectIdQuery,
		subjectId,
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (s subjectGradeComponentRepository) Upsert(ctx context.Context, tx *sqlx.Tx, data []models.CreateSubjectGradeComponent) *constants.ErrorResponse {
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

func (s subjectGradeComponentRepository) DeleteAllBySubjectIdExcludingNames(ctx context.Context, tx *sqlx.Tx, subjectId string, names []string) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		deleteAllBySubjectIdExcludingNamesQuery,
		subjectId,
		pq.Array(names),
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (s subjectGradeComponentRepository) DeleteAllBySubjectId(ctx context.Context, tx *sqlx.Tx, subjectId string) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		deleteAllBySubjectIdQuery,
		subjectId,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}
