package transcript

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type StudentTranscriptHandlerInterface interface {
	GetDetail(w http.ResponseWriter, r *http.Request)
}

func NewStudentTranscriptHandler(ctx *service.ServiceCtx) StudentTranscriptHandlerInterface {
	return &studentTranscriptHandler{
		ctx,
	}
}
