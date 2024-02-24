package lecturer_leave

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminLecturerLeaveHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	End(w http.ResponseWriter, r *http.Request)
}

func NewAdminLecturerLeaveHandler(ctx *service.ServiceCtx) AdminLecturerLeaveHandlerInterface {
	return &lecturerLeaveHandler{
		ctx,
	}
}
