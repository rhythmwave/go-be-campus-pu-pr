package lecturer_mutation

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

type lecturerMutationHandler struct {
	*service.ServiceCtx
}

func (l lecturerMutationHandler) GetList(w http.ResponseWriter, r *http.Request) {
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
	defer l.AdminActivityLogService.Create(ctx, r, time.Now(), "Lecturer Mutation", "Get List", nil)

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}
	data, errs := l.LecturerMutationService.GetList(ctx, paginationData, in.GetStudyProgramId(), in.GetIdNationalLecturer(), in.GetSemesterId())
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
			MutationDate:          v.MutationDate.Format(constants.DateRFC),
			DecisionNumber:        v.DecisionNumber,
			Destination:           v.Destination,
		})
	}

	result = GetListResponse{
		Meta: &Meta{
			Message: "Get List Lecturer Mutation",
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

func (l lecturerMutationHandler) Create(w http.ResponseWriter, r *http.Request) {
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
	defer l.AdminActivityLogService.Create(ctx, r, time.Now(), "Lecturer Mutation", "Create", &in)

	mutationDate, errs := utils.StringToTime(in.GetMutationDate())
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
	data := objects.CreateLecturerMutation{
		LecturerId:     in.GetLecturerId(),
		SemesterId:     in.GetSemesterId(),
		MutationDate:   mutationDate,
		DecisionNumber: in.GetDecisionNumber(),
		Destination:    in.GetDestination(),
	}
	errs = l.LecturerMutationService.Create(ctx, data)
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
			Message: "Create Lecturer Mutation",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
