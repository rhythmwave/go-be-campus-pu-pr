package accreditation

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminAccreditationHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewAdminAccreditationHandler(ctx *service.ServiceCtx) AdminAccreditationHandlerInterface {
	return &accreditationHandler{
		ctx,
	}
}
