package curriculum

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

type curriculumHandler struct {
	*service.ServiceCtx
}

func (a curriculumHandler) GetList(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Curriculum", "Get List", nil)

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}
	data, errs := a.CurriculumService.GetList(ctx, paginationData, in.GetStudyProgramId())
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
			Id:                           v.Id,
			StudyProgramId:               v.StudyProgramId,
			StudyProgramName:             v.StudyProgramName,
			DiktiStudyProgramCode:        v.DiktiStudyProgramCode,
			Name:                         v.Name,
			Year:                         v.Year,
			IdealStudyPeriod:             v.IdealStudyPeriod,
			MaximumStudyPeriod:           v.MaximumStudyPeriod,
			IsActive:                     v.IsActive,
			TotalSubject:                 v.TotalSubject,
			TotalSubjectWithPrerequisite: v.TotalSubjectWithPrerequisite,
			TotalSubjectWithEquivalency:  v.TotalSubjectWithEquivalence,
		})
	}

	result = GetListResponse{
		Meta: &Meta{
			Message: "Get List Curriculum",
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

func (a curriculumHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Curriculum", "Get Detail", nil)

	data, errs := a.CurriculumService.GetDetail(ctx, in.GetId())
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
			Message: "Get Detail Curriculum",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &GetDetailResponseData{
			Id:                    data.Id,
			StudyProgramId:        data.StudyProgramId,
			StudyProgramName:      data.StudyProgramName,
			StudyLevelShortName:   data.StudyLevelShortName,
			DiktiStudyProgramType: data.DiktiStudyProgramType,
			Name:                  data.Name,
			Year:                  data.Year,
			RectorDecisionNumber:  utils.NullStringScan(data.RectorDecisionNumber),
			RectorDecisionDate:    utils.SafetyDate(data.RectorDecisionDate),
			AggreeingParty:        utils.NullStringScan(data.AggreeingParty),
			AggreementDate:        utils.SafetyDate(data.AggreementDate),
			IdealStudyPeriod:      data.IdealStudyPeriod,
			MaximumStudyPeriod:    data.MaximumStudyPeriod,
			Remarks:               utils.NullStringScan(data.Remarks),
			IsActive:              data.IsActive,
			FinalScoreDeterminant: data.FinalScoreDeterminant,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a curriculumHandler) GetActiveByStudyProgramId(w http.ResponseWriter, r *http.Request) {
	var result GetActiveByStudyProgramIdResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetActiveByStudyProgramIdRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetActiveByStudyProgramIdResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Curriculum", "Get Detail", nil)

	data, errs := a.CurriculumService.GetActiveByStudyProgramId(ctx, in.GetStudyProgramId())
	if errs != nil {
		utils.PrintError(*errs)
		result = GetActiveByStudyProgramIdResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = GetActiveByStudyProgramIdResponse{
		Meta: &Meta{
			Message: "Get Detail Curriculum",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &GetActiveByStudyProgramIdResponseData{
			Id:                    data.Id,
			StudyProgramId:        data.StudyProgramId,
			StudyProgramName:      data.StudyProgramName,
			StudyLevelShortName:   data.StudyLevelShortName,
			DiktiStudyProgramType: data.DiktiStudyProgramType,
			Name:                  data.Name,
			Year:                  data.Year,
			RectorDecisionNumber:  utils.NullStringScan(data.RectorDecisionNumber),
			RectorDecisionDate:    utils.SafetyDate(data.RectorDecisionDate),
			AggreeingParty:        utils.NullStringScan(data.AggreeingParty),
			AggreementDate:        utils.SafetyDate(data.AggreementDate),
			IdealStudyPeriod:      data.IdealStudyPeriod,
			MaximumStudyPeriod:    data.MaximumStudyPeriod,
			Remarks:               utils.NullStringScan(data.Remarks),
			IsActive:              data.IsActive,
			FinalScoreDeterminant: data.FinalScoreDeterminant,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a curriculumHandler) Create(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Curriculum", "Create", &in)

	rectorDecisionDate, errs := utils.StringToTime(in.GetRectorDecisionDate())
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
	aggreementDate, errs := utils.StringToTime(in.GetAggreementDate())
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
	data := objects.CreateCurriculum{
		StudyProgramId:        in.GetStudyProgramId(),
		Name:                  in.GetName(),
		Year:                  in.GetYear(),
		RectorDecisionNumber:  in.GetRectorDecisionNumber(),
		RectorDecisionDate:    rectorDecisionDate,
		AggreeingParty:        in.GetAggreeingParty(),
		AggreementDate:        aggreementDate,
		IdealStudyPeriod:      in.GetIdealStudyPeriod(),
		MaximumStudyPeriod:    in.GetMaximumStudyPeriod(),
		Remarks:               in.GetRemarks(),
		IsActive:              in.GetIsActive(),
		FinalScoreDeterminant: in.GetFinalScoreDeterminant(),
	}

	errs = a.CurriculumService.Create(ctx, data)
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
			Message: "Create Curriculum",
			Status:  http.StatusCreated,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusCreated, &result)
}

func (a curriculumHandler) Update(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Curriculum", "Update", &in)

	rectorDecisionDate, errs := utils.StringToTime(in.GetRectorDecisionDate())
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
	aggreementDate, errs := utils.StringToTime(in.GetAggreementDate())
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
	data := objects.UpdateCurriculum{
		Id:                    in.GetId(),
		Name:                  in.GetName(),
		Year:                  in.GetYear(),
		RectorDecisionNumber:  in.GetRectorDecisionNumber(),
		RectorDecisionDate:    rectorDecisionDate,
		AggreeingParty:        in.GetAggreeingParty(),
		AggreementDate:        aggreementDate,
		IdealStudyPeriod:      in.GetIdealStudyPeriod(),
		MaximumStudyPeriod:    in.GetMaximumStudyPeriod(),
		Remarks:               in.GetRemarks(),
		IsActive:              in.GetIsActive(),
		FinalScoreDeterminant: in.GetFinalScoreDeterminant(),
	}
	errs = a.CurriculumService.Update(ctx, data)
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
			Message: "Update Curriculum",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a curriculumHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Curriculum", "Delete", nil)

	errs := a.CurriculumService.Delete(ctx, in.GetId())
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
			Message: "Delete Curriculum",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
