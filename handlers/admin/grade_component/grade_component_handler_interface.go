package grade_component

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminGradeComponentHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetListBySubjectCategory(w http.ResponseWriter, r *http.Request)
	BulkUpdatePercentage(w http.ResponseWriter, r *http.Request)
}

func NewAdminGradeComponentHandler(ctx *service.ServiceCtx) AdminGradeComponentHandlerInterface {
	return &gradeComponentHandler{
		ctx,
	}
}
