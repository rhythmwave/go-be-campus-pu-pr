package lecturer_leave

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

type lecturerLeaveHandler struct {
	*service.ServiceCtx
}

func (l lecturerLeaveHandler) GetList(w http.ResponseWriter, r *http.Request) {
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
	defer l.AdminActivityLogService.Create(ctx, r, time.Now(), "Lecturer Leave", "Get List", nil)

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}
	data, errs := l.LecturerLeaveService.GetList(ctx, paginationData, in.GetStudyProgramId(), in.GetIdNationalLecturer(), in.GetSemesterId(), in.GetIsActive())
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
			IdNationalLecturer:    v.IdNationalLecturer,
			FrontTitle:            utils.NullStringScan(v.FrontTitle),
			BackDegree:            utils.NullStringScan(v.BackDegree),
			SemesterSchoolYear:    v.SemesterSchoolYear,
			SemesterType:          v.SemesterType,
			DiktiStudyProgramCode: utils.NullStringScan(v.DiktiStudyProgramCode),
			StudyProgramName:      utils.NullStringScan(v.StudyProgramName),
			StudyLevelShortName:   utils.NullStringScan(v.StudyLevelShortName),
			DiktiStudyProgramType: utils.NullStringScan(v.DiktiStudyProgramType),
			StartDate:             v.StartDate.Format(constants.DateRFC),
			EndDate:               utils.SafetyDate(v.EndDate),
			PermitNumber:          v.PermitNumber,
			Purpose:               v.Purpose,
			Remarks:               v.Remarks,
			FileUrl:               v.FileUrl,
			FilePath:              utils.NullStringScan(v.FilePath),
			FilePathType:          utils.NullStringScan(v.FilePathType),
		})
	}

	result = GetListResponse{
		Meta: &Meta{
			Message: "Get List Lecturer Leave",
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

func (l lecturerLeaveHandler) Create(w http.ResponseWriter, r *http.Request) {
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
	defer l.AdminActivityLogService.Create(ctx, r, time.Now(), "Lecturer Leave", "Create", &in)

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
	data := objects.CreateLecturerLeave{
		LecturerId:   in.GetLecturerId(),
		SemesterId:   in.GetSemesterId(),
		StartDate:    startDate,
		PermitNumber: in.GetPermitNumber(),
		Purpose:      in.GetPurpose(),
		Remarks:      in.GetRemarks(),
		FilePath:     in.GetFilePath(),
		FilePathType: in.GetFilePathType(),
	}
	errs = l.LecturerLeaveService.Create(ctx, data)
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
			Message: "Create Lecturer Leave",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (l lecturerLeaveHandler) Update(w http.ResponseWriter, r *http.Request) {
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
	defer l.AdminActivityLogService.Create(ctx, r, time.Now(), "Lecturer Leave", "Update", &in)

	data := objects.UpdateLecturerLeave{
		Id:           in.GetId(),
		PermitNumber: in.GetPermitNumber(),
		Purpose:      in.GetPurpose(),
		Remarks:      in.GetRemarks(),
		FilePath:     in.GetFilePath(),
		FilePathType: in.GetFilePathType(),
	}
	errs := l.LecturerLeaveService.Update(ctx, data)
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
			Message: "Update Lecturer Leave",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (l lecturerLeaveHandler) End(w http.ResponseWriter, r *http.Request) {
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
	defer l.AdminActivityLogService.Create(ctx, r, time.Now(), "Lecturer Leave", "End", nil)

	errs := l.LecturerLeaveService.End(ctx, in.GetId())
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
			Message: "End Lecturer Leave",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
