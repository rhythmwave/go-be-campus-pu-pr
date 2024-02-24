package class_grade_component

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminClassGradeComponentHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Set(w http.ResponseWriter, r *http.Request)
}

func NewAdminClassGradeComponentHandler(ctx *service.ServiceCtx) AdminClassGradeComponentHandlerInterface {
	return &classGradeComponentHandler{
		ctx,
	}
}
