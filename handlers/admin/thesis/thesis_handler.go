package thesis

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

type thesisHandler struct {
	*service.ServiceCtx
}

func (t thesisHandler) GetList(w http.ResponseWriter, r *http.Request) {
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
	defer t.AdminActivityLogService.Create(ctx, r, time.Now(), "Thesis", "Get List", nil)

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}
	data, errs := t.ThesisService.GetList(ctx, paginationData, in.GetStudyProgramId(), in.GetNimNumber(), in.GetStartSemesterId(), in.GetStatus(), in.GetSupervisorLecturerId())
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
			Id:                        v.Id,
			Topic:                     v.Topic,
			Title:                     v.Title,
			Status:                    v.Status,
			StudentId:                 v.StudentId,
			StudentName:               v.StudentName,
			StudentNimNumber:          v.StudentNimNumber,
			StudentStatus:             utils.NullStringScan(v.StudentStatus),
			StudyProgramId:            v.StudyProgramId,
			StudyProgramName:          v.StudyProgramName,
			DiktiStudyProgramCode:     v.DiktiStudyProgramCode,
			DiktiStudyProgramType:     v.DiktiStudyProgramType,
			StudyLevelShortName:       v.StudyLevelShortName,
			StudentHasThesisStudyPlan: v.StudentHasThesisStudyPlan,
			StartSemesterId:           v.StartSemesterId,
			StartSemesterType:         v.StartSemesterType,
			StartSemesterSchoolYear:   v.StartSemesterSchoolYear,
		})
	}

	result = GetListResponse{
		Meta: &Meta{
			Message: "Get List Thesis",
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

func (t thesisHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
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
	defer t.AdminActivityLogService.Create(ctx, r, time.Now(), "Thesis", "Get Detail", nil)

	data, errs := t.ThesisService.GetDetail(ctx, in.GetId())
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

	var files []*GetDetailResponseDataFile
	for _, v := range data.Files {
		files = append(files, &GetDetailResponseDataFile{
			Id:              v.Id,
			FileUrl:         v.FileUrl,
			FilePath:        v.FilePath,
			FilePathType:    v.FilePathType,
			FileDescription: utils.NullStringScan(v.FileDescription),
		})
	}

	var thesisSupervisors []*GetDetailResponseDataThesisSupervisor
	for _, v := range data.ThesisSupervisors {
		thesisSupervisors = append(thesisSupervisors, &GetDetailResponseDataThesisSupervisor{
			Id:                       v.Id,
			LecturerId:               v.LecturerId,
			LecturerName:             v.LecturerName,
			LecturerFrontTitle:       utils.NullStringScan(v.LecturerFrontTitle),
			LecturerBackDegree:       utils.NullStringScan(v.LecturerBackDegree),
			ThesisSupervisorRoleId:   v.ThesisSupervisorRoleId,
			ThesisSupervisorRoleName: v.ThesisSupervisorRoleName,
			ThesisSupervisorRoleSort: v.ThesisSupervisorRoleSort,
		})
	}

	result = GetDetailResponse{
		Meta: &Meta{
			Message: "Get Detail Thesis",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &GetDetailResponseData{
			Id:                        data.Id,
			StudyProgramId:            utils.NullStringScan(data.StudyProgramId),
			StudentId:                 data.StudentId,
			StudentName:               data.StudentName,
			StudentNimNumber:          data.StudentNimNumber,
			StartSemesterId:           data.StartSemesterId,
			StartSemesterType:         data.StartSemesterType,
			StartSemesterSchoolYear:   data.StartSemesterSchoolYear,
			FinishSemesterId:          utils.NullStringScan(data.FinishSemesterId),
			FinishSemesterType:        utils.NullStringScan(data.FinishSemesterType),
			FinishSemesterSchoolYear:  data.FinishSemesterSchoolYear,
			Topic:                     data.Topic,
			Title:                     data.Title,
			EnglishTitle:              utils.NullStringScan(data.EnglishTitle),
			StartDate:                 data.StartDate.Format(constants.DateRFC),
			FinishDate:                utils.SafetyDate(data.FinishDate),
			Remarks:                   utils.NullStringScan(data.Remarks),
			IsJointThesis:             data.IsJointThesis,
			Status:                    data.Status,
			ProposalSeminarDate:       utils.SafetyDate(data.ProposalSeminarDate),
			ProposalCertificateNumber: utils.NullStringScan(data.ProposalCertificateNumber),
			ProposalCertificateDate:   utils.SafetyDate(data.ProposalCertificateDate),
			ThesisDefenseCount:        data.ThesisDefenseCount,
			GradePoint:                data.GradePoint,
			GradeCode:                 utils.NullStringScan(data.GradeCode),
			Files:                     files,
			ThesisSupervisors:         thesisSupervisors,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (t thesisHandler) Create(w http.ResponseWriter, r *http.Request) {
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
	defer t.AdminActivityLogService.Create(ctx, r, time.Now(), "Thesis", "Create", &in)

	startDate, errs := utils.StringToTime(in.GetStartDate())
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
	proposalSeminarDate, errs := utils.StringToTime(in.GetProposalSeminarDate())
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
	proposalCertificateDate, errs := utils.StringToTime(in.GetProposalCertificateDate())
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

	var thesisSupervisors []objects.CreateThesisSupervisor
	for _, v := range in.GetThesisSupervisors() {
		thesisSupervisors = append(thesisSupervisors, objects.CreateThesisSupervisor{
			LecturerId:             v.GetLecturerId(),
			ThesisSupervisorRoleId: v.GetThesisSupervisorRoleId(),
		})
	}

	data := objects.CreateThesis{
		StudentId:                 in.GetStudentId(),
		Topic:                     in.GetTopic(),
		Status:                    in.GetStatus(),
		Title:                     in.GetTitle(),
		EnglishTitle:              in.GetEnglishTitle(),
		StartSemesterId:           in.GetStartSemesterId(),
		StartDate:                 startDate,
		Remarks:                   in.GetRemarks(),
		IsJointThesis:             in.GetIsJointThesis(),
		FilePath:                  in.GetFilePath(),
		FilePathType:              in.GetFilePathType(),
		FileDescription:           in.GetFileDescription(),
		ProposalSeminarDate:       proposalSeminarDate,
		ProposalCertificateNumber: in.GetProposalCertificateNumber(),
		ProposalCertificateDate:   proposalCertificateDate,
		ThesisSupervisors:         thesisSupervisors,
	}
	errs = t.ThesisService.Create(ctx, data)
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
			Message: "Create Thesis",
			Status:  http.StatusCreated,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusCreated, &result)
}

func (t thesisHandler) Update(w http.ResponseWriter, r *http.Request) {
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
	defer t.AdminActivityLogService.Create(ctx, r, time.Now(), "Thesis", "Update", &in)

	startDate, errs := utils.StringToTime(in.GetStartDate())
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
	proposalSeminarDate, errs := utils.StringToTime(in.GetProposalSeminarDate())
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
	proposalCertificateDate, errs := utils.StringToTime(in.GetProposalCertificateDate())
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

	var files []objects.UpdateThesisFile
	for _, v := range in.GetFiles() {
		files = append(files, objects.UpdateThesisFile{
			FilePath:        v.GetFilePath(),
			FilePathType:    v.GetFilePathType(),
			FileDescription: v.GetFileDescription(),
		})
	}

	var thesisSupervisors []objects.UpdateThesisSupervisor
	for _, v := range in.GetThesisSupervisors() {
		thesisSupervisors = append(thesisSupervisors, objects.UpdateThesisSupervisor{
			LecturerId:             v.GetLecturerId(),
			ThesisSupervisorRoleId: v.GetThesisSupervisorRoleId(),
		})
	}
	data := objects.UpdateThesis{
		Id:                        in.GetId(),
		StudentId:                 in.GetStudentId(),
		Topic:                     in.GetTopic(),
		Status:                    in.GetStatus(),
		Title:                     in.GetTitle(),
		EnglishTitle:              in.GetEnglishTitle(),
		StartSemesterId:           in.GetStartSemesterId(),
		StartDate:                 startDate,
		Remarks:                   in.GetRemarks(),
		IsJointThesis:             in.GetIsJointThesis(),
		Files:                     files,
		ProposalSeminarDate:       proposalSeminarDate,
		ProposalCertificateNumber: in.GetProposalCertificateNumber(),
		ProposalCertificateDate:   proposalCertificateDate,
		ThesisSupervisors:         thesisSupervisors,
	}
	errs = t.ThesisService.Update(ctx, data)
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
			Message: "Update Thesis",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (t thesisHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
	defer t.AdminActivityLogService.Create(ctx, r, time.Now(), "Thesis", "Delete", nil)

	errs := t.ThesisService.Delete(ctx, in.GetId())
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
			Message: "Delete Thesis",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (t thesisHandler) GetListThesisDefenseRequest(w http.ResponseWriter, r *http.Request) {
	var result GetListThesisDefenseRequestResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetListThesisDefenseRequestRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetListThesisDefenseRequestResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer t.AdminActivityLogService.Create(ctx, r, time.Now(), "Thesis", "Get List", nil)

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}
	data, errs := t.ThesisService.GetListThesisDefenseRequest(ctx, paginationData, in.GetStudyProgramId(), in.GetNimNumber(), in.GetStartSemesterId())
	if errs != nil {
		utils.PrintError(*errs)
		result = GetListThesisDefenseRequestResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*GetListThesisDefenseRequestResponseData{}
	for _, v := range data.Data {
		examiners := []*GetListThesisDefenseRequestResponseDataExaminer{}
		for _, w := range v.Examiners {
			examiners = append(examiners, &GetListThesisDefenseRequestResponseDataExaminer{
				Id:                     w.Id,
				LecturerId:             w.LecturerId,
				LecturerName:           w.LecturerName,
				LecturerFrontTitle:     utils.NullStringScan(w.LecturerFrontTitle),
				LecturerBackDegree:     utils.NullStringScan(w.LecturerBackDegree),
				ThesisExaminerRoleId:   w.ThesisExaminerRoleId,
				ThesisExaminerRoleName: w.ThesisExaminerRoleName,
			})
		}

		resultData = append(resultData, &GetListThesisDefenseRequestResponseData{
			Id:                           v.Id,
			StudentId:                    v.StudentId,
			StudentName:                  v.StudentName,
			StudentNimNumber:             v.StudentNimNumber,
			StudentStatus:                v.StudentStatus,
			StudyProgramId:               v.StudyProgramId,
			StudyProgramName:             v.StudyProgramName,
			DiktiStudyProgramCode:        v.DiktiStudyProgramCode,
			DiktiStudyProgramType:        v.DiktiStudyProgramType,
			StudyLevelId:                 v.StudyLevelId,
			StudyLevelShortName:          v.StudyLevelShortName,
			ThesisId:                     v.ThesisId,
			ThesisTitle:                  v.ThesisTitle,
			ThesisStatus:                 v.ThesisStatus,
			ThesisDefenseCount:           v.ThesisDefenseCount,
			ThesisDefenseId:              utils.NullStringScan(v.ThesisDefenseId),
			ThesisDefensePlanDate:        utils.SafetyDate(v.ThesisDefensePlanDate),
			ThesisDefensePlanStartTime:   utils.NullUint32Scan(v.ThesisDefensePlanStartTime),
			ThesisDefensePlanEndTime:     utils.NullUint32Scan(v.ThesisDefensePlanEndTime),
			ThesisDefenseActualDate:      utils.SafetyDate(v.ThesisDefenseActualDate),
			ThesisDefenseActualStartTime: utils.NullUint32Scan(v.ThesisDefenseActualStartTime),
			ThesisDefenseActualEndTime:   utils.NullUint32Scan(v.ThesisDefenseActualEndTime),
			ThesisDefenseRoomId:          utils.NullStringScan(v.ThesisDefenseRoomId),
			ThesisDefenseRoomName:        utils.NullStringScan(v.ThesisDefenseRoomName),
			ThesisDefenseIsPassed:        utils.NullBooleanScan(v.ThesisDefenseIsPassed),
			ThesisDefenseRevision:        utils.NullStringScan(v.ThesisDefenseRevision),
			ThesisGradeCode:              utils.NullStringScan(v.ThesisGradeCode),
			CreatedAt:                    v.CreatedAt.Format(constants.DateRFC),
			Examiners:                    examiners,
		})
	}

	result = GetListThesisDefenseRequestResponse{
		Meta: &Meta{
			Message: "Get List Thesis Defense Request",
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

func (t thesisHandler) RegisterThesisDefense(w http.ResponseWriter, r *http.Request) {
	var result RegisterThesisDefenseResponse

	ctx := r.Context()
	var in RegisterThesisDefenseRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = RegisterThesisDefenseResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer t.AdminActivityLogService.Create(ctx, r, time.Now(), "Thesis", "Register Thesis Defense", &in)

	errs := t.ThesisService.RegisterThesisDefense(ctx, in.GetStudentId())
	if errs != nil {
		utils.PrintError(*errs)
		result = RegisterThesisDefenseResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = RegisterThesisDefenseResponse{
		Meta: &Meta{
			Message: "Register Thesis Defense Thesis",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (t thesisHandler) CreateThesisDefense(w http.ResponseWriter, r *http.Request) {
	var result CreateThesisDefenseResponse

	ctx := r.Context()
	var in CreateThesisDefenseRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = CreateThesisDefenseResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer t.AdminActivityLogService.Create(ctx, r, time.Now(), "Thesis", "Create Thesis Defense", &in)

	var examiners []objects.CreateThesisDefenseExaminer
	for _, v := range in.GetExaminers() {
		examiners = append(examiners, objects.CreateThesisDefenseExaminer{
			LecturerId:           v.GetLecturerId(),
			ThesisExaminerRoleId: v.GetThesisExaminerRoleId(),
		})
	}

	planDate, errs := utils.StringToTime(in.GetPlanDate())
	if errs != nil {
		utils.PrintError(*errs)
		result = CreateThesisDefenseResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	data := objects.CreateThesisDefense{
		ThesisId:      in.GetThesisId(),
		PlanDate:      planDate,
		PlanStartTime: in.GetPlanStartTime(),
		PlanEndTime:   in.GetPlanEndTime(),
		RoomId:        in.GetRoomId(),
		Examiners:     examiners,
	}
	errs = t.ThesisService.CreateThesisDefense(ctx, data)
	if errs != nil {
		utils.PrintError(*errs)
		result = CreateThesisDefenseResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = CreateThesisDefenseResponse{
		Meta: &Meta{
			Message: "Create Thesis Defense Thesis",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (t thesisHandler) UpdateThesisDefense(w http.ResponseWriter, r *http.Request) {
	var result UpdateThesisDefenseResponse

	ctx := r.Context()
	var in UpdateThesisDefenseRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = UpdateThesisDefenseResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer t.AdminActivityLogService.Create(ctx, r, time.Now(), "Thesis", "Update Thesis Defense", &in)

	var examiners []objects.UpdateThesisDefenseExaminer
	for _, v := range in.GetExaminers() {
		examiners = append(examiners, objects.UpdateThesisDefenseExaminer{
			LecturerId:           v.GetLecturerId(),
			ThesisExaminerRoleId: v.GetThesisExaminerRoleId(),
		})
	}

	planDate, errs := utils.StringToTime(in.GetPlanDate())
	if errs != nil {
		utils.PrintError(*errs)
		result = UpdateThesisDefenseResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	actualDate, errs := utils.StringToTime(in.GetActualDate())
	if errs != nil {
		utils.PrintError(*errs)
		result = UpdateThesisDefenseResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	data := objects.UpdateThesisDefense{
		Id:              in.GetId(),
		PlanDate:        planDate,
		PlanStartTime:   in.GetPlanStartTime(),
		PlanEndTime:     in.GetPlanEndTime(),
		RoomId:          in.GetRoomId(),
		ActualDate:      actualDate,
		ActualStartTime: in.GetActualStartTime(),
		ActualEndTime:   in.GetActualEndTime(),
		IsPassed:        in.GetIsPassed(),
		Revision:        in.GetRevision(),
		GradeCode:       in.GetGradeCode(),
		Examiners:       examiners,
	}
	errs = t.ThesisService.UpdateThesisDefense(ctx, data)
	if errs != nil {
		utils.PrintError(*errs)
		result = UpdateThesisDefenseResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = UpdateThesisDefenseResponse{
		Meta: &Meta{
			Message: "Update Thesis Defense Thesis",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (t thesisHandler) GetThesisSupervisorLog(w http.ResponseWriter, r *http.Request) {
	var result GetThesisSupervisorLogResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetThesisSupervisorLogRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetThesisSupervisorLogResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer t.AdminActivityLogService.Create(ctx, r, time.Now(), "Thesis", "Get Supervisor Log", nil)

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}
	data, errs := t.ThesisService.GetThesisSupervisorLog(ctx, paginationData, in.GetIdNationalLecturer(), in.GetSemesterId())
	if errs != nil {
		utils.PrintError(*errs)
		result = GetThesisSupervisorLogResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*GetThesisSupervisorLogResponseData{}
	for _, v := range data.Data {
		thesisSupervisorRoles := []*GetThesisSupervisorLogResponseDataThesisSupervisorRole{}
		for _, w := range v.ThesisSupervisorRoles {
			thesisSupervisorRoles = append(thesisSupervisorRoles, &GetThesisSupervisorLogResponseDataThesisSupervisorRole{
				Id:    w.Id,
				Name:  w.Name,
				Total: w.Total,
			})
		}

		resultData = append(resultData, &GetThesisSupervisorLogResponseData{
			Id:                     v.Id,
			IdNationalLecturer:     v.IdNationalLecturer,
			Name:                   v.Name,
			TotalSupervisedThesis:  v.TotalSupervisedThesis,
			ActiveSupervisedThesis: v.ActiveSupervisedThesis,
			ThesisSupervisorRoles:  thesisSupervisorRoles,
		})
	}

	result = GetThesisSupervisorLogResponse{
		Meta: &Meta{
			Message: "Get Supervisor Log Thesis",
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
