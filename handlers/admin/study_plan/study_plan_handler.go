package study_plan

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

type studyPlanHandler struct {
	*service.ServiceCtx
}

func (s studyPlanHandler) BulkCreate(w http.ResponseWriter, r *http.Request) {
	var result BulkCreateResponse

	ctx := r.Context()
	var in BulkCreateRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = BulkCreateResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer s.AdminActivityLogService.Create(ctx, r, time.Now(), "Study Plan", "Bulk Create", &in)

	data := objects.BulkCreateStudyPlan{
		SemesterId: in.GetSemesterId(),
		StudentIds: in.GetStudentIds(),
		ClassIds:   in.GetClassIds(),
		IsThesis:   in.GetIsThesis(),
	}

	errs := s.StudyPlanService.BulkCreate(ctx, data)
	if errs != nil {
		utils.PrintError(*errs)
		result = BulkCreateResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = BulkCreateResponse{
		Meta: &Meta{
			Message: "Bulk Create Study Plan",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (s studyPlanHandler) BulkApprove(w http.ResponseWriter, r *http.Request) {
	var result BulkApproveResponse

	ctx := r.Context()
	var in BulkApproveRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = BulkApproveResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer s.AdminActivityLogService.Create(ctx, r, time.Now(), "Study Plan", "Bulk Approve", &in)

	errs := s.StudyPlanService.BulkApprove(ctx, in.GetStudyPlanIds(), in.GetIsApproved())
	if errs != nil {
		utils.PrintError(*errs)
		result = BulkApproveResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = BulkApproveResponse{
		Meta: &Meta{
			Message: "Bulk Approve Study Plan",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (s studyPlanHandler) GetList(w http.ResponseWriter, r *http.Request) {
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
	defer s.AdminActivityLogService.Create(ctx, r, time.Now(), "Study Plan", "Get List", nil)

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}
	data, errs := s.StudyPlanService.GetList(ctx, paginationData, in.GetStudentId(), in.GetSemesterId())
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
			SemesterId:            v.SemesterId,
			SemesterStartYear:     v.SemesterStartYear,
			SchoolYear:            v.SchoolYear,
			SemesterType:          v.SemesterType,
			TotalMandatoryCredit:  v.TotalMandatoryCredit,
			TotalOptionalCredit:   v.TotalOptionalCredit,
			GradePoint:            v.GradePoint,
			StudentId:             v.StudentId,
			StudentNimNumber:      v.StudentNimNumber,
			StudentName:           v.StudentName,
			StudyProgramId:        v.StudyProgramId,
			StudyProgramName:      v.StudyProgramName,
			DiktiStudyProgramCode: v.DiktiStudyProgramCode,
			DiktiStudyProgramType: v.DiktiStudyProgramType,
			StudyLevelShortName:   v.StudyLevelShortName,
			IsThesis:              v.IsThesis,
		})
	}

	result = GetListResponse{
		Meta: &Meta{
			Message: "Get List Study Plan",
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
