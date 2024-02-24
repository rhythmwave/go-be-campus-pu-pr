package class_work

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type StudentClassWorkHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Submit(w http.ResponseWriter, r *http.Request)
}

func NewStudentClassWorkHandler(ctx *service.ServiceCtx) StudentClassWorkHandlerInterface {
	return &classWorkHandler{
		ctx,
	}
}
