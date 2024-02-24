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
	defer l.AdminActivityLogService.Create(ctx, r, time.Now(), "Student Leave", "Get List Request", nil)

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}
	isApproved, errs := utils.StringToBoolPointer(in.GetIsApproved())
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
	data, errs := l.StudentLeaveService.GetListRequest(ctx, paginationData, appConstants.AppTypeAdmin, in.GetStudyProgramId(), isApproved)
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
			NimNumber:                  v.NimNumber,
			Name:                       v.Name,
			DiktiStudyProgramCode:      v.DiktiStudyProgramCode,
			StudyProgramName:           v.StudyProgramName,
			StudyLevelShortName:        v.StudyLevelShortName,
			DiktiStudyProgramType:      v.DiktiStudyProgramType,
			StartDate:                  v.StartDate.Format(constants.DateRFC),
			TotalLeaveDurationSemester: v.TotalLeaveDurationSemester,
			PermitNumber:               utils.NullStringScan(v.PermitNumber),
			Purpose:                    v.Purpose,
			Remarks:                    v.Remarks,
			IsApproved:                 utils.NullBooleanScan(v.IsApproved),
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

func (l studentLeaveHandler) GetList(w http.ResponseWriter, r *http.Request) {
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
	defer l.AdminActivityLogService.Create(ctx, r, time.Now(), "Student Leave", "Get List", nil)

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}
	data, errs := l.StudentLeaveService.GetList(ctx, paginationData, in.GetStudyProgramId(), in.GetSemesterId())
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
			NimNumber:             v.NimNumber,
			Name:                  v.Name,
			DiktiStudyProgramCode: v.DiktiStudyProgramCode,
			StudyProgramName:      v.StudyProgramName,
			StudyLevelShortName:   v.StudyLevelShortName,
			DiktiStudyProgramType: v.DiktiStudyProgramType,
			SemesterSchoolYear:    v.SemesterSchoolYear,
			SemesterType:          v.SemesterType,
			PermitNumber:          utils.NullStringScan(v.PermitNumber),
			Purpose:               v.Purpose,
			Remarks:               v.Remarks,
		})
	}

	result = GetListResponse{
		Meta: &Meta{
			Message: "Get List Student Leave",
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
	defer l.AdminActivityLogService.Create(ctx, r, time.Now(), "Student Leave", "Create", &in)

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
		StudentId:                  in.GetStudentId(),
		TotalLeaveDurationSemester: in.GetTotalLeaveDurationSemester(),
		StartDate:                  startDate,
		PermitNumber:               in.GetPermitNumber(),
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
	defer l.AdminActivityLogService.Create(ctx, r, time.Now(), "Student Leave", "Update", &in)

	data := objects.UpdateStudentLeave{
		Id:           in.GetId(),
		PermitNumber: in.GetPermitNumber(),
		Purpose:      in.GetPurpose(),
		Remarks:      in.GetRemarks(),
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

func (l studentLeaveHandler) Approve(w http.ResponseWriter, r *http.Request) {
	var result ApproveResponse

	ctx := r.Context()
	var in ApproveRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = ApproveResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer l.AdminActivityLogService.Create(ctx, r, time.Now(), "Student Leave", "Approve", nil)

	errs := l.StudentLeaveService.Approve(ctx, in.GetId(), in.GetIsApproved())
	if errs != nil {
		utils.PrintError(*errs)
		result = ApproveResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = ApproveResponse{
		Meta: &Meta{
			Message: "Approve Student Leave",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (l studentLeaveHandler) End(w http.ResponseWriter, r *http.Request) {
	var result EndResponse

	ctx := r.Context()
	var in EndRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = EndResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer l.AdminActivityLogService.Create(ctx, r, time.Now(), "Student Leave", "End", nil)

	errs := l.StudentLeaveService.End(ctx, in.GetId())
	if errs != nil {
		utils.PrintError(*errs)
		result = EndResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = EndResponse{
		Meta: &Meta{
			Message: "End Student Leave",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
