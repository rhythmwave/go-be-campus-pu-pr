package report

import (
	"net/http"
	"time"

	"github.com/gorilla/schema"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/service"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
	appUtils "github.com/sccicitb/pupr-backend/utils"
	"github.com/sirupsen/logrus"
)

type reportHandler struct {
	*service.ServiceCtx
}

func (a reportHandler) StudentStatus(w http.ResponseWriter, r *http.Request) {
	var result StudentStatusResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in StudentStatusRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = StudentStatusResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Report", "Student Status", nil)

	data, errs := a.ReportService.StudentStatus(ctx, in.GetSemesterId())
	if errs != nil {
		utils.PrintError(*errs)
		result = StudentStatusResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*StudentStatusResponseData{}
	for _, v := range data {
		statuses := []*StudentStatusResponseDataStatus{}
		for _, w := range v.Statuses {
			statuses = append(statuses, &StudentStatusResponseDataStatus{
				Status: w.Status,
				Total:  w.Total,
			})
		}

		resultData = append(resultData, &StudentStatusResponseData{
			StudyProgramId:        v.StudyProgramId,
			StudyProgramName:      v.StudyProgramName,
			DiktiStudyProgramCode: v.DiktiStudyProgramCode,
			DiktiStudyProgramType: v.DiktiStudyProgramType,
			StudyLevelShortName:   v.StudyLevelShortName,
			Statuses:              statuses,
		})
	}

	result = StudentStatusResponse{
		Meta: &Meta{
			Message: "Student Status Report",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: resultData,
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (f reportHandler) StudentClassGrade(w http.ResponseWriter, r *http.Request) {
	var result StudentClassGradeResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in StudentClassGradeRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = StudentClassGradeResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer f.AdminActivityLogService.Create(ctx, r, time.Now(), "Report", "Student Class Grade", nil)

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}
	data, errs := f.ReportService.StudentClassGrade(ctx, paginationData, in.GetSemesterId(), in.GetStudyProgramId())
	if errs != nil {
		utils.PrintError(*errs)
		result = StudentClassGradeResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*StudentClassGradeResponseData{}
	for _, v := range data.Data {
		grades := []*StudentClassGradeResponseDataGrade{}
		for _, w := range v.Grades {
			grades = append(grades, &StudentClassGradeResponseDataGrade{
				GradeCode: w.GradeCode,
				Total:     w.Total,
			})
		}
		resultData = append(resultData, &StudentClassGradeResponseData{
			SubjectId:   v.SubjectId,
			SubjectName: v.SubjectName,
			Grades:      grades,
		})
	}

	result = StudentClassGradeResponse{
		Meta: &Meta{
			Message: "Student Class Grade Report",
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

func (a reportHandler) StudentProvince(w http.ResponseWriter, r *http.Request) {
	var result StudentProvinceResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in StudentProvinceRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = StudentProvinceResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Report", "Student Province", nil)

	data, errs := a.ReportService.StudentProvince(ctx, in.GetStudyProgramId(), in.GetStudentForceFrom(), in.GetStudentForceTo())
	if errs != nil {
		utils.PrintError(*errs)
		result = StudentProvinceResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*StudentProvinceResponseData{}
	for _, v := range data {
		studentForces := []*StudentProvinceResponseDataStudentForce{}
		for _, w := range v.StudentForces {
			studentForces = append(studentForces, &StudentProvinceResponseDataStudentForce{
				StudentForce: w.StudentForce,
				Total:        w.Total,
			})
		}

		resultData = append(resultData, &StudentProvinceResponseData{
			ProvinceId:    v.ProvinceId,
			ProvinceName:  v.ProvinceName,
			StudentForces: studentForces,
		})
	}

	result = StudentProvinceResponse{
		Meta: &Meta{
			Message: "Student Province Report",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: resultData,
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a reportHandler) StudentSchoolProvince(w http.ResponseWriter, r *http.Request) {
	var result StudentSchoolProvinceResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in StudentSchoolProvinceRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = StudentSchoolProvinceResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Report", "Student School Province", nil)

	data, errs := a.ReportService.StudentSchoolProvince(ctx, in.GetStudyProgramId(), in.GetStudentForceFrom(), in.GetStudentForceTo())
	if errs != nil {
		utils.PrintError(*errs)
		result = StudentSchoolProvinceResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*StudentSchoolProvinceResponseData{}
	for _, v := range data {
		studentForces := []*StudentSchoolProvinceResponseDataStudentForce{}
		for _, w := range v.StudentForces {
			studentForces = append(studentForces, &StudentSchoolProvinceResponseDataStudentForce{
				StudentForce: w.StudentForce,
				Total:        w.Total,
			})
		}

		resultData = append(resultData, &StudentSchoolProvinceResponseData{
			ProvinceId:    v.ProvinceId,
			ProvinceName:  v.ProvinceName,
			StudentForces: studentForces,
		})
	}

	result = StudentSchoolProvinceResponse{
		Meta: &Meta{
			Message: "Student School Province Report",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: resultData,
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a reportHandler) GpaDistribution(w http.ResponseWriter, r *http.Request) {
	result := GpaDistributionResponse{
		Meta: &Meta{
			Message: "Gpa Distribution Report",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: []*GpaDistributionResponseData{
			{
				GraduationSchoolYear:       appUtils.GenerateSchoolYear(2022),
				TotalLessThan_25:           0,
				PercentageLessThan_25:      0,
				TotalBetween_25And_30:      0,
				PercentageBetween_25And_30: 0,
				TotalMoreThan_30:           0,
				PercentageMoreThan_30:      0,
				TotalGraduates:             0,
				AverageGpa:                 0,
			},
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a reportHandler) StudyDurationDistribution(w http.ResponseWriter, r *http.Request) {
	result := StudyDurationDistributionResponse{
		Meta: &Meta{
			Message: "Study Duration Distribution Report",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: []*StudyDurationDistributionResponseData{
			{
				GraduationSchoolYear:      appUtils.GenerateSchoolYear(2022),
				TotalLessThan_3:           0,
				PercentageLessThan_3:      0,
				TotalBetween_3And_35:      0,
				PercentageBetween_3And_35: 0,
				TotalBetween_35And_4:      0,
				PercentageBetween_35And_4: 0,
				TotalMoreThan_4:           0,
				PercentageMoreThan_4:      0,
				TotalGraduates:            0,
				AverageStudyDuration:      0,
			},
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a reportHandler) ThesisDurationDistribution(w http.ResponseWriter, r *http.Request) {
	result := ThesisDurationDistributionResponse{
		Meta: &Meta{
			Message: "Thesis Duration Distribution Report",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: []*ThesisDurationDistributionResponseData{
			{
				GraduationSchoolYear:     appUtils.GenerateSchoolYear(2022),
				TotalLessThan_1:          0,
				PercentageLessThan_1:     0,
				TotalBetween_1And_2:      0,
				PercentageBetween_1And_2: 0,
				TotalMoreThan_2:          0,
				PercentageMoreThan_2:     0,
				TotalGraduates:           0,
				AverageThesisDuration:    0,
			},
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a reportHandler) StudentStatusSummary(w http.ResponseWriter, r *http.Request) {
	result := StudentStatusSummaryResponse{
		Meta: &Meta{
			Message: "Student Status Summary Report",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: []*StudentStatusSummaryResponseData{
			{
				StudentForce:       0,
				TotalActive:        0,
				PercentageActive:   0,
				TotalOut:           0,
				PercentageOut:      0,
				TotalGraduate:      0,
				PercentageGraduate: 0,
				TotalStudent:       0,
				TotalThesisStudent: 0,
			},
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a reportHandler) SubjectSummary(w http.ResponseWriter, r *http.Request) {
	result := SubjectSummaryResponse{
		Meta: &Meta{
			Message: "Subject Summary Report",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: []*SubjectSummaryResponseData{
			{
				SubjectId:              "",
				SubjectName:            "Bahasa Indonesia",
				TotalParticipant:       0,
				TotalRepeatParticipant: 0,
				TotalClass:             0,
				Grades: []*SubjectSummaryResponseDataGrade{
					{
						GradeCode: "A",
						Total:     0,
					},
				},
			},
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
