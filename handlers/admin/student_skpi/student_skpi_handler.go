package student_skpi

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

type studentSkpiHandler struct {
	*service.ServiceCtx
}

func (s studentSkpiHandler) GetList(w http.ResponseWriter, r *http.Request) {
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
	defer s.AdminActivityLogService.Create(ctx, r, time.Now(), "Student SKPI", "Get List", nil)

	paginationData := common.PaginationRequest{
		Limit: in.GetLimit(),
		Page:  in.GetPage(),
	}

	isApproved, errs := utils.StringToBoolPointer(in.GetIsApproved())
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
	paramsData := objects.GetStudentSkpiRequest{
		StudyProgramId: in.GetStudyProgramId(),
		Name:           in.GetName(),
		NimNumber:      in.GetNimNumber(),
		NimNumberFrom:  in.GetNimNumberFrom(),
		NimNumberTo:    in.GetNimNumberTo(),
		IsApproved:     isApproved,
	}
	data, errs := s.StudentSkpiService.GetList(ctx, paginationData, paramsData)
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
			Id:                           v.Id,
			StudentId:                    v.StudentId,
			StudentNimNumber:             v.StudentNimNumber,
			StudentName:                  v.StudentName,
			StudentStudyProgramId:        v.StudentStudyProgramId,
			StudentStudyProgramName:      v.StudentStudyProgramName,
			StudentDiktiStudyProgramCode: v.StudentDiktiStudyProgramCode,
			IsApproved:                   v.IsApproved,
		})
	}

	result = GetListResponse{
		Meta: &Meta{
			Message: "Get List Student SKPI",
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

func (s studentSkpiHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
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
	defer s.AdminActivityLogService.Create(ctx, r, time.Now(), "Student SKPI", "Get Detail", nil)

	data, errs := s.StudentSkpiService.GetDetail(ctx, in.GetId())
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

	achievements := []*GetDetailResponseDataAchievement{}
	organizations := []*GetDetailResponseDataOrganization{}
	certificates := []*GetDetailResponseDataCertificate{}
	characterBuildings := []*GetDetailResponseDataCharacterBuilding{}
	internships := []*GetDetailResponseDataInternship{}
	languages := []*GetDetailResponseDataLanguage{}

	for _, v := range data.Achievements {
		achievements = append(achievements, &GetDetailResponseDataAchievement{
			Id:   v.Id,
			Name: v.Name,
			Year: v.Year,
		})

	}
	for _, v := range data.Organizations {
		organizations = append(organizations, &GetDetailResponseDataOrganization{
			Id:            v.Id,
			Name:          v.Name,
			Position:      v.Position,
			ServiceLength: v.ServiceLength,
		})

	}
	for _, v := range data.Certificates {
		certificates = append(certificates, &GetDetailResponseDataCertificate{
			Id:   v.Id,
			Name: v.Name,
		})

	}
	for _, v := range data.CharacterBuildings {
		characterBuildings = append(characterBuildings, &GetDetailResponseDataCharacterBuilding{
			Id:   v.Id,
			Name: v.Name,
		})

	}
	for _, v := range data.Internships {
		internships = append(internships, &GetDetailResponseDataInternship{
			Id:   v.Id,
			Name: v.Name,
		})

	}
	for _, v := range data.Languages {
		languages = append(languages, &GetDetailResponseDataLanguage{
			Id:    v.Id,
			Name:  v.Name,
			Score: v.Score,
			Date:  v.Date.Format(constants.DateRFC),
		})

	}

	result = GetDetailResponse{
		Meta: &Meta{
			Message: "Get Detail Student SKPI",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &GetDetailResponseData{
			Id:                           data.Id,
			StudentId:                    data.StudentId,
			StudentNimNumber:             data.StudentNimNumber,
			StudentName:                  data.StudentName,
			StudentStudyProgramId:        data.StudentStudyProgramId,
			StudentStudyProgramName:      data.StudentStudyProgramName,
			StudentDiktiStudyProgramCode: data.StudentDiktiStudyProgramCode,
			SkpiNumber:                   utils.NullStringScan(data.SkpiNumber),
			IsApproved:                   data.IsApproved,
			AchievementPath:              utils.NullStringScan(data.AchievementPath),
			AchievementPathType:          utils.NullStringScan(data.AchievementPathType),
			AchievementUrl:               data.AchievementUrl,
			OrganizationPath:             utils.NullStringScan(data.OrganizationPath),
			OrganizationPathType:         utils.NullStringScan(data.OrganizationPathType),
			OrganizationUrl:              data.OrganizationUrl,
			CertificatePath:              utils.NullStringScan(data.CertificatePath),
			CertificatePathType:          utils.NullStringScan(data.CertificatePathType),
			CertificateUrl:               data.CertificateUrl,
			LanguagePath:                 utils.NullStringScan(data.LanguagePath),
			LanguagePathType:             utils.NullStringScan(data.LanguagePathType),
			LanguageUrl:                  data.LanguageUrl,
			Achievements:                 achievements,
			Organizations:                organizations,
			Certificates:                 certificates,
			CharacterBuildings:           characterBuildings,
			Internships:                  internships,
			Languages:                    languages,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (s studentSkpiHandler) Upsert(w http.ResponseWriter, r *http.Request) {
	var result UpsertResponse

	ctx := r.Context()
	var in UpsertRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = UpsertResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer s.AdminActivityLogService.Create(ctx, r, time.Now(), "Student SKPI", "Upsert", &in)

	achievements := []objects.UpsertStudentSkpiAchievement{}
	organizations := []objects.UpsertStudentSkpiOrganization{}
	certificates := []objects.UpsertStudentSkpiCertificate{}
	characterBuildings := []objects.UpsertStudentSkpiCharacterBuilding{}
	internships := []objects.UpsertStudentSkpiInternship{}
	languages := []objects.UpsertStudentSkpiLanguage{}
	for _, v := range in.GetAchievements() {
		achievements = append(achievements, objects.UpsertStudentSkpiAchievement{
			Name: v.GetName(),
			Year: v.GetYear(),
		})
	}
	for _, v := range in.GetOrganizations() {
		organizations = append(organizations, objects.UpsertStudentSkpiOrganization{
			Name:          v.GetName(),
			Position:      v.GetPosition(),
			ServiceLength: v.GetServiceLength(),
		})
	}
	for _, v := range in.GetCertificates() {
		certificates = append(certificates, objects.UpsertStudentSkpiCertificate{
			Name: v.GetName(),
		})
	}
	for _, v := range in.GetCharacterBuildings() {
		characterBuildings = append(characterBuildings, objects.UpsertStudentSkpiCharacterBuilding{
			Name: v.GetName(),
		})
	}
	for _, v := range in.GetInternships() {
		internships = append(internships, objects.UpsertStudentSkpiInternship{
			Name: v.GetName(),
		})
	}
	for _, v := range in.GetLanguages() {
		date, errs := utils.StringToTime(v.GetDate())
		if errs != nil {
			utils.PrintError(*errs)
			result = UpsertResponse{
				Meta: &Meta{
					Message: errs.Err.Error(),
					Status:  uint32(errs.HttpCode),
					Code:    errs.CustomCode,
				},
			}
			utils.JSONResponse(w, errs.HttpCode, &result)
			return
		}

		languages = append(languages, objects.UpsertStudentSkpiLanguage{
			Name:  v.GetName(),
			Score: v.GetScore(),
			Date:  date,
		})
	}

	data := objects.UpsertStudentSkpi{
		StudentId:            in.GetStudentId(),
		AchievementPath:      in.GetAchievementPath(),
		AchievementPathType:  in.GetAchievementPathType(),
		OrganizationPath:     in.GetOrganizationPath(),
		OrganizationPathType: in.GetOrganizationPathType(),
		CertificatePath:      in.GetCertificatePath(),
		CertificatePathType:  in.GetCertificatePathType(),
		LanguagePath:         in.GetLanguagePath(),
		LanguagePathType:     in.GetLanguagePathType(),
		Achievements:         achievements,
		Organizations:        organizations,
		Certificates:         certificates,
		CharacterBuildings:   characterBuildings,
		Internships:          internships,
		Languages:            languages,
	}
	errs := s.StudentSkpiService.Upsert(ctx, data)
	if errs != nil {
		utils.PrintError(*errs)
		result = UpsertResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = UpsertResponse{
		Meta: &Meta{
			Message: "Upsert Student SKPI",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (s studentSkpiHandler) Approve(w http.ResponseWriter, r *http.Request) {
	var result ApproveResponse

	ctx := r.Context()
	var in ApproveRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = ApproveResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer s.AdminActivityLogService.Create(ctx, r, time.Now(), "Student SKPI", "Approve", &in)

	data := objects.ApproveStudentSkpi{
		Id:         in.GetId(),
		SkpiNumber: in.GetSkpiNumber(),
	}
	errs := s.StudentSkpiService.Approve(ctx, data)
	if errs != nil {
		utils.PrintError(*errs)
		result = ApproveResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = ApproveResponse{
		Meta: &Meta{
			Message: "Approve Student SKPI",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (s studentSkpiHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
	defer s.AdminActivityLogService.Create(ctx, r, time.Now(), "Student SKPI", "Delete", nil)

	errs := s.StudentSkpiService.Delete(ctx, in.GetId())
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
			Message: "Delete Student SKPI",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
