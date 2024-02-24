package announcement

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminAnnouncementHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewAdminAnnouncementHandler(ctx *service.ServiceCtx) AdminAnnouncementHandlerInterface {
	return &announcementHandler{
		ctx,
	}
}
