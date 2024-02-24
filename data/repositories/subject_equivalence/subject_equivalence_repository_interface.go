package subject_equivalence

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
)

type SubjectEquivalenceRepositoryInterface interface {
	Upsert(ctx context.Context, tx *sqlx.Tx, data models.CreateSubjectEquivalence) *constants.ErrorResponse
	DeleteBySubjectIdAndEquivalentSubjectId(ctx context.Context, tx *sqlx.Tx, subjectId, equivalentSubjectId string) *constants.ErrorResponse
}

func NewSubjectEquivalenceRepository(db *db.DB) SubjectEquivalenceRepositoryInterface {
	return &subjectEquivalenceRepository{
		db,
	}
}
