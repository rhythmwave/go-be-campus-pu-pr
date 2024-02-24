package thesis_examiner_role

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminThesisExaminerRoleHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewAdminThesisExaminerRoleHandler(ctx *service.ServiceCtx) AdminThesisExaminerRoleHandlerInterface {
	return &thesisExaminerRoleHandler{
		ctx,
	}
}
