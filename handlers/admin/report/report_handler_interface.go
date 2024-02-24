package report

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminReportHandlerInterface interface {
	StudentStatus(w http.ResponseWriter, r *http.Request)
	StudentClassGrade(w http.ResponseWriter, r *http.Request)
	StudentProvince(w http.ResponseWriter, r *http.Request)
	StudentSchoolProvince(w http.ResponseWriter, r *http.Request)
	GpaDistribution(w http.ResponseWriter, r *http.Request)
	StudyDurationDistribution(w http.ResponseWriter, r *http.Request)
	ThesisDurationDistribution(w http.ResponseWriter, r *http.Request)
	StudentStatusSummary(w http.ResponseWriter, r *http.Request)
	SubjectSummary(w http.ResponseWriter, r *http.Request)
}

func NewAdminReportHandler(ctx *service.ServiceCtx) AdminReportHandlerInterface {
	return &reportHandler{
		ctx,
	}
}
