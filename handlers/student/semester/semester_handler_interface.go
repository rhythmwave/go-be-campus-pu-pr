package semester

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type StudentSemesterHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
}

func NewStudentSemesterHandler(ctx *service.ServiceCtx) StudentSemesterHandlerInterface {
	return &semesterHandler{
		ctx,
	}
}
