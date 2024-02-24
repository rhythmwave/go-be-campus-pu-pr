package lecture

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type StudentLectureHandlerInterface interface {
	AttendAutonomousLecture(w http.ResponseWriter, r *http.Request)
	GetHistory(w http.ResponseWriter, r *http.Request)
}

func NewStudentLectureHandler(ctx *service.ServiceCtx) StudentLectureHandlerInterface {
	return &studentLectureHandler{
		ctx,
	}
}
