package lecturer_student_activity_log

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type RootLecturerStudentActivityLogHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
}

func NewRootLecturerStudentActivityLogHandler(ctx *service.ServiceCtx) RootLecturerStudentActivityLogHandlerInterface {
	return &lecturerStudentActivityLogHandler{
		ctx,
	}
}
