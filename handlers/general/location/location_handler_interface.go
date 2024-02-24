package location

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/service"
)

type GeneralLocationHandlerInterface interface {
	GetListCountry(w http.ResponseWriter, r *http.Request)
	GetListProvince(w http.ResponseWriter, r *http.Request)
	GetListRegency(w http.ResponseWriter, r *http.Request)
	GetListDistrict(w http.ResponseWriter, r *http.Request)
	GetListVillage(w http.ResponseWriter, r *http.Request)
	TempCreateData(w http.ResponseWriter, r *http.Request)
	TempGetData(w http.ResponseWriter, r *http.Request)
}

func NewGeneralLocationHandler(ctx *service.ServiceCtx) GeneralLocationHandlerInterface {
	return &locationHandler{
		ctx,
	}
}
