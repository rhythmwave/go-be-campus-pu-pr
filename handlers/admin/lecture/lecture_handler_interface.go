package lecture

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminLectureHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	BulkCreate(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	ResetParticipation(w http.ResponseWriter, r *http.Request)
	GetStudentParticipation(w http.ResponseWriter, r *http.Request)
	GetCalendar(w http.ResponseWriter, r *http.Request)
}

func NewAdminLectureHandler(ctx *service.ServiceCtx) AdminLectureHandlerInterface {
	return &lectureHandler{
		ctx,
	}
}
