package student

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminStudentHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	GetDetail(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	BulkUpdateStatus(w http.ResponseWriter, r *http.Request)
	GetStatusSummary(w http.ResponseWriter, r *http.Request)
	GetSubjectGrade(w http.ResponseWriter, r *http.Request)
	BulkUpdatePayment(w http.ResponseWriter, r *http.Request)
	GetPaymentLog(w http.ResponseWriter, r *http.Request)
	ConvertGrade(w http.ResponseWriter, r *http.Request)
}

func NewAdminStudentHandler(ctx *service.ServiceCtx) AdminStudentHandlerInterface {
	return &studentHandler{
		ctx,
	}
}
