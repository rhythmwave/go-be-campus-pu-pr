package study_program

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type CareerStudyProgramHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	GetDetail(w http.ResponseWriter, r *http.Request)
}

func NewCareerStudyProgramHandler(ctx *service.ServiceCtx) CareerStudyProgramHandlerInterface {
	return &studyProgramHandler{
		ctx,
	}
}
