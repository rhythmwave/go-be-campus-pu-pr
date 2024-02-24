package study_program

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type RootStudyProgramHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	GetDetail(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewRootStudyProgramHandler(ctx *service.ServiceCtx) RootStudyProgramHandlerInterface {
	return &studyProgramHandler{
		ctx,
	}
}
