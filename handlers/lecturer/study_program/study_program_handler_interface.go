package study_program

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type LecturerStudyProgramHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
}

func NewLecturerStudyProgramHandler(ctx *service.ServiceCtx) LecturerStudyProgramHandlerInterface {
	return &studyProgramHandler{
		ctx,
	}
}
