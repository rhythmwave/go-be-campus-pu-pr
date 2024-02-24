package lecture

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/schema"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/service"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
	"github.com/sirupsen/logrus"
)

type studentLectureHandler struct {
	*service.ServiceCtx
}

func (l studentLectureHandler) AttendAutonomousLecture(w http.ResponseWriter, r *http.Request) {
	var result AttendAutonomousLectureResponse

	ctx := r.Context()
	var in AttendAutonomousLectureRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = AttendAutonomousLectureResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer l.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Lecture", "Attend Autonomous Lecture", nil)

	errs := l.LectureService.AttendAutonomousLecture(ctx, in.GetLectureId())
	if errs != nil {
		utils.PrintError(*errs)
		result = AttendAutonomousLectureResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = AttendAutonomousLectureResponse{
		Meta: &Meta{
			Message: "Attend Autonomous Lecture",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}

	utils.JSONResponse(w, http.StatusOK, &result)
}

func (f studentLectureHandler) GetHistory(w http.ResponseWriter, r *http.Request) {
	var result GetHistoryResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetHistoryRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetHistoryResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer f.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Lecture", "Get History", nil)

	paginationData := common.PaginationRequest{
		Limit: in.GetLimit(),
		Page:  in.GetPage(),
	}
	startDate, errs := utils.StringToDate(in.GetStartDate())
	if errs != nil {
		utils.PrintError(*errs)
		result = GetHistoryResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	endDate, errs := utils.StringToDate(in.GetEndDate())
	if errs != nil {
		utils.PrintError(*errs)
		result = GetHistoryResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	data, errs := f.LectureService.GetHistory(ctx, paginationData, "", startDate, endDate)
	if errs != nil {
		utils.PrintError(*errs)
		result = GetHistoryResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*GetHistoryResponseData{}
	for _, v := range data.Data {
		attendTimeStr := v.AttendTime.Format("1504")
		attendTime, err := strconv.Atoi(attendTimeStr)
		if err != nil {
			errs := constants.ErrorInternalServer(err.Error())
			utils.PrintError(*errs)
			result = GetHistoryResponse{
				Meta: &Meta{
					Message: errs.Err.Error(),
					Status:  uint32(errs.HttpCode),
					Code:    errs.CustomCode,
				},
			}
			utils.JSONResponse(w, errs.HttpCode, &result)
			return
		}
		resultData = append(resultData, &GetHistoryResponseData{
			Id:          v.Id,
			LectureDate: v.LectureDate.Format(constants.DateFormatStd),
			SubjectName: v.SubjectName,
			AttendTime:  uint32(attendTime),
		})
	}

	result = GetHistoryResponse{
		Meta: &Meta{
			Message: "Get History Lecture",
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
