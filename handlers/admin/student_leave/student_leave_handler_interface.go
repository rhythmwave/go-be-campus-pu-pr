package student_leave

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminStudentLeaveHandlerInterface interface {
	GetListRequests(w http.ResponseWriter, r *http.Request)
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Approve(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	End(w http.ResponseWriter, r *http.Request)
}

func NewAdminStudentLeaveHandler(ctx *service.ServiceCtx) AdminStudentLeaveHandlerInterface {
	return &studentLeaveHandler{
		ctx,
	}
}
