package study_plan

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type StudentStudyPlanHandlerInterface interface {
	GetDetail(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

func NewStudentStudyPlanHandler(ctx *service.ServiceCtx) StudentStudyPlanHandlerInterface {
	return &studyPlanHandler{
		ctx,
	}
}
