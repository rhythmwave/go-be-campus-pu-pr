package graduation_predicate

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminGraduationPredicateHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewAdminGraduationPredicateHandler(ctx *service.ServiceCtx) AdminGraduationPredicateHandlerInterface {
	return &graduationPredicateHandler{
		ctx,
	}
}
