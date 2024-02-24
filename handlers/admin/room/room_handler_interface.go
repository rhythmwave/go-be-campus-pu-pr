package room

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminRoomHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	GetDetail(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetSchedule(w http.ResponseWriter, r *http.Request)
}

func NewAdminRoomHandler(ctx *service.ServiceCtx) AdminRoomHandlerInterface {
	return &roomHandler{
		ctx,
	}
}
