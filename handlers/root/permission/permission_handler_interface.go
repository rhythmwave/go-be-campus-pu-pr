package permission

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type RootPermissionHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
}

func NewRootPermissionHandler(ctx *service.ServiceCtx) RootPermissionHandlerInterface {
	return &permissionHandler{
		ctx,
	}
}
