package lecture

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type LecturerLectureHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	GetDetail(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

func NewLecturerLectureHandler(ctx *service.ServiceCtx) LecturerLectureHandlerInterface {
	return &lectureHandler{
		ctx,
	}
}
