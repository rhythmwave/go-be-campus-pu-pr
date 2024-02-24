package class

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

type classHandler struct {
	*service.ServiceCtx
}

func (a classHandler) GetList(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Class", "Get List", nil)

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
		Sort:   in.GetSort(),
		SortBy: in.GetSortBy(),
	}
	isActive, errs := utils.StringToBoolPointer(in.GetIsActive())
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
	forOddSemester, errs := utils.StringToBoolPointer(in.GetForOddSemester())
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
	requestData := objects.GetClassListRequest{
		StudyProgramId: in.GetStudyProgramId(),
		SemesterId:     in.GetSemesterId(),
		IsActive:       isActive,
		ClassName:      in.GetClassName(),
		SubjectName:    in.GetSubjectName(),
		SubjectId:      in.GetSubjectId(),
		IsMbkm:         in.GetIsMbkm(),
		ForOddSemester: forOddSemester,
	}
	data, errs := a.ClassService.GetList(ctx, paginationData, requestData)
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
		lecturers := []*GetListResponseDataLecturer{}
		for _, w := range v.Lecturers {
			lecturers = append(lecturers, &GetListResponseDataLecturer{
				Id:         w.Id,
				Name:       w.Name,
				FrontTitle: utils.NullStringScan(w.FrontTitle),
				BackDegree: utils.NullStringScan(w.BackDegree),
			})
		}

		resultData = append(resultData, &GetListResponseData{
			Id:                          v.Id,
			Name:                        v.Name,
			SubjectId:                   v.SubjectId,
			SubjectCode:                 v.SubjectCode,
			SubjectName:                 v.SubjectName,
			SubjectIsMandatory:          v.SubjectIsMandatory,
			SubjectSemesterPackage:      v.SubjectSemesterPackage,
			TotalParticipant:            v.TotalParticipant,
			Lecturers:                   lecturers,
			SubjectTheoryCredit:         v.SubjectTheoryCredit,
			SubjectPracticumCredit:      v.SubjectPracticumCredit,
			SubjectFieldPracticumCredit: v.SubjectFieldPracticumCredit,
			IsActive:                    v.IsActive,
			MaximumParticipant:          utils.NullUint32Scan(v.MaximumParticipant),
			UnapprovedStudyPlan:         v.UnapprovedStudyPlan,
			TotalMaterial:               v.TotalMaterial,
			TotalWork:                   v.TotalWork,
			TotalDiscussion:             v.TotalDiscussion,
			TotalExam:                   v.TotalExam,
			TotalEvent:                  v.TotalEvent,
			TotalLecturePlan:            v.TotalLecturePlan,
			TotalLectureDone:            v.TotalLectureDone,
			TotalGradedParticipant:      v.TotalGradedParticipant,
			StudyLevelId:                v.StudyLevelId,
			ApplicationDeadline:         utils.SafetyDate(v.ApplicationDeadline),
			CurriculumId:                v.CurriculumId,
			CurriculumName:              v.CurriculumName,
			StudyProgramId:              v.StudyProgramId,
			StudyProgramName:            v.StudyProgramName,
			SemesterId:                  v.SemesterId,
			SemesterStartYear:           v.SemesterStartYear,
			SchoolYear:                  v.SchoolYear,
			SemesterType:                v.SemesterType,
		})
	}

	result = GetListResponse{
		Meta: &Meta{
			Message: "Get List Class",
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

func (a classHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Class", "Get Detail", nil)

	data, errs := a.ClassService.GetDetail(ctx, in.GetId())
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

	lecturers := []*GetDetailResponseDataLecturer{}
	gradeComponents := []*GetDetailResponseDataGradeComponent{}

	for _, v := range data.Lecturers {
		lecturers = append(lecturers, &GetDetailResponseDataLecturer{
			Id:                   v.Id,
			Name:                 v.Name,
			FrontTitle:           utils.NullStringScan(v.FrontTitle),
			BackDegree:           utils.NullStringScan(v.BackDegree),
			IsGradingResponsible: v.IsGradingResponsible,
		})
	}
	for _, v := range data.GradeComponents {
		gradeComponents = append(gradeComponents, &GetDetailResponseDataGradeComponent{
			Id:         v.Id,
			Name:       v.Name,
			Percentage: v.Percentage,
		})
	}

	result = GetDetailResponse{
		Meta: &Meta{
			Message: "Get Detail Class",
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
			CurriculumId:          data.CurriculumId,
			CurriculumName:        data.CurriculumName,
			CurriculumYear:        data.CurriculumYear,
			SubjectId:             data.SubjectId,
			SubjectCode:           data.SubjectCode,
			SubjectName:           data.SubjectName,
			SemesterId:            data.SemesterId,
			SemesterStartYear:     data.SemesterStartYear,
			SchoolYear:            data.SchoolYear,
			SemesterType:          data.SemesterType,
			Scope:                 utils.NullStringScan(data.Scope),
			IsOnline:              utils.NullBooleanScan(data.IsOnline),
			IsOffline:             utils.NullBooleanScan(data.IsOffline),
			MinimumParticipant:    utils.NullUint32Scan(data.MinimumParticipant),
			MaximumParticipant:    utils.NullUint32Scan(data.MaximumParticipant),
			TotalParticipant:      data.TotalParticipant,
			Remarks:               data.Remarks,
			IsActive:              data.IsActive,
			StudyLevelId:          data.StudyLevelId,
			ApplicationDeadline:   utils.SafetyDate(data.ApplicationDeadline),
			Lecturers:             lecturers,
			GradeComponents:       gradeComponents,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a classHandler) Create(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Class", "Create", &in)

	lecturers := []objects.CreateClassLecturer{}
	for _, v := range in.GetLecturers() {
		lecturers = append(lecturers, objects.CreateClassLecturer{
			Id:                   v.GetId(),
			IsGradingResponsible: v.GetIsGradingResponsible(),
		})
	}

	applicationDeadline, errs := utils.StringToTime(in.GetApplicationDeadline())
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
	data := objects.CreateClass{
		SubjectId:           in.GetSubjectId(),
		SemesterId:          in.GetSemesterId(),
		Name:                in.GetName(),
		Scope:               in.GetScope(),
		IsOnline:            in.GetIsOnline(),
		IsOffline:           in.GetIsOffline(),
		MinimumParticipant:  in.GetMinimumParticipant(),
		MaximumParticipant:  in.GetMaximumParticipant(),
		Remarks:             in.GetRemarks(),
		ApplicationDeadline: applicationDeadline,
		Lecturers:           lecturers,
	}
	errs = a.ClassService.Create(ctx, data)
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
			Message: "Create Class",
			Status:  http.StatusCreated,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusCreated, &result)
}

func (a classHandler) Update(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Class", "Update", &in)

	lecturers := []objects.UpdateClassLecturer{}
	for _, v := range in.GetLecturers() {
		lecturers = append(lecturers, objects.UpdateClassLecturer{
			Id:                   v.GetId(),
			IsGradingResponsible: v.GetIsGradingResponsible(),
		})
	}

	applicationDeadline, errs := utils.StringToTime(in.GetApplicationDeadline())
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
	data := objects.UpdateClass{
		Id:                  in.GetId(),
		SubjectId:           in.GetSubjectId(),
		Name:                in.GetName(),
		Scope:               in.GetScope(),
		IsOnline:            in.GetIsOnline(),
		IsOffline:           in.GetIsOffline(),
		MinimumParticipant:  in.GetMinimumParticipant(),
		MaximumParticipant:  in.GetMaximumParticipant(),
		Remarks:             in.GetRemarks(),
		ApplicationDeadline: applicationDeadline,
		Lecturers:           lecturers,
	}
	errs = a.ClassService.Update(ctx, data)
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
			Message: "Update Class",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a classHandler) UpdateActivation(w http.ResponseWriter, r *http.Request) {
	var result UpdateActivationResponse

	ctx := r.Context()
	var in UpdateActivationRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = UpdateActivationResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Class", "Update Activation", nil)

	errs := a.ClassService.UpdateActivation(ctx, in.GetId(), in.GetIsActive())
	if errs != nil {
		utils.PrintError(*errs)
		result = UpdateActivationResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = UpdateActivationResponse{
		Meta: &Meta{
			Message: "Update Activation Class",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a classHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Class", "Delete", nil)

	errs := a.ClassService.Delete(ctx, in.GetId())
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
			Message: "Delete Class",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a classHandler) Duplicate(w http.ResponseWriter, r *http.Request) {
	var result DuplicateResponse

	ctx := r.Context()
	var in DuplicateRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = DuplicateResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Class", "Duplicate", nil)

	errs := a.ClassService.Duplicate(ctx, in.GetFromSemesterId(), in.GetToSemesterId())
	if errs != nil {
		utils.PrintError(*errs)
		result = DuplicateResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = DuplicateResponse{
		Meta: &Meta{
			Message: "Duplicate Class",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a classHandler) BulkUpdateMaximumParticipant(w http.ResponseWriter, r *http.Request) {
	var result BulkUpdateMaximumParticipantResponse

	ctx := r.Context()
	var in BulkUpdateMaximumParticipantRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = BulkUpdateMaximumParticipantResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Class", "Bulk Update Maximum Participant", nil)

	data := []objects.UpdateClassMaximumParticipant{}
	for _, v := range in.GetData() {
		data = append(data, objects.UpdateClassMaximumParticipant{
			Id:                 v.GetClassId(),
			MaximumParticipant: v.GetMaximumParticipant(),
		})
	}

	errs := a.ClassService.BulkUpdateMaximumParticipant(ctx, data)
	if errs != nil {
		utils.PrintError(*errs)
		result = BulkUpdateMaximumParticipantResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = BulkUpdateMaximumParticipantResponse{
		Meta: &Meta{
			Message: "Bulk Update Maximum Participant Class",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a classHandler) GetClassParticipantList(w http.ResponseWriter, r *http.Request) {
	var result GetClassParticipantListResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetClassParticipantListRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetClassParticipantListResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Class", "Get Participant List", nil)

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}

	isGraded, errs := utils.StringToBoolPointer(in.GetIsGraded())
	if errs != nil {
		utils.PrintError(*errs)
		result = GetClassParticipantListResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	data, errs := a.ClassService.GetClassParticipantList(ctx, paginationData, in.GetClassId(), in.GetLectureId(), isGraded, in.GetStudentId())
	if errs != nil {
		utils.PrintError(*errs)
		result = GetClassParticipantListResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*GetClassParticipantListResponseData{}
	for _, v := range data.Data {
		grades := []*GetClassParticipantListResponseDataGrade{}
		for _, w := range v.Grades {
			grades = append(grades, &GetClassParticipantListResponseDataGrade{
				ClassGradeComponentId:   w.ClassGradeComponentId,
				ClassGradeComponentName: w.ClassGradeComponentName,
				InitialGrade:            w.InitialGrade,
				FinalGrade:              w.FinalGrade,
			})
		}

		resultData = append(resultData, &GetClassParticipantListResponseData{
			StudentId:             v.StudentId,
			StudentNimNumber:      v.StudentNimNumber,
			StudentName:           v.StudentName,
			StudyProgramId:        v.StudyProgramId,
			StudyProgramName:      v.StudyProgramName,
			DiktiStudyProgramCode: v.DiktiStudyProgramCode,
			DiktiStudyProgramType: v.DiktiStudyProgramType,
			StudyLevelShortName:   v.StudyLevelShortName,
			TotalAttendance:       v.TotalAttendance,
			AttendancePercentage:  v.AttendancePercentage,
			TotalSick:             v.TotalSick,
			TotalLeave:            v.TotalLeave,
			TotalAwol:             v.TotalAwol,
			IsAttend:              utils.NullBooleanScan(v.IsAttend),
			IsSick:                utils.NullBooleanScan(v.IsSick),
			IsLeave:               utils.NullBooleanScan(v.IsLeave),
			IsAwol:                utils.NullBooleanScan(v.IsAwol),
			GradePoint:            v.GradePoint,
			GradeCode:             utils.NullStringScan(v.GradeCode),
			GradedByAdminId:       utils.NullStringScan(v.GradedByAdminId),
			GradedByAdminName:     utils.NullStringScan(v.GradedByAdminName),
			GradedByLecturerId:    utils.NullStringScan(v.GradedByLecturerId),
			GradedByLecturerName:  utils.NullStringScan(v.GradedByLecturerName),
			GradedAt:              utils.SafetyDate(v.GradedAt),
			SubjectRepetition:     v.SubjectRepetition,
			Grades:                grades,
		})
	}

	result = GetClassParticipantListResponse{
		Meta: &Meta{
			Message: "Get Participant List Class",
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
