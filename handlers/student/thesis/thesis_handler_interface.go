package thesis

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type StudentThesisHandlerInterface interface {
	GetDetail(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

func NewStudentThesisHandler(ctx *service.ServiceCtx) StudentThesisHandlerInterface {
	return &thesisHandler{
		ctx,
	}
}
