package yudicium_session

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminYudiciumSessionHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewAdminYudiciumSessionHandler(ctx *service.ServiceCtx) AdminYudiciumSessionHandlerInterface {
	return &yudiciumSessionHandler{
		ctx,
	}
}
