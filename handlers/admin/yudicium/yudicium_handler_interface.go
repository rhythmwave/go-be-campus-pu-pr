package yudicium

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminYudiciumHandlerInterface interface {
	Apply(w http.ResponseWriter, r *http.Request)
	GetListStudent(w http.ResponseWriter, r *http.Request)
	Do(w http.ResponseWriter, r *http.Request)
}

func NewAdminYudiciumHandler(ctx *service.ServiceCtx) AdminYudiciumHandlerInterface {
	return &yudiciumHandler{
		ctx,
	}
}
