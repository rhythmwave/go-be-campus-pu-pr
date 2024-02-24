package yudicium_term

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminYudiciumTermHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewAdminYudiciumTermHandler(ctx *service.ServiceCtx) AdminYudiciumTermHandlerInterface {
	return &yudiciumTermHandler{
		ctx,
	}
}
