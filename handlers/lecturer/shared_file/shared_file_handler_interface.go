package shared_file

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type LecturerSharedFileHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewLecturerSharedFileHandler(ctx *service.ServiceCtx) LecturerSharedFileHandlerInterface {
	return &sharedFileHandler{
		ctx,
	}
}
