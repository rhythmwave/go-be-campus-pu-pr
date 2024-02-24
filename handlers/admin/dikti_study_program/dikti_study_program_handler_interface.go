package dikti_study_program

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminDiktiStudyProgramHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
}

func NewAdminDiktiStudyProgramHandler(ctx *service.ServiceCtx) AdminDiktiStudyProgramHandlerInterface {
	return &diktiStudyProgramHandler{
		ctx,
	}
}
