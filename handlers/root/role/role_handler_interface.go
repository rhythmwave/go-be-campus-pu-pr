package role

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type RootRoleHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewRootRoleHandler(ctx *service.ServiceCtx) RootRoleHandlerInterface {
	return &roleHandler{
		ctx,
	}
}
