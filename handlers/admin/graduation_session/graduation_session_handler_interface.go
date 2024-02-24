package graduation_session

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminGraduationSessionHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewAdminGraduationSessionHandler(ctx *service.ServiceCtx) AdminGraduationSessionHandlerInterface {
	return &graduationSessionHandler{
		ctx,
	}
}
