package class

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

type lecturerClassHandler struct {
	*service.ServiceCtx
}

func mapGetActiveSemesterClassList(data []objects.GetClass) []*GetActiveSemesterClassListResponseData {
	results := []*GetActiveSemesterClassListResponseData{}

	dataMap := make(map[uint32][]*GetActiveSemesterClassListResponseDataClass)
	for _, v := range data {
		lecturers := []*GetActiveSemesterClassListResponseDataClassLecturer{}
		for _, w := range v.Lecturers {
			lecturers = append(lecturers, &GetActiveSemesterClassListResponseDataClassLecturer{
				Name:       w.Name,
				FrontTitle: utils.NullStringScan(w.FrontTitle),
				BackDegree: utils.NullStringScan(w.BackDegree),
			})
		}

		dataMap[v.SubjectSemesterPackage] = append(dataMap[v.SubjectSemesterPackage], &GetActiveSemesterClassListResponseDataClass{
			Id:                          v.Id,
			Name:                        v.Name,
			SubjectCode:                 v.SubjectCode,
			SubjectName:                 v.SubjectName,
			SubjectIsMandatory:          v.SubjectIsMandatory,
			SubjectTheoryCredit:         v.SubjectTheoryCredit,
			SubjectPracticumCredit:      v.SubjectPracticumCredit,
			SubjectFieldPracticumCredit: v.SubjectFieldPracticumCredit,
			MaximumParticipant:          utils.NullUint32Scan(v.MaximumParticipant),
			Lecturers:                   lecturers,
		})
	}

	for k, v := range dataMap {
		results = append(results, &GetActiveSemesterClassListResponseData{
			SemesterPackage: k,
			Classes:         v,
		})
	}

	return results
}

func (l lecturerClassHandler) GetActiveSemesterClassList(w http.ResponseWriter, r *http.Request) {
	var result GetActiveSemesterClassListResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetActiveSemesterClassListRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetActiveSemesterClassListResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer l.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Class", "Get Active Semester Class List", nil)

	studyProgramId := in.GetStudyProgramId()
	if studyProgramId == "" {
		profileData, errs := l.LecturerService.GetProfile(ctx)
		if errs != nil {
			utils.PrintError(*errs)
			result = GetActiveSemesterClassListResponse{
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
			errs = appConstants.ErrLecturerHasNoStudyProgram
			utils.PrintError(*errs)
			result = GetActiveSemesterClassListResponse{
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
		StudyProgramId: studyProgramId,
		SemesterId:     "",
		IsActive:       &isActive,
		ClassName:      "",
		SubjectName:    "",
	}
	data, errs := l.ClassService.GetList(ctx, paginationData, requestData)
	if errs != nil {
		utils.PrintError(*errs)
		result = GetActiveSemesterClassListResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = GetActiveSemesterClassListResponse{
		Meta: &Meta{
			Message: "Get Active Semester Class List",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: mapGetActiveSemesterClassList(data.Data),
	}

	utils.JSONResponse(w, http.StatusOK, &result)
}

func (l lecturerClassHandler) GetAssignedClass(w http.ResponseWriter, r *http.Request) {
	var result GetAssignedClassResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetAssignedClassRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetAssignedClassResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer l.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Class", "Get Assigned Class", nil)

	classIsActive := true
	data, errs := l.LecturerService.GetAssignedClass(ctx, in.GetSemesterId(), "", &classIsActive)
	if errs != nil {
		utils.PrintError(*errs)
		result = GetAssignedClassResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*GetAssignedClassResponseData{}
	for _, v := range data {
		resultData = append(resultData, &GetAssignedClassResponseData{
			Id:                   v.Id,
			Name:                 v.Name,
			SubjectCode:          v.SubjectCode,
			SubjectName:          v.SubjectName,
			TheoryCredit:         v.TheoryCredit,
			PracticumCredit:      v.PracticumCredit,
			FieldPracticumCredit: v.FieldPracticumCredit,
			IsGradingResponsible: v.IsGradingResponsible,
			StudyProgramId:       v.StudyProgramId,
			StudyProgramName:     v.StudyProgramName,
		})
	}

	result = GetAssignedClassResponse{
		Meta: &Meta{
			Message: "Get Assigned Class",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: resultData,
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (l lecturerClassHandler) GetAssignedSchedule(w http.ResponseWriter, r *http.Request) {
	var result GetAssignedScheduleResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetAssignedScheduleRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetAssignedScheduleResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer l.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Class", "Get Assigned Schedule", nil)

	data, errs := l.LectureService.GetDetailByClassId(ctx, in.GetClassId(), appConstants.AppTypeLecturer)
	if errs != nil {
		utils.PrintError(*errs)
		result = GetAssignedScheduleResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	singleDaySchedules := []*GetAssignedScheduleResponseDataSingleDaySchedule{}
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
			singleDaySchedules = append(singleDaySchedules, &GetAssignedScheduleResponseDataSingleDaySchedule{
				Date:      v.LecturePlanDate.Format(constants.DateRFC),
				StartTime: v.StartTime,
				EndTime:   v.EndTime,
				RoomName:  utils.NullStringScan(v.RoomName),
			})
		}
	}

	result = GetAssignedScheduleResponse{
		Meta: &Meta{
			Message: "Get Assigned Schedule",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &GetAssignedScheduleResponseData{
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

func (l lecturerClassHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
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
	defer l.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Class", "Get Detail", nil)

	data, errs := l.ClassService.GetDetail(ctx, in.GetClassId())
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

	gradeComponents := []*GetDetailResponseDataGradeComponent{}
	for _, v := range data.GradeComponents {
		gradeComponents = append(gradeComponents, &GetDetailResponseDataGradeComponent{
			Id:         v.Id,
			Name:       v.Name,
			Percentage: v.Percentage,
		})
	}

	students := []*GetDetailResponseDataStudent{}
	for _, v := range data.Students {
		grades := []*GetDetailResponseDataStudentGrade{}
		for _, w := range v.Grades {
			grades = append(grades, &GetDetailResponseDataStudentGrade{
				ClassGradeComponentId:   w.ClassGradeComponentId,
				ClassGradeComponentName: w.ClassGradeComponentName,
				InitialGrade:            w.InitialGrade,
				FinalGrade:              w.FinalGrade,
			})
		}

		students = append(students, &GetDetailResponseDataStudent{
			Id:         v.Id,
			NimNumber:  v.NimNumber,
			Name:       v.Name,
			GradePoint: v.GradePoint,
			GradeCode:  utils.NullStringScan(v.GradeCode),
			Grades:     grades,
		})
	}

	gradeTypes := []*GetDetailResponseDataGradeType{}
	for _, v := range data.GradeTypes {
		gradeTypes = append(gradeTypes, &GetDetailResponseDataGradeType{
			Id:                  v.Id,
			StudyLevelId:        v.StudyLevelId,
			StudyLevelShortName: v.StudyLevelShortName,
			Code:                v.Code,
			GradePoint:          v.GradePoint,
			MinimumGrade:        v.MinimumGrade,
			MaximumGrade:        v.MaximumGrade,
			GradeCategory:       v.GradeCategory,
			GradePointCategory:  v.GradePointCategory,
			Label:               utils.NullStringScan(v.Label),
			EnglishLabel:        utils.NullStringScan(v.EnglishLabel),
			StartDate:           v.StartDate.Format(constants.DateRFC),
			EndDate:             v.EndDate.Format(constants.DateRFC),
		})
	}

	result = GetDetailResponse{
		Meta: &Meta{
			Message: "Get Detail",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &GetDetailResponseData{
			Id:                    data.Id,
			Name:                  data.Name,
			StudyProgramId:        data.StudyProgramId,
			StudyProgramName:      data.StudyProgramName,
			DiktiStudyProgramType: data.DiktiStudyProgramType,
			StudyLevelShortName:   data.StudyLevelShortName,
			SemesterId:            data.SemesterId,
			SemesterSchoolYear:    data.SchoolYear,
			SemesterType:          data.SemesterType,
			GradingStartDate:      utils.SafetyDate(data.GradingStartDate),
			GradingEndDate:        utils.SafetyDate(data.GradingEndDate),
			CurriculumId:          data.CurriculumId,
			CurriculumName:        data.CurriculumName,
			CurriculumYear:        data.CurriculumYear,
			SubjectId:             data.SubjectId,
			SubjectCode:           data.SubjectCode,
			SubjectName:           data.SubjectName,
			IsGradingResponsible:  utils.NullBooleanScan(data.IsGradingResponsible),
			GradeComponents:       gradeComponents,
			Students:              students,
			GradeTypes:            gradeTypes,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (l lecturerClassHandler) BulkGradeStudentClass(w http.ResponseWriter, r *http.Request) {
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
	defer l.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Class", "Get Detail", nil)

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
	errs := l.StudentClassService.BulkGradeStudentClass(ctx, in.GetClassId(), data)
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
