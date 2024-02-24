package subject_prerequisite

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
)

type subjectPrerequisiteRepository struct {
	*db.DB
}

func (s subjectPrerequisiteRepository) GetBySubjectId(ctx context.Context, tx *sqlx.Tx, subjectId string) ([]models.GetSubjectPrerequisite, *constants.ErrorResponse) {
	results := []models.GetSubjectPrerequisite{}

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

func (s subjectPrerequisiteRepository) Upsert(ctx context.Context, tx *sqlx.Tx, data []models.CreateSubjectPrerequisite) *constants.ErrorResponse {
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

func (s subjectPrerequisiteRepository) DeleteAllBySubjectIdExcludingPrerequisiteSubjectId(ctx context.Context, tx *sqlx.Tx, subjectId string, prerequisiteSubjectIds []string) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		deleteAllBySubjectIdExcludingPrerequisiteSubjectIdQuery,
		subjectId,
		pq.Array(prerequisiteSubjectIds),
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (s subjectPrerequisiteRepository) DeleteAllBySubjectId(ctx context.Context, tx *sqlx.Tx, subjectId string) *constants.ErrorResponse {
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
