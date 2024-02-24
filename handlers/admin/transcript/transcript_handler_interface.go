package transcript

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminTranscriptHandlerInterface interface {
	GetDetail(w http.ResponseWriter, r *http.Request)
}

func NewAdminTranscriptHandler(ctx *service.ServiceCtx) AdminTranscriptHandlerInterface {
	return &adminTranscriptHandler{
		ctx,
	}
}
