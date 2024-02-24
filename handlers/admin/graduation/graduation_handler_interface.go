package graduation

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminGraduationHandlerInterface interface {
	Apply(w http.ResponseWriter, r *http.Request)
	GetListStudent(w http.ResponseWriter, r *http.Request)
}

func NewAdminGraduationHandler(ctx *service.ServiceCtx) AdminGraduationHandlerInterface {
	return &graduationHandler{
		ctx,
	}
}
