package lecturer_resignation

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminLecturerResignationHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

func NewAdminLecturerResignationHandler(ctx *service.ServiceCtx) AdminLecturerResignationHandlerInterface {
	return &lecturerResignationHandler{
		ctx,
	}
}
