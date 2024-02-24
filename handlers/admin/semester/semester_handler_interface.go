package semester

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminSemesterHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	GetDetail(w http.ResponseWriter, r *http.Request)
	GetActive(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	UpdateActivation(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewAdminSemesterHandler(ctx *service.ServiceCtx) AdminSemesterHandlerInterface {
	return &semesterHandler{
		ctx,
	}
}
