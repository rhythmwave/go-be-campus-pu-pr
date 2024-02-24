package semester

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

type semesterHandler struct {
	*service.ServiceCtx
}

func (a semesterHandler) GetList(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Semester", "Get List", nil)

	paginationData := common.PaginationRequest{
		Limit: in.GetLimit(),
		Page:  in.GetPage(),
	}
	data, errs := a.SemesterService.GetList(ctx, paginationData, in.GetStudyProgramId(), in.GetExcludingSemesterId())
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
		curriculums := []*GetListResponseDataCurriculum{}
		for _, w := range v.Curriculums {
			curriculums = append(curriculums, &GetListResponseDataCurriculum{
				StudyProgramId:   w.StudyProgramId,
				StudyProgramName: w.StudyProgramName,
				CurriculumId:     w.CurriculumId,
				CurriculumName:   w.CurriculumName,
			})
		}

		resultData = append(resultData, &GetListResponseData{
			Id:                v.Id,
			SemesterStartYear: v.SemesterStartYear,
			SchoolYear:        v.SchoolYear,
			SemesterType:      v.SemesterType,
			IsActive:          v.IsActive,
			StartDate:         v.StartDate.Format(constants.DateRFC),
			EndDate:           v.EndDate.Format(constants.DateRFC),
			MidtermStartDate:  utils.SafetyDate(v.MidtermStartDate),
			MidtermEndDate:    utils.SafetyDate(v.MidtermEndDate),
			EndtermStartDate:  utils.SafetyDate(v.EndtermStartDate),
			EndtermEndDate:    utils.SafetyDate(v.EndtermEndDate),
			Curriculums:       curriculums,
		})
	}

	result = GetListResponse{
		Meta: &Meta{
			Message: "Get List Semester",
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

func (a semesterHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Semester", "Get Detail", nil)

	data, errs := a.SemesterService.GetDetail(ctx, in.GetId())
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

	curriculums := []*GetDetailResponseDataCurriculum{}
	for _, v := range data.Curriculums {
		curriculums = append(curriculums, &GetDetailResponseDataCurriculum{
			StudyProgramId:   v.StudyProgramId,
			StudyProgramName: v.StudyProgramName,
			CurriculumId:     v.CurriculumId,
			CurriculumName:   v.CurriculumName,
		})
	}

	result = GetDetailResponse{
		Meta: &Meta{
			Message: "Get Detail Semester",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &GetDetailResponseData{
			Id:                         data.Id,
			SemesterStartYear:          data.SemesterStartYear,
			SchoolYear:                 data.SchoolYear,
			SemesterType:               data.SemesterType,
			IsActive:                   data.IsActive,
			StartDate:                  data.StartDate.Format(constants.DateRFC),
			EndDate:                    data.EndDate.Format(constants.DateRFC),
			MidtermStartDate:           utils.SafetyDate(data.MidtermStartDate),
			MidtermEndDate:             utils.SafetyDate(data.MidtermEndDate),
			EndtermStartDate:           utils.SafetyDate(data.EndtermStartDate),
			EndtermEndDate:             utils.SafetyDate(data.EndtermEndDate),
			StudyPlanInputStartDate:    data.StudyPlanInputStartDate.Format(constants.DateRFC),
			StudyPlanInputEndDate:      data.StudyPlanInputEndDate.Format(constants.DateRFC),
			StudyPlanApprovalStartDate: data.StudyPlanApprovalStartDate.Format(constants.DateRFC),
			StudyPlanApprovalEndDate:   data.StudyPlanApprovalEndDate.Format(constants.DateRFC),
			ReferenceSemesterId:        utils.NullStringScan(data.ReferenceSemesterId),
			ReferenceSemesterStartYear: utils.NullUint32Scan(data.ReferenceSemesterStartYear),
			ReferenceSchoolYear:        utils.NullStringScan(data.ReferenceSchoolYear),
			ReferenceSemesterType:      utils.NullStringScan(data.ReferenceSemesterType),
			CheckMinimumGpa:            data.CheckMinimumGpa,
			CheckPassedCredit:          data.CheckPassedCredit,
			DefaultCredit:              data.DefaultCredit,
			Curriculums:                curriculums,
			GradingStartDate:           utils.SafetyDate(data.GradingStartDate),
			GradingEndDate:             utils.SafetyDate(data.GradingEndDate),
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a semesterHandler) GetActive(w http.ResponseWriter, r *http.Request) {
	var result GetActiveResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetActiveRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetActiveResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Semester", "Get Active", nil)

	data, errs := a.SemesterService.GetActive(ctx)
	if errs != nil {
		utils.PrintError(*errs)
		result = GetActiveResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	curriculums := []*GetActiveResponseDataCurriculum{}
	for _, v := range data.Curriculums {
		curriculums = append(curriculums, &GetActiveResponseDataCurriculum{
			StudyProgramId:   v.StudyProgramId,
			StudyProgramName: v.StudyProgramName,
			CurriculumId:     v.CurriculumId,
			CurriculumName:   v.CurriculumName,
		})
	}

	result = GetActiveResponse{
		Meta: &Meta{
			Message: "Get Active Semester",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &GetActiveResponseData{
			Id:                         data.Id,
			SemesterStartYear:          data.SemesterStartYear,
			SchoolYear:                 data.SchoolYear,
			SemesterType:               data.SemesterType,
			IsActive:                   data.IsActive,
			StartDate:                  data.StartDate.Format(constants.DateRFC),
			EndDate:                    data.EndDate.Format(constants.DateRFC),
			MidtermStartDate:           utils.SafetyDate(data.MidtermStartDate),
			MidtermEndDate:             utils.SafetyDate(data.MidtermEndDate),
			EndtermStartDate:           utils.SafetyDate(data.EndtermStartDate),
			EndtermEndDate:             utils.SafetyDate(data.EndtermEndDate),
			StudyPlanInputStartDate:    data.StudyPlanInputStartDate.Format(constants.DateRFC),
			StudyPlanInputEndDate:      data.StudyPlanInputEndDate.Format(constants.DateRFC),
			StudyPlanApprovalStartDate: data.StudyPlanApprovalStartDate.Format(constants.DateRFC),
			StudyPlanApprovalEndDate:   data.StudyPlanApprovalEndDate.Format(constants.DateRFC),
			ReferenceSemesterId:        utils.NullStringScan(data.ReferenceSemesterId),
			ReferenceSemesterStartYear: utils.NullUint32Scan(data.ReferenceSemesterStartYear),
			ReferenceSchoolYear:        utils.NullStringScan(data.ReferenceSchoolYear),
			ReferenceSemesterType:      utils.NullStringScan(data.ReferenceSemesterType),
			CheckMinimumGpa:            data.CheckMinimumGpa,
			CheckPassedCredit:          data.CheckPassedCredit,
			DefaultCredit:              data.DefaultCredit,
			Curriculums:                curriculums,
			GradingStartDate:           utils.SafetyDate(data.GradingStartDate),
			GradingEndDate:             utils.SafetyDate(data.GradingEndDate),
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a semesterHandler) Create(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Semester", "Create", &in)

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
	endDate, errs := utils.StringToTime(in.GetEndDate())
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
	midtermStartDate, errs := utils.StringToTime(in.GetMidtermStartDate())
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
	midtermEndDate, errs := utils.StringToTime(in.GetMidtermEndDate())
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
	endtermStartDate, errs := utils.StringToTime(in.GetEndtermStartDate())
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
	endtermEndDate, errs := utils.StringToTime(in.GetEndtermEndDate())
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
	studyPlanInputStartDate, errs := utils.StringToTime(in.GetStudyPlanInputStartDate())
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
	studyPlanInputEndDate, errs := utils.StringToTime(in.GetStudyPlanInputEndDate())
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
	studyPlanApprovalStartDate, errs := utils.StringToTime(in.GetStudyPlanApprovalStartDate())
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
	studyPlanApprovalEndDate, errs := utils.StringToTime(in.GetStudyPlanApprovalEndDate())
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
	gradingStartDate, errs := utils.StringToTime(in.GetGradingStartDate())
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
	gradingEndDate, errs := utils.StringToTime(in.GetGradingEndDate())
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

	curriculums := []objects.CreateSemesterCurriculum{}
	for _, v := range in.GetCurriculums() {
		curriculums = append(curriculums, objects.CreateSemesterCurriculum{
			CurriculumId: v.GetCurriculumId(),
		})
	}

	data := objects.CreateSemester{
		SemesterStartYear:          in.GetSemesterStartYear(),
		SemesterType:               in.GetSemesterType(),
		StartDate:                  startDate,
		EndDate:                    endDate,
		MidtermStartDate:           midtermStartDate,
		MidtermEndDate:             midtermEndDate,
		EndtermStartDate:           endtermStartDate,
		EndtermEndDate:             endtermEndDate,
		StudyPlanInputStartDate:    studyPlanInputStartDate,
		StudyPlanInputEndDate:      studyPlanInputEndDate,
		StudyPlanApprovalStartDate: studyPlanApprovalStartDate,
		StudyPlanApprovalEndDate:   studyPlanApprovalEndDate,
		ReferenceSemesterId:        in.GetReferenceSemesterId(),
		CheckMinimumGpa:            in.GetCheckMinimumGpa(),
		CheckPassedCredit:          in.GetCheckPassedCredit(),
		DefaultCredit:              in.GetDefaultCredit(),
		Curriculums:                curriculums,
		GradingStartDate:           gradingStartDate,
		GradingEndDate:             gradingEndDate,
	}
	errs = a.SemesterService.Create(ctx, data)
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
			Message: "Create Semester",
			Status:  http.StatusCreated,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusCreated, &result)
}

func (a semesterHandler) Update(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Semester", "Update", &in)

	startDate, errs := utils.StringToTime(in.GetStartDate())
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
	endDate, errs := utils.StringToTime(in.GetEndDate())
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
	midtermStartDate, errs := utils.StringToTime(in.GetMidtermStartDate())
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
	midtermEndDate, errs := utils.StringToTime(in.GetMidtermEndDate())
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
	endtermStartDate, errs := utils.StringToTime(in.GetEndtermStartDate())
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
	endtermEndDate, errs := utils.StringToTime(in.GetEndtermEndDate())
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
	studyPlanInputStartDate, errs := utils.StringToTime(in.GetStudyPlanInputStartDate())
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
	studyPlanInputEndDate, errs := utils.StringToTime(in.GetStudyPlanInputEndDate())
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
	studyPlanApprovalStartDate, errs := utils.StringToTime(in.GetStudyPlanApprovalStartDate())
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
	studyPlanApprovalEndDate, errs := utils.StringToTime(in.GetStudyPlanApprovalEndDate())
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
	gradingStartDate, errs := utils.StringToTime(in.GetGradingStartDate())
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
	gradingEndDate, errs := utils.StringToTime(in.GetGradingEndDate())
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

	curriculums := []objects.UpdateSemesterCurriculum{}
	for _, v := range in.GetCurriculums() {
		curriculums = append(curriculums, objects.UpdateSemesterCurriculum{
			CurriculumId: v.GetCurriculumId(),
		})
	}

	data := objects.UpdateSemester{
		Id:                         in.GetId(),
		SemesterStartYear:          in.GetSemesterStartYear(),
		SemesterType:               in.GetSemesterType(),
		StartDate:                  startDate,
		EndDate:                    endDate,
		MidtermStartDate:           midtermStartDate,
		MidtermEndDate:             midtermEndDate,
		EndtermStartDate:           endtermStartDate,
		EndtermEndDate:             endtermEndDate,
		StudyPlanInputStartDate:    studyPlanInputStartDate,
		StudyPlanInputEndDate:      studyPlanInputEndDate,
		StudyPlanApprovalStartDate: studyPlanApprovalStartDate,
		StudyPlanApprovalEndDate:   studyPlanApprovalEndDate,
		ReferenceSemesterId:        in.GetReferenceSemesterId(),
		CheckMinimumGpa:            in.GetCheckMinimumGpa(),
		CheckPassedCredit:          in.GetCheckPassedCredit(),
		DefaultCredit:              in.GetDefaultCredit(),
		Curriculums:                curriculums,
		GradingStartDate:           gradingStartDate,
		GradingEndDate:             gradingEndDate,
	}
	errs = a.SemesterService.Update(ctx, data)
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
			Message: "Update Semester",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a semesterHandler) UpdateActivation(w http.ResponseWriter, r *http.Request) {
	var result UpdateActivationResponse

	ctx := r.Context()
	var in UpdateActivationRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = UpdateActivationResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Semester", "Update Activation", &in)

	errs := a.SemesterService.UpdateActivation(ctx, in.GetId(), in.GetIsActive())
	if errs != nil {
		utils.PrintError(*errs)
		result = UpdateActivationResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = UpdateActivationResponse{
		Meta: &Meta{
			Message: "Update Activation Semester",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a semesterHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Semester", "Delete", nil)

	errs := a.SemesterService.Delete(ctx, in.GetId())
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
			Message: "Delete Semester",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
