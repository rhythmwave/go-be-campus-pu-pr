package class

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type LecturerClassHandlerInterface interface {
	GetActiveSemesterClassList(w http.ResponseWriter, r *http.Request)
	GetAssignedClass(w http.ResponseWriter, r *http.Request)
	GetAssignedSchedule(w http.ResponseWriter, r *http.Request)
	GetDetail(w http.ResponseWriter, r *http.Request)
	BulkGradeStudentClass(w http.ResponseWriter, r *http.Request)
}

func NewLecturerClassHandler(ctx *service.ServiceCtx) LecturerClassHandlerInterface {
	return &lecturerClassHandler{
		ctx,
	}
}
