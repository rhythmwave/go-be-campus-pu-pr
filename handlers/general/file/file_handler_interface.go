package file

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type GeneralFileHandlerInterface interface {
	UploadBase64Temp(w http.ResponseWriter, r *http.Request)
}

func NewGeneralFileHandler(ctx *service.ServiceCtx) GeneralFileHandlerInterface {
	return &fileHandler{
		ctx,
	}
}
