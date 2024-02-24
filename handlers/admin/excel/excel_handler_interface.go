package excel

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type ExcelHandlerInterface interface {
	StudyProgramDistributionDownload(w http.ResponseWriter, r *http.Request)
}

func NewExcelHandler(ctx *service.ServiceCtx) ExcelHandlerInterface {
	return &excelHandler{
		ctx,
	}
}
