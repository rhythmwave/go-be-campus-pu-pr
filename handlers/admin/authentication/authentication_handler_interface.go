package authentication

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminAuthenticationHandlerInterface interface {
	GetDetail(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	BulkCreate(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	UpdatePassword(w http.ResponseWriter, r *http.Request)
}

func NewAdminAuthentication(ctx *service.ServiceCtx) AdminAuthenticationHandlerInterface {
	return &authenticationHandler{
		ctx,
	}
}
