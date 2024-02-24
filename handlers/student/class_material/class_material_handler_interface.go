package class_material

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type StudentClassMaterialHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
}

func NewStudentClassMaterialHandler(ctx *service.ServiceCtx) StudentClassMaterialHandlerInterface {
	return &classMaterialHandler{
		ctx,
	}
}
