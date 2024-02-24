package lecture

import (
	"net/http"
	"time"

	"github.com/gorilla/schema"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/service"
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
	defer l.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Lecture", "Get List", nil)

	paginationData := common.PaginationRequest{
		Limit: constants.DefaultUnlimited,
		Page:  constants.DefaultPage,
	}
	data, errs := l.LectureService.GetList(ctx, paginationData, in.GetClassId(), "", nil, nil, "")
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
			Id:                               v.Id,
			LecturePlanDate:                  v.LecturePlanDate.Format(constants.DateRFC),
			LecturePlanDayOfWeek:             v.LecturePlanDayOfWeek,
			LecturePlanStartTime:             v.LecturePlanStartTime,
			LecturePlanEndTime:               v.LecturePlanEndTime,
			LectureActualDate:                utils.SafetyDate(v.LectureActualDate),
			LectureActualDayOfWeek:           utils.NullUint32Scan(v.LectureActualDayOfWeek),
			LectureActualStartTime:           utils.NullUint32Scan(v.LectureActualStartTime),
			LectureActualEndTime:             utils.NullUint32Scan(v.LectureActualEndTime),
			IsManualParticipation:            utils.NullBooleanScan(v.IsManualParticipation),
			AutonomousParticipationStartTime: utils.SafetyDate(v.AutonomousParticipationStartTime),
			AutonomousParticipationEndTime:   utils.SafetyDate(v.AutonomousParticipationEndTime),
			AttendingParticipant:             v.AttendingParticipant,
			UpdatedAt:                        utils.SafetyDate(v.UpdatedAt),
		})
	}

	result = GetListResponse{
		Meta: &Meta{
			Message: "Get List Lecture",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: resultData,
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (l lectureHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
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
	defer l.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Lecture", "Get Detail", nil)

	data, errs := l.LectureService.GetDetail(ctx, in.GetId())
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

	students := []*GetDetailResponseStudent{}
	for _, v := range data.Students {
		students = append(students, &GetDetailResponseStudent{
			Id:        v.Id,
			NimNumber: v.NimNumber,
			Name:      v.Name,
			IsAttend:  utils.NullBooleanScan(v.IsAttend),
			IsSick:    utils.NullBooleanScan(v.IsSick),
			IsLeave:   utils.NullBooleanScan(v.IsLeave),
			IsAwol:    utils.NullBooleanScan(v.IsAwol),
		})
	}

	result = GetDetailResponse{
		Meta: &Meta{
			Message: "Get Detail Lecture",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &GetDetailResponseData{
			Id:                 data.Id,
			SubjectId:          data.SubjectId,
			SubjectName:        data.SubjectName,
			SemesterId:         data.SemesterId,
			SemesterSchoolYear: data.SemesterSchoolYear,
			SemesterType:       data.SemesterType,
			LectureTheme:       utils.NullStringScan(data.LectureTheme),
			LectureSubject:     utils.NullStringScan(data.LectureSubject),
			Remarks:            utils.NullStringScan(data.Remarks),
			ClassId:            utils.NullStringScan(data.ClassId),
			StudyProgramId:     data.StudyProgramId,
			StudyProgramName:   data.StudyProgramName,
			Students:           students,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (l lectureHandler) Update(w http.ResponseWriter, r *http.Request) {
	// var result UpdateResponse

	// ctx := r.Context()
	// var in UpdateRequest
	// err := json.NewDecoder(r.Body).Decode(&in)
	// if err != nil {
	// 	logrus.Errorln(err)
	// 	result = UpdateResponse{
	// 		Meta: &Meta{
	// 			Message: err.Error(),
	// 			Status:  http.StatusInternalServerError,
	// 			Code:    constants.DefaultCustomErrorCode,
	// 		},
	// 	}
	// 	utils.JSONResponse(w, http.StatusInternalServerError, &result)
	// 	return
	// }
	// defer l.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Lecture", "Update", &in)

	// participantData := []objects.UpdateLectureParticipant{}
	// for _, v := range in.GetParticipants() {
	// 	participantData = append(participantData, objects.UpdateLectureParticipant{
	// 		StudentId: v.GetStudentId(),
	// 		IsAttend:  v.GetIsAttend(),
	// 		IsSick:    v.GetIsSick(),
	// 		IsLeave:   v.GetIsLeave(),
	// 		IsAwol:    v.GetIsAwol(),
	// 	})
	// }

	// autonomousParticipationEndTime, errs := utils.StringToTime(in.GetAutonomousParticipationEndTime())
	// if errs != nil {
	// 	utils.PrintError(*errs)
	// 	result = UpdateResponse{
	// 		Meta: &Meta{
	// 			Message: errs.Err.Error(),
	// 			Status:  uint32(errs.HttpCode),
	// 			Code:    errs.CustomCode,
	// 		},
	// 	}
	// 	utils.JSONResponse(w, errs.HttpCode, &result)
	// 	return
	// }

	// data := objects.UpdateLecture{
	// 	Id:                             in.GetId(),
	// 	LectureTheme:                   in.GetLectureTheme(),
	// 	LectureSubject:                 in.GetLectureSubject(),
	// 	Remarks:                        in.GetRemarks(),
	// 	IsManualParticipation:          in.GetIsManualParticipation(),
	// 	AutonomousParticipationEndTime: autonomousParticipationEndTime,
	// 	Participants:                   participantData,
	// }
	// errs = l.LectureService.Update(ctx, data)
	// if errs != nil {
	// 	utils.PrintError(*errs)
	// 	result = UpdateResponse{
	// 		Meta: &Meta{
	// 			Message: errs.Err.Error(),
	// 			Status:  uint32(errs.HttpCode),
	// 			Code:    errs.CustomCode,
	// 		},
	// 	}
	// 	utils.JSONResponse(w, errs.HttpCode, &result)
	// 	return
	// }

	// result = UpdateResponse{
	// 	Meta: &Meta{
	// 		Message: "Update Lecture",
	// 		Status:  http.StatusOK,
	// 		Code:    constants.SuccessCode,
	// 	},
	// }
	// utils.JSONResponse(w, http.StatusOK, &result)
}
