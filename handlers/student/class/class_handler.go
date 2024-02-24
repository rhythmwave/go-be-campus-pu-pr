package class

import (
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

func mapGetOfferedClassList(data []objects.GetClass) []*GetOfferedClassListResponseData {
	results := []*GetOfferedClassListResponseData{}

	dataMap := make(map[uint32][]*GetOfferedClassListResponseDataClass)
	for _, v := range data {
		lecturers := []*GetOfferedClassListResponseDataClassLecturer{}
		for _, w := range v.Lecturers {
			lecturers = append(lecturers, &GetOfferedClassListResponseDataClassLecturer{
				Name:       w.Name,
				FrontTitle: utils.NullStringScan(w.FrontTitle),
				BackDegree: utils.NullStringScan(w.BackDegree),
			})
		}

		dataMap[v.SubjectSemesterPackage] = append(dataMap[v.SubjectSemesterPackage], &GetOfferedClassListResponseDataClass{
			Id:                          v.Id,
			Name:                        v.Name,
			SubjectCode:                 v.SubjectCode,
			SubjectName:                 v.SubjectName,
			SubjectIsMandatory:          v.SubjectIsMandatory,
			SubjectTheoryCredit:         v.SubjectTheoryCredit,
			SubjectPracticumCredit:      v.SubjectPracticumCredit,
			SubjectFieldPracticumCredit: v.SubjectFieldPracticumCredit,
			MaximumParticipant:          utils.NullUint32Scan(v.MaximumParticipant),
			SubjectTotalLessonPlan:      v.SubjectTotalLessonPlan,
			Lecturers:                   lecturers,
		})
	}

	for k, v := range dataMap {
		results = append(results, &GetOfferedClassListResponseData{
			SemesterPackage: k,
			Classes:         v,
		})
	}

	return results
}

func (l studentClassHandler) GetOfferedClassList(w http.ResponseWriter, r *http.Request) {
	var result GetOfferedClassListResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetOfferedClassListRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetOfferedClassListResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer l.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Class", "Get Offered Class List", nil)

	studyProgramId := in.GetStudyProgramId()
	if studyProgramId == "" {
		profileData, errs := l.StudentService.GetSemesterSummary(ctx)
		if errs != nil {
			utils.PrintError(*errs)
			result = GetOfferedClassListResponse{
				Meta: &Meta{
					Message: errs.Err.Error(),
					Status:  uint32(errs.HttpCode),
					Code:    errs.CustomCode,
				},
			}
			utils.JSONResponse(w, errs.HttpCode, &result)
			return
		}
		if profileData.StudyProgramId == nil {
			errs = appConstants.ErrStudentHasNoStudyProgram
			utils.PrintError(*errs)
			result = GetOfferedClassListResponse{
				Meta: &Meta{
					Message: errs.Err.Error(),
					Status:  uint32(errs.HttpCode),
					Code:    errs.CustomCode,
				},
			}
			utils.JSONResponse(w, errs.HttpCode, &result)
			return
		}
		studyProgramId = utils.NullStringScan(profileData.StudyProgramId)
	}

	paginationData := common.PaginationRequest{
		Limit:  constants.DefaultUnlimited,
		Page:   constants.DefaultPage,
		SortBy: "s.semester_package",
		Sort:   constants.Ascending,
	}
	isActive := true

	requestData := objects.GetClassListRequest{
		StudyProgramId:         studyProgramId,
		SemesterId:             "",
		IsActive:               &isActive,
		ClassName:              "",
		SubjectName:            "",
		FollowSemesterIdParity: true,
	}
	data, errs := l.ClassService.GetList(ctx, paginationData, requestData)
	if errs != nil {
		utils.PrintError(*errs)
		result = GetOfferedClassListResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = GetOfferedClassListResponse{
		Meta: &Meta{
			Message: "Get Offered Class List",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: mapGetOfferedClassList(data.Data),
	}

	utils.JSONResponse(w, http.StatusOK, &result)
}

func (l studentClassHandler) GetOfferedSchedule(w http.ResponseWriter, r *http.Request) {
	var result GetOfferedScheduleResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetOfferedScheduleRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetOfferedScheduleResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer l.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Class", "Get Offered Schedule", nil)

	data, errs := l.LectureService.GetDetailByClassId(ctx, in.GetClassId(), appConstants.AppTypeStudent)
	if errs != nil {
		utils.PrintError(*errs)
		result = GetOfferedScheduleResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	singleDaySchedules := []*GetOfferedScheduleResponseDataSingleDaySchedule{}
	var midtermExamDate *time.Time
	var midtermStartTime uint32
	var midtermEndTime uint32
	var midtermRoomName string
	var endtermExamDate *time.Time
	var endtermStartTime uint32
	var endtermEndTime uint32
	var endtermRoomName string
	for _, v := range data.Lectures {
		if utils.NullBooleanScan(v.IsMidtermExam) {
			midtermExamDate = &v.LecturePlanDate
			midtermStartTime = v.StartTime
			midtermEndTime = v.EndTime
			midtermRoomName = utils.NullStringScan(v.RoomName)
		} else if utils.NullBooleanScan(v.IsEndtermExam) {
			endtermExamDate = &v.LecturePlanDate
			endtermStartTime = v.StartTime
			endtermEndTime = v.EndTime
			endtermRoomName = utils.NullStringScan(v.RoomName)
		} else {
			singleDaySchedules = append(singleDaySchedules, &GetOfferedScheduleResponseDataSingleDaySchedule{
				Date:      v.LecturePlanDate.Format(constants.DateRFC),
				StartTime: v.StartTime,
				EndTime:   v.EndTime,
				RoomName:  utils.NullStringScan(v.RoomName),
			})
		}
	}

	result = GetOfferedScheduleResponse{
		Meta: &Meta{
			Message: "Get Offered Schedule",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &GetOfferedScheduleResponseData{
			StudyProgramName:                data.StudyProgramName,
			SubjectCode:                     data.SubjectCode,
			SubjectName:                     data.SubjectName,
			SemesterPackage:                 data.SemesterPackage,
			TheoryCredit:                    data.TheoryCredit,
			PracticumCredit:                 data.PracticumCredit,
			FieldPracticumCredit:            data.FieldPracticumCredit,
			SubjectMinimumPassingGradePoint: data.SubjectMinimumPassingGradePoint,
			SubjectIsMandatory:              data.SubjectIsMandatory,
			MaximumParticipant:              utils.NullUint32Scan(data.MaximumParticipant),
			PrerequisiteSubjects:            data.PrerequisiteSubjects,
			SingleDaySchedules:              singleDaySchedules,
			MidtermExamDate:                 utils.SafetyDate(midtermExamDate),
			MidtermStartTime:                midtermStartTime,
			MidtermEndTime:                  midtermEndTime,
			MidtermRoomName:                 midtermRoomName,
			EndtermExamDate:                 utils.SafetyDate(endtermExamDate),
			EndtermStartTime:                endtermStartTime,
			EndtermEndTime:                  endtermEndTime,
			EndtermRoomName:                 endtermRoomName,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (l studentClassHandler) GetTakenClass(w http.ResponseWriter, r *http.Request) {
	var result GetTakenClassResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetTakenClassRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetTakenClassResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer l.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Class", "Get Taken Class", nil)

	data, errs := l.StudentClassService.GetList(ctx, common.PaginationRequest{Page: constants.DefaultPage, Limit: constants.DefaultUnlimited}, "", "", in.GetSemesterId(), appConstants.AppTypeStudent, nil)
	if errs != nil {
		utils.PrintError(*errs)
		result = GetTakenClassResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	// data, errs := l.LecturerService.GetTakenClass(ctx, in.GetSemesterId(), "", &classIsActive)

	resultData := []*GetTakenClassResponseData{}
	for _, v := range data.Data {
		resultData = append(resultData, &GetTakenClassResponseData{
			Id:                   v.ClassId,
			Name:                 v.ClassName,
			SubjectCode:          v.SubjectCode,
			SubjectName:          v.SubjectName,
			TheoryCredit:         v.SubjectTheoryCredit,
			PracticumCredit:      v.SubjectPracticumCredit,
			FieldPracticumCredit: v.SubjectFieldPracticumCredit,
		})
	}

	result = GetTakenClassResponse{
		Meta: &Meta{
			Message: "Get Taken Class",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: resultData,
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
