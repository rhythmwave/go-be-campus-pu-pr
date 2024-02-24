package excel

import (
	"net/http"
	"time"

	"github.com/sccicitb/pupr-backend/infra/context/service"
	"github.com/sccicitb/pupr-backend/utils"
)

type excelHandler struct {
	*service.ServiceCtx
}

func (f excelHandler) StudyProgramDistributionDownload(w http.ResponseWriter, r *http.Request) {
	semesterID := r.URL.Query().Get("semester_id")
	studyProgramIds := r.URL.Query().Get("study_program_id")
	ctx := r.Context()

	data, errs := f.ExcelService.StudyProgramDistributionDownload(ctx, studyProgramIds, semesterID)
	if errs != nil {
		utils.PrintError(*errs)

		utils.JSONResponse(w, errs.HttpCode, map[string]string{
			"message": errs.Err.Error(),
		})
		return
	}

	filename := time.Now().Format("02012006150405") + ".xlsx"
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Description", filename)
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	w.Write(data)

}
