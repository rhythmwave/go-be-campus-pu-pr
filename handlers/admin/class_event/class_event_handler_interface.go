package class_event

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminClassEventHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
}

func NewAdminClassEventHandler(ctx *service.ServiceCtx) AdminClassEventHandlerInterface {
	return &classEventHandler{
		ctx,
	}
}
