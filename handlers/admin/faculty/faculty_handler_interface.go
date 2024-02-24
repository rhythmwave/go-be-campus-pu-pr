package faculty

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminFacultyHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
}

func NewAdminFacultyHandler(ctx *service.ServiceCtx) AdminFacultyHandlerInterface {
	return &facultyHandler{
		ctx,
	}
}
