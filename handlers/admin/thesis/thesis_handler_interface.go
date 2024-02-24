package thesis

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminThesisHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	GetDetail(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetListThesisDefenseRequest(w http.ResponseWriter, r *http.Request)
	RegisterThesisDefense(w http.ResponseWriter, r *http.Request)
	CreateThesisDefense(w http.ResponseWriter, r *http.Request)
	UpdateThesisDefense(w http.ResponseWriter, r *http.Request)
	GetThesisSupervisorLog(w http.ResponseWriter, r *http.Request)
}

func NewAdminThesisHandler(ctx *service.ServiceCtx) AdminThesisHandlerInterface {
	return &thesisHandler{
		ctx,
	}
}
