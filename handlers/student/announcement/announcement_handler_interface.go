package announcement

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type StudentAnnouncementHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
}

func NewStudentAnnouncementHandler(ctx *service.ServiceCtx) StudentAnnouncementHandlerInterface {
	return &announcementHandler{
		ctx,
	}
}
