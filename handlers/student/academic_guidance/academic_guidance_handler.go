package academic_guidance

import (
	"net/http"
	"time"

	"github.com/gorilla/schema"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/service"
	"github.com/sccicitb/pupr-backend/utils"
	"github.com/sirupsen/logrus"
)

type academicGuidanceHandler struct {
	*service.ServiceCtx
}

func (f academicGuidanceHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
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
	defer f.AdminActivityLogService.Create(ctx, r, time.Now(), "Academic Guidance", "Get Detail", nil)

	data, errs := f.AcademicGuidanceService.GetDetail(ctx, "", in.GetSemesterId())
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
			Message: "Get Detail Academic Guidance",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &GetDetailResponseData{
			Id:                 data.Id,
			SemesterId:         data.SemesterId,
			LecturerId:         data.LecturerId,
			LecturerName:       data.LecturerName,
			LecturerFrontTitle: utils.NullStringScan(data.LecturerFrontTitle),
			LecturerBackDegree: utils.NullStringScan(data.LecturerBackDegree),
			DecisionNumber:     utils.NullStringScan(data.DecisionNumber),
			DecisionDate:       utils.SafetyDate(data.DecisionDate),
			TotalStudent:       data.TotalStudent,
		},
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

	data, errs := f.AcademicGuidanceService.GetSessionList(ctx, in.GetAcademicGuidanceId(), "", "")
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
