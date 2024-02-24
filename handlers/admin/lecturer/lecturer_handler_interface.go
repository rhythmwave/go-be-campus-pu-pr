package lecturer

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminLecturerHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	GetDetail(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetSchedule(w http.ResponseWriter, r *http.Request)
	GetAssignedClass(w http.ResponseWriter, r *http.Request)
}

func NewAdminLecturerHandler(ctx *service.ServiceCtx) AdminLecturerHandlerInterface {
	return &lecturerHandler{
		ctx,
	}
}
