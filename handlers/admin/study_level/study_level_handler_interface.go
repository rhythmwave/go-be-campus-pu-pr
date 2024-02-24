package study_level

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type AdminStudyLevelHandlerInterface interface {
	GetList(w http.ResponseWriter, r *http.Request)
	UpdateSkpi(w http.ResponseWriter, r *http.Request)
}

func NewAdminStudyLevelHandler(ctx *service.ServiceCtx) AdminStudyLevelHandlerInterface {
	return &studyLevelHandler{
		ctx,
	}
}
