package academic_guidance

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type LecturerAcademicGuidanceHandlerInterface interface {
	GetListStudent(w http.ResponseWriter, r *http.Request)
	GetSessionList(w http.ResponseWriter, r *http.Request)
	CreateSession(w http.ResponseWriter, r *http.Request)
	UpdateSession(w http.ResponseWriter, r *http.Request)
	DeleteSession(w http.ResponseWriter, r *http.Request)
}

func NewLecturerAcademicGuidanceHandler(ctx *service.ServiceCtx) LecturerAcademicGuidanceHandlerInterface {
	return &academicGuidanceHandler{
		ctx,
	}
}
