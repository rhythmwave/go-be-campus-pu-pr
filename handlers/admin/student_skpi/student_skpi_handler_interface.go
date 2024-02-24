package student_skpi

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminStudentSkpiHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	GetDetail(w http.ResponseWriter, r *http.Request)
	Upsert(w http.ResponseWriter, r *http.Request)
	Approve(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewAdminStudentSkpiHandler(ctx *service.ServiceCtx) AdminStudentSkpiHandlerInterface {
	return &studentSkpiHandler{
		ctx,
	}
}
