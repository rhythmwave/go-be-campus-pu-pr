package student_skpi

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type StudentSkpiRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, paramsData objects.GetStudentSkpiRequest) ([]models.GetStudentSkpi, common.Pagination, *constants.ErrorResponse)
	GetById(ctx context.Context, tx *sqlx.Tx, id string) (models.GetStudentSkpiDetail, *constants.ErrorResponse)
	GetByStudentId(ctx context.Context, tx *sqlx.Tx, studentId string) (models.GetStudentSkpiDetail, *constants.ErrorResponse)
	GetStudentSkpiAchievementByStudentSkpiId(ctx context.Context, tx *sqlx.Tx, studentSkpiId string) ([]models.GetStudentSkpiAchievement, *constants.ErrorResponse)
	GetStudentSkpiOrganizationByStudentSkpiId(ctx context.Context, tx *sqlx.Tx, studentSkpiId string) ([]models.GetStudentSkpiOrganization, *constants.ErrorResponse)
	GetStudentSkpiCertificateByStudentSkpiId(ctx context.Context, tx *sqlx.Tx, studentSkpiId string) ([]models.GetStudentSkpiCertificate, *constants.ErrorResponse)
	GetStudentSkpiCharacterBuildingByStudentSkpiId(ctx context.Context, tx *sqlx.Tx, studentSkpiId string) ([]models.GetStudentSkpiCharacterBuilding, *constants.ErrorResponse)
	GetStudentSkpiInternshipByStudentSkpiId(ctx context.Context, tx *sqlx.Tx, studentSkpiId string) ([]models.GetStudentSkpiInternship, *constants.ErrorResponse)
	GetStudentSkpiLanguageByStudentSkpiId(ctx context.Context, tx *sqlx.Tx, studentSkpiId string) ([]models.GetStudentSkpiLanguage, *constants.ErrorResponse)
	UpsertStudentSkpi(ctx context.Context, tx *sqlx.Tx, data models.UpsertStudentSkpi) (string, *constants.ErrorResponse)
	DeleteStudentSkpiAchievementExcludingName(ctx context.Context, tx *sqlx.Tx, studentSkpiId string, excludedName []string) *constants.ErrorResponse
	DeleteStudentSkpiOrganizationExcludingName(ctx context.Context, tx *sqlx.Tx, studentSkpiId string, excludedName []string) *constants.ErrorResponse
	DeleteStudentSkpiCertificateExcludingName(ctx context.Context, tx *sqlx.Tx, studentSkpiId string, excludedName []string) *constants.ErrorResponse
	DeleteStudentSkpiCharacterBuildingExcludingName(ctx context.Context, tx *sqlx.Tx, studentSkpiId string, excludedName []string) *constants.ErrorResponse
	DeleteStudentSkpiInternshipExcludingName(ctx context.Context, tx *sqlx.Tx, studentSkpiId string, excludedName []string) *constants.ErrorResponse
	DeleteStudentSkpiLanguageExcludingName(ctx context.Context, tx *sqlx.Tx, studentSkpiId string, excludedName []string) *constants.ErrorResponse
	UpsertStudentSkpiAchievement(ctx context.Context, tx *sqlx.Tx, data []models.UpsertStudentSkpiAchievement) *constants.ErrorResponse
	UpsertStudentSkpiOrganization(ctx context.Context, tx *sqlx.Tx, data []models.UpsertStudentSkpiOrganization) *constants.ErrorResponse
	UpsertStudentSkpiCertificate(ctx context.Context, tx *sqlx.Tx, data []models.UpsertStudentSkpiCertificate) *constants.ErrorResponse
	UpsertStudentSkpiCharacterBuilding(ctx context.Context, tx *sqlx.Tx, data []models.UpsertStudentSkpiCharacterBuilding) *constants.ErrorResponse
	UpsertStudentSkpiInternship(ctx context.Context, tx *sqlx.Tx, data []models.UpsertStudentSkpiInternship) *constants.ErrorResponse
	UpsertStudentSkpiLanguage(ctx context.Context, tx *sqlx.Tx, data []models.UpsertStudentSkpiLanguage) *constants.ErrorResponse
	Approve(ctx context.Context, tx *sqlx.Tx, data models.ApproveStudentSkpi) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
}

func NewStudentSkpiRepository(db *db.DB) StudentSkpiRepositoryInterface {
	return &studentSkpiRepository{
		db,
	}

}
