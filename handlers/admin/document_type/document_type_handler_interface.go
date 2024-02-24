package document_type

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminDocumentTypeHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewAdminDocumentTypeHandler(ctx *service.ServiceCtx) AdminDocumentTypeHandlerInterface {
	return &documentTypeHandler{
		ctx,
	}
}
