package student

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

type studentHandler struct {
	*service.ServiceCtx
}

func (a studentHandler) GetList(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Student", "Get List", nil)

	paginationData := common.PaginationRequest{
		Limit: in.GetLimit(),
		Page:  in.GetPage(),
	}

	hasAuthentication, errs := utils.StringToBoolPointer(in.GetHasAuthentication())
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

	studyPlanIsSubmitted, errs := utils.StringToBoolPointer(in.GetStudyPlanIsSubmitted())
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
	studyPlanIsApproved, errs := utils.StringToBoolPointer(in.GetStudyPlanIsApproved())
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
	hasStudyPlan, errs := utils.StringToBoolPointer(in.GetHasStudyPlan())
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
	isRegistered, errs := utils.StringToBoolPointer(in.GetIsRegistered())
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
	hasPaid, errs := utils.StringToBoolPointer(in.GetHasPaid())
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
	isGraduateEligible, errs := utils.StringToBoolPointer(in.GetIsGraduateEligible())
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
	isThesisEligible, errs := utils.StringToBoolPointer(in.GetIsThesisEligible())
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

	requestData := objects.GetStudentRequest{
		StudyProgramId:       in.GetStudyProgramId(),
		StudentForceFrom:     in.GetStudentForceFrom(),
		StudentForceTo:       in.GetStudentForceTo(),
		NimNumberFrom:        in.GetNimNumberFrom(),
		NimNumberTo:          in.GetNimNumberTo(),
		Name:                 in.GetName(),
		Address:              in.GetAddress(),
		RegencyId:            in.GetRegencyId(),
		Status:               in.GetStatus(),
		GetAcademicGuidance:  in.GetGetAcademicGuidance(),
		HasAuthentication:    hasAuthentication,
		StudyPlanSemesterId:  in.GetStudyPlanSemesterId(),
		StudyPlanIsSubmitted: studyPlanIsSubmitted,
		StudyPlanIsApproved:  studyPlanIsApproved,
		HasStudyPlan:         hasStudyPlan,
		StatusSemesterId:     in.GetStatusSemesterId(),
		IsRegistered:         isRegistered,
		HasPaid:              hasPaid,
		IsGraduateEligible:   isGraduateEligible,
		IsThesisEligible:     isThesisEligible,
		YudiciumSessionId:    in.GetYudiciumSessionId(),
	}

	data, errs := a.StudentService.GetList(ctx, paginationData, requestData)
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
			Id:                                 v.Id,
			Name:                               v.Name,
			Sex:                                utils.NullStringScan(v.Sex),
			MaritalStatus:                      utils.NullStringScan(v.MaritalStatus),
			BirthRegencyId:                     utils.NullUint32Scan(v.BirthRegencyId),
			BirthRegencyName:                   utils.NullStringScan(v.BirthRegencyName),
			BirthDate:                          utils.SafetyDate(v.BirthDate),
			Religion:                           utils.NullStringScan(v.Religion),
			Address:                            utils.NullStringScan(v.Address),
			Rt:                                 utils.NullStringScan(v.Rt),
			Rw:                                 utils.NullStringScan(v.Rw),
			VillageId:                          utils.NullUint32Scan(v.VillageId),
			VillageName:                        utils.NullStringScan(v.VillageName),
			DistrictId:                         utils.NullUint32Scan(v.DistrictId),
			DistrictName:                       utils.NullStringScan(v.DistrictName),
			RegencyId:                          utils.NullUint32Scan(v.RegencyId),
			RegencyName:                        utils.NullStringScan(v.RegencyName),
			ProvinceId:                         utils.NullUint32Scan(v.ProvinceId),
			ProvinceName:                       utils.NullStringScan(v.ProvinceName),
			CountryId:                          utils.NullUint32Scan(v.CountryId),
			CountryName:                        utils.NullStringScan(v.CountryName),
			PostalCode:                         utils.NullStringScan(v.PostalCode),
			PreviousAddress:                    utils.NullStringScan(v.PreviousAddress),
			IdNumber:                           utils.NullStringScan(v.IdNumber),
			NpwpNumber:                         utils.NullStringScan(v.NpwpNumber),
			NisnNumber:                         utils.NullStringScan(v.NisnNumber),
			ResidenceType:                      utils.NullStringScan(v.ResidenceType),
			TransportationMean:                 utils.NullStringScan(v.TransportationMean),
			KpsReceiver:                        utils.NullStringScan(v.KpsReceiver),
			PhoneNumber:                        utils.NullStringScan(v.PhoneNumber),
			MobilePhoneNumber:                  utils.NullStringScan(v.MobilePhoneNumber),
			Email:                              utils.NullStringScan(v.Email),
			Homepage:                           utils.NullStringScan(v.Homepage),
			WorkType:                           utils.NullStringScan(v.WorkType),
			WorkPlace:                          utils.NullStringScan(v.WorkPlace),
			Nationality:                        utils.NullStringScan(v.Nationality),
			AskesNumber:                        utils.NullStringScan(v.AskesNumber),
			TotalBrother:                       utils.NullUint32Scan(v.TotalBrother),
			TotalSister:                        utils.NullUint32Scan(v.TotalSister),
			Hobby:                              utils.NullStringScan(v.Hobby),
			Experience:                         utils.NullStringScan(v.Experience),
			TotalDependent:                     utils.NullUint32Scan(v.TotalDependent),
			NimNumber:                          v.NimNumber,
			StudentForce:                       utils.NullUint32Scan(v.StudentForce),
			AdmittanceSemester:                 utils.NullStringScan(v.AdmittanceSemester),
			StudyProgramId:                     utils.NullStringScan(v.StudyProgramId),
			StudyProgramName:                   utils.NullStringScan(v.StudyProgramName),
			CurriculumId:                       utils.NullStringScan(v.CurriculumId),
			CurriculumName:                     utils.NullStringScan(v.CurriculumName),
			AdmittanceTestNumber:               utils.NullStringScan(v.AdmittanceTestNumber),
			AdmittanceDate:                     utils.SafetyDate(v.AdmittanceDate),
			AdmittanceStatus:                   utils.NullStringScan(v.AdmittanceStatus),
			TotalTransferCredit:                utils.NullUint32Scan(v.TotalTransferCredit),
			PreviousCollege:                    utils.NullStringScan(v.PreviousCollege),
			PreviousStudyProgram:               utils.NullStringScan(v.PreviousStudyProgram),
			PreviousNimNumber:                  utils.NullInt64Scan(v.PreviousNimNumber),
			PreviousNimAdmittanceYear:          utils.NullStringScan(v.PreviousNimAdmittanceYear),
			Status:                             utils.NullStringScan(v.Status),
			IsForeignStudent:                   utils.NullBooleanScan(v.IsForeignStudent),
			CollegeEntranceType:                utils.NullStringScan(v.CollegeEntranceType),
			ClassTime:                          utils.NullStringScan(v.ClassTime),
			FundSource:                         utils.NullStringScan(v.FundSource),
			IsScholarshipGrantee:               utils.NullBooleanScan(v.IsScholarshipGrantee),
			HasCompleteRequirement:             utils.NullBooleanScan(v.HasCompleteRequirement),
			SchoolCertificateType:              utils.NullStringScan(v.SchoolCertificateType),
			SchoolGraduationYear:               utils.NullStringScan(v.SchoolGraduationYear),
			SchoolName:                         utils.NullStringScan(v.SchoolName),
			SchoolAccreditation:                utils.NullStringScan(v.SchoolAccreditation),
			SchoolAddress:                      utils.NullStringScan(v.SchoolAddress),
			SchoolMajor:                        utils.NullStringScan(v.SchoolMajor),
			SchoolCertificateNumber:            utils.NullStringScan(v.SchoolCertificateNumber),
			SchoolCertificateDate:              utils.SafetyDate(v.SchoolCertificateDate),
			TotalSchoolFinalExamSubject:        utils.NullUint32Scan(v.TotalSchoolFinalExamSubject),
			SchoolFinalExamScore:               utils.NullFloatScan(v.SchoolFinalExamScore),
			GuardianName:                       utils.NullStringScan(v.GuardianName),
			GuardianBirthDate:                  utils.SafetyDate(v.GuardianBirthDate),
			GuardianDeathDate:                  utils.SafetyDate(v.GuardianDeathDate),
			GuardianAddress:                    utils.NullStringScan(v.GuardianAddress),
			GuardianRegencyId:                  utils.NullUint32Scan(v.GuardianRegencyId),
			GuardianRegencyName:                utils.NullStringScan(v.GuardianRegencyName),
			GuardianPostalCode:                 utils.NullStringScan(v.GuardianPostalCode),
			GuardianPhoneNumber:                utils.NullStringScan(v.GuardianPhoneNumber),
			GuardianEmail:                      utils.NullStringScan(v.GuardianEmail),
			GuardianFinalAcademicBackground:    utils.NullStringScan(v.GuardianFinalAcademicBackground),
			GuardianOccupation:                 utils.NullStringScan(v.GuardianOccupation),
			FatherIdNumber:                     utils.NullStringScan(v.FatherIdNumber),
			FatherName:                         utils.NullStringScan(v.FatherName),
			FatherBirthDate:                    utils.SafetyDate(v.FatherBirthDate),
			FatherDeathDate:                    utils.SafetyDate(v.FatherDeathDate),
			MotherIdNumber:                     utils.NullStringScan(v.MotherIdNumber),
			MotherName:                         utils.NullStringScan(v.MotherName),
			MotherBirthDate:                    utils.SafetyDate(v.MotherBirthDate),
			MotherDeathDate:                    utils.SafetyDate(v.MotherDeathDate),
			ParentAddress:                      utils.NullStringScan(v.ParentAddress),
			ParentRegencyId:                    utils.NullUint32Scan(v.ParentRegencyId),
			ParentRegencyName:                  utils.NullStringScan(v.ParentRegencyName),
			ParentPostalCode:                   utils.NullStringScan(v.ParentPostalCode),
			ParentPhoneNumber:                  utils.NullStringScan(v.ParentPhoneNumber),
			ParentEmail:                        utils.NullStringScan(v.ParentEmail),
			FatherFinalAcademicBackground:      utils.NullStringScan(v.FatherFinalAcademicBackground),
			FatherOccupation:                   utils.NullStringScan(v.FatherOccupation),
			MotherFinalAcademicBackground:      utils.NullStringScan(v.MotherFinalAcademicBackground),
			MotherOccupation:                   utils.NullStringScan(v.MotherOccupation),
			ParentIncome:                       utils.NullFloatScan(v.ParentIncome),
			IsFinanciallyCapable:               utils.NullBooleanScan(v.IsFinanciallyCapable),
			AuthenticationId:                   utils.NullStringScan(v.AuthenticationId),
			AuthenticationIsActive:             utils.NullBooleanScan(v.AuthenticationIsActive),
			AuthenticationSuspensionRemarks:    utils.NullStringScan(v.AuthenticationSuspensionRemarks),
			DiktiStudyProgramCode:              utils.NullStringScan(v.DiktiStudyProgramCode),
			AcademicGuidanceLecturerId:         utils.NullStringScan(v.AcademicGuidanceLecturerId),
			AcademicGuidanceLecturerName:       utils.NullStringScan(v.AcademicGuidanceLecturerName),
			AcademicGuidanceSemesterId:         utils.NullStringScan(v.AcademicGuidanceSemesterId),
			AcademicGuidanceSemesterSchoolYear: utils.NullStringScan(v.AcademicGuidanceSemesterSchoolYear),
			StudyPlanTotalMandatoryCredit:      utils.NullUint32Scan(v.StudyPlanTotalMandatoryCredit),
			StudyPlanTotalOptionalCredit:       utils.NullUint32Scan(v.StudyPlanTotalOptionalCredit),
			StudyPlanMaximumCredit:             utils.NullUint32Scan(v.StudyPlanMaximumCredit),
			StudyPlanIsApproved:                utils.NullBooleanScan(v.StudyPlanIsApproved),
			StudyPlanId:                        utils.NullStringScan(v.StudyPlanId),
			TotalStudyPlan:                     v.TotalStudyPlan,
			CurrentSemesterPackage:             v.CurrentSemesterPackage,
			StatusSemesterId:                   utils.NullStringScan(v.StatusSemesterId),
			StatusSemesterSchoolYear:           v.StatusSemesterSchoolYear,
			StatusSemesterType:                 utils.NullStringScan(v.StatusSemesterType),
			StatusDate:                         utils.SafetyDate(v.StatusDate),
			StatusReferenceNumber:              utils.NullStringScan(v.StatusReferenceNumber),
			StatusPurpose:                      utils.NullStringScan(v.StatusPurpose),
			StatusRemarks:                      utils.NullStringScan(v.StatusRemarks),
			Gpa:                                utils.NullFloatScan(v.Gpa),
			TotalCredit:                        utils.NullUint32Scan(v.TotalCredit),
			TranscriptIsArchived:               v.TranscriptIsArchived,
			HasPaid:                            v.HasPaid,
			StudyDurationMonth:                 v.StudyDurationMonth,
			ThesisDurationMonth:                v.ThesisDurationMonth,
			ThesisDurationSemester:             v.ThesisDurationSemester,
			GraduationPredicate:                utils.NullStringScan(v.GraduationPredicate),
		})
	}

	result = GetListResponse{
		Meta: &Meta{
			Message: "Get List Student",
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

func (a studentHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Student", "Get Detail", nil)

	data, errs := a.StudentService.GetDetail(ctx, in.GetId())
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

	result = GetDetailResponse{
		Meta: &Meta{
			Message: "Get Detail Student",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &GetDetailResponseData{
			Id:                                 data.Id,
			Name:                               data.Name,
			Sex:                                utils.NullStringScan(data.Sex),
			MaritalStatus:                      utils.NullStringScan(data.MaritalStatus),
			BirthRegencyId:                     utils.NullUint32Scan(data.BirthRegencyId),
			BirthRegencyName:                   utils.NullStringScan(data.BirthRegencyName),
			BirthDate:                          utils.SafetyDate(data.BirthDate),
			Religion:                           utils.NullStringScan(data.Religion),
			Address:                            utils.NullStringScan(data.Address),
			Rt:                                 utils.NullStringScan(data.Rt),
			Rw:                                 utils.NullStringScan(data.Rw),
			VillageId:                          utils.NullUint32Scan(data.VillageId),
			VillageName:                        utils.NullStringScan(data.VillageName),
			DistrictId:                         utils.NullUint32Scan(data.DistrictId),
			DistrictName:                       utils.NullStringScan(data.DistrictName),
			RegencyId:                          utils.NullUint32Scan(data.RegencyId),
			RegencyName:                        utils.NullStringScan(data.RegencyName),
			ProvinceId:                         utils.NullUint32Scan(data.ProvinceId),
			ProvinceName:                       utils.NullStringScan(data.ProvinceName),
			CountryId:                          utils.NullUint32Scan(data.CountryId),
			CountryName:                        utils.NullStringScan(data.CountryName),
			PostalCode:                         utils.NullStringScan(data.PostalCode),
			PreviousAddress:                    utils.NullStringScan(data.PreviousAddress),
			IdNumber:                           utils.NullStringScan(data.IdNumber),
			NpwpNumber:                         utils.NullStringScan(data.NpwpNumber),
			NisnNumber:                         utils.NullStringScan(data.NisnNumber),
			ResidenceType:                      utils.NullStringScan(data.ResidenceType),
			TransportationMean:                 utils.NullStringScan(data.TransportationMean),
			KpsReceiver:                        utils.NullStringScan(data.KpsReceiver),
			PhoneNumber:                        utils.NullStringScan(data.PhoneNumber),
			MobilePhoneNumber:                  utils.NullStringScan(data.MobilePhoneNumber),
			Email:                              utils.NullStringScan(data.Email),
			Homepage:                           utils.NullStringScan(data.Homepage),
			WorkType:                           utils.NullStringScan(data.WorkType),
			WorkPlace:                          utils.NullStringScan(data.WorkPlace),
			Nationality:                        utils.NullStringScan(data.Nationality),
			AskesNumber:                        utils.NullStringScan(data.AskesNumber),
			TotalBrother:                       utils.NullUint32Scan(data.TotalBrother),
			TotalSister:                        utils.NullUint32Scan(data.TotalSister),
			Hobby:                              utils.NullStringScan(data.Hobby),
			Experience:                         utils.NullStringScan(data.Experience),
			TotalDependent:                     utils.NullUint32Scan(data.TotalDependent),
			NimNumber:                          data.NimNumber,
			StudentForce:                       utils.NullUint32Scan(data.StudentForce),
			AdmittanceSemester:                 utils.NullStringScan(data.AdmittanceSemester),
			StudyProgramId:                     utils.NullStringScan(data.StudyProgramId),
			StudyProgramName:                   utils.NullStringScan(data.StudyProgramName),
			CurriculumId:                       utils.NullStringScan(data.CurriculumId),
			CurriculumName:                     utils.NullStringScan(data.CurriculumName),
			AdmittanceTestNumber:               utils.NullStringScan(data.AdmittanceTestNumber),
			AdmittanceDate:                     utils.SafetyDate(data.AdmittanceDate),
			AdmittanceStatus:                   utils.NullStringScan(data.AdmittanceStatus),
			TotalTransferCredit:                utils.NullUint32Scan(data.TotalTransferCredit),
			PreviousCollege:                    utils.NullStringScan(data.PreviousCollege),
			PreviousStudyProgram:               utils.NullStringScan(data.PreviousStudyProgram),
			PreviousNimNumber:                  utils.NullInt64Scan(data.PreviousNimNumber),
			PreviousNimAdmittanceYear:          utils.NullStringScan(data.PreviousNimAdmittanceYear),
			Status:                             utils.NullStringScan(data.Status),
			IsForeignStudent:                   utils.NullBooleanScan(data.IsForeignStudent),
			CollegeEntranceType:                utils.NullStringScan(data.CollegeEntranceType),
			ClassTime:                          utils.NullStringScan(data.ClassTime),
			FundSource:                         utils.NullStringScan(data.FundSource),
			IsScholarshipGrantee:               utils.NullBooleanScan(data.IsScholarshipGrantee),
			HasCompleteRequirement:             utils.NullBooleanScan(data.HasCompleteRequirement),
			SchoolCertificateType:              utils.NullStringScan(data.SchoolCertificateType),
			SchoolGraduationYear:               utils.NullStringScan(data.SchoolGraduationYear),
			SchoolName:                         utils.NullStringScan(data.SchoolName),
			SchoolAccreditation:                utils.NullStringScan(data.SchoolAccreditation),
			SchoolAddress:                      utils.NullStringScan(data.SchoolAddress),
			SchoolMajor:                        utils.NullStringScan(data.SchoolMajor),
			SchoolCertificateNumber:            utils.NullStringScan(data.SchoolCertificateNumber),
			SchoolCertificateDate:              utils.SafetyDate(data.SchoolCertificateDate),
			TotalSchoolFinalExamSubject:        utils.NullUint32Scan(data.TotalSchoolFinalExamSubject),
			SchoolFinalExamScore:               utils.NullFloatScan(data.SchoolFinalExamScore),
			GuardianName:                       utils.NullStringScan(data.GuardianName),
			GuardianBirthDate:                  utils.SafetyDate(data.GuardianBirthDate),
			GuardianDeathDate:                  utils.SafetyDate(data.GuardianDeathDate),
			GuardianAddress:                    utils.NullStringScan(data.GuardianAddress),
			GuardianRegencyId:                  utils.NullUint32Scan(data.GuardianRegencyId),
			GuardianRegencyName:                utils.NullStringScan(data.GuardianRegencyName),
			GuardianPostalCode:                 utils.NullStringScan(data.GuardianPostalCode),
			GuardianPhoneNumber:                utils.NullStringScan(data.GuardianPhoneNumber),
			GuardianEmail:                      utils.NullStringScan(data.GuardianEmail),
			GuardianFinalAcademicBackground:    utils.NullStringScan(data.GuardianFinalAcademicBackground),
			GuardianOccupation:                 utils.NullStringScan(data.GuardianOccupation),
			FatherIdNumber:                     utils.NullStringScan(data.FatherIdNumber),
			FatherName:                         utils.NullStringScan(data.FatherName),
			FatherBirthDate:                    utils.SafetyDate(data.FatherBirthDate),
			FatherDeathDate:                    utils.SafetyDate(data.FatherDeathDate),
			MotherIdNumber:                     utils.NullStringScan(data.MotherIdNumber),
			MotherName:                         utils.NullStringScan(data.MotherName),
			MotherBirthDate:                    utils.SafetyDate(data.MotherBirthDate),
			MotherDeathDate:                    utils.SafetyDate(data.MotherDeathDate),
			ParentAddress:                      utils.NullStringScan(data.ParentAddress),
			ParentRegencyId:                    utils.NullUint32Scan(data.ParentRegencyId),
			ParentRegencyName:                  utils.NullStringScan(data.ParentRegencyName),
			ParentPostalCode:                   utils.NullStringScan(data.ParentPostalCode),
			ParentPhoneNumber:                  utils.NullStringScan(data.ParentPhoneNumber),
			ParentEmail:                        utils.NullStringScan(data.ParentEmail),
			FatherFinalAcademicBackground:      utils.NullStringScan(data.FatherFinalAcademicBackground),
			FatherOccupation:                   utils.NullStringScan(data.FatherOccupation),
			MotherFinalAcademicBackground:      utils.NullStringScan(data.MotherFinalAcademicBackground),
			MotherOccupation:                   utils.NullStringScan(data.MotherOccupation),
			ParentIncome:                       utils.NullFloatScan(data.ParentIncome),
			IsFinanciallyCapable:               utils.NullBooleanScan(data.IsFinanciallyCapable),
			AuthenticationId:                   utils.NullStringScan(data.AuthenticationId),
			AuthenticationIsActive:             utils.NullBooleanScan(data.AuthenticationIsActive),
			AuthenticationSuspensionRemarks:    utils.NullStringScan(data.AuthenticationSuspensionRemarks),
			DiktiStudyProgramCode:              utils.NullStringScan(data.DiktiStudyProgramCode),
			AcademicGuidanceLecturerId:         utils.NullStringScan(data.AcademicGuidanceLecturerId),
			AcademicGuidanceLecturerName:       utils.NullStringScan(data.AcademicGuidanceLecturerName),
			AcademicGuidanceSemesterId:         utils.NullStringScan(data.AcademicGuidanceSemesterId),
			AcademicGuidanceSemesterSchoolYear: utils.NullStringScan(data.AcademicGuidanceSemesterSchoolYear),
			StudyPlanTotalMandatoryCredit:      utils.NullUint32Scan(data.StudyPlanTotalMandatoryCredit),
			StudyPlanTotalOptionalCredit:       utils.NullUint32Scan(data.StudyPlanTotalOptionalCredit),
			StudyPlanMaximumCredit:             utils.NullUint32Scan(data.StudyPlanMaximumCredit),
			StudyPlanIsApproved:                utils.NullBooleanScan(data.StudyPlanIsApproved),
			StudyPlanId:                        utils.NullStringScan(data.StudyPlanId),
			CurrentSemesterPackage:             data.CurrentSemesterPackage,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a studentHandler) Create(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Student", "Create", &in)

	var birthDate time.Time
	var admittanceDate time.Time
	var schoolCertificateDate time.Time
	var guardianBirthDate time.Time
	var guardianDeathDate time.Time
	var fatherBirthDate time.Time
	var fatherDeathDate time.Time
	var motherBirthDate time.Time
	var motherDeathDate time.Time
	var errs *constants.ErrorResponse
	if in.GetBirthDate() != "" {
		birthDate, errs = utils.StringToTime(in.GetBirthDate())
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
	}
	if in.GetAdmittanceDate() != "" {
		admittanceDate, errs = utils.StringToTime(in.GetAdmittanceDate())
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
	}
	if in.GetSchoolCertificateDate() != "" {
		schoolCertificateDate, errs = utils.StringToTime(in.GetSchoolCertificateDate())
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
	}
	if in.GetGuardianBirthDate() != "" {
		guardianBirthDate, errs = utils.StringToTime(in.GetGuardianBirthDate())
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
	}
	if in.GetGuardianDeathDate() != "" {
		guardianDeathDate, errs = utils.StringToTime(in.GetGuardianDeathDate())
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
	}
	if in.GetFatherBirthDate() != "" {
		fatherBirthDate, errs = utils.StringToTime(in.GetFatherBirthDate())
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
	}
	if in.GetFatherDeathDate() != "" {
		fatherDeathDate, errs = utils.StringToTime(in.GetFatherDeathDate())
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
	}
	if in.GetMotherBirthDate() != "" {
		motherBirthDate, errs = utils.StringToTime(in.GetMotherBirthDate())
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
	}
	if in.GetMotherDeathDate() != "" {
		motherDeathDate, errs = utils.StringToTime(in.GetMotherDeathDate())
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
	}
	data := objects.CreateStudent{
		Name:                            in.GetName(),
		Sex:                             in.GetSex(),
		MaritalStatus:                   in.GetMaritalStatus(),
		BirthRegencyId:                  in.GetBirthRegencyId(),
		BirthDate:                       birthDate,
		Religion:                        in.GetReligion(),
		Address:                         in.GetAddress(),
		Rt:                              in.GetRt(),
		Rw:                              in.GetRw(),
		VillageId:                       in.GetVillageId(),
		PostalCode:                      in.GetPostalCode(),
		PreviousAddress:                 in.GetPreviousAddress(),
		IdNumber:                        in.GetIdNumber(),
		NpwpNumber:                      in.GetNpwpNumber(),
		NisnNumber:                      in.GetNisnNumber(),
		ResidenceType:                   in.GetResidenceType(),
		TransportationMean:              in.GetTransportationMean(),
		KpsReceiver:                     in.GetKpsReceiver(),
		PhoneNumber:                     in.GetPhoneNumber(),
		MobilePhoneNumber:               in.GetMobilePhoneNumber(),
		Email:                           in.GetEmail(),
		Homepage:                        in.GetHomepage(),
		WorkType:                        in.GetWorkType(),
		WorkPlace:                       in.GetWorkPlace(),
		Nationality:                     in.GetNationality(),
		AskesNumber:                     in.GetAskesNumber(),
		TotalBrother:                    in.GetTotalBrother(),
		TotalSister:                     in.GetTotalSister(),
		Hobby:                           in.GetHobby(),
		Experience:                      in.GetExperience(),
		TotalDependent:                  in.GetTotalDependent(),
		NimNumber:                       in.GetNimNumber(),
		StudentForce:                    in.GetStudentForce(),
		AdmittanceSemester:              in.GetAdmittanceSemester(),
		StudyProgramId:                  in.GetStudyProgramId(),
		CurriculumId:                    in.GetCurriculumId(),
		AdmittanceTestNumber:            in.GetAdmittanceTestNumber(),
		AdmittanceDate:                  admittanceDate,
		AdmittanceStatus:                in.GetAdmittanceStatus(),
		TotalTransferCredit:             in.GetTotalTransferCredit(),
		PreviousCollege:                 in.GetPreviousCollege(),
		PreviousStudyProgram:            in.GetPreviousStudyProgram(),
		PreviousNimNumber:               in.GetPreviousNimNumber(),
		PreviousNimAdmittanceYear:       in.GetPreviousNimAdmittanceYear(),
		Status:                          in.GetStatus(),
		IsForeignStudent:                in.GetIsForeignStudent(),
		CollegeEntranceType:             in.GetCollegeEntranceType(),
		ClassTime:                       in.GetClassTime(),
		FundSource:                      in.GetFundSource(),
		IsScholarshipGrantee:            in.GetIsScholarshipGrantee(),
		HasCompleteRequirement:          in.GetHasCompleteRequirement(),
		SchoolCertificateType:           in.GetSchoolCertificateType(),
		SchoolGraduationYear:            in.GetSchoolGraduationYear(),
		SchoolName:                      in.GetSchoolName(),
		SchoolAccreditation:             in.GetSchoolAccreditation(),
		SchoolAddress:                   in.GetSchoolAddress(),
		SchoolMajor:                     in.GetSchoolMajor(),
		SchoolCertificateNumber:         in.GetSchoolCertificateNumber(),
		SchoolCertificateDate:           schoolCertificateDate,
		TotalSchoolFinalExamSubject:     in.GetTotalSchoolFinalExamSubject(),
		SchoolFinalExamScore:            in.GetSchoolFinalExamScore(),
		GuardianName:                    in.GetGuardianName(),
		GuardianBirthDate:               guardianBirthDate,
		GuardianDeathDate:               guardianDeathDate,
		GuardianAddress:                 in.GetGuardianAddress(),
		GuardianRegencyId:               in.GetGuardianRegencyId(),
		GuardianPostalCode:              in.GetGuardianPostalCode(),
		GuardianPhoneNumber:             in.GetGuardianPhoneNumber(),
		GuardianEmail:                   in.GetGuardianEmail(),
		GuardianFinalAcademicBackground: in.GetGuardianFinalAcademicBackground(),
		GuardianOccupation:              in.GetGuardianOccupation(),
		FatherIdNumber:                  in.GetFatherIdNumber(),
		FatherName:                      in.GetFatherName(),
		FatherBirthDate:                 fatherBirthDate,
		FatherDeathDate:                 fatherDeathDate,
		MotherIdNumber:                  in.GetMotherIdNumber(),
		MotherName:                      in.GetMotherName(),
		MotherBirthDate:                 motherBirthDate,
		MotherDeathDate:                 motherDeathDate,
		ParentAddress:                   in.GetParentAddress(),
		ParentRegencyId:                 in.GetParentRegencyId(),
		ParentPostalCode:                in.GetParentPostalCode(),
		ParentPhoneNumber:               in.GetParentPhoneNumber(),
		ParentEmail:                     in.GetParentEmail(),
		FatherFinalAcademicBackground:   in.GetFatherFinalAcademicBackground(),
		FatherOccupation:                in.GetFatherOccupation(),
		MotherFinalAcademicBackground:   in.GetMotherFinalAcademicBackground(),
		MotherOccupation:                in.GetMotherOccupation(),
		ParentIncome:                    in.GetParentIncome(),
		IsFinanciallyCapable:            in.GetIsFinanciallyCapable(),
	}

	errs = a.StudentService.Create(ctx, data)
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
			Message: "Create Student",
			Status:  http.StatusCreated,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusCreated, &result)
}

func (a studentHandler) Update(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Student", "Update", &in)

	var birthDate time.Time
	var admittanceDate time.Time
	var schoolCertificateDate time.Time
	var guardianBirthDate time.Time
	var guardianDeathDate time.Time
	var fatherBirthDate time.Time
	var fatherDeathDate time.Time
	var motherBirthDate time.Time
	var motherDeathDate time.Time
	var errs *constants.ErrorResponse
	if in.GetBirthDate() != "" {
		birthDate, errs = utils.StringToTime(in.GetBirthDate())
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
	}
	if in.GetAdmittanceDate() != "" {
		admittanceDate, errs = utils.StringToTime(in.GetAdmittanceDate())
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
	}
	if in.GetSchoolCertificateDate() != "" {
		schoolCertificateDate, errs = utils.StringToTime(in.GetSchoolCertificateDate())
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
	}
	if in.GetGuardianBirthDate() != "" {
		guardianBirthDate, errs = utils.StringToTime(in.GetGuardianBirthDate())
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
	}
	if in.GetGuardianDeathDate() != "" {
		guardianDeathDate, errs = utils.StringToTime(in.GetGuardianDeathDate())
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
	}
	if in.GetFatherBirthDate() != "" {
		fatherBirthDate, errs = utils.StringToTime(in.GetFatherBirthDate())
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
	}
	if in.GetFatherDeathDate() != "" {
		fatherDeathDate, errs = utils.StringToTime(in.GetFatherDeathDate())
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
	}
	if in.GetMotherBirthDate() != "" {
		motherBirthDate, errs = utils.StringToTime(in.GetMotherBirthDate())
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
	}
	if in.GetMotherDeathDate() != "" {
		motherDeathDate, errs = utils.StringToTime(in.GetMotherDeathDate())
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
	}
	data := objects.UpdateStudent{
		Id:                              in.GetId(),
		Name:                            in.GetName(),
		Sex:                             in.GetSex(),
		MaritalStatus:                   in.GetMaritalStatus(),
		BirthRegencyId:                  in.GetBirthRegencyId(),
		BirthDate:                       birthDate,
		Religion:                        in.GetReligion(),
		Address:                         in.GetAddress(),
		Rt:                              in.GetRt(),
		Rw:                              in.GetRw(),
		VillageId:                       in.GetVillageId(),
		PostalCode:                      in.GetPostalCode(),
		PreviousAddress:                 in.GetPreviousAddress(),
		IdNumber:                        in.GetIdNumber(),
		NpwpNumber:                      in.GetNpwpNumber(),
		NisnNumber:                      in.GetNisnNumber(),
		ResidenceType:                   in.GetResidenceType(),
		TransportationMean:              in.GetTransportationMean(),
		KpsReceiver:                     in.GetKpsReceiver(),
		PhoneNumber:                     in.GetPhoneNumber(),
		MobilePhoneNumber:               in.GetMobilePhoneNumber(),
		Email:                           in.GetEmail(),
		Homepage:                        in.GetHomepage(),
		WorkType:                        in.GetWorkType(),
		WorkPlace:                       in.GetWorkPlace(),
		Nationality:                     in.GetNationality(),
		AskesNumber:                     in.GetAskesNumber(),
		TotalBrother:                    in.GetTotalBrother(),
		TotalSister:                     in.GetTotalSister(),
		Hobby:                           in.GetHobby(),
		Experience:                      in.GetExperience(),
		TotalDependent:                  in.GetTotalDependent(),
		NimNumber:                       in.GetNimNumber(),
		StudentForce:                    in.GetStudentForce(),
		AdmittanceSemester:              in.GetAdmittanceSemester(),
		StudyProgramId:                  in.GetStudyProgramId(),
		CurriculumId:                    in.GetCurriculumId(),
		AdmittanceTestNumber:            in.GetAdmittanceTestNumber(),
		AdmittanceDate:                  admittanceDate,
		AdmittanceStatus:                in.GetAdmittanceStatus(),
		TotalTransferCredit:             in.GetTotalTransferCredit(),
		PreviousCollege:                 in.GetPreviousCollege(),
		PreviousStudyProgram:            in.GetPreviousStudyProgram(),
		PreviousNimNumber:               in.GetPreviousNimNumber(),
		PreviousNimAdmittanceYear:       in.GetPreviousNimAdmittanceYear(),
		Status:                          in.GetStatus(),
		IsForeignStudent:                in.GetIsForeignStudent(),
		CollegeEntranceType:             in.GetCollegeEntranceType(),
		ClassTime:                       in.GetClassTime(),
		FundSource:                      in.GetFundSource(),
		IsScholarshipGrantee:            in.GetIsScholarshipGrantee(),
		HasCompleteRequirement:          in.GetHasCompleteRequirement(),
		SchoolCertificateType:           in.GetSchoolCertificateType(),
		SchoolGraduationYear:            in.GetSchoolGraduationYear(),
		SchoolName:                      in.GetSchoolName(),
		SchoolAccreditation:             in.GetSchoolAccreditation(),
		SchoolAddress:                   in.GetSchoolAddress(),
		SchoolMajor:                     in.GetSchoolMajor(),
		SchoolCertificateNumber:         in.GetSchoolCertificateNumber(),
		SchoolCertificateDate:           schoolCertificateDate,
		TotalSchoolFinalExamSubject:     in.GetTotalSchoolFinalExamSubject(),
		SchoolFinalExamScore:            in.GetSchoolFinalExamScore(),
		GuardianName:                    in.GetGuardianName(),
		GuardianBirthDate:               guardianBirthDate,
		GuardianDeathDate:               guardianDeathDate,
		GuardianAddress:                 in.GetGuardianAddress(),
		GuardianRegencyId:               in.GetGuardianRegencyId(),
		GuardianPostalCode:              in.GetGuardianPostalCode(),
		GuardianPhoneNumber:             in.GetGuardianPhoneNumber(),
		GuardianEmail:                   in.GetGuardianEmail(),
		GuardianFinalAcademicBackground: in.GetGuardianFinalAcademicBackground(),
		GuardianOccupation:              in.GetGuardianOccupation(),
		FatherIdNumber:                  in.GetFatherIdNumber(),
		FatherName:                      in.GetFatherName(),
		FatherBirthDate:                 fatherBirthDate,
		FatherDeathDate:                 fatherDeathDate,
		MotherIdNumber:                  in.GetMotherIdNumber(),
		MotherName:                      in.GetMotherName(),
		MotherBirthDate:                 motherBirthDate,
		MotherDeathDate:                 motherDeathDate,
		ParentAddress:                   in.GetParentAddress(),
		ParentRegencyId:                 in.GetParentRegencyId(),
		ParentPostalCode:                in.GetParentPostalCode(),
		ParentPhoneNumber:               in.GetParentPhoneNumber(),
		ParentEmail:                     in.GetParentEmail(),
		FatherFinalAcademicBackground:   in.GetFatherFinalAcademicBackground(),
		FatherOccupation:                in.GetFatherOccupation(),
		MotherFinalAcademicBackground:   in.GetMotherFinalAcademicBackground(),
		MotherOccupation:                in.GetMotherOccupation(),
		ParentIncome:                    in.GetParentIncome(),
		IsFinanciallyCapable:            in.GetIsFinanciallyCapable(),
	}
	errs = a.StudentService.Update(ctx, data)
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
			Message: "Update Student",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a studentHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Student", "Delete", nil)

	errs := a.StudentService.Delete(ctx, in.GetId())
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
			Message: "Delete Student",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a studentHandler) BulkUpdateStatus(w http.ResponseWriter, r *http.Request) {
	var result BulkUpdateStatusResponse

	ctx := r.Context()
	var in BulkUpdateStatusRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = BulkUpdateStatusResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Student", "Bulk Update Status", &in)

	statusDate, errs := utils.StringToTime(in.GetStatusDate())
	if errs != nil {
		utils.PrintError(*errs)
		result = BulkUpdateStatusResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	data := objects.BulkUpdateStatusStudent{
		StudentIds:            in.GetStudentIds(),
		Status:                in.GetStatus(),
		StatusDate:            statusDate,
		StatusReferenceNumber: in.GetStatusReferenceNumber(),
		StatusPurpose:         in.GetStatusPurpose(),
		StatusRemarks:         in.GetStatusRemarks(),
	}
	errs = a.StudentService.BulkUpdateStatus(ctx, data)
	if errs != nil {
		utils.PrintError(*errs)
		result = BulkUpdateStatusResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = BulkUpdateStatusResponse{
		Meta: &Meta{
			Message: "Bulk Update Status Student",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a studentHandler) GetStatusSummary(w http.ResponseWriter, r *http.Request) {
	var result GetStatusSummaryResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetStatusSummaryRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetStatusSummaryResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Student", "Get Status Summary", nil)

	data, errs := a.StudentService.GetStatusSummary(ctx, in.GetSemesterId())
	if errs != nil {
		utils.PrintError(*errs)
		result = GetStatusSummaryResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*GetStatusSummaryResponseData{}
	for _, v := range data {
		statuses := []*GetStatusSummaryResponseDataStatus{}
		for _, s := range v.Statuses {
			statuses = append(statuses, &GetStatusSummaryResponseDataStatus{
				Status: s.Status,
				Total:  s.Total,
			})
		}
		resultData = append(resultData, &GetStatusSummaryResponseData{
			StudyProgramId:        v.StudyProgramId,
			StudyProgramName:      v.StudyProgramName,
			DiktiStudyProgramType: v.DiktiStudyProgramType,
			StudyLevelShortName:   v.StudyLevelShortName,
			Statuses:              statuses,
		})
	}

	result = GetStatusSummaryResponse{
		Meta: &Meta{
			Message: "Get Status Summary Student",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: resultData,
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (s studentHandler) GetSubjectGrade(w http.ResponseWriter, r *http.Request) {
	var result GetSubjectGradeResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetSubjectGradeRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetSubjectGradeResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}
	data, errs := s.StudentService.GetSubjectGrade(ctx, paginationData, in.GetStudentId())
	if errs != nil {
		utils.PrintError(*errs)
		result = GetSubjectGradeResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*GetSubjectGradeResponseData{}
	for _, v := range data.Data {
		resultData = append(resultData, &GetSubjectGradeResponseData{
			SubjectId:               v.SubjectId,
			SubjectCode:             v.SubjectCode,
			SubjectName:             v.SubjectName,
			GradeSemesterId:         v.GradeSemesterId,
			GradeSemesterSchoolYear: v.GradeSemesterSchoolYear,
			GradeSemesterType:       v.GradeSemesterType,
			GradePoint:              v.GradePoint,
			GradeCode:               utils.NullStringScan(v.GradeCode),
			SubjectIsMandatory:      v.SubjectIsMandatory, // SubjectIsMandatory      bool    `protobuf:"varint,9,opt,name=subject_is_mandatory,json=subjectIsMandatory,proto3" json:"subject_is_mandatory"`
			SemesterPackage:         v.SemesterPackage,    // SemesterPackage         uint32  `protobuf:"varint,10,opt,name=semester_package,json=semesterPackage,proto3" json:"semester_package"`
			SubjectTotalCredit:      v.SubjectTotalCredit, // SubjectTotalCredit      uint32  `protobuf:"varint,11,opt,name=subject_total_credit,json=subjectTotalCredit,proto3" json:"subject_total_credit"`
			SubjectType:             v.SubjectType,        // SubjectType             string  `protobuf:"bytes,12,opt,name=subject_type,json=subjectType,proto3" json:"subject_type"`
		})
	}

	result = GetSubjectGradeResponse{
		Meta: &Meta{
			Message: "Get Student Subject Grade",
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

func (a studentHandler) BulkUpdatePayment(w http.ResponseWriter, r *http.Request) {
	var result BulkUpdatePaymentResponse

	ctx := r.Context()
	var in BulkUpdatePaymentRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = BulkUpdatePaymentResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Student", "Bulk Update Payment", nil)

	errs := a.StudentService.BulkUpdatePayment(ctx, in.GetStudentIds())
	if errs != nil {
		utils.PrintError(*errs)
		result = BulkUpdatePaymentResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = BulkUpdatePaymentResponse{
		Meta: &Meta{
			Message: "Bulk Update Payment Student",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a studentHandler) GetPaymentLog(w http.ResponseWriter, r *http.Request) {
	var result GetPaymentLogResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetPaymentLogRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetPaymentLogResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Student", "Get Payment Log", nil)

	data, errs := a.StudentService.GetPaymentLog(ctx, in.GetStudentId())
	if errs != nil {
		utils.PrintError(*errs)
		result = GetPaymentLogResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*GetPaymentLogResponseData{}
	for _, v := range data {
		resultData = append(resultData, &GetPaymentLogResponseData{
			SemesterId:         v.SemesterId,
			SemesterType:       v.SemesterType,
			SemesterStartYear:  v.SemesterStartYear,
			SemesterSchoolYear: v.SemesterSchoolYear,
			CreatedAt:          v.CreatedAt.Format(constants.DateRFC),
		})
	}

	result = GetPaymentLogResponse{
		Meta: &Meta{
			Message: "Get Payment Log Student",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: resultData,
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a studentHandler) ConvertGrade(w http.ResponseWriter, r *http.Request) {
	var result ConvertGradeResponse

	ctx := r.Context()
	var in ConvertGradeRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = ConvertGradeResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Student", "Convert Grade", nil)

	data := objects.ConvertMbkmGrade{
		StudentId:             in.GetStudentId(),
		MbkmClassId:           in.GetMbkmClassId(),
		DestinationSubjectIds: in.GetDestinationSubjectIds(),
	}
	errs := a.StudentSubjectService.ConvertMbkmGrade(ctx, data)
	if errs != nil {
		utils.PrintError(*errs)
		result = ConvertGradeResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = ConvertGradeResponse{
		Meta: &Meta{
			Message: "Convert Grade Student",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
