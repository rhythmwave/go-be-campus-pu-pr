package class_work

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type LecturerClassWorkHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	GetSubmission(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GradeSubmission(w http.ResponseWriter, r *http.Request)
}

func NewLecturerClassWorkHandler(ctx *service.ServiceCtx) LecturerClassWorkHandlerInterface {
	return &classWorkHandler{
		ctx,
	}
}
