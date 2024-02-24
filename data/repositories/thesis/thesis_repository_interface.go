package thesis

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type ThesisRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studyProgramId string, nimNumber int64, startSemesterId, status, supervisorLecturerId string) ([]models.GetListThesis, common.Pagination, *constants.ErrorResponse)
	GetById(ctx context.Context, tx *sqlx.Tx, id string) (models.GetDetailThesis, *constants.ErrorResponse)
	GetByStudentIdStatus(ctx context.Context, tx *sqlx.Tx, studentId, status string) (models.GetDetailThesis, *constants.ErrorResponse)
	GetNonCancelled(ctx context.Context, tx *sqlx.Tx, studentId string) (models.GetDetailThesis, *constants.ErrorResponse)
	GetFileByThesisId(ctx context.Context, tx *sqlx.Tx, thesisId string) ([]models.GetThesisFile, *constants.ErrorResponse)
	GetSupervisorByThesisId(ctx context.Context, tx *sqlx.Tx, thesisId string) ([]models.GetThesisSupervisor, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateThesis) *constants.ErrorResponse
	UpsertFile(ctx context.Context, tx *sqlx.Tx, data []models.UpsertThesisFile) *constants.ErrorResponse
	UpsertSupervisor(ctx context.Context, tx *sqlx.Tx, data []models.UpsertThesisSupervisor) *constants.ErrorResponse
	DeleteFileExcludingPaths(ctx context.Context, tx *sqlx.Tx, thesisId string, excludedPaths []string) *constants.ErrorResponse
	DeleteSupervisorExcludingLecturerIds(ctx context.Context, tx *sqlx.Tx, thesisId string, excludedLecturerIds []string) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateThesis) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
	CreateDefenseRequest(ctx context.Context, tx *sqlx.Tx, thesisId string) *constants.ErrorResponse
	GetListDefenseRequest(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studyProgramId string, nimNumber int64, startSemesterId string) ([]models.GetThesisDefenseRequest, common.Pagination, *constants.ErrorResponse)
	GetActiveDefenseRequest(ctx context.Context, tx *sqlx.Tx, thesisId string) (models.GetThesisDefenseRequest, *constants.ErrorResponse)
	GetDefenseById(ctx context.Context, tx *sqlx.Tx, id string) (models.GetThesisDefense, *constants.ErrorResponse)
	CreateDefense(ctx context.Context, tx *sqlx.Tx, data models.CreateThesisDefense) *constants.ErrorResponse
	DeleteExaminerExcludingLecturerIds(ctx context.Context, tx *sqlx.Tx, thesisDefenseId string, excludedLecturerIds []string) *constants.ErrorResponse
	UpsertDefenseExaminer(ctx context.Context, tx *sqlx.Tx, data []models.UpsertThesisDefenseExaminer) *constants.ErrorResponse
	UpdateDefense(ctx context.Context, tx *sqlx.Tx, data models.UpdateThesisDefense) *constants.ErrorResponse
	FinishDefense(ctx context.Context, tx *sqlx.Tx, data models.FinishThesisDefense) *constants.ErrorResponse
	GetThesisDefenseExaminerByThesisDefenseIds(ctx context.Context, tx *sqlx.Tx, thesisDefenseIds []string) ([]models.GetThesisDefenseExaminer, *constants.ErrorResponse)
	GetActiveSemesterThesisSupervisorLog(ctx context.Context, tx *sqlx.Tx, lecturerIds []string) ([]models.GetThesisSupervisorLog, *constants.ErrorResponse)
	GetThesisSupervisorLog(ctx context.Context, tx *sqlx.Tx, semesterId string, lecturerIds []string) ([]models.GetThesisSupervisorLog, *constants.ErrorResponse)
}

func NewThesisRepository(db *db.DB) ThesisRepositoryInterface {
	return &thesisRepository{
		db,
	}
}
