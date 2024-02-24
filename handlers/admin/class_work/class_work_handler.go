package class_work

import (
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
	defer f.AdminActivityLogService.Create(ctx, r, time.Now(), "Class Work", "Get List", nil)

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}
	classIds := []string{}
	if in.GetClassId() != "" {
		classIds = append(classIds, in.GetClassId())
	}
	data, errs := f.ClassWorkService.GetList(ctx, paginationData, appConstants.AppTypeAdmin, classIds)
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
			Id:                 v.Id,
			Title:              v.Title,
			Abstraction:        utils.NullStringScan(v.Abstraction),
			FileUrl:            v.FileUrl,
			LecturerId:         v.LecturerId,
			LecturerName:       v.LecturerName,
			LecturerFrontTitle: utils.NullStringScan(v.LecturerFrontTitle),
			LecturerBackDegree: utils.NullStringScan(v.LecturerBackDegree),
			StartTime:          v.StartTime.Format(constants.DateRFC),
			EndTime:            v.EndTime.Format(constants.DateRFC),
			TotalSubmission:    v.TotalSubmission,
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

func (f classWorkHandler) GetSubmission(w http.ResponseWriter, r *http.Request) {
	var result GetSubmissionResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetSubmissionRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetSubmissionResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer f.AdminActivityLogService.Create(ctx, r, time.Now(), "Class Work", "Get Submission", nil)

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}
	data, errs := f.ClassWorkService.GetSubmission(ctx, paginationData, in.GetClassWorkId())
	if errs != nil {
		utils.PrintError(*errs)
		result = GetSubmissionResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*GetSubmissionResponseData{}
	for _, v := range data.Data {
		resultData = append(resultData, &GetSubmissionResponseData{
			Id:               utils.NullStringScan(v.Id),
			StudentId:        v.StudentId,
			NimNumber:        v.NimNumber,
			Name:             v.Name,
			StudyProgramName: v.StudyProgramName,
			FileUrl:          v.FileUrl,
			Point:            utils.NullFloatScan(v.Point),
		})
	}

	result = GetSubmissionResponse{
		Meta: &Meta{
			Message: "Get Submission Class Work",
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
