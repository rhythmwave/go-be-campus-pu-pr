package class

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminClassHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	GetDetail(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	UpdateActivation(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Duplicate(w http.ResponseWriter, r *http.Request)
	BulkUpdateMaximumParticipant(w http.ResponseWriter, r *http.Request)
	GetClassParticipantList(w http.ResponseWriter, r *http.Request)
}

func NewAdminClassHandler(ctx *service.ServiceCtx) AdminClassHandlerInterface {
	return &classHandler{
		ctx,
	}
}
