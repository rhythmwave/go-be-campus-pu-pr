package announcement

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type AnnouncementRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, appType, announcementType string, ids []string) ([]models.GetAnnouncement, common.Pagination, *constants.ErrorResponse)
	GetAnnouncementStudyProgramByAnnouncementIds(ctx context.Context, tx *sqlx.Tx, announcementIds []string) ([]models.GetAnnouncementStudyProgram, *constants.ErrorResponse)
	GetAnnouncementStudyProgramByStudyProgramIds(ctx context.Context, tx *sqlx.Tx, studyProgramIds []string) ([]models.GetAnnouncementStudyProgram, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetAnnouncement, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateAnnouncement) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateAnnouncement) *constants.ErrorResponse
	DeleteAnnouncementStudyProgramExcludingStudyProgramIds(ctx context.Context, tx *sqlx.Tx, announcementId string, excludedStudyProgramIds []string) *constants.ErrorResponse
	UpsertAnnouncementStudyProgram(ctx context.Context, tx *sqlx.Tx, data []models.UpsertAnnouncementStudyProgram) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
}

func NewAnnouncementRepository(db *db.DB) AnnouncementRepositoryInterface {
	return &announcementRepository{
		db,
	}
}
