package major

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type RootMajorHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	GetDetail(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewRootMajorHandler(ctx *service.ServiceCtx) RootMajorHandlerInterface {
	return &majorHandler{
		ctx,
	}
}
