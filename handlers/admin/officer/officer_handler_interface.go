package officer

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminOfficerHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewAdminOfficerHandler(ctx *service.ServiceCtx) AdminOfficerHandlerInterface {
	return &officerHandler{
		ctx,
	}
}
