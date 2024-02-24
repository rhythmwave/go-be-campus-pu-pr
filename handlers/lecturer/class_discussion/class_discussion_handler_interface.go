package class_discussion

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type LecturerClassDiscussionHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	GetComment(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	CreateComment(w http.ResponseWriter, r *http.Request)
	DeleteComment(w http.ResponseWriter, r *http.Request)
}

func NewLecturerClassDiscussionHandler(ctx *service.ServiceCtx) LecturerClassDiscussionHandlerInterface {
	return &classDiscussionHandler{
		ctx,
	}
}
