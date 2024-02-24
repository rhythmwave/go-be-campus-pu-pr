package excel

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
)

type ExcelRepositoryInterface interface {
	StudyProgramDistributionDownload(ctx context.Context, tx *sqlx.Tx, studyProgramIds []string, semesterID string) ([]models.StudyProgramDistributionDownload, *constants.ErrorResponse)
}

func NewExcelRepository(db *db.DB) ExcelRepositoryInterface {
	return &excelRepository{
		db,
	}
}
