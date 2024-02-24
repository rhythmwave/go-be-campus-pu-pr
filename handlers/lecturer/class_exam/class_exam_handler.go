package class_exam

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

type classExamHandler struct {
	*service.ServiceCtx
}

func (f classExamHandler) GetList(w http.ResponseWriter, r *http.Request) {
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
	defer f.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Class Exam", "Get List", nil)

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}
	classIds := []string{}
	if in.GetClassId() != "" {
		classIds = append(classIds, in.GetClassId())
	}
	data, errs := f.ClassExamService.GetList(ctx, paginationData, appConstants.AppTypeLecturer, classIds)
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
			Id:                 v.Id,
			Title:              v.Title,
			Abstraction:        utils.NullStringScan(v.Abstraction),
			FileUrl:            v.FileUrl,
			LecturerId:         v.LecturerId,
			LecturerName:       v.LecturerName,
			LecturerFrontTitle: utils.NullStringScan(v.LecturerFrontTitle),
			LecturerBackDegree: utils.NullStringScan(v.LecturerBackDegree),
			StartTime:          v.StartTime.Format(constants.DateRFC),
			EndTime:            v.EndTime.Format(constants.DateRFC),
			TotalSubmission:    v.TotalSubmission,
		})
	}

	result = GetListResponse{
		Meta: &Meta{
			Message: "Get List Class Exam",
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

func (f classExamHandler) GetSubmission(w http.ResponseWriter, r *http.Request) {
	var result GetSubmissionResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetSubmissionRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetSubmissionResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer f.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Class Exam", "Get Submission", nil)

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}
	data, errs := f.ClassExamService.GetSubmission(ctx, paginationData, in.GetClassExamId())
	if errs != nil {
		utils.PrintError(*errs)
		result = GetSubmissionResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*GetSubmissionResponseData{}
	for _, v := range data.Data {
		resultData = append(resultData, &GetSubmissionResponseData{
			Id:               utils.NullStringScan(v.Id),
			StudentId:        v.StudentId,
			NimNumber:        v.NimNumber,
			Name:             v.Name,
			StudyProgramName: v.StudyProgramName,
			FileUrl:          v.FileUrl,
			Point:            utils.NullFloatScan(v.Point),
		})
	}

	result = GetSubmissionResponse{
		Meta: &Meta{
			Message: "Get Submission Class Exam",
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

func (c classExamHandler) Create(w http.ResponseWriter, r *http.Request) {
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
	defer c.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Class Exam", "Create", nil)

	startTime, errs := utils.StringToTime(in.GetStartTime())
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
	endTime, errs := utils.StringToTime(in.GetEndTime())
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
	data := objects.CreateClassExam{
		ClassId:      in.GetClassId(),
		Title:        in.GetTitle(),
		Abstraction:  in.GetAbstraction(),
		FilePath:     in.GetFilePath(),
		FilePathType: in.GetFilePathType(),
		StartTime:    startTime,
		EndTime:      endTime,
	}
	errs = c.ClassExamService.Create(ctx, data)
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
			Message: "Create Class Exam",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (c classExamHandler) Update(w http.ResponseWriter, r *http.Request) {
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
	defer c.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Class Exam", "Update", nil)

	startTime, errs := utils.StringToTime(in.GetStartTime())
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
	endTime, errs := utils.StringToTime(in.GetEndTime())
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
	data := objects.UpdateClassExam{
		Id:           in.GetId(),
		Title:        in.GetTitle(),
		Abstraction:  in.GetAbstraction(),
		FilePath:     in.GetFilePath(),
		FilePathType: in.GetFilePathType(),
		StartTime:    startTime,
		EndTime:      endTime,
	}
	errs = c.ClassExamService.Update(ctx, data)
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
			Message: "Update Class Exam",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (c classExamHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
	defer c.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Class Exam", "Delete", nil)

	errs := c.ClassExamService.Delete(ctx, in.GetIds())
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
			Message: "Delete Class Exam",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (c classExamHandler) GradeSubmission(w http.ResponseWriter, r *http.Request) {
	var result GradeSubmissionResponse

	ctx := r.Context()
	var in GradeSubmissionRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = GradeSubmissionResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer c.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Class Exam", "Grade Submission", nil)

	data := []objects.GradeClassExamSubmission{}
	for _, v := range in.GetData() {
		data = append(data, objects.GradeClassExamSubmission{
			Id:    v.GetId(),
			Point: v.GetPoint(),
		})
	}

	errs := c.ClassExamService.GradeSubmission(ctx, in.GetClassExamId(), data)
	if errs != nil {
		utils.PrintError(*errs)
		result = GradeSubmissionResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = GradeSubmissionResponse{
		Meta: &Meta{
			Message: "Grade Submission Class Exam",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
