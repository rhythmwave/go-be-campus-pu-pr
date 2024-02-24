package student_skpi

import (
	"context"

	appConstants "github.com/sccicitb/pupr-backend/constants"
	"google.golang.org/grpc/codes"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
)

type studentSkpiService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (s studentSkpiService) GetList(ctx context.Context, paginationData common.PaginationRequest, paramsData objects.GetStudentSkpiRequest) (objects.StudentSkpiListWithPagination, *constants.ErrorResponse) {
	var result objects.StudentSkpiListWithPagination

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	if paramsData.StudyProgramId != "" {
		_, errs := s.StudyProgramRepo.GetDetail(ctx, tx, paramsData.StudyProgramId, claims.Role, claims.ID)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	}

	modelResult, paginationResult, errs := s.StudentSkpiRepo.GetList(ctx, tx, paginationData, paramsData)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetStudentSkpi{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetStudentSkpi{
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

	result = objects.StudentSkpiListWithPagination{
		Pagination: paginationResult,
		Data:       resultData,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (s studentSkpiService) GetDetail(ctx context.Context, id string) (objects.GetStudentSkpiDetail, *constants.ErrorResponse) {
	var result objects.GetStudentSkpiDetail

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	var resultData models.GetStudentSkpiDetail
	if id != "" {
		resultData, errs = s.StudentSkpiRepo.GetById(ctx, tx, id)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	} else {
		resultData, errs = s.StudentSkpiRepo.GetByStudentId(ctx, tx, claims.ID)
		if errs != nil {
			if errs.GrpcCode == codes.NotFound {
				_ = tx.Rollback()
				return result, nil
			}
			_ = tx.Rollback()
			return result, errs
		}
	}

	achievementData, errs := s.StudentSkpiRepo.GetStudentSkpiAchievementByStudentSkpiId(ctx, tx, resultData.Id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}
	organizationData, errs := s.StudentSkpiRepo.GetStudentSkpiOrganizationByStudentSkpiId(ctx, tx, resultData.Id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}
	certificateData, errs := s.StudentSkpiRepo.GetStudentSkpiCertificateByStudentSkpiId(ctx, tx, resultData.Id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}
	characterBuildingData, errs := s.StudentSkpiRepo.GetStudentSkpiCharacterBuildingByStudentSkpiId(ctx, tx, resultData.Id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}
	internshipData, errs := s.StudentSkpiRepo.GetStudentSkpiInternshipByStudentSkpiId(ctx, tx, resultData.Id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}
	languageData, errs := s.StudentSkpiRepo.GetStudentSkpiLanguageByStudentSkpiId(ctx, tx, resultData.Id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result, errs = s.mapGetDetail(resultData, achievementData, organizationData, certificateData, characterBuildingData, internshipData, languageData)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (s studentSkpiService) Upsert(ctx context.Context, data objects.UpsertStudentSkpi) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	studentId := data.StudentId
	if studentId == "" {
		studentId = claims.ID
	}

	studentSkpiData := models.UpsertStudentSkpi{
		StudentId:            studentId,
		AchievementPath:      utils.NewNullString(data.AchievementPath),
		AchievementPathType:  utils.NewNullString(data.AchievementPathType),
		OrganizationPath:     utils.NewNullString(data.OrganizationPath),
		OrganizationPathType: utils.NewNullString(data.OrganizationPathType),
		CertificatePath:      utils.NewNullString(data.CertificatePath),
		CertificatePathType:  utils.NewNullString(data.CertificatePathType),
		LanguagePath:         utils.NewNullString(data.LanguagePath),
		LanguagePathType:     utils.NewNullString(data.LanguagePathType),
	}
	studentSkpiId, errs := s.StudentSkpiRepo.UpsertStudentSkpi(ctx, tx, studentSkpiData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	studentSkpiDetail, errs := s.StudentSkpiRepo.GetById(ctx, tx, studentSkpiId)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	if studentSkpiDetail.IsApproved {
		_ = tx.Rollback()
		return appConstants.ErrUneditableApprovedStudentSkpi
	}

	deleteStudentSkpiAchievementExcludedName := []string{}
	deleteStudentSkpiOrganizationExcludedName := []string{}
	deleteStudentSkpiCertificateExcludedName := []string{}
	deleteStudentSkpiCharacterBuildingExcludedName := []string{}
	deleteStudentSkpiInternshipExcludedName := []string{}
	deleteStudentSkpiLanguageExcludedName := []string{}

	studentSkpiAchievementData := []models.UpsertStudentSkpiAchievement{}
	studentSkpiOrganizationData := []models.UpsertStudentSkpiOrganization{}
	studentSkpiCertificateData := []models.UpsertStudentSkpiCertificate{}
	studentSkpiCharacterBuildingData := []models.UpsertStudentSkpiCharacterBuilding{}
	studentSkpiInternshipData := []models.UpsertStudentSkpiInternship{}
	studentSkpiLanguageData := []models.UpsertStudentSkpiLanguage{}
	for _, v := range data.Achievements {
		deleteStudentSkpiAchievementExcludedName = append(deleteStudentSkpiAchievementExcludedName, v.Name)
		studentSkpiAchievementData = append(studentSkpiAchievementData, models.UpsertStudentSkpiAchievement{
			StudentSkpiId: studentSkpiId,
			Name:          v.Name,
			Year:          v.Year,
		})
	}
	for _, v := range data.Organizations {
		deleteStudentSkpiOrganizationExcludedName = append(deleteStudentSkpiOrganizationExcludedName, v.Name)
		studentSkpiOrganizationData = append(studentSkpiOrganizationData, models.UpsertStudentSkpiOrganization{
			StudentSkpiId: studentSkpiId,
			Name:          v.Name,
			Position:      v.Position,
			ServiceLength: v.ServiceLength,
		})
	}
	for _, v := range data.Certificates {
		deleteStudentSkpiCertificateExcludedName = append(deleteStudentSkpiCertificateExcludedName, v.Name)
		studentSkpiCertificateData = append(studentSkpiCertificateData, models.UpsertStudentSkpiCertificate{
			StudentSkpiId: studentSkpiId,
			Name:          v.Name,
		})
	}
	for _, v := range data.CharacterBuildings {
		deleteStudentSkpiCharacterBuildingExcludedName = append(deleteStudentSkpiCharacterBuildingExcludedName, v.Name)
		studentSkpiCharacterBuildingData = append(studentSkpiCharacterBuildingData, models.UpsertStudentSkpiCharacterBuilding{
			StudentSkpiId: studentSkpiId,
			Name:          v.Name,
		})
	}
	for _, v := range data.Internships {
		deleteStudentSkpiInternshipExcludedName = append(deleteStudentSkpiInternshipExcludedName, v.Name)
		studentSkpiInternshipData = append(studentSkpiInternshipData, models.UpsertStudentSkpiInternship{
			StudentSkpiId: studentSkpiId,
			Name:          v.Name,
		})
	}
	for _, v := range data.Languages {
		deleteStudentSkpiLanguageExcludedName = append(deleteStudentSkpiLanguageExcludedName, v.Name)
		studentSkpiLanguageData = append(studentSkpiLanguageData, models.UpsertStudentSkpiLanguage{
			StudentSkpiId: studentSkpiId,
			Name:          v.Name,
			Score:         v.Score,
			Date:          v.Date,
		})
	}

	errs = s.StudentSkpiRepo.DeleteStudentSkpiAchievementExcludingName(ctx, tx, studentSkpiId, deleteStudentSkpiAchievementExcludedName)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	errs = s.StudentSkpiRepo.DeleteStudentSkpiOrganizationExcludingName(ctx, tx, studentSkpiId, deleteStudentSkpiOrganizationExcludedName)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	errs = s.StudentSkpiRepo.DeleteStudentSkpiCertificateExcludingName(ctx, tx, studentSkpiId, deleteStudentSkpiCertificateExcludedName)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	errs = s.StudentSkpiRepo.DeleteStudentSkpiCharacterBuildingExcludingName(ctx, tx, studentSkpiId, deleteStudentSkpiCharacterBuildingExcludedName)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	errs = s.StudentSkpiRepo.DeleteStudentSkpiInternshipExcludingName(ctx, tx, studentSkpiId, deleteStudentSkpiInternshipExcludedName)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	errs = s.StudentSkpiRepo.DeleteStudentSkpiLanguageExcludingName(ctx, tx, studentSkpiId, deleteStudentSkpiLanguageExcludedName)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	if len(studentSkpiAchievementData) != 0 {
		errs = s.StudentSkpiRepo.UpsertStudentSkpiAchievement(ctx, tx, studentSkpiAchievementData)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	}
	if len(studentSkpiOrganizationData) != 0 {
		errs = s.StudentSkpiRepo.UpsertStudentSkpiOrganization(ctx, tx, studentSkpiOrganizationData)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	}
	if len(studentSkpiCertificateData) != 0 {
		errs = s.StudentSkpiRepo.UpsertStudentSkpiCertificate(ctx, tx, studentSkpiCertificateData)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	}
	if len(studentSkpiCharacterBuildingData) != 0 {
		errs = s.StudentSkpiRepo.UpsertStudentSkpiCharacterBuilding(ctx, tx, studentSkpiCharacterBuildingData)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	}
	if len(studentSkpiInternshipData) != 0 {
		errs = s.StudentSkpiRepo.UpsertStudentSkpiInternship(ctx, tx, studentSkpiInternshipData)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	}
	if len(studentSkpiLanguageData) != 0 {
		errs = s.StudentSkpiRepo.UpsertStudentSkpiLanguage(ctx, tx, studentSkpiLanguageData)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (s studentSkpiService) Approve(ctx context.Context, data objects.ApproveStudentSkpi) *constants.ErrorResponse {
	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := s.StudentSkpiRepo.GetById(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = s.StudentSkpiRepo.Approve(ctx, tx, models.ApproveStudentSkpi{
		Id:         data.Id,
		SkpiNumber: data.SkpiNumber,
	})
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (s studentSkpiService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := s.StudentSkpiRepo.GetById(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = s.StudentSkpiRepo.Delete(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}
