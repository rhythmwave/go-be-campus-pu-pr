package student_leave

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type StudentStudentLeaveHandlerInterface interface {
	GetListRequests(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

func NewStudentStudentLeaveHandler(ctx *service.ServiceCtx) StudentStudentLeaveHandlerInterface {
	return &studentLeaveHandler{
		ctx,
	}
}
