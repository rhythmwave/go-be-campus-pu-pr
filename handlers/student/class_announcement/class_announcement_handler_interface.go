package class_announcement

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type StudentClassAnnouncementHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
}

func NewStudentClassAnnouncementHandler(ctx *service.ServiceCtx) StudentClassAnnouncementHandlerInterface {
	return &classAnnouncementHandler{
		ctx,
	}
}
