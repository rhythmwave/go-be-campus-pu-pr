package lecturer

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type RootLecturerHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
}

func NewRootLecturerHandler(ctx *service.ServiceCtx) RootLecturerHandlerInterface {
	return &lecturerHandler{
		ctx,
	}
}
