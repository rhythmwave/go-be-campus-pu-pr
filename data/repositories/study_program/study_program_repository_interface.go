package study_program

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type StudyProgramRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, majorId, appType, userId string) ([]models.GetStudyProgramList, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id, appType, userId string) (models.GetStudyProgramDetail, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateStudyProgram) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateStudyProgram) *constants.ErrorResponse
	UpdateDegree(ctx context.Context, tx *sqlx.Tx, data models.UpdateDegreeStudyProgram) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
	GetByRoleIds(ctx context.Context, tx *sqlx.Tx, roleIds []string) ([]models.GetStudyProgramByRoleIds, *constants.ErrorResponse)
}

func NewStudyProgramRepository(db *db.DB) StudyProgramRepositoryInterface {
	return &studyProgramRepository{
		db,
	}
}
