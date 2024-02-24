package learning_achievement_category

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminLearningAchievementCategoryHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewAdminLearningAchievementCategoryHandler(ctx *service.ServiceCtx) AdminLearningAchievementCategoryHandlerInterface {
	return &learningAchievementCategoryHandler{
		ctx,
	}
}
