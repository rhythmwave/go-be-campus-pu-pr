package auth

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type GeneralAuthHandlerInterface interface {
	Login(w http.ResponseWriter, r *http.Request)
	RefreshToken(w http.ResponseWriter, r *http.Request)
	UpdatePassword(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	GetSsoAuth(w http.ResponseWriter, r *http.Request)
	LoginWithSso(w http.ResponseWriter, r *http.Request)
}

func NewGeneralAuthHandler(ctx *service.ServiceCtx) GeneralAuthHandlerInterface {
	return &authHandler{
		ctx,
	}
}
