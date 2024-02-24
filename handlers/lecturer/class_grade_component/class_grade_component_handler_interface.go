package class_grade_component

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type LecturerClassGradeComponentHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Set(w http.ResponseWriter, r *http.Request)
}

func NewLecturerClassGradeComponentHandler(ctx *service.ServiceCtx) LecturerClassGradeComponentHandlerInterface {
	return &classGradeComponentHandler{
		ctx,
	}
}
