package class_work

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/schema"
	"github.com/sccicitb/pupr-backend/constants"
	appConstants "github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/service"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
	"github.com/sirupsen/logrus"
)

type classWorkHandler struct {
	*service.ServiceCtx
}

func (f classWorkHandler) GetList(w http.ResponseWriter, r *http.Request) {
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
	defer f.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Class Work", "Get List", nil)

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}
	data, errs := f.ClassWorkService.GetList(ctx, paginationData, appConstants.AppTypeStudent, nil)
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
			Id:                     v.Id,
			Title:                  v.Title,
			Abstraction:            utils.NullStringScan(v.Abstraction),
			FileUrl:                v.FileUrl,
			LecturerId:             v.LecturerId,
			LecturerName:           v.LecturerName,
			LecturerFrontTitle:     utils.NullStringScan(v.LecturerFrontTitle),
			LecturerBackDegree:     utils.NullStringScan(v.LecturerBackDegree),
			StartTime:              v.StartTime.Format(constants.DateRFC),
			EndTime:                v.EndTime.Format(constants.DateRFC),
			SubmissionFileUrl:      v.SubmissionFileUrl,
			SubmissionFilePath:     utils.NullStringScan(v.SubmissionFilePath),
			SubmissionFilePathType: utils.NullStringScan(v.SubmissionFilePathType),
			SubmissionPoint:        utils.NullFloatScan(v.SubmissionPoint),
		})
	}

	result = GetListResponse{
		Meta: &Meta{
			Message: "Get List Class Work",
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

func (c classWorkHandler) Submit(w http.ResponseWriter, r *http.Request) {
	var result SubmitResponse

	ctx := r.Context()
	var in SubmitRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = SubmitResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer c.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Class Work", "Submit", nil)

	errs := c.ClassWorkService.Submit(ctx, in.GetClassWorkId(), in.GetFilePath(), in.GetFilePathType())
	if errs != nil {
		utils.PrintError(*errs)
		result = SubmitResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = SubmitResponse{
		Meta: &Meta{
			Message: "Submit Class Work",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
