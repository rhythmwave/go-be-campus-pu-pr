package class_discussion

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type StudentClassDiscussionHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	GetComment(w http.ResponseWriter, r *http.Request)
	CreateComment(w http.ResponseWriter, r *http.Request)
	DeleteComment(w http.ResponseWriter, r *http.Request)
}

func NewStudentClassDiscussionHandler(ctx *service.ServiceCtx) StudentClassDiscussionHandlerInterface {
	return &classDiscussionHandler{
		ctx,
	}
}
