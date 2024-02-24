package announcement

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type LecturerAnnouncementHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
}

func NewLecturerAnnouncementHandler(ctx *service.ServiceCtx) LecturerAnnouncementHandlerInterface {
	return &announcementHandler{
		ctx,
	}
}
