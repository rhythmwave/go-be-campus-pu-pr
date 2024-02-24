package admin_activity_log

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type RootAdminActivityLogHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
}

func NewRootAdminActivityLogHandler(ctx *service.ServiceCtx) RootAdminActivityLogHandlerInterface {
	return &adminActivityLogHandler{
		ctx,
	}
}
