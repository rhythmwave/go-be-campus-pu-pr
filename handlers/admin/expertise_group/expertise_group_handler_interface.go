package expertise_group

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminExpertiseGroupHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewAdminExpertiseGroupHandler(ctx *service.ServiceCtx) AdminExpertiseGroupHandlerInterface {
	return &expertiseGroupHandler{
		ctx,
	}
}
