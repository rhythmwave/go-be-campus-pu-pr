package faculty

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type RootFacultyHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	GetDetail(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewRootFacultyHandler(ctx *service.ServiceCtx) RootFacultyHandlerInterface {
	return &facultyHandler{
		ctx,
	}
}
