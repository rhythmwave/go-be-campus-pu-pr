package study_program

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminStudyProgramHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	GetDetail(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	UpdateDegree(w http.ResponseWriter, r *http.Request)
}

func NewAdminStudyProgramHandler(ctx *service.ServiceCtx) AdminStudyProgramHandlerInterface {
	return &studyProgramHandler{
		ctx,
	}
}
