package class_exam

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type StudentClassExamHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Submit(w http.ResponseWriter, r *http.Request)
}

func NewStudentClassExamHandler(ctx *service.ServiceCtx) StudentClassExamHandlerInterface {
	return &classExamHandler{
		ctx,
	}
}
