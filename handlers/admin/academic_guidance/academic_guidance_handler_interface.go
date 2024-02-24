package academic_guidance

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminAcademicGuidanceHandlerInterface interface {
	GetListStudent(w http.ResponseWriter, r *http.Request)
	Upsert(w http.ResponseWriter, r *http.Request)
	UpsertDecision(w http.ResponseWriter, r *http.Request)
	GetSessionList(w http.ResponseWriter, r *http.Request)
	CreateSession(w http.ResponseWriter, r *http.Request)
	UpdateSession(w http.ResponseWriter, r *http.Request)
	DeleteSession(w http.ResponseWriter, r *http.Request)
}

func NewAdminAcademicGuidanceHandler(ctx *service.ServiceCtx) AdminAcademicGuidanceHandlerInterface {
	return &academicGuidanceHandler{
		ctx,
	}
}
