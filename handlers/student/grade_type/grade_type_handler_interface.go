package grade_type

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type StudentGradeTypeHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
}

func NewStudentGradeTypeHandler(ctx *service.ServiceCtx) StudentGradeTypeHandlerInterface {
	return &gradeTypeHandler{
		ctx,
	}
}
