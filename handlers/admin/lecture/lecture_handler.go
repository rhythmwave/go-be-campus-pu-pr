package lecture

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

type lectureHandler struct {
	*service.ServiceCtx
}

func (l lectureHandler) GetList(w http.ResponseWriter, r *http.Request) {
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
	defer l.AdminActivityLogService.Create(ctx, r, time.Now(), "Lecture", "Get List", nil)

	paginationData := common.PaginationRequest{
		Limit: in.GetLimit(),
		Page:  in.GetPage(),
	}

	hasActualLecture, errs := utils.StringToBoolPointer(in.GetHasActualLecture())
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
	isExam, errs := utils.StringToBoolPointer(in.GetIsExam())
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
	data, errs := l.LectureService.GetList(ctx, paginationData, in.GetClassId(), in.GetSemesterId(), hasActualLecture, isExam, in.GetExamType())
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

		examSupervisors := []*GetListResponseDataExamSupervisor{}
		for _, w := range v.ExamSupervisors {
			examSupervisors = append(examSupervisors, &GetListResponseDataExamSupervisor{
				Id:                     w.Id,
				Name:                   w.Name,
				FrontTitle:             utils.NullStringScan(w.FrontTitle),
				BackDegree:             utils.NullStringScan(w.BackDegree),
				ExamSupervisorRoleId:   w.ExamSupervisorRoleId,
				ExamSupervisorRoleName: w.ExamSupervisorRoleName,
			})
		}
		resultData = append(resultData, &GetListResponseData{
			Id:                            v.Id,
			LecturePlanDate:               v.LecturePlanDate.Format(constants.DateRFC),
			LecturePlanDayOfWeek:          v.LecturePlanDayOfWeek,
			LecturePlanStartTime:          v.LecturePlanStartTime,
			LecturePlanEndTime:            v.LecturePlanEndTime,
			LectureActualDate:             utils.SafetyDate(v.LectureActualDate),
			LectureActualDayOfWeek:        utils.NullUint32Scan(v.LectureActualDayOfWeek),
			LectureActualStartTime:        utils.NullUint32Scan(v.LectureActualStartTime),
			LectureActualEndTime:          utils.NullUint32Scan(v.LectureActualEndTime),
			LecturerId:                    utils.NullStringScan(v.LecturerId),
			LecturerName:                  utils.NullStringScan(v.LecturerName),
			ForeignLecturerName:           utils.NullStringScan(v.ForeignLecturerName),
			ForeignLecturerSourceInstance: utils.NullStringScan(v.ForeignLecturerSourceInstance),
			IsOriginalLecturer:            utils.NullBooleanScan(v.IsOriginalLecturer),
			ClassId:                       utils.NullStringScan(v.ClassId),
			ClassName:                     utils.NullStringScan(v.ClassName),
			RoomId:                        v.RoomId,
			RoomName:                      utils.NullStringScan(v.RoomName),
			IsMidtermExam:                 utils.NullBooleanScan(v.IsMidtermExam),
			IsEndtermExam:                 utils.NullBooleanScan(v.IsEndtermExam),
			IsTheoryExam:                  utils.NullBooleanScan(v.IsTheoryExam),
			IsPracticumExam:               utils.NullBooleanScan(v.IsPracticumExam),
			IsFieldPracticumExam:          utils.NullBooleanScan(v.IsFieldPracticumExam),
			BuildingId:                    v.BuildingId,
			BuildingName:                  v.BuildingName,
			SubjectCode:                   v.SubjectCode,
			SubjectName:                   v.SubjectName,
			TotalParticipant:              v.TotalParticipant,
			ExamSupervisors:               examSupervisors,
		})
	}

	result = GetListResponse{
		Meta: &Meta{
			Message: "Get List Lecture",
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

func (l lectureHandler) BulkCreate(w http.ResponseWriter, r *http.Request) {
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
	defer l.AdminActivityLogService.Create(ctx, r, time.Now(), "Lecture", "Bulk Create", &in)

	var lecturePlans []objects.CreateLecturePlan
	for _, v := range in.GetLecturePlans() {
		lecturePlanDate, errs := utils.StringToTime(v.GetLecturePlanDate())
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

		var examSupervisors []objects.CreateLecturePlanExamSupervisor
		for _, w := range v.ExamSupervisors {
			examSupervisors = append(examSupervisors, objects.CreateLecturePlanExamSupervisor{
				ExamSupervisorId:     w.GetExamSupervisorId(),
				ExamSupervisorRoleId: w.GetExamSupervisorRoleId(),
			})
		}

		lecturePlans = append(lecturePlans, objects.CreateLecturePlan{
			LecturePlanDate:      lecturePlanDate,
			LecturePlanStartTime: v.GetLecturePlanStartTime(),
			LecturePlanEndTime:   v.GetLecturePlanEndTime(),
			RoomId:               v.GetRoomId(),
			LecturerId:           v.GetLecturerId(),
			IsExam:               v.GetIsExam(),
			IsTheoryExam:         v.GetIsTheoryExam(),
			IsPracticumExam:      v.GetIsPracticumExam(),
			IsFieldPracticumExam: v.GetIsFieldPracticumExam(),
			IsMidtermExam:        v.GetIsMidtermExam(),
			IsEndtermExam:        v.GetIsEndtermExam(),
			ExamSupervisors:      examSupervisors,
		})
	}

	errs := l.LectureService.BulkCreate(ctx, objects.CreateLecture{
		ClassId:      in.GetClassId(),
		LecturePlans: lecturePlans,
	})
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
			Message: "Bulk Create Lecture",
			Status:  http.StatusCreated,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusCreated, &result)
}

func (l lectureHandler) Update(w http.ResponseWriter, r *http.Request) {
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
	defer l.AdminActivityLogService.Create(ctx, r, time.Now(), "Lecture", "Update", &in)

	participantData := []objects.UpdateLectureParticipant{}
	for _, v := range in.GetParticipants() {
		participantData = append(participantData, objects.UpdateLectureParticipant{
			StudentId: v.GetStudentId(),
			IsAttend:  v.GetIsAttend(),
			IsSick:    v.GetIsSick(),
			IsLeave:   v.GetIsLeave(),
			IsAwol:    v.GetIsAwol(),
		})
	}

	examSupervisorData := []objects.UpdateLectureExamSupervisor{}
	for _, v := range in.GetExamSupervisors() {
		examSupervisorData = append(examSupervisorData, objects.UpdateLectureExamSupervisor{
			ExamSupervisorId:     v.GetExamSupervisorId(),
			ExamSupervisorRoleId: v.GetExamSupervisorRoleId(),
		})
	}

	lecturePlanDate, errs := utils.StringToTime(in.GetLecturePlanDate())
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

	data := objects.UpdateLecture{
		Id:                            in.GetId(),
		RoomId:                        in.GetRoomId(),
		LecturerId:                    in.GetLecturerId(),
		ForeignLecturerName:           in.GetForeignLecturerName(),
		ForeignLecturerSourceInstance: in.GetForeignLecturerSourceInstance(),
		LecturePlanDate:               lecturePlanDate,
		LecturePlanStartTime:          in.GetLecturePlanStartTime(),
		LecturePlanEndTime:            in.GetLecturePlanEndTime(),
		LectureTheme:                  in.GetLectureTheme(),
		LectureSubject:                in.GetLectureSubject(),
		Remarks:                       in.GetRemarks(),
		Participants:                  participantData,
		ExamSupervisors:               examSupervisorData,
	}
	errs = l.LectureService.Update(ctx, data)
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
			Message: "Update Lecture",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (l lectureHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
	defer l.AdminActivityLogService.Create(ctx, r, time.Now(), "Lecture", "Delete", &in)

	errs := l.LectureService.Delete(ctx, in.GetId())
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
			Message: "Delete Lecture",
			Status:  http.StatusCreated,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusCreated, &result)
}

func (l lectureHandler) ResetParticipation(w http.ResponseWriter, r *http.Request) {
	var result ResetParticipationResponse

	ctx := r.Context()
	var in ResetParticipationRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = ResetParticipationResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer l.AdminActivityLogService.Create(ctx, r, time.Now(), "Lecture", "Reset Participation", &in)

	errs := l.LectureService.ResetParticipation(ctx, in.GetId())
	if errs != nil {
		utils.PrintError(*errs)
		result = ResetParticipationResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = ResetParticipationResponse{
		Meta: &Meta{
			Message: "Reset Participation Lecture",
			Status:  http.StatusCreated,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusCreated, &result)
}

func (l lectureHandler) GetStudentParticipation(w http.ResponseWriter, r *http.Request) {
	var result GetStudentParticipationResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetStudentParticipationRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetStudentParticipationResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer l.AdminActivityLogService.Create(ctx, r, time.Now(), "Lecture", "Get Student Participation", nil)

	paginationData := common.PaginationRequest{
		Limit: in.GetLimit(),
		Page:  in.GetPage(),
	}
	data, errs := l.LectureService.GetStudentParticipation(ctx, paginationData, in.GetClassId(), in.GetStudentId())
	if errs != nil {
		utils.PrintError(*errs)
		result = GetStudentParticipationResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*GetStudentParticipationResponseData{}
	for _, v := range data.Data {
		resultData = append(resultData, &GetStudentParticipationResponseData{
			Id:                     v.Id,
			LecturePlanDate:        v.LecturePlanDate.Format(constants.DateRFC),
			LecturePlanDayOfWeek:   v.LecturePlanDayOfWeek,
			LecturePlanStartTime:   v.LecturePlanStartTime,
			LecturePlanEndTime:     v.LecturePlanEndTime,
			LectureActualDate:      utils.SafetyDate(v.LectureActualDate),
			LectureActualDayOfWeek: utils.NullUint32Scan(v.LectureActualDayOfWeek),
			LectureActualStartTime: utils.NullUint32Scan(v.LectureActualStartTime),
			LectureActualEndTime:   utils.NullUint32Scan(v.LectureActualEndTime),
			IsAttend:               utils.NullBooleanScan(v.IsAttend),
			IsSick:                 utils.NullBooleanScan(v.IsSick),
			IsLeave:                utils.NullBooleanScan(v.IsLeave),
			IsAwol:                 utils.NullBooleanScan(v.IsAwol),
		})
	}

	result = GetStudentParticipationResponse{
		Meta: &Meta{
			Message: "Get Student Participation Lecture",
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

func (l lectureHandler) GetCalendar(w http.ResponseWriter, r *http.Request) {
	var result GetCalendarResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetCalendarRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetCalendarResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer l.AdminActivityLogService.Create(ctx, r, time.Now(), "Lecture", "Get Calendar", nil)

	req := objects.GetLectureCalendarRequest{
		Year:       in.GetYear(),
		Month:      in.GetMonth(),
		RoomId:     in.GetRoomId(),
		LecturerId: in.GetLecturerId(),
		ClassId:    in.GetClassId(),
	}
	data, errs := l.LectureService.GetLectureCalendar(ctx, req)
	if errs != nil {
		utils.PrintError(*errs)
		result = GetCalendarResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*GetCalendarResponseData{}
	for _, v := range data {
		lectures := []*GetCalendarResponseDataLecture{}

		for _, w := range v.Lectures {
			lectures = append(lectures, &GetCalendarResponseDataLecture{
				LecturePlanStartTime:          w.LecturePlanStartTime,
				LecturePlanEndTime:            w.LecturePlanEndTime,
				ClassId:                       w.ClassId,
				ClassName:                     w.ClassName,
				RoomId:                        w.RoomId,
				RoomName:                      w.RoomName,
				LecturerId:                    utils.NullStringScan(w.LecturerId),
				LecturerName:                  utils.NullStringScan(w.LecturerName),
				LecturerFrontTitle:            utils.NullStringScan(w.LecturerFrontTitle),
				LecturerBackDegree:            utils.NullStringScan(w.LecturerBackDegree),
				ForeignLecturerName:           utils.NullStringScan(w.ForeignLecturerName),
				ForeignLecturerSourceInstance: utils.NullStringScan(w.ForeignLecturerSourceInstance),
			})
		}

		resultData = append(resultData, &GetCalendarResponseData{
			Date:     v.Date.Format(constants.DateFormatStd),
			Lectures: lectures,
		})
	}

	result = GetCalendarResponse{
		Meta: &Meta{
			Message: "Get Calendar Lecture",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: resultData,
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
