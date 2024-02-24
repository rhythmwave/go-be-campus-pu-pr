package lesson_plan

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminLessonPlanHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewAdminLessonPlanHandler(ctx *service.ServiceCtx) AdminLessonPlanHandlerInterface {
	return &lessonPlanHandler{
		ctx,
	}
}
