package class_work

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminClassWorkHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	GetSubmission(w http.ResponseWriter, r *http.Request)
}

func NewAdminClassWorkHandler(ctx *service.ServiceCtx) AdminClassWorkHandlerInterface {
	return &classWorkHandler{
		ctx,
	}
}
