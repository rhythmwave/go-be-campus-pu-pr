package study_plan

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminStudyPlanHandlerInterface interface {
	BulkCreate(w http.ResponseWriter, r *http.Request)
	BulkApprove(w http.ResponseWriter, r *http.Request)
	GetList(w http.ResponseWriter, r *http.Request)
}

func NewAdminStudyPlanHandler(ctx *service.ServiceCtx) AdminStudyPlanHandlerInterface {
	return &studyPlanHandler{
		ctx,
	}
}
