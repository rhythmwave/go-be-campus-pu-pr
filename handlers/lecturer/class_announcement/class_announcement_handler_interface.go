package class_announcement

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type LecturerClassAnnouncementHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewLecturerClassAnnouncementHandler(ctx *service.ServiceCtx) LecturerClassAnnouncementHandlerInterface {
	return &classAnnouncementHandler{
		ctx,
	}
}
