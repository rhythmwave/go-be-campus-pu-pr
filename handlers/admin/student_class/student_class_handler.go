package student_class

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

type studentClassHandler struct {
	*service.ServiceCtx
}

func (s studentClassHandler) GetList(w http.ResponseWriter, r *http.Request) {
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
	defer s.AdminActivityLogService.Create(ctx, r, time.Now(), "Student Class", "Get List", nil)

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}

	isMbkm, errs := utils.StringToBoolPointer(in.GetIsMbkm())
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
	data, errs := s.StudentClassService.GetList(ctx, paginationData, in.GetStudyPlanId(), in.GetStudentId(), in.GetSemesterId(), appConstants.AppTypeAdmin, isMbkm)
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
		schedules := []*GetListResponseDataSchedule{}
		for _, w := range v.Schedules {
			schedules = append(schedules, &GetListResponseDataSchedule{
				Date:      w.Date.Format(constants.DateRFC),
				StartTime: w.StartTime,
				EndTime:   w.EndTime,
				RoomId:    w.RoomId,
				RoomName:  utils.NullStringScan(w.RoomName),
			})
		}

		resultData = append(resultData, &GetListResponseData{
			Id:                   v.Id,
			ClassId:              v.ClassId,
			ClassName:            v.ClassName,
			SubjectId:            v.SubjectId,
			SubjectCode:          v.SubjectCode,
			SubjectName:          v.SubjectName,
			SubjectTotalCredit:   v.SubjectTheoryCredit + v.SubjectPracticumCredit + v.SubjectFieldPracticumCredit,
			SubjectRepetition:    v.SubjectRepetition,
			SubjectIsMandatory:   v.SubjectIsMandatory,
			TotalAttendance:      v.TotalAttendance,
			TotalSick:            v.TotalSick,
			TotalLeave:           v.TotalLeave,
			TotalAwol:            v.TotalAwol,
			GradePoint:           v.GradePoint,
			GradeCode:            utils.NullStringScan(v.GradeCode),
			GradedByAdminId:      utils.NullStringScan(v.GradedByAdminId),
			GradedByAdminName:    utils.NullStringScan(v.GradedByAdminName),
			GradedByLecturerId:   utils.NullStringScan(v.GradedByLecturerId),
			GradedByLecturerName: utils.NullStringScan(v.GradedByLecturerName),
			GradedAt:             utils.NullStringScan(v.GradedAt),
			AttendancePercentage: v.AttendancePercentage,
			TotalLecture:         v.TotalLecture,
			Schedules:            schedules,
		})
	}

	result = GetListResponse{
		Meta: &Meta{
			Message: "Get List Student Class",
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

func (s studentClassHandler) TransferStudentClass(w http.ResponseWriter, r *http.Request) {
	var result TransferStudentClassResponse

	ctx := r.Context()
	var in TransferStudentClassRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = TransferStudentClassResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer s.AdminActivityLogService.Create(ctx, r, time.Now(), "Student Class", "Transfer Student Class", &in)

	transferData := []objects.TransferStudentClassData{}
	for _, v := range in.GetData() {
		transferData = append(transferData, objects.TransferStudentClassData{
			StudentId:          v.GetStudentId(),
			DestinationClassId: v.GetDestinationClassId(),
		})
	}

	data := objects.TransferStudentClass{
		SourceClassId: in.GetSourceClassId(),
		Data:          transferData,
	}
	errs := s.StudentClassService.TransferStudentClass(ctx, data)
	if errs != nil {
		utils.PrintError(*errs)
		result = TransferStudentClassResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = TransferStudentClassResponse{
		Meta: &Meta{
			Message: "Transfer Student Class",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (s studentClassHandler) ReshuffleStudentClass(w http.ResponseWriter, r *http.Request) {
	var result ReshuffleStudentClassResponse

	ctx := r.Context()
	var in ReshuffleStudentClassRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = ReshuffleStudentClassResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer s.AdminActivityLogService.Create(ctx, r, time.Now(), "Student Class", "Reshuffle Student Class", &in)

	data := []objects.ReshuffleStudentClass{}
	for _, v := range in.GetData() {
		studentData := []objects.ReshuffleStudentClassStudent{}
		for _, w := range v.GetStudents() {
			studentData = append(studentData, objects.ReshuffleStudentClassStudent{
				SourceClassId: w.GetSourceClassId(),
				StudentId:     w.GetStudentId(),
			})
		}
		data = append(data, objects.ReshuffleStudentClass{
			DestinationClassId: v.GetDestinationClassId(),
			Students:           studentData,
		})
	}
	errs := s.StudentClassService.ReshuffleStudentClass(ctx, data)
	if errs != nil {
		utils.PrintError(*errs)
		result = ReshuffleStudentClassResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = ReshuffleStudentClassResponse{
		Meta: &Meta{
			Message: "Reshuffle Student Class",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (s studentClassHandler) MergeStudentClass(w http.ResponseWriter, r *http.Request) {
	var result MergeStudentClassResponse

	ctx := r.Context()
	var in MergeStudentClassRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = MergeStudentClassResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer s.AdminActivityLogService.Create(ctx, r, time.Now(), "Student Class", "Merge Student Class", &in)

	data := objects.MergeStudentClass{
		SourceClassIds:     in.GetSourceClassIds(),
		DestinationClassId: in.GetDestinationClassId(),
	}
	errs := s.StudentClassService.MergeStudentClass(ctx, data)
	if errs != nil {
		utils.PrintError(*errs)
		result = MergeStudentClassResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = MergeStudentClassResponse{
		Meta: &Meta{
			Message: "Merge Student Class",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (s studentClassHandler) BulkGradeStudentClass(w http.ResponseWriter, r *http.Request) {
	var result BulkGradeStudentClassResponse

	ctx := r.Context()
	var in BulkGradeStudentClassRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = BulkGradeStudentClassResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer s.AdminActivityLogService.Create(ctx, r, time.Now(), "Student Class", "Bulk Grade Student Class", &in)

	data := []objects.GradeStudentClass{}
	for _, v := range in.GetStudents() {
		for _, w := range v.GetGrades() {
			data = append(data, objects.GradeStudentClass{
				StudentId:             v.GetStudentId(),
				ClassGradeComponentId: w.GetClassGradeComponentId(),
				InitialGrade:          w.GetInitialGrade(),
			})
		}
	}
	errs := s.StudentClassService.BulkGradeStudentClass(ctx, in.GetClassId(), data)
	if errs != nil {
		utils.PrintError(*errs)
		result = BulkGradeStudentClassResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = BulkGradeStudentClassResponse{
		Meta: &Meta{
			Message: "Bulk Grade Student Class",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
