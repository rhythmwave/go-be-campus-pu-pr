package class_exam

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type LecturerClassExamHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	GetSubmission(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GradeSubmission(w http.ResponseWriter, r *http.Request)
}

func NewLecturerClassExamHandler(ctx *service.ServiceCtx) LecturerClassExamHandlerInterface {
	return &classExamHandler{
		ctx,
	}
}
