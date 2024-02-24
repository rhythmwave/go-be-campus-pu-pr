package general

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/service"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/utils"
	"github.com/sirupsen/logrus"
)

type studentGeneralHandler struct {
	*service.ServiceCtx
}

func (l studentGeneralHandler) GetSemesterSummary(w http.ResponseWriter, r *http.Request) {
	var result GetSemesterSummaryResponse

	ctx := r.Context()

	defer l.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "General", "Get Semester Summary", nil)

	data, errs := l.StudentService.GetSemesterSummary(ctx)
	if errs != nil {
		utils.PrintError(*errs)
		result = GetSemesterSummaryResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = GetSemesterSummaryResponse{
		Meta: &Meta{
			Message: "Get Semester Summary Student",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &GetSemesterSummaryResponseData{
			SemesterId:                         data.SemesterId,
			SemesterSchoolYear:                 data.SemesterSchoolYear,
			SemesterType:                       data.SemesterType,
			Status:                             utils.NullStringScan(data.Status),
			StudyProgramId:                     utils.NullStringScan(data.StudyProgramId),
			StudyProgramName:                   utils.NullStringScan(data.StudyProgramName),
			HasPaid:                            data.HasPaid,
			AcademicGuidanceLecturerId:         utils.NullStringScan(data.AcademicGuidanceLecturerId),
			AcademicGuidanceLecturerName:       utils.NullStringScan(data.AcademicGuidanceLecturerName),
			AcademicGuidanceLecturerFrontTitle: utils.NullStringScan(data.AcademicGuidanceLecturerFrontTitle),
			AcademicGuidanceLecturerBackDegree: utils.NullStringScan(data.AcademicGuidanceLecturerBackDegree),
			MaximumCredit:                      data.MaximumCredit,
			StudyPlanInputStartDate:            data.StudyPlanInputStartDate.Format(constants.DateRFC),
			StudyPlanInputEndDate:              data.StudyPlanInputEndDate.Format(constants.DateRFC),
			StudyPlanApprovalStartDate:         data.StudyPlanApprovalStartDate.Format(constants.DateRFC),
			StudyPlanApprovalEndDate:           data.StudyPlanApprovalEndDate.Format(constants.DateRFC),
			TotalMandatoryCreditTaken:          data.TotalMandatoryCreditTaken,
			TotalOptionalCreditTaken:           data.TotalOptionalCreditTaken,
			Gpa:                                data.Gpa,
		},
	}

	utils.JSONResponse(w, http.StatusOK, &result)
}

func (l studentGeneralHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	var result GetProfileResponse

	ctx := r.Context()

	defer l.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "General", "Get Profile", nil)

	data, errs := l.StudentService.GetProfile(ctx)
	if errs != nil {
		utils.PrintError(*errs)
		result = GetProfileResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = GetProfileResponse{
		Meta: &Meta{
			Message: "Get Profile Student",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &GetProfileResponseData{
			Id:                    data.Id,
			ProfilePhotoPath:      utils.NullStringScan(data.ProfilePhotoPath),
			ProfilePhotoPathType:  utils.NullStringScan(data.ProfilePhotoPathType),
			ProfilePhotoUrl:       data.ProfilePhotoUrl,
			Name:                  data.Name,
			Sex:                   utils.NullStringScan(data.Sex),
			BirthProvinceId:       utils.NullUint32Scan(data.BirthProvinceId),
			BirthProvinceName:     utils.NullStringScan(data.BirthProvinceName),
			BirthRegencyId:        utils.NullUint32Scan(data.BirthRegencyId),
			BirthRegencyName:      utils.NullStringScan(data.BirthRegencyName),
			BirthDate:             utils.SafetyDate(data.BirthDate),
			BloodType:             utils.NullStringScan(data.BloodType),
			Height:                utils.NullFloatScan(data.Height),
			Weight:                utils.NullFloatScan(data.Weight),
			IsColorBlind:          utils.NullBooleanScan(data.IsColorBlind),
			UseGlasses:            utils.NullBooleanScan(data.UseGlasses),
			HasCompleteTeeth:      utils.NullBooleanScan(data.HasCompleteTeeth),
			StudyProgramId:        utils.NullStringScan(data.StudyProgramId),
			StudyProgramName:      utils.NullStringScan(data.StudyProgramName),
			DiktiStudyProgramType: utils.NullStringScan(data.DiktiStudyProgramType),
			StudyLevelShortName:   utils.NullStringScan(data.StudyLevelShortName),
			AdmittanceSemester:    utils.NullStringScan(data.AdmittanceSemester),
			StudentForce:          utils.NullUint32Scan(data.StudentForce),
			AdmittanceTestNumber:  utils.NullStringScan(data.AdmittanceTestNumber),
			CollegeEntranceType:   utils.NullStringScan(data.CollegeEntranceType),
			AdmittanceDate:        utils.SafetyDate(data.AdmittanceDate),
			AdmittanceStatus:      utils.NullStringScan(data.AdmittanceStatus),
			NpwpNumber:            utils.NullStringScan(data.NpwpNumber),
			NisnNumber:            utils.NullStringScan(data.NisnNumber),
			Religion:              utils.NullStringScan(data.Religion),
			MaritalStatus:         utils.NullStringScan(data.MaritalStatus),
			Nationality:           utils.NullStringScan(data.Nationality),
			ProvinceId:            utils.NullUint32Scan(data.ProvinceId),
			ProvinceName:          utils.NullStringScan(data.ProvinceName),
			RegencyId:             utils.NullUint32Scan(data.RegencyId),
			RegencyName:           utils.NullStringScan(data.RegencyName),
			DistrictId:            utils.NullUint32Scan(data.DistrictId),
			DistrictName:          utils.NullStringScan(data.DistrictName),
			VillageId:             utils.NullUint32Scan(data.VillageId),
			VillageName:           utils.NullStringScan(data.VillageName),
			Rt:                    utils.NullStringScan(data.Rt),
			Rw:                    utils.NullStringScan(data.Rw),
			PostalCode:            utils.NullStringScan(data.PostalCode),
			Address:               utils.NullStringScan(data.Address),
			PhoneNumber:           utils.NullStringScan(data.PhoneNumber),
			MobilePhoneNumber:     utils.NullStringScan(data.MobilePhoneNumber),
			Email:                 utils.NullStringScan(data.Email),
			TransportationMean:    utils.NullStringScan(data.TransportationMean),
			IsKpsRecipient:        utils.NullBooleanScan(data.IsKpsRecipient),
			FundSource:            utils.NullStringScan(data.FundSource),
			IsScholarshipGrantee:  utils.NullBooleanScan(data.IsScholarshipGrantee),
			TotalBrother:          utils.NullUint32Scan(data.TotalBrother),
			TotalSister:           utils.NullUint32Scan(data.TotalSister),
			WorkType:              utils.NullStringScan(data.WorkType),
			WorkPlace:             utils.NullStringScan(data.WorkPlace),
			WorkAddress:           utils.NullStringScan(data.WorkAddress),
			AssuranceNumber:       utils.NullStringScan(data.AssuranceNumber),
			Hobby:                 utils.NullStringScan(data.Hobby),
		},
	}

	utils.JSONResponse(w, http.StatusOK, &result)
}

func (l studentGeneralHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var result UpdateProfileResponse

	ctx := r.Context()
	var in UpdateProfileRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = UpdateProfileResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}

	defer l.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "General", "Update Profile", nil)

	data := objects.UpdateStudentProfile{
		ProfilePhotoPath:     in.GetProfilePhotoPath(),
		ProfilePhotoPathType: in.GetProfilePhotoPathType(),
		Sex:                  in.GetSex(),
		BirthRegencyId:       in.GetBirthRegencyId(),
		BloodType:            in.GetBloodType(),
		Height:               in.GetHeight(),
		Weight:               in.GetWeight(),
		IsColorBlind:         in.GetIsColorBlind(),
		UseGlasses:           in.GetUseGlasses(),
		HasCompleteTeeth:     in.GetHasCompleteTeeth(),
		IdNumber:             in.GetIdNumber(),
		NpwpNumber:           in.GetNpwpNumber(),
		NisnNumber:           in.GetNisnNumber(),
		Religion:             in.GetReligion(),
		MaritalStatus:        in.GetMaritalStatus(),
		Nationality:          in.GetNationality(),
		VillageId:            in.GetVillageId(),
		Rt:                   in.GetRt(),
		Rw:                   in.GetRw(),
		PostalCode:           in.GetPostalCode(),
		Address:              in.GetAddress(),
		PhoneNumber:          in.GetPhoneNumber(),
		MobilePhoneNumber:    in.GetMobilePhoneNumber(),
		Email:                in.GetEmail(),
		TransportationMean:   in.GetTransportationMean(),
		IsKpsRecipient:       in.GetIsKpsRecipient(),
		FundSource:           in.GetFundSource(),
		IsScholarshipGrantee: in.GetIsScholarshipGrantee(),
		TotalBrother:         in.GetTotalBrother(),
		TotalSister:          in.GetTotalSister(),
		WorkType:             in.GetWorkType(),
		WorkPlace:            in.GetWorkPlace(),
		WorkAddress:          in.GetWorkAddress(),
		AssuranceNumber:      in.GetAssuranceNumber(),
		Hobby:                in.GetHobby(),
	}
	errs := l.StudentService.UpdateProfile(ctx, data)
	if errs != nil {
		utils.PrintError(*errs)
		result = UpdateProfileResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = UpdateProfileResponse{
		Meta: &Meta{
			Message: "Update Profile Student",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}

	utils.JSONResponse(w, http.StatusOK, &result)
}

func (l studentGeneralHandler) GetParentProfile(w http.ResponseWriter, r *http.Request) {
	var result GetParentProfileResponse

	ctx := r.Context()

	defer l.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "General", "Get Parent Profile", nil)

	data, errs := l.StudentService.GetProfile(ctx)
	if errs != nil {
		utils.PrintError(*errs)
		result = GetParentProfileResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = GetParentProfileResponse{
		Meta: &Meta{
			Message: "Get Parent Profile Student",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &GetParentProfileResponseData{
			FatherIdNumber:                  utils.NullStringScan(data.FatherIdNumber),
			FatherName:                      utils.NullStringScan(data.FatherName),
			FatherBirthDate:                 utils.SafetyDate(data.FatherBirthDate),
			FatherDeathDate:                 utils.SafetyDate(data.FatherDeathDate),
			FatherFinalAcademicBackground:   utils.NullStringScan(data.FatherFinalAcademicBackground),
			FatherOccupation:                utils.NullStringScan(data.FatherOccupation),
			MotherIdNumber:                  utils.NullStringScan(data.MotherIdNumber),
			MotherName:                      utils.NullStringScan(data.MotherName),
			MotherBirthDate:                 utils.SafetyDate(data.MotherBirthDate),
			MotherDeathDate:                 utils.SafetyDate(data.MotherDeathDate),
			MotherFinalAcademicBackground:   utils.NullStringScan(data.MotherFinalAcademicBackground),
			MotherOccupation:                utils.NullStringScan(data.MotherOccupation),
			ParentReligion:                  utils.NullStringScan(data.ParentReligion),
			ParentNationality:               utils.NullStringScan(data.ParentNationality),
			ParentAddress:                   utils.NullStringScan(data.ParentAddress),
			FatherWorkAddress:               utils.NullStringScan(data.FatherWorkAddress),
			ParentProvinceId:                utils.NullUint32Scan(data.ParentProvinceId),
			ParentProvinceName:              utils.NullStringScan(data.ParentProvinceName),
			ParentRegencyId:                 utils.NullUint32Scan(data.ParentRegencyId),
			ParentRegencyName:               utils.NullStringScan(data.ParentRegencyName),
			ParentPostalCode:                utils.NullStringScan(data.ParentPostalCode),
			ParentPhoneNumber:               utils.NullStringScan(data.ParentPhoneNumber),
			ParentEmail:                     utils.NullStringScan(data.ParentEmail),
			IsFinanciallyCapable:            utils.NullBooleanScan(data.IsFinanciallyCapable),
			ParentIncome:                    utils.NullFloatScan(data.ParentIncome),
			TotalDependent:                  utils.NullUint32Scan(data.TotalDependent),
			GuardianName:                    utils.NullStringScan(data.GuardianName),
			GuardianBirthDate:               utils.SafetyDate(data.GuardianBirthDate),
			GuardianDeathDate:               utils.SafetyDate(data.GuardianDeathDate),
			GuardianAddress:                 utils.NullStringScan(data.GuardianAddress),
			GuardianProvinceId:              utils.NullUint32Scan(data.GuardianProvinceId),
			GuardianProvinceName:            utils.NullStringScan(data.GuardianProvinceName),
			GuardianRegencyId:               utils.NullUint32Scan(data.GuardianRegencyId),
			GuardianRegencyName:             utils.NullStringScan(data.GuardianRegencyName),
			GuardianPostalCode:              utils.NullStringScan(data.GuardianPostalCode),
			GuardianPhoneNumber:             utils.NullStringScan(data.GuardianPhoneNumber),
			GuardianEmail:                   utils.NullStringScan(data.GuardianEmail),
			GuardianFinalAcademicBackground: utils.NullStringScan(data.GuardianFinalAcademicBackground),
			GuardianOccupation:              utils.NullStringScan(data.GuardianOccupation),
		},
	}

	utils.JSONResponse(w, http.StatusOK, &result)
}

func (l studentGeneralHandler) UpdateParentProfile(w http.ResponseWriter, r *http.Request) {
	var result UpdateParentProfileResponse

	ctx := r.Context()
	var in UpdateParentProfileRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = UpdateParentProfileResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}

	defer l.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "General", "Update Parent Profile", nil)

	data := objects.UpdateStudentParentProfile{
		FatherIdNumber:                  in.GetFatherIdNumber(),
		FatherName:                      in.GetFatherName(),
		FatherBirthDate:                 in.GetFatherBirthDate(),
		FatherDeathDate:                 in.GetFatherDeathDate(),
		FatherFinalAcademicBackground:   in.GetFatherFinalAcademicBackground(),
		FatherOccupation:                in.GetFatherOccupation(),
		MotherIdNumber:                  in.GetMotherIdNumber(),
		MotherName:                      in.GetMotherName(),
		MotherBirthDate:                 in.GetMotherBirthDate(),
		MotherDeathDate:                 in.GetMotherDeathDate(),
		MotherFinalAcademicBackground:   in.GetMotherFinalAcademicBackground(),
		MotherOccupation:                in.GetMotherOccupation(),
		ParentReligion:                  in.GetParentReligion(),
		ParentNationality:               in.GetParentNationality(),
		ParentAddress:                   in.GetParentAddress(),
		FatherWorkAddress:               in.GetFatherWorkAddress(),
		ParentRegencyId:                 in.GetParentRegencyId(),
		ParentPostalCode:                in.GetParentPostalCode(),
		ParentPhoneNumber:               in.GetParentPhoneNumber(),
		ParentEmail:                     in.GetParentEmail(),
		IsFinanciallyCapable:            in.GetIsFinanciallyCapable(),
		ParentIncome:                    in.GetParentIncome(),
		TotalDependent:                  in.GetTotalDependent(),
		GuardianName:                    in.GetGuardianName(),
		GuardianBirthDate:               in.GetGuardianBirthDate(),
		GuardianDeathDate:               in.GetGuardianDeathDate(),
		GuardianAddress:                 in.GetGuardianAddress(),
		GuardianRegencyId:               in.GetGuardianRegencyId(),
		GuardianPostalCode:              in.GetGuardianPostalCode(),
		GuardianPhoneNumber:             in.GetGuardianPhoneNumber(),
		GuardianEmail:                   in.GetGuardianEmail(),
		GuardianFinalAcademicBackground: in.GetGuardianFinalAcademicBackground(),
		GuardianOccupation:              in.GetGuardianOccupation(),
	}
	errs := l.StudentService.UpdateParentProfile(ctx, data)
	if errs != nil {
		utils.PrintError(*errs)
		result = UpdateParentProfileResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = UpdateParentProfileResponse{
		Meta: &Meta{
			Message: "Update Parent Profile Student",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}

	utils.JSONResponse(w, http.StatusOK, &result)
}

func (l studentGeneralHandler) GetSchoolProfile(w http.ResponseWriter, r *http.Request) {
	var result GetSchoolProfileResponse

	ctx := r.Context()

	defer l.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "General", "Get School Profile", nil)

	data, errs := l.StudentService.GetProfile(ctx)
	if errs != nil {
		utils.PrintError(*errs)
		result = GetSchoolProfileResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	preHighSchoolHistories := []*GetSchoolProfileResponseDataPreHighSchoolHistory{}
	for _, v := range data.PreHighSchoolHistories {
		preHighSchoolHistories = append(preHighSchoolHistories, &GetSchoolProfileResponseDataPreHighSchoolHistory{
			Id:             v.Id,
			Level:          v.Level,
			Name:           v.Name,
			GraduationYear: v.GraduationYear,
		})
	}

	result = GetSchoolProfileResponse{
		Meta: &Meta{
			Message: "Get School Profile Student",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &GetSchoolProfileResponseData{
			SchoolEnrollmentYear:            utils.NullStringScan(data.SchoolEnrollmentYear),
			SchoolGraduationYear:            utils.NullStringScan(data.SchoolGraduationYear),
			SchoolEnrollmentClass:           utils.NullStringScan(data.SchoolEnrollmentClass),
			SchoolMajor:                     utils.NullStringScan(data.SchoolMajor),
			SchoolType:                      utils.NullStringScan(data.SchoolType),
			SchoolName:                      utils.NullStringScan(data.SchoolName),
			SchoolProvinceId:                utils.NullUint32Scan(data.SchoolProvinceId),
			SchoolProvinceName:              utils.NullStringScan(data.SchoolProvinceName),
			SchoolAddress:                   utils.NullStringScan(data.SchoolAddress),
			SchoolCertificateNumber:         utils.NullStringScan(data.SchoolCertificateNumber),
			SchoolCertificateDate:           utils.SafetyDate(data.SchoolCertificateDate),
			SchoolStatus:                    utils.NullStringScan(data.SchoolStatus),
			SchoolAccreditation:             utils.NullStringScan(data.SchoolAccreditation),
			SchoolFinalExamScore:            utils.NullFloatScan(data.SchoolFinalExamScore),
			SchoolMathematicsFinalExamScore: utils.NullFloatScan(data.SchoolMathematicsFinalExamScore),
			SchoolIndonesianFinalExamScore:  utils.NullFloatScan(data.SchoolIndonesianFinalExamScore),
			SchoolEnglishFinalExamScore:     utils.NullFloatScan(data.SchoolEnglishFinalExamScore),
			SchoolMathematicsReportScore:    utils.NullFloatScan(data.SchoolMathematicsReportScore),
			SchoolIndonesianReportScore:     utils.NullFloatScan(data.SchoolIndonesianReportScore),
			SchoolEnglishReportScore:        utils.NullFloatScan(data.SchoolEnglishReportScore),
			PreHighSchoolHistories:          preHighSchoolHistories,
		},
	}

	utils.JSONResponse(w, http.StatusOK, &result)
}

func (l studentGeneralHandler) UpdateSchoolProfile(w http.ResponseWriter, r *http.Request) {
	var result UpdateSchoolProfileResponse

	ctx := r.Context()
	var in UpdateSchoolProfileRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = UpdateSchoolProfileResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}

	defer l.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "General", "Update School Profile", nil)

	preHighSchoolHistoryData := []objects.UpdateStudentSchoolProfilePreHighSchoolHistory{}
	for _, v := range in.GetPreHighSchoolHistories() {
		preHighSchoolHistoryData = append(preHighSchoolHistoryData, objects.UpdateStudentSchoolProfilePreHighSchoolHistory{
			Level:          v.GetLevel(),
			Name:           v.GetName(),
			GraduationYear: v.GetGraduationYear(),
		})
	}

	data := objects.UpdateStudentSchoolProfile{
		SchoolEnrollmentYear:            in.GetSchoolEnrollmentYear(),
		SchoolGraduationYear:            in.GetSchoolGraduationYear(),
		SchoolEnrollmentClass:           in.GetSchoolEnrollmentClass(),
		SchoolMajor:                     in.GetSchoolMajor(),
		SchoolType:                      in.GetSchoolType(),
		SchoolName:                      in.GetSchoolName(),
		SchoolProvinceId:                in.GetSchoolProvinceId(),
		SchoolAddress:                   in.GetSchoolAddress(),
		SchoolCertificateNumber:         in.GetSchoolCertificateNumber(),
		SchoolCertificateDate:           in.GetSchoolCertificateDate(),
		SchoolStatus:                    in.GetSchoolStatus(),
		SchoolAccreditation:             in.GetSchoolAccreditation(),
		SchoolFinalExamScore:            in.GetSchoolFinalExamScore(),
		SchoolMathematicsFinalExamScore: in.GetSchoolMathematicsFinalExamScore(),
		SchoolIndonesianFinalExamScore:  in.GetSchoolIndonesianFinalExamScore(),
		SchoolEnglishFinalExamScore:     in.GetSchoolEnglishFinalExamScore(),
		SchoolMathematicsReportScore:    in.GetSchoolMathematicsReportScore(),
		SchoolIndonesianReportScore:     in.GetSchoolIndonesianReportScore(),
		SchoolEnglishReportScore:        in.GetSchoolEnglishReportScore(),
		PreHighSchoolHistories:          preHighSchoolHistoryData,
	}
	errs := l.StudentService.UpdateSchoolProfile(ctx, data)
	if errs != nil {
		utils.PrintError(*errs)
		result = UpdateSchoolProfileResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = UpdateSchoolProfileResponse{
		Meta: &Meta{
			Message: "Update School Profile Student",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}

	utils.JSONResponse(w, http.StatusOK, &result)
}
