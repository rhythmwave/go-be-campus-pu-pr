package class

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type StudentClassHandlerInterface interface {
	GetOfferedClassList(w http.ResponseWriter, r *http.Request)
	GetOfferedSchedule(w http.ResponseWriter, r *http.Request)
	GetTakenClass(w http.ResponseWriter, r *http.Request)
}

func NewStudentClassHandler(ctx *service.ServiceCtx) StudentClassHandlerInterface {
	return &studentClassHandler{
		ctx,
	}
}
