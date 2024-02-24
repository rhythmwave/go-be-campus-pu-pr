package subject_grade_component

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/schema"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/service"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/utils"
	"github.com/sirupsen/logrus"
)

type subjectGradeComponentHandler struct {
	*service.ServiceCtx
}

func (f subjectGradeComponentHandler) GetList(w http.ResponseWriter, r *http.Request) {
	var result GetListResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetListRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetListResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer f.AdminActivityLogService.Create(ctx, r, time.Now(), "Subject Grade Component", "Get List", nil)

	data, errs := f.SubjectGradeComponentService.GetList(ctx, in.GetSubjectId())
	if errs != nil {
		utils.PrintError(*errs)
		result = GetListResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*GetListResponseData{}
	for _, v := range data {
		resultData = append(resultData, &GetListResponseData{
			Id:         v.Id,
			Name:       v.Name,
			Percentage: v.Percentage,
			IsActive:   v.IsActive,
		})
	}

	result = GetListResponse{
		Meta: &Meta{
			Message: "Get List Subject Grade Component",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: resultData,
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (f subjectGradeComponentHandler) Set(w http.ResponseWriter, r *http.Request) {
	var result SetResponse

	ctx := r.Context()
	var in SetRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = SetResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer f.AdminActivityLogService.Create(ctx, r, time.Now(), "Subject Grade Component", "Set", &in)

	data := []objects.SetSubjectGradeComponent{}
	for _, v := range in.GetData() {
		data = append(data, objects.SetSubjectGradeComponent{
			Name:       v.GetName(),
			Percentage: v.GetPercentage(),
			IsActive:   v.GetIsActive(),
		})
	}
	errs := f.SubjectGradeComponentService.Set(ctx, in.GetSubjectId(), data)
	if errs != nil {
		utils.PrintError(*errs)
		result = SetResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = SetResponse{
		Meta: &Meta{
			Message: "Set Subject Grade Component",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
