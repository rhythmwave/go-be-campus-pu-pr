package credit_quota

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminCreditQuotaHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewAdminCreditQuotaHandler(ctx *service.ServiceCtx) AdminCreditQuotaHandlerInterface {
	return &creditQuotaHandler{
		ctx,
	}
}
