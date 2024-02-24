package room

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

type roomHandler struct {
	*service.ServiceCtx
}

func (a roomHandler) GetList(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Room", "Get List", nil)

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}
	isLaboratory, errs := utils.StringToBoolPointer(in.GetIsLaboratory())
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
	excludeLectureDate, errs := utils.StringToDate(in.GetExcludeLectureDate())
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
	req := objects.GetRoomRequest{
		BuildingId:            in.GetBuildingId(),
		IsLaboratory:          isLaboratory,
		ExcludeLectureDate:    excludeLectureDate,
		ExcludeStartTime:      in.GetExcludeStartTime(),
		ExcludeEndTime:        in.GetExcludeEndTime(),
		MaximumParticipant:    in.GetMaximumParticipant(),
		ForExam:               in.GetForExam(),
		ForceIncludeLectureId: in.GetForceIncludeLectureId(),
	}
	data, errs := a.RoomService.GetList(ctx, paginationData, req)
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
			Id:           v.Id,
			Code:         v.Code,
			Name:         utils.NullStringScan(v.Name),
			Capacity:     utils.NullUint32Scan(v.Capacity),
			IsLaboratory: v.IsLaboratory,
		})
	}

	result = GetListResponse{
		Meta: &Meta{
			Message: "Get List Room",
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

func (a roomHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
	var result GetDetailResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetDetailRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetDetailResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Room", "Get Detail", nil)

	data, errs := a.RoomService.GetDetail(ctx, in.GetId())
	if errs != nil {
		utils.PrintError(*errs)
		result = GetDetailResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = GetDetailResponse{
		Meta: &Meta{
			Message: "Get Detail Room",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &GetDetailResponseData{
			Id:               data.Id,
			BuildingId:       data.BuildingId,
			BuildingCode:     data.BuildingCode,
			BuildingName:     data.BuildingName,
			Code:             data.Code,
			Name:             utils.NullStringScan(data.Name),
			Capacity:         utils.NullUint32Scan(data.Capacity),
			ExamCapacity:     utils.NullUint32Scan(data.ExamCapacity),
			Purpose:          data.Purpose,
			IsUsable:         data.IsUsable,
			Area:             utils.NullFloatScan(data.Area),
			PhoneNumber:      utils.NullStringScan(data.PhoneNumber),
			Facility:         utils.NullStringScan(data.Facility),
			Remarks:          utils.NullStringScan(data.Remarks),
			Owner:            utils.NullStringScan(data.Owner),
			Location:         utils.NullStringScan(data.Location),
			StudyProgramId:   utils.NullStringScan(data.StudyProgramId),
			StudyProgramName: utils.NullStringScan(data.StudyProgramName),
			IsLaboratory:     data.IsLaboratory,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a roomHandler) Create(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Room", "Create", &in)

	data := objects.CreateRoom{
		BuildingId:     in.GetBuildingId(),
		Code:           in.GetCode(),
		Name:           in.GetName(),
		Capacity:       in.GetCapacity(),
		ExamCapacity:   in.GetExamCapacity(),
		IsUsable:       in.GetIsUsable(),
		Area:           in.GetArea(),
		PhoneNumber:    in.GetPhoneNumber(),
		Facility:       in.GetFacility(),
		Remarks:        in.GetRemarks(),
		Purpose:        in.GetPurpose(),
		Owner:          in.GetOwner(),
		Location:       in.GetLocation(),
		StudyProgramId: in.GetStudyProgramId(),
		IsLaboratory:   in.GetIsLaboratory(),
	}

	errs := a.RoomService.Create(ctx, data)
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
			Message: "Create Room",
			Status:  http.StatusCreated,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusCreated, &result)
}

func (a roomHandler) Update(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Room", "Update", &in)

	data := objects.UpdateRoom{
		Id:             in.GetId(),
		Code:           in.GetCode(),
		Name:           in.GetName(),
		Capacity:       in.GetCapacity(),
		ExamCapacity:   in.GetExamCapacity(),
		IsUsable:       in.GetIsUsable(),
		Area:           in.GetArea(),
		PhoneNumber:    in.GetPhoneNumber(),
		Facility:       in.GetFacility(),
		Remarks:        in.GetRemarks(),
		Purpose:        in.GetPurpose(),
		Owner:          in.GetOwner(),
		Location:       in.GetLocation(),
		StudyProgramId: in.GetStudyProgramId(),
	}
	errs := a.RoomService.Update(ctx, data)
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
			Message: "Update Room",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a roomHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var result DeleteResponse

	ctx := r.Context()
	var in DeleteRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = DeleteResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Room", "Delete", nil)

	errs := a.RoomService.Delete(ctx, in.GetId())
	if errs != nil {
		utils.PrintError(*errs)
		result = DeleteResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = DeleteResponse{
		Meta: &Meta{
			Message: "Delete Room",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a roomHandler) GetSchedule(w http.ResponseWriter, r *http.Request) {
	var result GetScheduleResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetScheduleRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetScheduleResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Room", "Get Schedule", nil)

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}
	data, errs := a.RoomService.GetSchedule(ctx, paginationData, in.GetRoomId(), in.GetDayOfWeek(), in.GetSemesterId())
	if errs != nil {
		utils.PrintError(*errs)
		result = GetScheduleResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*GetScheduleResponseData{}
	for _, v := range data.Data {
		dates := []*GetScheduleResponseDataDate{}
		for _, w := range v.Dates {
			schedules := []*GetScheduleResponseDataDateSchedule{}
			for _, x := range w.Schedules {
				schedules = append(schedules, &GetScheduleResponseDataDateSchedule{
					StartTime:        x.StartTime,
					EndTime:          x.EndTime,
					SubjectName:      x.SubjectName,
					ClassName:        x.ClassName,
					StudyProgramName: x.StudyProgramName,
				})
			}
			dates = append(dates, &GetScheduleResponseDataDate{
				Date:      w.Date.Format(constants.DateRFC),
				Schedules: schedules,
			})
		}

		resultData = append(resultData, &GetScheduleResponseData{
			RoomId:   v.RoomId,
			RoomName: utils.NullStringScan(v.RoomName),
			Dates:    dates,
		})
	}

	result = GetScheduleResponse{
		Meta: &Meta{
			Message: "Get Schedule Room",
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
