package academic_guidance

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type StudentAcademicGuidanceHandlerInterface interface {
	GetDetail(w http.ResponseWriter, r *http.Request)
	GetSessionList(w http.ResponseWriter, r *http.Request)
}

func NewStudentAcademicGuidanceHandler(ctx *service.ServiceCtx) StudentAcademicGuidanceHandlerInterface {
	return &academicGuidanceHandler{
		ctx,
	}
}
