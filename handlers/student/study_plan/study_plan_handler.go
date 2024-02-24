package study_plan

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/schema"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/service"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/utils"
	"github.com/sirupsen/logrus"
)

type studyPlanHandler struct {
	*service.ServiceCtx
}

func (f studyPlanHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
	var result GetDetailResponse
	now := time.Now()

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
	defer f.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Study Plan", "Get Detail", nil)

	data, errs := f.StudyPlanService.GetDetail(ctx, "", in.GetSemesterId())
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

	classes := []*GetDetailResponseDataClass{}
	for _, v := range data.Classes {
		schedules := []*GetDetailResponseDataClassSchedule{}
		var nearestSchedule GetDetailResponseDataClassSchedule
		var nearestScheduleDiff time.Duration = time.Hour * 24 * 365
		for _, w := range v.Schedules {

			hour := int(w.EndTime / 100)
			minute := int(w.EndTime) - (hour * 100)
			scheduleTime := time.Date(w.Date.Year(), w.Date.Month(), w.Date.Day(), hour, minute, 0, 0, time.UTC)
			currentScheduleDiff := scheduleTime.Sub(now)
			if currentScheduleDiff > 0 && currentScheduleDiff < nearestScheduleDiff {
				nearestScheduleDiff = currentScheduleDiff
				nearestSchedule = GetDetailResponseDataClassSchedule{
					Date:      w.Date.Format(constants.DateRFC),
					StartTime: w.StartTime,
					EndTime:   w.EndTime,
					RoomId:    w.RoomId,
					RoomName:  utils.NullStringScan(w.RoomName),
				}
			}
			schedules = append(schedules, &GetDetailResponseDataClassSchedule{
				Date:      w.Date.Format(constants.DateRFC),
				StartTime: w.StartTime,
				EndTime:   w.EndTime,
				RoomId:    w.RoomId,
				RoomName:  utils.NullStringScan(w.RoomName),
			})
		}

		classes = append(classes, &GetDetailResponseDataClass{
			Id:                          v.Id,
			Name:                        v.Name,
			SubjectId:                   v.SubjectId,
			SubjectName:                 v.SubjectName,
			SubjectCode:                 v.SubjectCode,
			SubjectTheoryCredit:         v.SubjectTheoryCredit,
			SubjectFieldPracticumCredit: v.SubjectFieldPracticumCredit,
			SubjectPracticumCredit:      v.SubjectPracticumCredit,
			TotalLectureDone:            v.TotalLectureDone,
			TotalAttendance:             v.TotalAttendance,
			ActiveLectureId:             utils.NullStringScan(v.ActiveLectureId),
			GradePoint:                  v.GradePoint,
			GradeCode:                   utils.NullStringScan(v.GradeCode),
			ActiveLectureHasAttend:      utils.NullBooleanScan(v.ActiveLectureHasAttend),
			SubjectIsMandatory:          v.SubjectIsMandatory,
			Schedules:                   schedules,
			NearestSchedule:             &nearestSchedule,
		})
	}

	result = GetDetailResponse{
		Meta: &Meta{
			Message: "Get Detail Study Plan",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &GetDetailResponseData{
			StudyPlanInputStartDate:            data.StudyPlanInputStartDate.Format(constants.DateRFC),
			StudyPlanInputEndDate:              data.StudyPlanInputEndDate.Format(constants.DateRFC),
			Id:                                 data.Id,
			IsSubmitted:                        data.IsSubmitted,
			IsApproved:                         data.IsApproved,
			StudentId:                          data.StudentId,
			StudentNimNumber:                   data.StudentNimNumber,
			StudentName:                        data.StudentName,
			StudyProgramId:                     utils.NullStringScan(data.StudyProgramId),
			StudyProgramName:                   utils.NullStringScan(data.StudyProgramName),
			SemesterId:                         data.SemesterId,
			SemesterSchoolYear:                 data.SemesterSchoolYear,
			SemesterType:                       data.SemesterType,
			MaximumCredit:                      data.MaximumCredit,
			AcademicGuidanceLecturerId:         utils.NullStringScan(data.AcademicGuidanceLecturerId),
			AcademicGuidanceLecturerName:       utils.NullStringScan(data.AcademicGuidanceLecturerName),
			AcademicGuidanceLecturerFrontTitle: utils.NullStringScan(data.AcademicGuidanceLecturerFrontTitle),
			AcademicGuidanceLecturerBackDegree: utils.NullStringScan(data.AcademicGuidanceLecturerBackDegree),
			TotalMandatoryCredit:               data.TotalMandatoryCredit,
			TotalOptionalCredit:                data.TotalOptionalCredit,
			GradePoint:                         data.GradePoint,
			Gpa:                                utils.NullFloatScan(data.Gpa),
			IsThesis:                           data.IsThesis,
			Classes:                            classes,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (c studyPlanHandler) Create(w http.ResponseWriter, r *http.Request) {
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
	defer c.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Study Plan", "Create", nil)

	errs := c.StudyPlanService.BulkCreate(ctx, objects.BulkCreateStudyPlan{ClassIds: in.GetClassIds(), IsThesis: in.GetIsThesis()})
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
			Message: "Create Study Plan",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
