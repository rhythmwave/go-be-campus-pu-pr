package general

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type StudentGeneralHandlerInterface interface {
	GetSemesterSummary(w http.ResponseWriter, r *http.Request)
	GetProfile(w http.ResponseWriter, r *http.Request)
	UpdateProfile(w http.ResponseWriter, r *http.Request)
	GetParentProfile(w http.ResponseWriter, r *http.Request)
	UpdateParentProfile(w http.ResponseWriter, r *http.Request)
	GetSchoolProfile(w http.ResponseWriter, r *http.Request)
	UpdateSchoolProfile(w http.ResponseWriter, r *http.Request)
}

func NewStudentGeneralHandler(ctx *service.ServiceCtx) StudentGeneralHandlerInterface {
	return &studentGeneralHandler{
		ctx,
	}
}
