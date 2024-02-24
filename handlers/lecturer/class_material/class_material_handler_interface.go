package class_material

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type LecturerClassMaterialHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	BulkUpdateActivation(w http.ResponseWriter, r *http.Request)
	BulkDelete(w http.ResponseWriter, r *http.Request)
}

func NewLecturerClassMaterialHandler(ctx *service.ServiceCtx) LecturerClassMaterialHandlerInterface {
	return &classMaterialHandler{
		ctx,
	}
}
