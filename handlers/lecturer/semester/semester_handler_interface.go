package semester

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type LecturerSemesterHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
}

func NewLecturerSemesterHandler(ctx *service.ServiceCtx) LecturerSemesterHandlerInterface {
	return &semesterHandler{
		ctx,
	}
}
