package student_skpi

import (
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/objects"
)

func (s studentSkpiService) mapGetDetail(resultData models.GetStudentSkpiDetail, achievementData []models.GetStudentSkpiAchievement, organizationData []models.GetStudentSkpiOrganization, certificateData []models.GetStudentSkpiCertificate, characterBuildingData []models.GetStudentSkpiCharacterBuilding, internshipData []models.GetStudentSkpiInternship, languageData []models.GetStudentSkpiLanguage) (objects.GetStudentSkpiDetail, *constants.ErrorResponse) {
	var result objects.GetStudentSkpiDetail
	var errs *constants.ErrorResponse

	achievements := []objects.GetStudentSkpiDetailAchievement{}
	organizations := []objects.GetStudentSkpiDetailOrganization{}
	certificates := []objects.GetStudentSkpiDetailCertificate{}
	characterBuildings := []objects.GetStudentSkpiDetailCharacterBuilding{}
	internships := []objects.GetStudentSkpiDetailInternship{}
	languages := []objects.GetStudentSkpiDetailLanguage{}

	for _, v := range achievementData {
		achievements = append(achievements, objects.GetStudentSkpiDetailAchievement{
			Id:   v.Id,
			Name: v.Name,
			Year: v.Year,
		})
	}
	for _, v := range organizationData {
		organizations = append(organizations, objects.GetStudentSkpiDetailOrganization{
			Id:            v.Id,
			Name:          v.Name,
			Position:      v.Position,
			ServiceLength: v.ServiceLength,
		})
	}
	for _, v := range certificateData {
		certificates = append(certificates, objects.GetStudentSkpiDetailCertificate{
			Id:   v.Id,
			Name: v.Name,
		})
	}
	for _, v := range characterBuildingData {
		characterBuildings = append(characterBuildings, objects.GetStudentSkpiDetailCharacterBuilding{
			Id:   v.Id,
			Name: v.Name,
		})
	}
	for _, v := range internshipData {
		internships = append(internships, objects.GetStudentSkpiDetailInternship{
			Id:   v.Id,
			Name: v.Name,
		})
	}
	for _, v := range languageData {
		languages = append(languages, objects.GetStudentSkpiDetailLanguage{
			Id:    v.Id,
			Name:  v.Name,
			Score: v.Score,
			Date:  v.Date,
		})
	}

	var achievementUrl string
	if resultData.AchievementPath != nil && resultData.AchievementPathType != nil {
		achievementUrl, errs = s.Storage.GetURL(*resultData.AchievementPath, *resultData.AchievementPathType, nil)
		if errs != nil {
			return result, errs
		}
	}

	var organizationUrl string
	if resultData.OrganizationPath != nil && resultData.OrganizationPathType != nil {
		organizationUrl, errs = s.Storage.GetURL(*resultData.OrganizationPath, *resultData.OrganizationPathType, nil)
		if errs != nil {
			return result, errs
		}
	}

	var certificateUrl string
	if resultData.CertificatePath != nil && resultData.CertificatePathType != nil {
		certificateUrl, errs = s.Storage.GetURL(*resultData.CertificatePath, *resultData.CertificatePathType, nil)
		if errs != nil {
			return result, errs
		}
	}

	var languageUrl string
	if resultData.LanguagePath != nil && resultData.LanguagePathType != nil {
		languageUrl, errs = s.Storage.GetURL(*resultData.LanguagePath, *resultData.LanguagePathType, nil)
		if errs != nil {
			return result, errs
		}
	}

	result = objects.GetStudentSkpiDetail{
		Id:                           resultData.Id,
		StudentId:                    resultData.StudentId,
		StudentNimNumber:             resultData.StudentNimNumber,
		StudentName:                  resultData.StudentName,
		StudentStudyProgramId:        resultData.StudentStudyProgramId,
		StudentStudyProgramName:      resultData.StudentStudyProgramName,
		StudentDiktiStudyProgramCode: resultData.StudentDiktiStudyProgramCode,
		SkpiNumber:                   resultData.SkpiNumber,
		IsApproved:                   resultData.IsApproved,
		AchievementPath:              resultData.AchievementPath,
		AchievementPathType:          resultData.AchievementPathType,
		AchievementUrl:               achievementUrl,
		OrganizationPath:             resultData.OrganizationPath,
		OrganizationPathType:         resultData.OrganizationPathType,
		OrganizationUrl:              organizationUrl,
		CertificatePath:              resultData.CertificatePath,
		CertificatePathType:          resultData.CertificatePathType,
		CertificateUrl:               certificateUrl,
		LanguagePath:                 resultData.LanguagePath,
		LanguagePathType:             resultData.LanguagePathType,
		LanguageUrl:                  languageUrl,
		Achievements:                 achievements,
		Organizations:                organizations,
		Certificates:                 certificates,
		CharacterBuildings:           characterBuildings,
		Internships:                  internships,
		Languages:                    languages,
	}

	return result, nil
}
