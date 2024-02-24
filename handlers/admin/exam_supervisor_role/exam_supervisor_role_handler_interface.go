package exam_supervisor_role

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminExamSupervisorRoleHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewAdminExamSupervisorRoleHandler(ctx *service.ServiceCtx) AdminExamSupervisorRoleHandlerInterface {
	return &examSupervisorRoleHandler{
		ctx,
	}
}
