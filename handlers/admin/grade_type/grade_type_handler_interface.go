package grade_type

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminGradeTypeHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewAdminGradeTypeHandler(ctx *service.ServiceCtx) AdminGradeTypeHandlerInterface {
	return &gradeTypeHandler{
		ctx,
	}
}
