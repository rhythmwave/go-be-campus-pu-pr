package class_discussion

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminClassDiscussionHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	GetComment(w http.ResponseWriter, r *http.Request)
}

func NewAdminClassDiscussionHandler(ctx *service.ServiceCtx) AdminClassDiscussionHandlerInterface {
	return &classDiscussionHandler{
		ctx,
	}
}
