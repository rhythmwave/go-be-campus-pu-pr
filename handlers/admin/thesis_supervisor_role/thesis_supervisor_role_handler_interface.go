package thesis_supervisor_role

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminThesisSupervisorRoleHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewAdminThesisSupervisorRoleHandler(ctx *service.ServiceCtx) AdminThesisSupervisorRoleHandlerInterface {
	return &thesisSupervisorRoleHandler{
		ctx,
	}
}
