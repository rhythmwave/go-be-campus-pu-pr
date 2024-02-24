package major

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminMajorHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
}

func NewAdminMajorHandler(ctx *service.ServiceCtx) AdminMajorHandlerInterface {
	return &majorHandler{
		ctx,
	}
}
