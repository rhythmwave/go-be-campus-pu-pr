package class_event

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type LecturerClassEventHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	BulkUpdateActivation(w http.ResponseWriter, r *http.Request)
	BulkDelete(w http.ResponseWriter, r *http.Request)
}

func NewLecturerClassEventHandler(ctx *service.ServiceCtx) LecturerClassEventHandlerInterface {
	return &classEventHandler{
		ctx,
	}
}
