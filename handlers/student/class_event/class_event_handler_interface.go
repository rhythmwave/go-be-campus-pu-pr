package class_event

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type StudentClassEventHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
}

func NewStudentClassEventHandler(ctx *service.ServiceCtx) StudentClassEventHandlerInterface {
	return &classEventHandler{
		ctx,
	}
}
