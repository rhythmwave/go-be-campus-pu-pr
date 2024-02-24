package file

import (
	"encoding/json"
	"net/http"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/service"
	"github.com/sccicitb/pupr-backend/utils"
	"github.com/sirupsen/logrus"
)

// fileHandler fileHandler
type fileHandler struct {
	*service.ServiceCtx
}

func (f fileHandler) UploadBase64Temp(w http.ResponseWriter, r *http.Request) {
	var result FileResponse

	ctx := r.Context()
	var in FileBase64Request
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = FileResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	dataRes, errs := f.FileService.UploadBase64Temp(ctx, in.GetFile(), "")
	if errs != nil {
		utils.PrintError(*errs)
		result = FileResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = FileResponse{
		Meta: &Meta{
			Message: "File uploaded",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &PathResponse{
			Path:     dataRes.Path,
			MimeType: dataRes.MimeType,
			Url:      dataRes.URL,
			Size:     dataRes.Size,
			PathType: dataRes.PathType,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
