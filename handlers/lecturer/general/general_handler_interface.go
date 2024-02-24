package general

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type LecturerGeneralHandlerInterface interface {
	GetSemesterSummary(w http.ResponseWriter, r *http.Request)
	GetProfile(w http.ResponseWriter, r *http.Request)
}

func NewLecturerGeneralHandler(ctx *service.ServiceCtx) LecturerGeneralHandlerInterface {
	return &lecturerGeneralHandler{
		ctx,
	}
}
