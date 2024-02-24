package subject

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminSubjectHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	GetDetail(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	SetPrerequisiteSubject(w http.ResponseWriter, r *http.Request)
	SetEquivalentSubject(w http.ResponseWriter, r *http.Request)
	DeleteEquivalentSubject(w http.ResponseWriter, r *http.Request)
}

func NewAdminSubjectHandler(ctx *service.ServiceCtx) AdminSubjectHandlerInterface {
	return &subjectHandler{
		ctx,
	}
}
