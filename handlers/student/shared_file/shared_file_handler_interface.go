package shared_file

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type StudentSharedFileHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
}

func NewStudentSharedFileHandler(ctx *service.ServiceCtx) StudentSharedFileHandlerInterface {
	return &sharedFileHandler{
		ctx,
	}
}
