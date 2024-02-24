package subject_prerequisite

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
)

type SubjectPrerequisiteRepositoryInterface interface {
	GetBySubjectId(ctx context.Context, tx *sqlx.Tx, subjectId string) ([]models.GetSubjectPrerequisite, *constants.ErrorResponse)
	Upsert(ctx context.Context, tx *sqlx.Tx, data []models.CreateSubjectPrerequisite) *constants.ErrorResponse
	DeleteAllBySubjectIdExcludingPrerequisiteSubjectId(ctx context.Context, tx *sqlx.Tx, subjectId string, prerequisiteSubjectIds []string) *constants.ErrorResponse
	DeleteAllBySubjectId(ctx context.Context, tx *sqlx.Tx, subjectId string) *constants.ErrorResponse
}

func NewSubjectPrerequisiteRepository(db *db.DB) SubjectPrerequisiteRepositoryInterface {
	return &subjectPrerequisiteRepository{
		db,
	}
}
