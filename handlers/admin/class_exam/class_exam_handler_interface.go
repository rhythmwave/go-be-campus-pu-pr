package class_exam

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminClassExamHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	GetSubmission(w http.ResponseWriter, r *http.Request)
}

func NewAdminClassExamHandler(ctx *service.ServiceCtx) AdminClassExamHandlerInterface {
	return &classExamHandler{
		ctx,
	}
}
