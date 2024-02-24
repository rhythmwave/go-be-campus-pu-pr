package curriculum

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminCurriculumHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	GetDetail(w http.ResponseWriter, r *http.Request)
	GetActiveByStudyProgramId(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewAdminCurriculumHandler(ctx *service.ServiceCtx) AdminCurriculumHandlerInterface {
	return &curriculumHandler{
		ctx,
	}
}
