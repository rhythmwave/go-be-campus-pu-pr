package document_action

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminDocumentActionHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewAdminDocumentActionHandler(ctx *service.ServiceCtx) AdminDocumentActionHandlerInterface {
	return &documentActionHandler{
		ctx,
	}
}
