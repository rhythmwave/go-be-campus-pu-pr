package learning_achievement

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminLearningAchievementHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewAdminLearningAchievementHandler(ctx *service.ServiceCtx) AdminLearningAchievementHandlerInterface {
	return &learningAchievementHandler{
		ctx,
	}
}
