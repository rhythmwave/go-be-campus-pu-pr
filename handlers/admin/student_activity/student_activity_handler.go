package student_activity

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

type studentActivityHandler struct {
	*service.ServiceCtx
}

func (s studentActivityHandler) GetList(w http.ResponseWriter, r *http.Request) {
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
	defer s.AdminActivityLogService.Create(ctx, r, time.Now(), "Student Activity", "Get List", nil)

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}
	data, errs := s.StudentActivityService.GetList(ctx, paginationData, in.GetActivityType(), in.GetStudyProgramId(), in.GetSemesterId(), in.GetIsMbkm())
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
			StudyProgramId:     v.StudyProgramId,
			StudyProgramName:   v.StudyProgramName,
			SemesterId:         v.SemesterId,
			SemesterSchoolYear: v.SemesterSchoolYear,
			SemesterType:       v.SemesterType,
			ActivityType:       v.ActivityType,
			Title:              v.Title,
		})
	}

	result = GetListResponse{
		Meta: &Meta{
			Message: "Get List Student Activity",
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

func (s studentActivityHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
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
	defer s.AdminActivityLogService.Create(ctx, r, time.Now(), "Student Activity", "Get Detail", nil)

	data, errs := s.StudentActivityService.GetDetail(ctx, in.GetId())
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

	participants := []*GetDetailResponseDataParticipant{}
	mentors := []*GetDetailResponseDataMentor{}
	examiners := []*GetDetailResponseDataExaminer{}
	for _, v := range data.Participants {
		participants = append(participants, &GetDetailResponseDataParticipant{
			StudentId:        v.StudentId,
			NimNumber:        v.NimNumber,
			Name:             v.Name,
			StudyProgramId:   utils.NullStringScan(v.StudyProgramId),
			StudyProgramName: utils.NullStringScan(v.StudyProgramName),
			Role:             v.Role,
		})

	}
	for _, v := range data.Mentors {
		mentors = append(mentors, &GetDetailResponseDataMentor{
			LecturerId:         v.LecturerId,
			IdNationalLecturer: v.IdNationalLecturer,
			Name:               v.Name,
			FrontTitle:         utils.NullStringScan(v.FrontTitle),
			BackDegree:         utils.NullStringScan(v.BackDegree),
			ActivityCategory:   v.ActivityCategory,
			Sort:               v.Sort,
		})

	}
	for _, v := range data.Examiners {
		examiners = append(examiners, &GetDetailResponseDataExaminer{
			LecturerId:         v.LecturerId,
			IdNationalLecturer: v.IdNationalLecturer,
			Name:               v.Name,
			FrontTitle:         utils.NullStringScan(v.FrontTitle),
			BackDegree:         utils.NullStringScan(v.BackDegree),
			ActivityCategory:   v.ActivityCategory,
			Sort:               v.Sort,
		})

	}

	result = GetDetailResponse{
		Meta: &Meta{
			Message: "Get Detail Student Activity",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &GetDetailResponseData{
			Id:                 data.Id,
			StudyProgramId:     data.StudyProgramId,
			StudyProgramName:   data.StudyProgramName,
			SemesterId:         data.SemesterId,
			SemesterSchoolYear: data.SemesterSchoolYear,
			SemesterType:       data.SemesterType,
			ActivityType:       data.ActivityType,
			Title:              data.Title,
			Location:           utils.NullStringScan(data.Location),
			DecisionNumber:     utils.NullStringScan(data.DecisionNumber),
			DecisionDate:       utils.SafetyDate(data.DecisionDate),
			IsGroupActivity:    data.IsGroupActivity,
			Remarks:            utils.NullStringScan(data.Remarks),
			Participants:       participants,
			Mentors:            mentors,
			Examiners:          examiners,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (s studentActivityHandler) Create(w http.ResponseWriter, r *http.Request) {
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
	defer s.AdminActivityLogService.Create(ctx, r, time.Now(), "Student Activity", "Create", &in)

	participants := []objects.CreateStudentActivityParticipant{}
	mentors := []objects.CreateStudentActivityLecturer{}
	examiners := []objects.CreateStudentActivityLecturer{}
	for _, v := range in.GetParticipants() {
		participants = append(participants, objects.CreateStudentActivityParticipant{
			StudentId: v.GetStudentId(),
			Role:      v.GetRole(),
		})
	}
	for _, v := range in.GetMentors() {
		mentors = append(mentors, objects.CreateStudentActivityLecturer{
			LecturerId:       v.GetLecturerId(),
			ActivityCategory: v.GetActivityCategory(),
			Sort:             v.GetSort(),
		})
	}
	for _, v := range in.GetExaminers() {
		examiners = append(examiners, objects.CreateStudentActivityLecturer{
			LecturerId:       v.GetLecturerId(),
			ActivityCategory: v.GetActivityCategory(),
			Sort:             v.GetSort(),
		})
	}

	data := objects.CreateStudentActivity{
		StudyProgramId:  in.GetStudyProgramId(),
		SemesterId:      in.GetSemesterId(),
		ActivityType:    in.GetActivityType(),
		Title:           in.GetTitle(),
		Location:        in.GetLocation(),
		DecisionNumber:  in.GetDecisionNumber(),
		DecisionDate:    in.GetDecisionDate(),
		IsGroupActivity: in.GetIsGroupActivity(),
		Remarks:         in.GetRemarks(),
		IsMbkm:          in.GetIsMbkm(),
		Participants:    participants,
		Mentors:         mentors,
		Examiners:       examiners,
	}
	errs := s.StudentActivityService.Create(ctx, data)
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
			Message: "Create Student Activity",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (s studentActivityHandler) Update(w http.ResponseWriter, r *http.Request) {
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
	defer s.AdminActivityLogService.Create(ctx, r, time.Now(), "Student Activity", "Update", &in)

	participants := []objects.UpdateStudentActivityParticipant{}
	mentors := []objects.UpdateStudentActivityLecturer{}
	examiners := []objects.UpdateStudentActivityLecturer{}
	for _, v := range in.GetParticipants() {
		participants = append(participants, objects.UpdateStudentActivityParticipant{
			StudentId: v.GetStudentId(),
			Role:      v.GetRole(),
		})
	}
	for _, v := range in.GetMentors() {
		mentors = append(mentors, objects.UpdateStudentActivityLecturer{
			LecturerId:       v.GetLecturerId(),
			ActivityCategory: v.GetActivityCategory(),
			Sort:             v.GetSort(),
		})
	}
	for _, v := range in.GetExaminers() {
		examiners = append(examiners, objects.UpdateStudentActivityLecturer{
			LecturerId:       v.GetLecturerId(),
			ActivityCategory: v.GetActivityCategory(),
			Sort:             v.GetSort(),
		})
	}

	data := objects.UpdateStudentActivity{
		Id:              in.GetId(),
		StudyProgramId:  in.GetStudyProgramId(),
		SemesterId:      in.GetSemesterId(),
		ActivityType:    in.GetActivityType(),
		Title:           in.GetTitle(),
		Location:        in.GetLocation(),
		DecisionNumber:  in.GetDecisionNumber(),
		DecisionDate:    in.GetDecisionDate(),
		IsGroupActivity: in.GetIsGroupActivity(),
		Remarks:         in.GetRemarks(),
		IsMbkm:          in.GetIsMbkm(),
		Participants:    participants,
		Mentors:         mentors,
		Examiners:       examiners,
	}
	errs := s.StudentActivityService.Update(ctx, data)
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
			Message: "Update Student Activity",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (s studentActivityHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
	defer s.AdminActivityLogService.Create(ctx, r, time.Now(), "Student Activity", "Delete", nil)

	errs := s.StudentActivityService.Delete(ctx, in.GetId())
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
			Message: "Delete Student Activity",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
