package shared_file

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminSharedFileHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Approve(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewAdminSharedFileHandler(ctx *service.ServiceCtx) AdminSharedFileHandlerInterface {
	return &sharedFileHandler{
		ctx,
	}
}
