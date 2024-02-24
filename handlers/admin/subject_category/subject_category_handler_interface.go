package subject_category

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminSubjectCategoryHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewAdminSubjectCategoryHandler(ctx *service.ServiceCtx) AdminSubjectCategoryHandlerInterface {
	return &subjectCategoryHandler{
		ctx,
	}
}
