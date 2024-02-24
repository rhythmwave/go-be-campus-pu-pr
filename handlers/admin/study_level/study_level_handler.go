package study_level

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/schema"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/service"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
	"github.com/sirupsen/logrus"
)

type studyLevelHandler struct {
	*service.ServiceCtx
}

func (d studyLevelHandler) GetList(w http.ResponseWriter, r *http.Request) {
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

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}
	data, errs := d.StudyLevelService.GetList(ctx, paginationData)
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
	for _, v := range data.Data {
		resultData = append(resultData, &GetListResponseData{
			Id:                    v.Id,
			Name:                  v.Name,
			ShortName:             v.ShortName,
			KkniQualification:     utils.NullStringScan(v.KkniQualification),
			AcceptanceRequirement: utils.NullStringScan(v.AcceptanceRequirement),
			FurtherEducationLevel: utils.NullStringScan(v.FurtherEducationLevel),
			ProfessionalStatus:    utils.NullStringScan(v.ProfessionalStatus),
			CourseLanguage:        utils.NullStringScan(v.CourseLanguage),
		})
	}

	result = GetListResponse{
		Meta: &Meta{
			Message: "Get List Study Level",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Pagination: &Pagination{
			Page:         uint32(data.Pagination.Page),
			Limit:        uint32(data.Pagination.Limit),
			Prev:         uint32(*data.Pagination.Prev),
			Next:         uint32(*data.Pagination.Next),
			TotalPages:   uint32(*data.Pagination.TotalPages),
			TotalRecords: uint32(*data.Pagination.TotalRecords),
		},
		Data: resultData,
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a studyLevelHandler) UpdateSkpi(w http.ResponseWriter, r *http.Request) {
	var result UpdateSkpiResponse

	ctx := r.Context()
	var in UpdateSkpiRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = UpdateSkpiResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Study Level SKPI", "Update", &in)

	data := objects.UpdateStudyLevelSkpi{
		Id:                    in.GetId(),
		KkniQualification:     in.GetKkniQualification(),
		AcceptanceRequirement: in.GetAcceptanceRequirement(),
		FurtherEducationLevel: in.GetFurtherEducationLevel(),
		ProfessionalStatus:    in.GetProfessionalStatus(),
		CourseLanguage:        in.GetCourseLanguage(),
	}
	errs := a.StudyLevelService.UpdateSkpi(ctx, data)
	if errs != nil {
		utils.PrintError(*errs)
		result = UpdateSkpiResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = UpdateSkpiResponse{
		Meta: &Meta{
			Message: "Update Study Level SKPI",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
