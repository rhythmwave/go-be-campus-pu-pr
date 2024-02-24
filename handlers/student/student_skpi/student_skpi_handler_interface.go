package student_skpi

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type StudentStudentSkpiHandlerInterface interface {
	GetDetail(w http.ResponseWriter, r *http.Request)
	Upsert(w http.ResponseWriter, r *http.Request)
}

func NewStudentStudentSkpiHandler(ctx *service.ServiceCtx) StudentStudentSkpiHandlerInterface {
	return &studentSkpiHandler{
		ctx,
	}
}
