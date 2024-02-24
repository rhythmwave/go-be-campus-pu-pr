package admin

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type RootAdminHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewRootAdminHandler(ctx *service.ServiceCtx) RootAdminHandlerInterface {
	return &adminHandler{
		ctx,
	}
}
