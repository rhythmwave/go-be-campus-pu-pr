package excel

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
)

type excelRepository struct {
	*db.DB
}

func (f excelRepository) StudyProgramDistributionDownload(ctx context.Context, tx *sqlx.Tx, studyProgramIds []string, semesterID string) ([]models.StudyProgramDistributionDownload, *constants.ErrorResponse) {
	var resultData []models.StudyProgramDistributionDownload

	err := tx.SelectContext(
		ctx,
		&resultData,
		getStudyProgramDistributionData,
		pq.Array(studyProgramIds),
		semesterID,
	)
	if err != nil {
		return []models.StudyProgramDistributionDownload{}, constants.ErrorInternalServer(err.Error())
	}

	return resultData, nil
}
