package student_leave

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/schema"
	"github.com/sccicitb/pupr-backend/constants"
	appConstants "github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/service"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
	"github.com/sirupsen/logrus"
)

type studentLeaveHandler struct {
	*service.ServiceCtx
}

func (l studentLeaveHandler) GetListRequests(w http.ResponseWriter, r *http.Request) {
	var result GetListRequestsResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetListRequestsRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetListRequestsResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer l.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Student Leave", "Get List Request", nil)

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}
	data, errs := l.StudentLeaveService.GetListRequest(ctx, paginationData, appConstants.AppTypeStudent, "", nil)
	if errs != nil {
		utils.PrintError(*errs)
		result = GetListRequestsResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*GetListRequestsResponseData{}
	for _, v := range data.Data {
		resultData = append(resultData, &GetListRequestsResponseData{
			Id:                         v.Id,
			StartDate:                  v.StartDate.Format(constants.DateRFC),
			TotalLeaveDurationSemester: v.TotalLeaveDurationSemester,
			PermitNumber:               utils.NullStringScan(v.PermitNumber),
			Purpose:                    v.Purpose,
			Remarks:                    v.Remarks,
			IsApproved:                 utils.NullBooleanScan(v.IsApproved),
			SemesterType:               v.SemesterType,
			SemesterSchoolYear:         v.SemesterSchoolYear,
		})
	}

	result = GetListRequestsResponse{
		Meta: &Meta{
			Message: "Get List Request Student Leave",
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

func (l studentLeaveHandler) Create(w http.ResponseWriter, r *http.Request) {
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
	defer l.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Student Leave", "Create", &in)

	startDate, errs := utils.StringToTime(in.GetStartDate())
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
	data := objects.CreateStudentLeave{
		TotalLeaveDurationSemester: in.GetTotalLeaveDurationSemester(),
		StartDate:                  startDate,
		Purpose:                    in.GetPurpose(),
		Remarks:                    in.GetRemarks(),
	}
	errs = l.StudentLeaveService.Create(ctx, data)
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
			Message: "Create Student Leave",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (l studentLeaveHandler) Update(w http.ResponseWriter, r *http.Request) {
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
	defer l.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Student Leave", "Update", &in)

	data := objects.UpdateStudentLeave{
		Id:      in.GetId(),
		Purpose: in.GetPurpose(),
		Remarks: in.GetRemarks(),
	}
	errs := l.StudentLeaveService.Update(ctx, data)
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
			Message: "Update Student Leave",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
