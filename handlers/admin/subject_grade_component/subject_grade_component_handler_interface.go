package subject_grade_component

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminSubjectGradeComponentHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Set(w http.ResponseWriter, r *http.Request)
}

func NewAdminSubjectGradeComponentHandler(ctx *service.ServiceCtx) AdminSubjectGradeComponentHandlerInterface {
	return &subjectGradeComponentHandler{
		ctx,
	}
}
