package excel

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
)

type ExcelServiceInterface interface {
	StudyProgramDistributionDownload(ctx context.Context, studyProgramIds string, semesterID string) ([]byte, *constants.ErrorResponse)
}

func NewExcelServiceInterface(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) ExcelServiceInterface {
	return &excelService{
		repoCtx,
		infraCtx,
	}
}
