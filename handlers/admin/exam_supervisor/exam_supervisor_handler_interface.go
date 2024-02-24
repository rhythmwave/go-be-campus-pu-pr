package exam_supervisor

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminExamSupervisorHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	GetDetail(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewAdminExamSupervisorHandler(ctx *service.ServiceCtx) AdminExamSupervisorHandlerInterface {
	return &examSupervisorHandler{
		ctx,
	}
}
