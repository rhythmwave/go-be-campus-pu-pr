package subject_equivalence

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
)

type subjectEquivalenceRepository struct {
	*db.DB
}

func (s subjectEquivalenceRepository) Upsert(ctx context.Context, tx *sqlx.Tx, data models.CreateSubjectEquivalence) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		upsertQuery,
		data.SubjectId,
		data.EquivalentSubjectId,
		data.CreatedBy,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (s subjectEquivalenceRepository) DeleteBySubjectIdAndEquivalentSubjectId(ctx context.Context, tx *sqlx.Tx, subjectId, equivalentSubjectId string) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		deleteBySubjectIdAndEquivalentSubjectIdQuery,
		subjectId,
		equivalentSubjectId,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}
