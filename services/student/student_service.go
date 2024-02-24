package student

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/sccicitb/pupr-backend/constants"
	appConstants "github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
	appUtils "github.com/sccicitb/pupr-backend/utils"
)

type studentService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (s studentService) mapGetList(studentData []models.GetStudent, preHighSchoolHistoryData []models.GetStudentPreHighSchoolHistory) ([]objects.GetStudent, *constants.ErrorResponse) {
	results := []objects.GetStudent{}

	preHighSchoolHistoryMap := make(map[string][]objects.GetStudentPreHighSchoolHistory)
	for _, v := range preHighSchoolHistoryData {
		preHighSchoolHistoryMap[v.StudentId] = append(preHighSchoolHistoryMap[v.StudentId], objects.GetStudentPreHighSchoolHistory{
			Id:             v.Id,
			Level:          v.Level,
			Name:           v.Name,
			GraduationYear: v.GraduationYear,
		})
	}

	for _, v := range studentData {
		status := v.Status
		if v.StatusLog != nil {
			status = v.StatusLog
		}

		var academicGuidanceSemesterSchoolYear string
		if v.AcademicGuidanceSemesterStartYear != nil {
			academicGuidanceSemesterSchoolYear = appUtils.GenerateSchoolYear(*v.AcademicGuidanceSemesterStartYear)
		}

		var statusSemesterSchoolYear string
		if v.StatusSemesterStartYear != nil {
			statusSemesterSchoolYear = appUtils.GenerateSchoolYear(*v.StatusSemesterStartYear)
		}

		var profilePhotoUrl string
		var errs *constants.ErrorResponse
		if v.ProfilePhotoPath != nil && v.ProfilePhotoPathType != nil {
			profilePhotoUrl, errs = s.Storage.GetURL(*v.ProfilePhotoPath, *v.ProfilePhotoPathType, nil)
			if errs != nil {
				return results, errs
			}
		}

		var studyDurationMonth uint32
		var thesisDurationMonth uint32
		var thesisDurationSemester uint32
		if v.GraduationDate != nil {
			graduationDate := *v.GraduationDate
			diff := graduationDate.Sub(v.CreatedAt)
			studyDurationMonth = uint32(diff.Hours() / 24 / 30)
		}
		if v.ThesisStartDate != nil && v.ThesisFinishDate != nil {
			thesisStartDate := *v.ThesisStartDate
			thesisFinishDate := *v.ThesisFinishDate

			diff := thesisFinishDate.Sub(thesisStartDate)
			thesisDurationMonth = uint32(diff.Hours() / 24 / 30)
		}
		if v.ThesisStartSemesterId != nil && v.ThesisFinishSemesterId != nil {
			thesisStartSemesterId := *v.ThesisStartSemesterId
			thesisFinishSemesterId := *v.ThesisFinishSemesterId

			if thesisStartSemesterId != thesisFinishSemesterId {
				thesisDurationSemester = 2
			} else {
				thesisDurationSemester = 1
			}
		}

		results = append(results, objects.GetStudent{
			Id:                                 v.Id,
			Name:                               v.Name,
			Sex:                                v.Sex,
			MaritalStatus:                      v.MaritalStatus,
			BirthRegencyId:                     v.BirthRegencyId,
			BirthRegencyName:                   v.BirthRegencyName,
			BirthDate:                          v.BirthDate,
			Religion:                           v.Religion,
			Address:                            v.Address,
			Rt:                                 v.Rt,
			Rw:                                 v.Rw,
			VillageId:                          v.VillageId,
			VillageName:                        v.VillageName,
			DistrictId:                         v.DistrictId,
			DistrictName:                       v.DistrictName,
			RegencyId:                          v.RegencyId,
			RegencyName:                        v.RegencyName,
			ProvinceId:                         v.ProvinceId,
			ProvinceName:                       v.ProvinceName,
			CountryId:                          v.CountryId,
			CountryName:                        v.CountryName,
			PostalCode:                         v.PostalCode,
			PreviousAddress:                    v.PreviousAddress,
			IdNumber:                           v.IdNumber,
			NpwpNumber:                         v.NpwpNumber,
			NisnNumber:                         v.NisnNumber,
			ResidenceType:                      v.ResidenceType,
			TransportationMean:                 v.TransportationMean,
			KpsReceiver:                        v.KpsReceiver,
			PhoneNumber:                        v.PhoneNumber,
			MobilePhoneNumber:                  v.MobilePhoneNumber,
			Email:                              v.Email,
			Homepage:                           v.Homepage,
			WorkType:                           v.WorkType,
			WorkPlace:                          v.WorkPlace,
			Nationality:                        v.Nationality,
			AskesNumber:                        v.AskesNumber,
			TotalBrother:                       v.TotalBrother,
			TotalSister:                        v.TotalSister,
			Hobby:                              v.Hobby,
			Experience:                         v.Experience,
			TotalDependent:                     v.TotalDependent,
			NimNumber:                          v.NimNumber,
			DiktiStudyProgramType:              v.DiktiStudyProgramType,
			StudyLevelShortName:                v.StudyLevelShortName,
			StudentForce:                       v.StudentForce,
			AdmittanceSemester:                 v.AdmittanceSemester,
			StudyProgramId:                     v.StudyProgramId,
			StudyProgramName:                   v.StudyProgramName,
			CurriculumId:                       v.CurriculumId,
			CurriculumName:                     v.CurriculumName,
			AdmittanceTestNumber:               v.AdmittanceTestNumber,
			AdmittanceDate:                     v.AdmittanceDate,
			AdmittanceStatus:                   v.AdmittanceStatus,
			TotalTransferCredit:                v.TotalTransferCredit,
			PreviousCollege:                    v.PreviousCollege,
			PreviousStudyProgram:               v.PreviousStudyProgram,
			PreviousNimNumber:                  v.PreviousNimNumber,
			PreviousNimAdmittanceYear:          v.PreviousNimAdmittanceYear,
			Status:                             status,
			IsForeignStudent:                   v.IsForeignStudent,
			CollegeEntranceType:                v.CollegeEntranceType,
			ClassTime:                          v.ClassTime,
			FundSource:                         v.FundSource,
			IsScholarshipGrantee:               v.IsScholarshipGrantee,
			HasCompleteRequirement:             v.HasCompleteRequirement,
			SchoolCertificateType:              v.SchoolCertificateType,
			SchoolGraduationYear:               v.SchoolGraduationYear,
			SchoolName:                         v.SchoolName,
			SchoolAccreditation:                v.SchoolAccreditation,
			SchoolAddress:                      v.SchoolAddress,
			SchoolMajor:                        v.SchoolMajor,
			SchoolCertificateNumber:            v.SchoolCertificateNumber,
			SchoolCertificateDate:              v.SchoolCertificateDate,
			TotalSchoolFinalExamSubject:        v.TotalSchoolFinalExamSubject,
			SchoolFinalExamScore:               v.SchoolFinalExamScore,
			GuardianName:                       v.GuardianName,
			GuardianBirthDate:                  v.GuardianBirthDate,
			GuardianDeathDate:                  v.GuardianDeathDate,
			GuardianAddress:                    v.GuardianAddress,
			GuardianRegencyId:                  v.GuardianRegencyId,
			GuardianRegencyName:                v.GuardianRegencyName,
			GuardianPostalCode:                 v.GuardianPostalCode,
			GuardianPhoneNumber:                v.GuardianPhoneNumber,
			GuardianEmail:                      v.GuardianEmail,
			GuardianFinalAcademicBackground:    v.GuardianFinalAcademicBackground,
			GuardianOccupation:                 v.GuardianOccupation,
			FatherIdNumber:                     v.FatherIdNumber,
			FatherName:                         v.FatherName,
			FatherBirthDate:                    v.FatherBirthDate,
			FatherDeathDate:                    v.FatherDeathDate,
			MotherIdNumber:                     v.MotherIdNumber,
			MotherName:                         v.MotherName,
			MotherBirthDate:                    v.MotherBirthDate,
			MotherDeathDate:                    v.MotherDeathDate,
			ParentAddress:                      v.ParentAddress,
			ParentRegencyId:                    v.ParentRegencyId,
			ParentRegencyName:                  v.ParentRegencyName,
			ParentPostalCode:                   v.ParentPostalCode,
			ParentPhoneNumber:                  v.ParentPhoneNumber,
			ParentEmail:                        v.ParentEmail,
			FatherFinalAcademicBackground:      v.FatherFinalAcademicBackground,
			FatherOccupation:                   v.FatherOccupation,
			MotherFinalAcademicBackground:      v.MotherFinalAcademicBackground,
			MotherOccupation:                   v.MotherOccupation,
			ParentIncome:                       v.ParentIncome,
			IsFinanciallyCapable:               v.IsFinanciallyCapable,
			AuthenticationId:                   v.AuthenticationId,
			AuthenticationIsActive:             v.AuthenticationIsActive,
			AuthenticationSuspensionRemarks:    v.AuthenticationSuspensionRemarks,
			DiktiStudyProgramCode:              v.DiktiStudyProgramCode,
			AcademicGuidanceLecturerId:         v.AcademicGuidanceLecturerId,
			AcademicGuidanceLecturerName:       v.AcademicGuidanceLecturerName,
			AcademicGuidanceSemesterId:         v.AcademicGuidanceSemesterId,
			AcademicGuidanceSemesterSchoolYear: &academicGuidanceSemesterSchoolYear,
			StudyPlanId:                        v.StudyPlanId,
			StudyPlanTotalMandatoryCredit:      v.StudyPlanTotalMandatoryCredit,
			StudyPlanTotalOptionalCredit:       v.StudyPlanTotalOptionalCredit,
			StudyPlanMaximumCredit:             v.StudyPlanMaximumCredit,
			StudyPlanIsApproved:                v.StudyPlanIsApproved,
			CurrentSemesterPackage:             v.CurrentSemesterPackage,
			TotalStudyPlan:                     v.TotalStudyPlan,
			StatusSemesterId:                   v.StatusSemesterId,
			StatusSemesterSchoolYear:           statusSemesterSchoolYear,
			StatusSemesterType:                 v.StatusSemesterType,
			StatusDate:                         v.StatusDate,
			StatusReferenceNumber:              v.StatusReferenceNumber,
			StatusPurpose:                      v.StatusPurpose,
			StatusRemarks:                      v.StatusRemarks,
			ProfilePhotoPath:                   v.ProfilePhotoPath,
			ProfilePhotoPathType:               v.ProfilePhotoPathType,
			ProfilePhotoUrl:                    profilePhotoUrl,
			BirthProvinceId:                    v.BirthProvinceId,
			BirthProvinceName:                  v.BirthProvinceName,
			Height:                             v.Height,
			Weight:                             v.Weight,
			IsColorBlind:                       v.IsColorBlind,
			UseGlasses:                         v.UseGlasses,
			HasCompleteTeeth:                   v.HasCompleteTeeth,
			IsKpsRecipient:                     v.IsKpsRecipient,
			WorkAddress:                        v.WorkAddress,
			AssuranceNumber:                    v.AssuranceNumber,
			ParentReligion:                     v.ParentReligion,
			ParentNationality:                  v.ParentNationality,
			FatherWorkAddress:                  v.FatherWorkAddress,
			ParentProvinceId:                   v.ParentProvinceId,
			ParentProvinceName:                 v.ParentProvinceName,
			GuardianProvinceId:                 v.GuardianProvinceId,
			GuardianProvinceName:               v.GuardianProvinceName,
			BloodType:                          v.BloodType,
			SchoolEnrollmentYear:               v.SchoolEnrollmentYear,
			SchoolEnrollmentClass:              v.SchoolEnrollmentClass,
			SchoolType:                         v.SchoolType,
			SchoolProvinceId:                   v.SchoolProvinceId,
			SchoolProvinceName:                 v.SchoolProvinceName,
			SchoolStatus:                       v.SchoolStatus,
			SchoolMathematicsFinalExamScore:    v.SchoolMathematicsFinalExamScore,
			SchoolIndonesianFinalExamScore:     v.SchoolIndonesianFinalExamScore,
			SchoolEnglishFinalExamScore:        v.SchoolEnglishFinalExamScore,
			SchoolMathematicsReportScore:       v.SchoolMathematicsReportScore,
			SchoolIndonesianReportScore:        v.SchoolIndonesianReportScore,
			SchoolEnglishReportScore:           v.SchoolEnglishReportScore,
			Gpa:                                v.Gpa,
			TotalCredit:                        v.TotalCredit,
			TranscriptIsArchived:               v.TranscriptIsArchived,
			HasPaid:                            v.HasPaid,
			GraduationPredicate:                v.GraduationPredicate,
			StudyDurationMonth:                 studyDurationMonth,
			ThesisDurationMonth:                thesisDurationMonth,
			ThesisDurationSemester:             thesisDurationSemester,
			PreHighSchoolHistories:             preHighSchoolHistoryMap[v.Id],
		})
	}

	return results, nil
}

func mapGetStatusSummary(studyProgramData []models.GetStudyProgramList, resultData []models.StudentStatusSummary) []objects.GetStatusSummaryStudent {
	results := []objects.GetStatusSummaryStudent{}

	resultMap := make(map[string]uint32)
	for _, v := range resultData {
		key := fmt.Sprintf("%s-%s", utils.NullStringScan(v.StudyProgramId), v.Status)
		resultMap[key] = v.Total
	}

	for _, v := range studyProgramData {
		statuses := []objects.GetStatusSummaryStudentStatus{}
		for _, w := range appConstants.ValidStudentStatus() {
			key := fmt.Sprintf("%s-%s", v.Id, w)
			statuses = append(statuses, objects.GetStatusSummaryStudentStatus{
				Status: w,
				Total:  resultMap[key],
			})

		}
		results = append(results, objects.GetStatusSummaryStudent{
			StudyProgramId:        v.Id,
			StudyProgramName:      v.Name,
			DiktiStudyProgramType: v.DiktiStudyProgramType,
			StudyLevelShortName:   v.StudyLevelShortName,
			Statuses:              statuses,
		})
	}

	return results
}

func (a studentService) GetList(ctx context.Context, paginationData common.PaginationRequest, requestData objects.GetStudentRequest) (objects.StudentListWithPagination, *constants.ErrorResponse) {
	var result objects.StudentListWithPagination

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	if requestData.StudyProgramId != "" {
		_, errs = a.StudyProgramRepo.GetDetail(ctx, tx, requestData.StudyProgramId, claims.Role, claims.ID)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	}
	modelResult, paginationResult, errs := a.StudentRepo.GetList(ctx, tx, paginationData, requestData)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	studentIds := []string{}
	for _, v := range modelResult {
		studentIds = append(studentIds, v.Id)
	}
	preHighSchoolHistoryData := []models.GetStudentPreHighSchoolHistory{}
	if len(studentIds) != 0 {
		preHighSchoolHistoryData, errs = a.StudentRepo.GetPreHighSchoolHistoryByStudentIds(ctx, tx, studentIds)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	}

	resultData, errs := a.mapGetList(modelResult, preHighSchoolHistoryData)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}
	result = objects.StudentListWithPagination{
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

func (a studentService) GetDetail(ctx context.Context, id string) (objects.GetStudent, *constants.ErrorResponse) {
	var result objects.GetStudent

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	resultData, errs := a.StudentRepo.GetDetail(ctx, tx, id, "")
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	preHighSchoolHistoryData, errs := a.StudentRepo.GetPreHighSchoolHistoryByStudentIds(ctx, tx, []string{id})
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	results, errs := a.mapGetList([]models.GetStudent{resultData}, preHighSchoolHistoryData)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return results[0], nil
}

func (a studentService) Create(ctx context.Context, data objects.CreateStudent) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	if data.StudyProgramId != "" {
		_, errs = a.StudyProgramRepo.GetDetail(ctx, tx, data.StudyProgramId, claims.Role, claims.ID)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	}
	createData := models.CreateStudent{
		Name:                            data.Name,
		Sex:                             utils.NewNullString(data.Sex),
		MaritalStatus:                   utils.NewNullString(data.MaritalStatus),
		BirthRegencyId:                  utils.NewNullInt32(int32(data.BirthRegencyId)),
		BirthDate:                       utils.NewNullTime(data.BirthDate),
		Religion:                        utils.NewNullString(data.Religion),
		Address:                         utils.NewNullString(data.Address),
		Rt:                              utils.NewNullString(data.Rt),
		Rw:                              utils.NewNullString(data.Rw),
		VillageId:                       utils.NewNullInt32(int32(data.VillageId)),
		PostalCode:                      utils.NewNullString(data.PostalCode),
		PreviousAddress:                 utils.NewNullString(data.PreviousAddress),
		IdNumber:                        utils.NewNullString(data.IdNumber),
		NpwpNumber:                      utils.NewNullString(data.NpwpNumber),
		NisnNumber:                      utils.NewNullString(data.NisnNumber),
		ResidenceType:                   utils.NewNullString(data.ResidenceType),
		TransportationMean:              utils.NewNullString(data.TransportationMean),
		KpsReceiver:                     utils.NewNullString(data.KpsReceiver),
		PhoneNumber:                     utils.NewNullString(data.PhoneNumber),
		MobilePhoneNumber:               utils.NewNullString(data.MobilePhoneNumber),
		Email:                           utils.NewNullString(data.Email),
		Homepage:                        utils.NewNullString(data.Homepage),
		WorkType:                        utils.NewNullString(data.WorkType),
		WorkPlace:                       utils.NewNullString(data.WorkPlace),
		Nationality:                     utils.NewNullString(data.Nationality),
		AskesNumber:                     utils.NewNullString(data.AskesNumber),
		TotalBrother:                    utils.NewNullInt32(int32(data.TotalBrother)),
		TotalSister:                     utils.NewNullInt32(int32(data.TotalSister)),
		Hobby:                           utils.NewNullString(data.Hobby),
		Experience:                      utils.NewNullString(data.Experience),
		TotalDependent:                  utils.NewNullInt32(int32(data.TotalDependent)),
		NimNumber:                       data.NimNumber,
		StudentForce:                    utils.NewNullInt32(int32(data.StudentForce)),
		AdmittanceSemester:              utils.NewNullString(data.AdmittanceSemester),
		StudyProgramId:                  utils.NewNullString(data.StudyProgramId),
		CurriculumId:                    utils.NewNullString(data.CurriculumId),
		AdmittanceTestNumber:            utils.NewNullString(data.AdmittanceTestNumber),
		AdmittanceDate:                  utils.NewNullTime(data.AdmittanceDate),
		AdmittanceStatus:                utils.NewNullString(data.AdmittanceStatus),
		TotalTransferCredit:             utils.NewNullInt32(int32(data.TotalTransferCredit)),
		PreviousCollege:                 utils.NewNullString(data.PreviousCollege),
		PreviousStudyProgram:            utils.NewNullString(data.PreviousStudyProgram),
		PreviousNimNumber:               utils.NewNullInt64(data.PreviousNimNumber),
		PreviousNimAdmittanceYear:       utils.NewNullString(data.PreviousNimAdmittanceYear),
		Status:                          utils.NewNullString(data.Status),
		IsForeignStudent:                utils.NewNullBoolean(data.IsForeignStudent),
		CollegeEntranceType:             utils.NewNullString(data.CollegeEntranceType),
		ClassTime:                       utils.NewNullString(data.ClassTime),
		FundSource:                      utils.NewNullString(data.FundSource),
		IsScholarshipGrantee:            utils.NewNullBoolean(data.IsScholarshipGrantee),
		HasCompleteRequirement:          utils.NewNullBoolean(data.HasCompleteRequirement),
		SchoolCertificateType:           utils.NewNullString(data.SchoolCertificateType),
		SchoolGraduationYear:            utils.NewNullString(data.SchoolGraduationYear),
		SchoolName:                      utils.NewNullString(data.SchoolName),
		SchoolAccreditation:             utils.NewNullString(data.SchoolAccreditation),
		SchoolAddress:                   utils.NewNullString(data.SchoolAddress),
		SchoolMajor:                     utils.NewNullString(data.SchoolMajor),
		SchoolCertificateNumber:         utils.NewNullString(data.SchoolCertificateNumber),
		SchoolCertificateDate:           utils.NewNullTime(data.SchoolCertificateDate),
		TotalSchoolFinalExamSubject:     utils.NewNullInt32(int32(data.TotalSchoolFinalExamSubject)),
		SchoolFinalExamScore:            utils.NewNullFloat64(&data.SchoolFinalExamScore),
		GuardianName:                    utils.NewNullString(data.GuardianName),
		GuardianBirthDate:               utils.NewNullTime(data.GuardianBirthDate),
		GuardianDeathDate:               utils.NewNullTime(data.GuardianDeathDate),
		GuardianAddress:                 utils.NewNullString(data.GuardianAddress),
		GuardianRegencyId:               utils.NewNullInt32(int32(data.GuardianRegencyId)),
		GuardianPostalCode:              utils.NewNullString(data.GuardianPostalCode),
		GuardianPhoneNumber:             utils.NewNullString(data.GuardianPhoneNumber),
		GuardianEmail:                   utils.NewNullString(data.GuardianEmail),
		GuardianFinalAcademicBackground: utils.NewNullString(data.GuardianFinalAcademicBackground),
		GuardianOccupation:              utils.NewNullString(data.GuardianOccupation),
		FatherIdNumber:                  utils.NewNullString(data.FatherIdNumber),
		FatherName:                      utils.NewNullString(data.FatherName),
		FatherBirthDate:                 utils.NewNullTime(data.FatherBirthDate),
		FatherDeathDate:                 utils.NewNullTime(data.FatherDeathDate),
		MotherIdNumber:                  utils.NewNullString(data.MotherIdNumber),
		MotherName:                      utils.NewNullString(data.MotherName),
		MotherBirthDate:                 utils.NewNullTime(data.MotherBirthDate),
		MotherDeathDate:                 utils.NewNullTime(data.MotherDeathDate),
		ParentAddress:                   utils.NewNullString(data.ParentAddress),
		ParentRegencyId:                 utils.NewNullInt32(int32(data.ParentRegencyId)),
		ParentPostalCode:                utils.NewNullString(data.ParentPostalCode),
		ParentPhoneNumber:               utils.NewNullString(data.ParentPhoneNumber),
		ParentEmail:                     utils.NewNullString(data.ParentEmail),
		FatherFinalAcademicBackground:   utils.NewNullString(data.FatherFinalAcademicBackground),
		FatherOccupation:                utils.NewNullString(data.FatherOccupation),
		MotherFinalAcademicBackground:   utils.NewNullString(data.MotherFinalAcademicBackground),
		MotherOccupation:                utils.NewNullString(data.MotherOccupation),
		ParentIncome:                    utils.NewNullFloat64(&data.ParentIncome),
		IsFinanciallyCapable:            utils.NewNullBoolean(data.IsFinanciallyCapable),
	}
	errs = a.StudentRepo.Create(ctx, tx, createData)
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

func (a studentService) Update(ctx context.Context, data objects.UpdateStudent) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := a.StudentRepo.GetDetail(ctx, tx, data.Id, "")
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateStudent{
		Id:                              data.Id,
		Name:                            data.Name,
		Sex:                             utils.NewNullString(data.Sex),
		MaritalStatus:                   utils.NewNullString(data.MaritalStatus),
		BirthRegencyId:                  utils.NewNullInt32(int32(data.BirthRegencyId)),
		BirthDate:                       utils.NewNullTime(data.BirthDate),
		Religion:                        utils.NewNullString(data.Religion),
		Address:                         utils.NewNullString(data.Address),
		Rt:                              utils.NewNullString(data.Rt),
		Rw:                              utils.NewNullString(data.Rw),
		VillageId:                       utils.NewNullInt32(int32(data.VillageId)),
		PostalCode:                      utils.NewNullString(data.PostalCode),
		PreviousAddress:                 utils.NewNullString(data.PreviousAddress),
		IdNumber:                        utils.NewNullString(data.IdNumber),
		NpwpNumber:                      utils.NewNullString(data.NpwpNumber),
		NisnNumber:                      utils.NewNullString(data.NisnNumber),
		ResidenceType:                   utils.NewNullString(data.ResidenceType),
		TransportationMean:              utils.NewNullString(data.TransportationMean),
		KpsReceiver:                     utils.NewNullString(data.KpsReceiver),
		PhoneNumber:                     utils.NewNullString(data.PhoneNumber),
		MobilePhoneNumber:               utils.NewNullString(data.MobilePhoneNumber),
		Email:                           utils.NewNullString(data.Email),
		Homepage:                        utils.NewNullString(data.Homepage),
		WorkType:                        utils.NewNullString(data.WorkType),
		WorkPlace:                       utils.NewNullString(data.WorkPlace),
		Nationality:                     utils.NewNullString(data.Nationality),
		AskesNumber:                     utils.NewNullString(data.AskesNumber),
		TotalBrother:                    utils.NewNullInt32(int32(data.TotalBrother)),
		TotalSister:                     utils.NewNullInt32(int32(data.TotalSister)),
		Hobby:                           utils.NewNullString(data.Hobby),
		Experience:                      utils.NewNullString(data.Experience),
		TotalDependent:                  utils.NewNullInt32(int32(data.TotalDependent)),
		NimNumber:                       data.NimNumber,
		StudentForce:                    utils.NewNullInt32(int32(data.StudentForce)),
		AdmittanceSemester:              utils.NewNullString(data.AdmittanceSemester),
		StudyProgramId:                  utils.NewNullString(data.StudyProgramId),
		CurriculumId:                    utils.NewNullString(data.CurriculumId),
		AdmittanceTestNumber:            utils.NewNullString(data.AdmittanceTestNumber),
		AdmittanceDate:                  utils.NewNullTime(data.AdmittanceDate),
		AdmittanceStatus:                utils.NewNullString(data.AdmittanceStatus),
		TotalTransferCredit:             utils.NewNullInt32(int32(data.TotalTransferCredit)),
		PreviousCollege:                 utils.NewNullString(data.PreviousCollege),
		PreviousStudyProgram:            utils.NewNullString(data.PreviousStudyProgram),
		PreviousNimNumber:               utils.NewNullInt64(data.PreviousNimNumber),
		PreviousNimAdmittanceYear:       utils.NewNullString(data.PreviousNimAdmittanceYear),
		Status:                          utils.NewNullString(data.Status),
		IsForeignStudent:                utils.NewNullBoolean(data.IsForeignStudent),
		CollegeEntranceType:             utils.NewNullString(data.CollegeEntranceType),
		ClassTime:                       utils.NewNullString(data.ClassTime),
		FundSource:                      utils.NewNullString(data.FundSource),
		IsScholarshipGrantee:            utils.NewNullBoolean(data.IsScholarshipGrantee),
		HasCompleteRequirement:          utils.NewNullBoolean(data.HasCompleteRequirement),
		SchoolCertificateType:           utils.NewNullString(data.SchoolCertificateType),
		SchoolGraduationYear:            utils.NewNullString(data.SchoolGraduationYear),
		SchoolName:                      utils.NewNullString(data.SchoolName),
		SchoolAccreditation:             utils.NewNullString(data.SchoolAccreditation),
		SchoolAddress:                   utils.NewNullString(data.SchoolAddress),
		SchoolMajor:                     utils.NewNullString(data.SchoolMajor),
		SchoolCertificateNumber:         utils.NewNullString(data.SchoolCertificateNumber),
		SchoolCertificateDate:           utils.NewNullTime(data.SchoolCertificateDate),
		TotalSchoolFinalExamSubject:     utils.NewNullInt32(int32(data.TotalSchoolFinalExamSubject)),
		SchoolFinalExamScore:            utils.NewNullFloat64(&data.SchoolFinalExamScore),
		GuardianName:                    utils.NewNullString(data.GuardianName),
		GuardianBirthDate:               utils.NewNullTime(data.GuardianBirthDate),
		GuardianDeathDate:               utils.NewNullTime(data.GuardianDeathDate),
		GuardianAddress:                 utils.NewNullString(data.GuardianAddress),
		GuardianRegencyId:               utils.NewNullInt32(int32(data.GuardianRegencyId)),
		GuardianPostalCode:              utils.NewNullString(data.GuardianPostalCode),
		GuardianPhoneNumber:             utils.NewNullString(data.GuardianPhoneNumber),
		GuardianEmail:                   utils.NewNullString(data.GuardianEmail),
		GuardianFinalAcademicBackground: utils.NewNullString(data.GuardianFinalAcademicBackground),
		GuardianOccupation:              utils.NewNullString(data.GuardianOccupation),
		FatherIdNumber:                  utils.NewNullString(data.FatherIdNumber),
		FatherName:                      utils.NewNullString(data.FatherName),
		FatherBirthDate:                 utils.NewNullTime(data.FatherBirthDate),
		FatherDeathDate:                 utils.NewNullTime(data.FatherDeathDate),
		MotherIdNumber:                  utils.NewNullString(data.MotherIdNumber),
		MotherName:                      utils.NewNullString(data.MotherName),
		MotherBirthDate:                 utils.NewNullTime(data.MotherBirthDate),
		MotherDeathDate:                 utils.NewNullTime(data.MotherDeathDate),
		ParentAddress:                   utils.NewNullString(data.ParentAddress),
		ParentRegencyId:                 utils.NewNullInt32(int32(data.ParentRegencyId)),
		ParentPostalCode:                utils.NewNullString(data.ParentPostalCode),
		ParentPhoneNumber:               utils.NewNullString(data.ParentPhoneNumber),
		ParentEmail:                     utils.NewNullString(data.ParentEmail),
		FatherFinalAcademicBackground:   utils.NewNullString(data.FatherFinalAcademicBackground),
		FatherOccupation:                utils.NewNullString(data.FatherOccupation),
		MotherFinalAcademicBackground:   utils.NewNullString(data.MotherFinalAcademicBackground),
		MotherOccupation:                utils.NewNullString(data.MotherOccupation),
		ParentIncome:                    utils.NewNullFloat64(&data.ParentIncome),
		IsFinanciallyCapable:            utils.NewNullBoolean(data.IsFinanciallyCapable),
	}
	errs = a.StudentRepo.Update(ctx, tx, updateData)
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

func (a studentService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := a.StudentRepo.GetDetail(ctx, tx, id, "")
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = a.StudentRepo.Delete(ctx, tx, id)
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

func (a studentService) BulkUpdateStatus(ctx context.Context, data objects.BulkUpdateStatusStudent) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := a.StudentRepo.GetDetailByIds(ctx, tx, data.StudentIds)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	var validStatus bool
	for _, v := range appConstants.ManualEditableStatus() {
		if v == data.Status {
			validStatus = true
			break
		}
	}
	if !validStatus {
		return appConstants.ErrInvalidStudentStatus
	}

	updateData := models.BulkUpdateStatusStudent{
		Ids:                   data.StudentIds,
		Status:                data.Status,
		StatusReferenceNumber: utils.NewNullString(data.StatusReferenceNumber),
		StatusDate:            utils.NewNullTime(data.StatusDate),
		StatusPurpose:         utils.NewNullString(data.StatusPurpose),
		StatusRemarks:         utils.NewNullString(data.StatusRemarks),
	}
	errs = a.StudentRepo.BulkUpdateStatus(ctx, tx, updateData)
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

func (a studentService) GetStatusSummary(ctx context.Context, semesterId string) ([]objects.GetStatusSummaryStudent, *constants.ErrorResponse) {
	result := []objects.GetStatusSummaryStudent{}
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	paginationData := common.PaginationRequest{
		Page:  constants.DefaultPage,
		Limit: constants.DefaultUnlimited,
	}
	studyProgramData, _, errs := a.StudyProgramRepo.GetList(ctx, tx, paginationData, "", claims.Role, claims.ID)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	studyProgramIds := []string{}
	for _, v := range studyProgramData {
		studyProgramIds = append(studyProgramIds, v.Id)
	}

	if len(studyProgramIds) != 0 {
		resultData, errs := a.StudentRepo.GetStatusSummary(ctx, tx, studyProgramIds, semesterId)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}

		result = mapGetStatusSummary(studyProgramData, resultData)
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (s studentService) GetSemesterSummary(ctx context.Context) (objects.GetStudentSemesterSummary, *constants.ErrorResponse) {
	var result objects.GetStudentSemesterSummary

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	activeSemesterData, errs := s.SemesterRepo.GetActive(ctx, tx)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}
	previousSemesterData, errs := s.SemesterRepo.GetPreviousSemester(ctx, tx, activeSemesterData.Id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}
	studentData, errs := s.StudentRepo.GetDetail(ctx, tx, claims.ID, previousSemesterData.Id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}
	creditQuotaData, _, errs := s.CreditQuotaRepo.GetList(ctx, tx, common.PaginationRequest{Page: constants.DefaultPage, Limit: constants.DefaultUnlimited})
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}
	studyPlanData, errs := s.StudyPlanRepo.GetApprovedByStudentId(ctx, tx, claims.ID)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	var totalMandatoryCreditTaken uint32
	var totalOptionalCreditTaken uint32
	var gpa float64
	var i float64
	for _, v := range studyPlanData {
		totalMandatoryCreditTaken += v.TotalMandatoryCredit
		totalOptionalCreditTaken += v.TotalOptionalCredit
		if v.GradePoint != 0 {
			gpa += v.GradePoint
			i++
		}
	}
	if i != 0 {
		gpa = gpa / i
	}

	result = objects.GetStudentSemesterSummary{
		SemesterId:                         activeSemesterData.Id,
		SemesterSchoolYear:                 appUtils.GenerateSchoolYear(activeSemesterData.SemesterStartYear),
		SemesterType:                       activeSemesterData.SemesterType,
		Status:                             studentData.Status,
		StudyProgramId:                     studentData.StudyProgramId,
		StudyProgramName:                   studentData.StudyProgramName,
		HasPaid:                            studentData.HasPaid,
		AcademicGuidanceLecturerId:         studentData.AcademicGuidanceLecturerId,
		AcademicGuidanceLecturerName:       studentData.AcademicGuidanceLecturerName,
		AcademicGuidanceLecturerFrontTitle: studentData.AcademicGuidanceLecturerFrontTitle,
		AcademicGuidanceLecturerBackDegree: studentData.AcademicGuidanceLecturerBackDegree,
		MaximumCredit:                      appUtils.GetMaximumCredit(creditQuotaData, studentData.PreviousSemesterGradePoint),
		StudyPlanInputStartDate:            activeSemesterData.StudyPlanInputStartDate,
		StudyPlanInputEndDate:              activeSemesterData.StudyPlanInputEndDate,
		StudyPlanApprovalStartDate:         activeSemesterData.StudyPlanApprovalStartDate,
		StudyPlanApprovalEndDate:           activeSemesterData.StudyPlanApprovalEndDate,
		TotalMandatoryCreditTaken:          totalMandatoryCreditTaken,
		TotalOptionalCreditTaken:           totalOptionalCreditTaken,
		Gpa:                                gpa,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (a studentService) GetProfile(ctx context.Context) (objects.GetStudent, *constants.ErrorResponse) {
	var result objects.GetStudent

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	resultData, errs := a.StudentRepo.GetDetail(ctx, tx, claims.ID, "")
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}
	preHighSchoolHistoryData, errs := a.StudentRepo.GetPreHighSchoolHistoryByStudentIds(ctx, tx, []string{claims.ID})
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	data, errs := a.mapGetList([]models.GetStudent{resultData}, preHighSchoolHistoryData)
	if errs != nil {
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return data[0], nil
}

func (a studentService) UpdateProfile(ctx context.Context, data objects.UpdateStudentProfile) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs = a.StudentRepo.GetDetail(ctx, tx, claims.ID, "")
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateStudentProfile{
		Id:                   claims.ID,
		ProfilePhotoPath:     utils.NewNullString(data.ProfilePhotoPath),
		ProfilePhotoPathType: utils.NewNullString(data.ProfilePhotoPathType),
		Sex:                  utils.NewNullString(data.Sex),
		BirthRegencyId:       utils.NewNullInt32(int32(data.BirthRegencyId)),
		BloodType:            utils.NewNullString(data.BloodType),
		Height:               utils.NewNullFloat64(&data.Height),
		Weight:               utils.NewNullFloat64(&data.Weight),
		IsColorBlind:         utils.NewNullBoolean(data.IsColorBlind),
		UseGlasses:           utils.NewNullBoolean(data.UseGlasses),
		HasCompleteTeeth:     utils.NewNullBoolean(data.HasCompleteTeeth),
		IdNumber:             utils.NewNullString(data.IdNumber),
		NpwpNumber:           utils.NewNullString(data.NpwpNumber),
		NisnNumber:           utils.NewNullString(data.NisnNumber),
		Religion:             utils.NewNullString(data.Religion),
		MaritalStatus:        utils.NewNullString(data.MaritalStatus),
		Nationality:          utils.NewNullString(data.Nationality),
		VillageId:            utils.NewNullInt32(int32(data.VillageId)),
		Rt:                   utils.NewNullString(data.Rt),
		Rw:                   utils.NewNullString(data.Rw),
		PostalCode:           utils.NewNullString(data.PostalCode),
		Address:              utils.NewNullString(data.Address),
		PhoneNumber:          utils.NewNullString(data.PhoneNumber),
		MobilePhoneNumber:    utils.NewNullString(data.MobilePhoneNumber),
		Email:                utils.NewNullString(data.Email),
		TransportationMean:   utils.NewNullString(data.TransportationMean),
		IsKpsRecipient:       utils.NewNullBoolean(data.IsKpsRecipient),
		FundSource:           utils.NewNullString(data.FundSource),
		IsScholarshipGrantee: utils.NewNullBoolean(data.IsScholarshipGrantee),
		TotalBrother:         utils.NewNullInt32(int32(data.TotalBrother)),
		TotalSister:          utils.NewNullInt32(int32(data.TotalSister)),
		WorkType:             utils.NewNullString(data.WorkType),
		WorkPlace:            utils.NewNullString(data.WorkPlace),
		WorkAddress:          utils.NewNullString(data.WorkAddress),
		AssuranceNumber:      utils.NewNullString(data.AssuranceNumber),
		Hobby:                utils.NewNullString(data.Hobby),
	}
	errs = a.StudentRepo.UpdateProfile(ctx, tx, updateData)
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

func (a studentService) UpdateParentProfile(ctx context.Context, data objects.UpdateStudentParentProfile) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs = a.StudentRepo.GetDetail(ctx, tx, claims.ID, "")
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateStudentParentProfile{
		Id:                              claims.ID,
		FatherIdNumber:                  utils.NewNullString(data.FatherIdNumber),
		FatherName:                      utils.NewNullString(data.FatherName),
		FatherBirthDate:                 utils.NewNullString(data.FatherBirthDate),
		FatherDeathDate:                 utils.NewNullString(data.FatherDeathDate),
		FatherFinalAcademicBackground:   utils.NewNullString(data.FatherFinalAcademicBackground),
		FatherOccupation:                utils.NewNullString(data.FatherOccupation),
		MotherIdNumber:                  utils.NewNullString(data.MotherIdNumber),
		MotherName:                      utils.NewNullString(data.MotherName),
		MotherBirthDate:                 utils.NewNullString(data.MotherBirthDate),
		MotherDeathDate:                 utils.NewNullString(data.MotherDeathDate),
		MotherFinalAcademicBackground:   utils.NewNullString(data.MotherFinalAcademicBackground),
		MotherOccupation:                utils.NewNullString(data.MotherOccupation),
		ParentReligion:                  utils.NewNullString(data.ParentReligion),
		ParentNationality:               utils.NewNullString(data.ParentNationality),
		ParentAddress:                   utils.NewNullString(data.ParentAddress),
		FatherWorkAddress:               utils.NewNullString(data.FatherWorkAddress),
		ParentRegencyId:                 utils.NewNullInt32(int32(data.ParentRegencyId)),
		ParentPostalCode:                utils.NewNullString(data.ParentPostalCode),
		ParentPhoneNumber:               utils.NewNullString(data.ParentPhoneNumber),
		ParentEmail:                     utils.NewNullString(data.ParentEmail),
		IsFinanciallyCapable:            utils.NewNullBoolean(data.IsFinanciallyCapable),
		ParentIncome:                    utils.NewNullFloat64(&data.ParentIncome),
		TotalDependent:                  utils.NewNullInt32(int32(data.TotalDependent)),
		GuardianName:                    utils.NewNullString(data.GuardianName),
		GuardianBirthDate:               utils.NewNullString(data.GuardianBirthDate),
		GuardianDeathDate:               utils.NewNullString(data.GuardianDeathDate),
		GuardianAddress:                 utils.NewNullString(data.GuardianAddress),
		GuardianRegencyId:               utils.NewNullInt32(int32(data.GuardianRegencyId)),
		GuardianPostalCode:              utils.NewNullString(data.GuardianPostalCode),
		GuardianPhoneNumber:             utils.NewNullString(data.GuardianPhoneNumber),
		GuardianEmail:                   utils.NewNullString(data.GuardianEmail),
		GuardianFinalAcademicBackground: utils.NewNullString(data.GuardianFinalAcademicBackground),
		GuardianOccupation:              utils.NewNullString(data.GuardianOccupation),
	}
	errs = a.StudentRepo.UpdateParentProfile(ctx, tx, updateData)
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

func (a studentService) UpdateSchoolProfile(ctx context.Context, data objects.UpdateStudentSchoolProfile) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs = a.StudentRepo.GetDetail(ctx, tx, claims.ID, "")
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateStudentSchoolProfile{
		Id:                              claims.ID,
		SchoolEnrollmentYear:            utils.NewNullString(data.SchoolEnrollmentYear),
		SchoolGraduationYear:            utils.NewNullString(data.SchoolGraduationYear),
		SchoolEnrollmentClass:           utils.NewNullString(data.SchoolEnrollmentClass),
		SchoolMajor:                     utils.NewNullString(data.SchoolMajor),
		SchoolType:                      utils.NewNullString(data.SchoolType),
		SchoolName:                      utils.NewNullString(data.SchoolName),
		SchoolProvinceId:                utils.NewNullInt32(int32(data.SchoolProvinceId)),
		SchoolAddress:                   utils.NewNullString(data.SchoolAddress),
		SchoolCertificateNumber:         utils.NewNullString(data.SchoolCertificateNumber),
		SchoolCertificateDate:           utils.NewNullString(data.SchoolCertificateDate),
		SchoolStatus:                    utils.NewNullString(data.SchoolStatus),
		SchoolAccreditation:             utils.NewNullString(data.SchoolAccreditation),
		SchoolFinalExamScore:            utils.NewNullFloat64(&data.SchoolFinalExamScore),
		SchoolMathematicsFinalExamScore: utils.NewNullFloat64(&data.SchoolMathematicsFinalExamScore),
		SchoolIndonesianFinalExamScore:  utils.NewNullFloat64(&data.SchoolIndonesianFinalExamScore),
		SchoolEnglishFinalExamScore:     utils.NewNullFloat64(&data.SchoolEnglishFinalExamScore),
		SchoolMathematicsReportScore:    utils.NewNullFloat64(&data.SchoolMathematicsReportScore),
		SchoolIndonesianReportScore:     utils.NewNullFloat64(&data.SchoolIndonesianReportScore),
		SchoolEnglishReportScore:        utils.NewNullFloat64(&data.SchoolEnglishReportScore),
	}
	errs = a.StudentRepo.UpdateSchoolProfile(ctx, tx, updateData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	preHighSchoollevels := []string{}
	preHighSchoolData := []models.UpsertStudentPreHighSchoolHistory{}
	for _, v := range data.PreHighSchoolHistories {
		preHighSchoollevels = append(preHighSchoollevels, v.Level)
		preHighSchoolData = append(preHighSchoolData, models.UpsertStudentPreHighSchoolHistory{
			StudentId:      claims.ID,
			Level:          v.Level,
			Name:           v.Name,
			GraduationYear: v.GraduationYear,
		})
	}

	errs = a.StudentRepo.DeletePreHighSchoolHistoryExcludingLevel(ctx, tx, claims.ID, preHighSchoollevels)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = a.StudentRepo.UpsertPreHighSchoolHistory(ctx, tx, preHighSchoolData)
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

func (s studentService) GetSubjectGrade(ctx context.Context, paginationData common.PaginationRequest, studentId string) (objects.StudentSubjectWithPagination, *constants.ErrorResponse) {
	var result objects.StudentSubjectWithPagination

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := s.StudentRepo.GetStudentSubject(ctx, tx, paginationData, studentId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetStudentSubject{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetStudentSubject{
			SubjectId:               v.SubjectId,
			SubjectCode:             v.SubjectCode,
			SubjectName:             v.SubjectName,
			GradeSemesterId:         v.GradeSemesterId,
			GradeSemesterSchoolYear: appUtils.GenerateSchoolYear(v.GradeSemesterStartYear),
			GradeSemesterType:       v.GradeSemesterType,
			GradePoint:              v.GradePoint,
			GradeCode:               v.GradeCode,
			SubjectIsMandatory:      v.SubjectIsMandatory,
			SemesterPackage:         v.SemesterPackage,
			SubjectTotalCredit:      v.SubjectTotalCredit,
			SubjectType:             v.SubjectType,
		})
	}

	result = objects.StudentSubjectWithPagination{
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

func (s studentService) BulkUpdatePayment(ctx context.Context, studentIds []string) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	errs = s.StudentRepo.UpdatePayment(ctx, tx, studentIds, claims.ID)
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

func (s studentService) GetPaymentLog(ctx context.Context, studentId string) ([]objects.GetStudentPaymentLog, *constants.ErrorResponse) {
	var result []objects.GetStudentPaymentLog

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	resultData, errs := s.StudentRepo.GetPaymentLog(ctx, tx, studentId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	for _, v := range resultData {
		result = append(result, objects.GetStudentPaymentLog{
			SemesterId:         v.SemesterId,
			SemesterType:       v.SemesterType,
			SemesterStartYear:  v.SemesterStartYear,
			SemesterSchoolYear: appUtils.GenerateSchoolYear(v.SemesterStartYear),
			CreatedAt:          v.CreatedAt,
		})
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (s studentService) BulkCreate(ctx context.Context, data []objects.BulkCreateStudent) ([]objects.BulkCreateAuthenticationResponse, *constants.ErrorResponse) {
	var result []objects.BulkCreateAuthenticationResponse

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	studyProgramData, _, errs := s.StudyProgramRepo.GetList(ctx, tx, common.PaginationRequest{Page: constants.DefaultPage, Limit: constants.DefaultUnlimited}, "", "", "")
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	studyProgramMap := make(map[string]string)
	for _, v := range studyProgramData {
		studyProgramMap[v.DiktiStudyProgramCode] = v.Id
	}

	var createData []models.BulkCreateStudent
	var authData []models.CreateAuthentication
	for _, v := range data {
		if v.NimNumber == 0 {
			_ = tx.Rollback()
			return result, appConstants.ErrInvalidNimNumber
		}
		studyProgramId := studyProgramMap[v.DiktiStudyProgramCode]
		if studyProgramId == "" {
			_ = tx.Rollback()
			return result, appConstants.ErrInvalidDiktiStudyProgramCode
		}

		studentId := uuid.New().String()
		password := utils.Uid(8)
		hashedPassword, err := utils.HashPassword(password)
		if err != nil {
			return result, constants.ErrorInternalServer(err.Error())
		}
		createData = append(createData, models.BulkCreateStudent{
			Id:                              studentId,
			NimNumber:                       v.NimNumber,
			Name:                            v.Name,
			Sex:                             utils.NewNullString(v.Sex),
			MaritalStatus:                   utils.NewNullString(v.MaritalStatus),
			BirthRegencyId:                  utils.NewNullInt32(int32(v.BirthRegencyId)),
			BirthDate:                       utils.NewNullTime(v.BirthDate),
			Religion:                        utils.NewNullString(v.Religion),
			Address:                         utils.NewNullString(v.Address),
			Rt:                              utils.NewNullString(v.Rt),
			Rw:                              utils.NewNullString(v.Rw),
			VillageId:                       utils.NewNullInt32(int32(v.VillageId)),
			PostalCode:                      utils.NewNullString(v.PostalCode),
			IdNumber:                        utils.NewNullString(v.IdNumber),
			NisnNumber:                      utils.NewNullString(v.NisnNumber),
			MobilePhoneNumber:               utils.NewNullString(v.MobilePhoneNumber),
			Nationality:                     utils.NewNullString(v.Nationality),
			StudyProgramId:                  utils.NewNullString(studyProgramId),
			SchoolName:                      utils.NewNullString(v.SchoolName),
			SchoolAddress:                   utils.NewNullString(v.SchoolAddress),
			SchoolProvinceId:                utils.NewNullInt32(int32(v.SchoolProvinceId)),
			SchoolMajor:                     utils.NewNullString(v.SchoolMajor),
			SchoolType:                      utils.NewNullString(v.SchoolType),
			SchoolGraduationYear:            utils.NewNullString(v.SchoolGraduationYear),
			FatherName:                      utils.NewNullString(v.FatherName),
			FatherIdNumber:                  utils.NewNullString(v.FatherIdNumber),
			FatherBirthDate:                 utils.NewNullTime(v.FatherBirthDate),
			FatherFinalAcademicBackground:   utils.NewNullString(v.FatherFinalAcademicBackground),
			FatherOccupation:                utils.NewNullString(v.FatherOccupation),
			MotherName:                      utils.NewNullString(v.MotherName),
			MotherIdNumber:                  utils.NewNullString(v.MotherIdNumber),
			MotherBirthDate:                 utils.NewNullTime(v.MotherBirthDate),
			MotherFinalAcademicBackground:   utils.NewNullString(v.MotherFinalAcademicBackground),
			MotherOccupation:                utils.NewNullString(v.MotherOccupation),
			GuardianName:                    utils.NewNullString(v.GuardianName),
			GuardianIdNumber:                utils.NewNullString(v.GuardianIdNumber),
			GuardianBirthDate:               utils.NewNullTime(v.GuardianBirthDate),
			GuardianFinalAcademicBackground: utils.NewNullString(v.GuardianFinalAcademicBackground),
			GuardianOccupation:              utils.NewNullString(v.GuardianOccupation),
		})

		username := strconv.Itoa(int(v.NimNumber))
		result = append(result, objects.BulkCreateAuthenticationResponse{
			UserId:   studentId,
			Name:     v.Name,
			Username: username,
			Password: password,
		})
		authData = append(authData, models.CreateAuthentication{
			Username:  username,
			Password:  hashedPassword,
			StudentId: utils.NewNullString(studentId),
		})
	}

	if len(createData) != 0 {
		errs = s.StudentRepo.BulkCreate(ctx, tx, createData)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}

		errs = s.AuthenticationRepo.Create(ctx, tx, authData)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}
