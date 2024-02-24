package class_material

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminClassMaterialHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
}

func NewAdminClassMaterialHandler(ctx *service.ServiceCtx) AdminClassMaterialHandlerInterface {
	return &classMaterialHandler{
		ctx,
	}
}
