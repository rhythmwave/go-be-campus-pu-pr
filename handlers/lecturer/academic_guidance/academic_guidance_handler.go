package academic_guidance

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

type academicGuidanceHandler struct {
	*service.ServiceCtx
}

func (f academicGuidanceHandler) GetListStudent(w http.ResponseWriter, r *http.Request) {
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
	defer f.AdminActivityLogService.Create(ctx, r, time.Now(), "Academic Guidance", "Get List Student", nil)

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}
	data, errs := f.AcademicGuidanceService.GetListStudent(ctx, paginationData, "", in.GetSemesterId())
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
			Id:                      v.Id,
			NimNumber:               v.NimNumber,
			StudentForce:            utils.NullStringScan(v.StudentForce),
			Name:                    v.Name,
			Status:                  utils.NullStringScan(v.Status),
			StudyPlanFormIsApproved: utils.NullBooleanScan(v.StudyPlanFormIsApproved),
		})
	}

	result = GetListStudentResponse{
		Meta: &Meta{
			Message: "Get List Student Academic Guidance",
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

func (f academicGuidanceHandler) GetSessionList(w http.ResponseWriter, r *http.Request) {
	var result GetSessionListResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetSessionListRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetSessionListResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer f.AdminActivityLogService.Create(ctx, r, time.Now(), "Academic Guidance", "Get Session List", nil)

	data, errs := f.AcademicGuidanceService.GetSessionList(ctx, "", in.GetSemesterId(), "")
	if errs != nil {
		utils.PrintError(*errs)
		result = GetSessionListResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*GetSessionListResponseData{}
	for _, v := range data {
		files := []*GetSessionListResponseDataFile{}
		for _, w := range v.Files {
			files = append(files, &GetSessionListResponseDataFile{
				Id:           w.Id,
				Title:        w.Title,
				FileUrl:      w.FileUrl,
				FilePath:     w.FilePath,
				FilePathType: w.FilePathType,
			})
		}

		students := []*GetSessionListResponseDataStudent{}
		for _, w := range v.Students {
			students = append(students, &GetSessionListResponseDataStudent{
				Id:        w.Id,
				Name:      w.Name,
				NimNumber: w.NimNumber,
			})
		}

		resultData = append(resultData, &GetSessionListResponseData{
			Id:                 v.Id,
			AcademicGuidanceId: v.AcademicGuidanceId,
			Subject:            v.Subject,
			SessionDate:        v.SessionDate.Format(constants.DateRFC),
			Summary:            v.Summary,
			Files:              files,
			Students:           students,
		})
	}

	result = GetSessionListResponse{
		Meta: &Meta{
			Message: "Get Session List Academic Guidance",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: resultData,
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (f academicGuidanceHandler) CreateSession(w http.ResponseWriter, r *http.Request) {
	var result CreateSessionResponse

	ctx := r.Context()
	var in CreateSessionRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = CreateSessionResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer f.AdminActivityLogService.Create(ctx, r, time.Now(), "Academic Guidance", "Create Session", &in)

	files := []objects.CreateAcademicGuidanceSessionFile{}
	for _, v := range in.GetFiles() {
		files = append(files, objects.CreateAcademicGuidanceSessionFile{
			Title:        v.GetTitle(),
			FilePath:     v.GetFilePath(),
			FilePathType: v.GetFilePathType(),
		})
	}

	sessionDate, errs := utils.StringToTime(in.GetSessionDate())
	if errs != nil {
		utils.PrintError(*errs)
		result = CreateSessionResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	data := objects.CreateAcademicGuidanceSession{
		SemesterId:  in.GetSemesterId(),
		Subject:     in.GetSubject(),
		SessionDate: sessionDate,
		Summary:     in.GetSummary(),
		Files:       files,
		StudentIds:  in.GetStudentIds(),
	}
	errs = f.AcademicGuidanceService.CreateSession(ctx, data)
	if errs != nil {
		utils.PrintError(*errs)
		result = CreateSessionResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = CreateSessionResponse{
		Meta: &Meta{
			Message: "Create Session Academic Guidance",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (f academicGuidanceHandler) UpdateSession(w http.ResponseWriter, r *http.Request) {
	var result UpdateSessionResponse

	ctx := r.Context()
	var in UpdateSessionRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = UpdateSessionResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer f.AdminActivityLogService.Create(ctx, r, time.Now(), "Academic Guidance", "Update Session", &in)

	files := []objects.UpdateAcademicGuidanceSessionFile{}
	for _, v := range in.GetFiles() {
		files = append(files, objects.UpdateAcademicGuidanceSessionFile{
			Title:        v.GetTitle(),
			FilePath:     v.GetFilePath(),
			FilePathType: v.GetFilePathType(),
		})
	}

	sessionDate, errs := utils.StringToTime(in.GetSessionDate())
	if errs != nil {
		utils.PrintError(*errs)
		result = UpdateSessionResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	data := objects.UpdateAcademicGuidanceSession{
		Id:          in.GetId(),
		Subject:     in.GetSubject(),
		SessionDate: sessionDate,
		Summary:     in.GetSummary(),
		Files:       files,
		StudentIds:  in.GetStudentIds(),
	}
	errs = f.AcademicGuidanceService.UpdateSession(ctx, data)
	if errs != nil {
		utils.PrintError(*errs)
		result = UpdateSessionResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = UpdateSessionResponse{
		Meta: &Meta{
			Message: "Update Session Academic Guidance",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (f academicGuidanceHandler) DeleteSession(w http.ResponseWriter, r *http.Request) {
	var result DeleteSessionResponse

	ctx := r.Context()
	var in DeleteSessionRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = DeleteSessionResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer f.AdminActivityLogService.Create(ctx, r, time.Now(), "Academic Guidance", "Delete Session", &in)

	errs := f.AcademicGuidanceService.DeleteSession(ctx, in.GetId())
	if errs != nil {
		utils.PrintError(*errs)
		result = DeleteSessionResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = DeleteSessionResponse{
		Meta: &Meta{
			Message: "Delete Session Academic Guidance",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
