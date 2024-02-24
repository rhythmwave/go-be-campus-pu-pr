package lecturer_mutation

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminLecturerMutationHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

func NewAdminLecturerMutationHandler(ctx *service.ServiceCtx) AdminLecturerMutationHandlerInterface {
	return &lecturerMutationHandler{
		ctx,
	}
}
