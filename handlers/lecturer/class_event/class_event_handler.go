package class_event

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

type classEventHandler struct {
	*service.ServiceCtx
}

func (c classEventHandler) GetList(w http.ResponseWriter, r *http.Request) {
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
	defer c.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Class Event", "Get List", nil)

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}
	data, errs := c.ClassEventService.GetList(ctx, paginationData, in.GetClassId(), "", false, nil)
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
			Frequency:          v.Frequency,
			EventTime:          v.EventTime.Format(constants.DateRFC),
			LecturerId:         v.LecturerId,
			LecturerName:       v.LecturerName,
			LecturerFrontTitle: utils.NullStringScan(v.LecturerFrontTitle),
			LecturerBackDegree: utils.NullStringScan(v.LecturerBackDegree),
			Remarks:            utils.NullStringScan(v.Remarks),
			IsActive:           v.IsActive,
			CreatedAt:          v.CreatedAt,
		})
	}

	result = GetListResponse{
		Meta: &Meta{
			Message: "Get List Class Event",
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

func (c classEventHandler) Create(w http.ResponseWriter, r *http.Request) {
	var result CreateResponse

	ctx := r.Context()
	var in CreateRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = CreateResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer c.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Class Event", "Create", nil)

	eventTime, errs := utils.StringToTime(in.GetEventTime())
	if errs != nil {
		utils.PrintError(*errs)
		result = CreateResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	data := objects.CreateClassEvent{
		ClassId:   in.GetClassId(),
		Title:     in.GetTitle(),
		Frequency: in.GetFrequency(),
		EventTime: eventTime,
		Remarks:   in.GetRemarks(),
		IsActive:  in.GetIsActive(),
	}
	errs = c.ClassEventService.Create(ctx, data)
	if errs != nil {
		utils.PrintError(*errs)
		result = CreateResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = CreateResponse{
		Meta: &Meta{
			Message: "Create Class Event",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (c classEventHandler) Update(w http.ResponseWriter, r *http.Request) {
	var result UpdateResponse

	ctx := r.Context()
	var in UpdateRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = UpdateResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer c.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Class Event", "Update", nil)

	eventTime, errs := utils.StringToTime(in.GetEventTime())
	if errs != nil {
		utils.PrintError(*errs)
		result = UpdateResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	data := objects.UpdateClassEvent{
		Id:        in.GetId(),
		Title:     in.GetTitle(),
		Frequency: in.GetFrequency(),
		EventTime: eventTime,
		Remarks:   in.GetRemarks(),
		IsActive:  in.GetIsActive(),
	}
	errs = c.ClassEventService.Update(ctx, data)
	if errs != nil {
		utils.PrintError(*errs)
		result = UpdateResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = UpdateResponse{
		Meta: &Meta{
			Message: "Update Class Event",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (c classEventHandler) BulkUpdateActivation(w http.ResponseWriter, r *http.Request) {
	var result BulkUpdateActivationResponse

	ctx := r.Context()
	var in BulkUpdateActivationRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = BulkUpdateActivationResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer c.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Class Event", "Bulk Update Activation", nil)

	errs := c.ClassEventService.BulkUpdateActivation(ctx, in.GetIds(), in.GetIsActive())
	if errs != nil {
		utils.PrintError(*errs)
		result = BulkUpdateActivationResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = BulkUpdateActivationResponse{
		Meta: &Meta{
			Message: "Bulk Update Activation Class Event",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (c classEventHandler) BulkDelete(w http.ResponseWriter, r *http.Request) {
	var result BulkDeleteResponse

	ctx := r.Context()
	var in BulkDeleteRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = BulkDeleteResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer c.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Class Event", "Bulk Delete", nil)

	errs := c.ClassEventService.BulkDelete(ctx, in.GetIds())
	if errs != nil {
		utils.PrintError(*errs)
		result = BulkDeleteResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = BulkDeleteResponse{
		Meta: &Meta{
			Message: "Bulk Delete Class Event",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
