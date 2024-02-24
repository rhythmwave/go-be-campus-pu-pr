package student_class

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminStudentClassHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	TransferStudentClass(w http.ResponseWriter, r *http.Request)
	ReshuffleStudentClass(w http.ResponseWriter, r *http.Request)
	MergeStudentClass(w http.ResponseWriter, r *http.Request)
	BulkGradeStudentClass(w http.ResponseWriter, r *http.Request)
}

func NewAdminStudentClassHandler(ctx *service.ServiceCtx) AdminStudentClassHandlerInterface {
	return &studentClassHandler{
		ctx,
	}
}
