package yudicium

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

type yudiciumHandler struct {
	*service.ServiceCtx
}

func (y yudiciumHandler) Apply(w http.ResponseWriter, r *http.Request) {
	var result ApplyResponse

	ctx := r.Context()
	var in ApplyRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = ApplyResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer y.AdminActivityLogService.Create(ctx, r, time.Now(), "Yudicium", "Apply", &in)

	applicationDate, errs := utils.StringToTime(in.GetApplicationDate())
	if errs != nil {
		utils.PrintError(*errs)
		result = ApplyResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	data := objects.ApplyYudicium{
		StudentId:       in.GetStudentId(),
		ApplicationDate: applicationDate,
		WithThesis:      in.GetWithThesis(),
	}
	errs = y.YudiciumService.Apply(ctx, data)
	if errs != nil {
		utils.PrintError(*errs)
		result = ApplyResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = ApplyResponse{
		Meta: &Meta{
			Message: "Apply Yudicium",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (y yudiciumHandler) GetListStudent(w http.ResponseWriter, r *http.Request) {
	var result GetListStudentResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetListStudentRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetListStudentResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer y.AdminActivityLogService.Create(ctx, r, time.Now(), "Yudicium", "Get List", nil)

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}

	req := objects.GetListStudentYudiciumRequest{
		StudyProgramId:   in.GetStudyProgramId(),
		NimNumberFrom:    in.GetNimNumberFrom(),
		NimNumberTo:      in.GetNimNumberTo(),
		SemesterId:       in.GetSemesterId(),
		StudentForceFrom: in.GetStudentForceFrom(),
		StudentForceTo:   in.GetStudentForceTo(),
	}
	data, errs := y.YudiciumService.GetListStudent(ctx, paginationData, req)
	if errs != nil {
		utils.PrintError(*errs)
		result = GetListStudentResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*GetListStudentResponseData{}
	for _, v := range data.Data {
		resultData = append(resultData, &GetListStudentResponseData{
			Id:                    v.Id,
			NimNumber:             v.NimNumber,
			Name:                  v.Name,
			DiktiStudyProgramCode: utils.NullStringScan(v.DiktiStudyProgramCode),
			TotalCredit:           utils.NullUint32Scan(v.TotalCredit),
			Gpa:                   utils.NullFloatScan(v.Gpa),
			Status:                utils.NullStringScan(v.Status),
			ApplicationDate:       v.ApplicationDate.Format(constants.DateRFC),
			DoneYudicium:          v.DoneYudicium,
			StudyProgramName:      utils.NullStringScan(v.StudyProgramName),
			StudyLevelShortName:   utils.NullStringScan(v.StudyLevelShortName),
			DiktiStudyProgramType: utils.NullStringScan(v.DiktiStudyProgramType),
		})
	}

	result = GetListStudentResponse{
		Meta: &Meta{
			Message: "Get List Yudicium",
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

func (y yudiciumHandler) Do(w http.ResponseWriter, r *http.Request) {
	var result DoResponse

	ctx := r.Context()
	var in DoRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = DoResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer y.AdminActivityLogService.Create(ctx, r, time.Now(), "Yudicium", "Do", &in)

	yudiciumDate, errs := utils.StringToTime(in.GetYudiciumDate())
	if errs != nil {
		utils.PrintError(*errs)
		result = DoResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	data := objects.DoYudicium{
		YudiciumSessionId: in.GetYudiciumSessionId(),
		YudiciumNumber:    in.GetYudiciumNumber(),
		YudiciumDate:      yudiciumDate,
		StudentIds:        in.GetStudentIds(),
	}
	errs = y.YudiciumService.Do(ctx, data)
	if errs != nil {
		utils.PrintError(*errs)
		result = DoResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = DoResponse{
		Meta: &Meta{
			Message: "Do Yudicium",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
