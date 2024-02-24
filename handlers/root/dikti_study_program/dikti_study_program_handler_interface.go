package dikti_study_program

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type RootDiktiStudyProgramHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
}

func NewRootDiktiStudyProgramHandler(ctx *service.ServiceCtx) RootDiktiStudyProgramHandlerInterface {
	return &diktiStudyProgramHandler{
		ctx,
	}
}
