package student

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type PmbStudentHandlerInterface interface {
	BulkCreate(w http.ResponseWriter, r *http.Request)
}

func NewPmbStudentHandler(ctx *service.ServiceCtx) PmbStudentHandlerInterface {
	return &studentHandler{
		ctx,
	}
}
