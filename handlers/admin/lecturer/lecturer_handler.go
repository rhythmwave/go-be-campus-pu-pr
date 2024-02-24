package lecturer

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

type lecturerHandler struct {
	*service.ServiceCtx
}

func (p lecturerHandler) GetList(w http.ResponseWriter, r *http.Request) {
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
	excludeLectureDate, errs := utils.StringToDate(in.GetExcludeLectureDate())
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

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}
	req := objects.GetLecturerRequest{
		StudyProgramId:             in.GetStudyProgramId(),
		IdNationalLecturer:         in.GetIdNationalLecturer(),
		EmploymentStatus:           in.GetEmploymentStatus(),
		AcademicGuidanceSemesterId: in.GetAcademicGuidanceSemesterId(),
		Status:                     in.GetStatus(),
		HasAuthentication:          hasAuthentication,
		ClassId:                    in.GetClassId(),
		ExcludeLectureDate:         excludeLectureDate,
		ExcludeStartTime:           in.GetExcludeStartTime(),
		ExcludeEndTime:             in.GetExcludeEndTime(),
		ForceIncludeLectureId:      in.GetForceIncludeLectureId(),
	}
	data, errs := p.LecturerService.GetList(ctx, paginationData, req)
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
			Id:                              v.Id,
			Name:                            v.Name,
			PhoneNumber:                     utils.NullStringScan(v.PhoneNumber),
			MobilePhoneNumber:               utils.NullStringScan(v.MobilePhoneNumber),
			OfficePhoneNumber:               utils.NullStringScan(v.OfficePhoneNumber),
			IdNationalLecturer:              v.IdNationalLecturer,
			FrontTitle:                      utils.NullStringScan(v.FrontTitle),
			BackDegree:                      utils.NullStringScan(v.BackDegree),
			DiktiStudyProgramCode:           utils.NullStringScan(v.DiktiStudyProgramCode),
			StudyProgramName:                utils.NullStringScan(v.StudyProgramName),
			EmploymentStatus:                utils.NullStringScan(v.EmploymentStatus),
			Status:                          utils.NullStringScan(v.Status),
			AuthenticationId:                utils.NullStringScan(v.AuthenticationId),
			AuthenticationIsActive:          utils.NullBooleanScan(v.AuthenticationIsActive),
			AuthenticationSuspensionRemarks: utils.NullStringScan(v.AuthenticationSuspensionRemarks),
			AcademicGuidanceTotalStudent:    utils.NullUint32Scan(v.AcademicGuidanceTotalStudent),
			AcademicGuidanceId:              utils.NullStringScan(v.AcademicGuidanceId),
			AcademicGuidanceDecisionNumber:  utils.NullStringScan(v.AcademicGuidanceDecisionNumber),
			AcademicGuidanceDecisionDate:    utils.SafetyDate(v.AcademicGuidanceDecisionDate),
		})
	}

	result = GetListResponse{
		Meta: &Meta{
			Message: "Get List Lecturer",
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

func (a lecturerHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Lecturer", "Get Detail", nil)

	data, errs := a.LecturerService.GetDetail(ctx, in.GetId())
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
			Message: "Get Detail Lecturer",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &GetDetailResponseData{
			Id:                        data.Id,
			IdNationalLecturer:        data.IdNationalLecturer,
			Name:                      data.Name,
			FrontTitle:                utils.NullStringScan(data.FrontTitle),
			BackDegree:                utils.NullStringScan(data.BackDegree),
			StudyProgramId:            utils.NullStringScan(data.StudyProgramId),
			StudyProgramName:          utils.NullStringScan(data.StudyProgramName),
			IdNumber:                  utils.NullStringScan(data.IdNumber),
			BirthDate:                 utils.SafetyDate(data.BirthDate),
			BirthRegencyId:            utils.NullUint32Scan(data.BirthRegencyId),
			BirthRegencyName:          utils.NullStringScan(data.BirthRegencyName),
			BirthCountryId:            utils.NullUint32Scan(data.BirthCountryId),
			BirthCountryName:          utils.NullStringScan(data.BirthCountryName),
			IdEmployee:                utils.NullStringScan(data.IdEmployee),
			Stambuk:                   utils.NullStringScan(data.Stambuk),
			Sex:                       utils.NullStringScan(data.Sex),
			BloodType:                 utils.NullStringScan(data.BloodType),
			Religion:                  utils.NullStringScan(data.Religion),
			MaritalStatus:             utils.NullStringScan(data.MaritalStatus),
			Address:                   utils.NullStringScan(data.Address),
			RegencyId:                 utils.NullUint32Scan(data.RegencyId),
			RegencyName:               utils.NullStringScan(data.RegencyName),
			CountryId:                 utils.NullUint32Scan(data.CountryId),
			CountryName:               utils.NullStringScan(data.CountryName),
			PostalCode:                utils.NullStringScan(data.PostalCode),
			PhoneNumber:               utils.NullStringScan(data.PhoneNumber),
			Fax:                       utils.NullStringScan(data.Fax),
			MobilePhoneNumber:         utils.NullStringScan(data.MobilePhoneNumber),
			OfficePhoneNumber:         utils.NullStringScan(data.OfficePhoneNumber),
			EmployeeType:              utils.NullStringScan(data.EmployeeType),
			EmployeeStatus:            utils.NullStringScan(data.EmployeeStatus),
			SkCpnsNumber:              utils.NullStringScan(data.SkCpnsNumber),
			SkCpnsDate:                utils.SafetyDate(data.SkCpnsDate),
			TmtCpnsDate:               utils.SafetyDate(data.TmtCpnsDate),
			CpnsCategory:              utils.NullStringScan(data.CpnsCategory),
			CpnsDurationMonth:         utils.NullUint32Scan(data.CpnsDurationMonth),
			PrePositionDate:           utils.SafetyDate(data.PrePositionDate),
			SkPnsNumber:               utils.NullStringScan(data.SkPnsNumber),
			SkPnsDate:                 utils.SafetyDate(data.SkPnsDate),
			TmtPnsDate:                utils.SafetyDate(data.TmtPnsDate),
			PnsCategory:               utils.NullStringScan(data.PnsCategory),
			PnsOathDate:               utils.SafetyDate(data.PnsOathDate),
			JoinDate:                  utils.SafetyDate(data.JoinDate),
			EndDate:                   utils.SafetyDate(data.EndDate),
			TaspenNumber:              utils.NullStringScan(data.TaspenNumber),
			FormerInstance:            utils.NullStringScan(data.FormerInstance),
			Remarks:                   utils.NullStringScan(data.Remarks),
			LecturerNumber:            utils.NullStringScan(data.LecturerNumber),
			AcademicPosition:          utils.NullStringScan(data.AcademicPosition),
			EmploymentStatus:          utils.NullStringScan(data.EmploymentStatus),
			Expertise:                 utils.NullStringScan(data.Expertise),
			HighestDegree:             utils.NullStringScan(data.HighestDegree),
			InstanceCode:              utils.NullStringScan(data.InstanceCode),
			TeachingCertificateNumber: utils.NullStringScan(data.TeachingCertificateNumber),
			TeachingPermitNumber:      utils.NullStringScan(data.TeachingPermitNumber),
			Status:                    utils.NullStringScan(data.Status),
			ResignSemester:            utils.NullStringScan(data.ResignSemester),
			ExpertiseGroupId:          utils.NullStringScan(data.ExpertiseGroupId),
			ExpertiseGroupName:        utils.NullStringScan(data.ExpertiseGroupName),
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a lecturerHandler) Create(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Lecturer", "Create", &in)

	birthDate, errs := utils.StringToTime(in.GetBirthDate())
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
	skCpnsDate, errs := utils.StringToTime(in.GetSkCpnsDate())
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
	tmtCpnsDate, errs := utils.StringToTime(in.GetTmtCpnsDate())
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
	prePositionDate, errs := utils.StringToTime(in.GetPrePositionDate())
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
	skPnsDate, errs := utils.StringToTime(in.GetSkPnsDate())
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
	tmtPnsDate, errs := utils.StringToTime(in.GetTmtPnsDate())
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
	pnsOathDate, errs := utils.StringToTime(in.GetPnsOathDate())
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
	joinDate, errs := utils.StringToTime(in.GetJoinDate())
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
	endDate, errs := utils.StringToTime(in.GetEndDate())
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
	data := objects.CreateLecturer{
		IdNationalLecturer:        in.GetIdNationalLecturer(),
		Name:                      in.GetName(),
		FrontTitle:                in.GetFrontTitle(),
		BackDegree:                in.GetBackDegree(),
		StudyProgramId:            in.GetStudyProgramId(),
		IdNumber:                  in.GetIdNumber(),
		BirthDate:                 birthDate,
		BirthRegencyId:            in.GetBirthRegencyId(),
		IdEmployee:                in.GetIdEmployee(),
		Stambuk:                   in.GetStambuk(),
		Sex:                       in.GetSex(),
		BloodType:                 in.GetBloodType(),
		Religion:                  in.GetReligion(),
		MaritalStatus:             in.GetMaritalStatus(),
		Address:                   in.GetAddress(),
		RegencyId:                 in.GetRegencyId(),
		PostalCode:                in.GetPostalCode(),
		PhoneNumber:               in.GetPhoneNumber(),
		Fax:                       in.GetFax(),
		MobilePhoneNumber:         in.GetMobilePhoneNumber(),
		OfficePhoneNumber:         in.GetOfficePhoneNumber(),
		EmployeeType:              in.GetEmployeeType(),
		EmployeeStatus:            in.GetEmployeeStatus(),
		SkCpnsNumber:              in.GetSkCpnsNumber(),
		SkCpnsDate:                skCpnsDate,
		TmtCpnsDate:               tmtCpnsDate,
		CpnsCategory:              in.GetCpnsCategory(),
		CpnsDurationMonth:         in.GetCpnsDurationMonth(),
		PrePositionDate:           prePositionDate,
		SkPnsNumber:               in.GetSkPnsNumber(),
		SkPnsDate:                 skPnsDate,
		TmtPnsDate:                tmtPnsDate,
		PnsCategory:               in.GetPnsCategory(),
		PnsOathDate:               pnsOathDate,
		JoinDate:                  joinDate,
		EndDate:                   endDate,
		TaspenNumber:              in.GetTaspenNumber(),
		FormerInstance:            in.GetFormerInstance(),
		Remarks:                   in.GetRemarks(),
		LecturerNumber:            in.GetLecturerNumber(),
		AcademicPosition:          in.GetAcademicPosition(),
		EmploymentStatus:          in.GetEmploymentStatus(),
		Expertise:                 in.GetExpertise(),
		HighestDegree:             in.GetHighestDegree(),
		InstanceCode:              in.GetInstanceCode(),
		TeachingCertificateNumber: in.GetTeachingCertificateNumber(),
		TeachingPermitNumber:      in.GetTeachingPermitNumber(),
		Status:                    in.GetStatus(),
		ResignSemester:            in.GetResignSemester(),
		ExpertiseGroupId:          in.GetExpertiseGroupId(),
	}

	errs = a.LecturerService.Create(ctx, data)
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
			Message: "Create Lecturer",
			Status:  http.StatusCreated,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusCreated, &result)
}

func (a lecturerHandler) Update(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Lecturer", "Update", &in)

	birthDate, errs := utils.StringToTime(in.GetBirthDate())
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
	skCpnsDate, errs := utils.StringToTime(in.GetSkCpnsDate())
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
	tmtCpnsDate, errs := utils.StringToTime(in.GetTmtCpnsDate())
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
	prePositionDate, errs := utils.StringToTime(in.GetPrePositionDate())
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
	skPnsDate, errs := utils.StringToTime(in.GetSkPnsDate())
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
	tmtPnsDate, errs := utils.StringToTime(in.GetTmtPnsDate())
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
	pnsOathDate, errs := utils.StringToTime(in.GetPnsOathDate())
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
	joinDate, errs := utils.StringToTime(in.GetJoinDate())
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
	endDate, errs := utils.StringToTime(in.GetEndDate())
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
	data := objects.UpdateLecturer{
		Id:                        in.GetId(),
		IdNationalLecturer:        in.GetIdNationalLecturer(),
		Name:                      in.GetName(),
		FrontTitle:                in.GetFrontTitle(),
		BackDegree:                in.GetBackDegree(),
		StudyProgramId:            in.GetStudyProgramId(),
		IdNumber:                  in.GetIdNumber(),
		BirthDate:                 birthDate,
		BirthRegencyId:            in.GetBirthRegencyId(),
		IdEmployee:                in.GetIdEmployee(),
		Stambuk:                   in.GetStambuk(),
		Sex:                       in.GetSex(),
		BloodType:                 in.GetBloodType(),
		Religion:                  in.GetReligion(),
		MaritalStatus:             in.GetMaritalStatus(),
		Address:                   in.GetAddress(),
		RegencyId:                 in.GetRegencyId(),
		PostalCode:                in.GetPostalCode(),
		PhoneNumber:               in.GetPhoneNumber(),
		Fax:                       in.GetFax(),
		MobilePhoneNumber:         in.GetMobilePhoneNumber(),
		OfficePhoneNumber:         in.GetOfficePhoneNumber(),
		EmployeeType:              in.GetEmployeeType(),
		EmployeeStatus:            in.GetEmployeeStatus(),
		SkCpnsNumber:              in.GetSkCpnsNumber(),
		SkCpnsDate:                skCpnsDate,
		TmtCpnsDate:               tmtCpnsDate,
		CpnsCategory:              in.GetCpnsCategory(),
		CpnsDurationMonth:         in.GetCpnsDurationMonth(),
		PrePositionDate:           prePositionDate,
		SkPnsNumber:               in.GetSkPnsNumber(),
		SkPnsDate:                 skPnsDate,
		TmtPnsDate:                tmtPnsDate,
		PnsCategory:               in.GetPnsCategory(),
		PnsOathDate:               pnsOathDate,
		JoinDate:                  joinDate,
		EndDate:                   endDate,
		TaspenNumber:              in.GetTaspenNumber(),
		FormerInstance:            in.GetFormerInstance(),
		Remarks:                   in.GetRemarks(),
		LecturerNumber:            in.GetLecturerNumber(),
		AcademicPosition:          in.GetAcademicPosition(),
		EmploymentStatus:          in.GetEmploymentStatus(),
		Expertise:                 in.GetExpertise(),
		HighestDegree:             in.GetHighestDegree(),
		InstanceCode:              in.GetInstanceCode(),
		TeachingCertificateNumber: in.GetTeachingCertificateNumber(),
		TeachingPermitNumber:      in.GetTeachingPermitNumber(),
		Status:                    in.GetStatus(),
		ResignSemester:            in.GetResignSemester(),
		ExpertiseGroupId:          in.GetExpertiseGroupId(),
	}
	errs = a.LecturerService.Update(ctx, data)
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
			Message: "Update Lecturer",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a lecturerHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Lecturer", "Delete", nil)

	errs := a.LecturerService.Delete(ctx, in.GetId())
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
			Message: "Delete Lecturer",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (p lecturerHandler) GetSchedule(w http.ResponseWriter, r *http.Request) {
	var result GetScheduleResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetScheduleRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetScheduleResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer p.AdminActivityLogService.Create(ctx, r, time.Now(), "Lecturer", "Get Schedule", nil)

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}
	data, errs := p.LecturerService.GetSchedule(ctx, paginationData, in.GetStudyProgramId(), in.GetIdNationalLecturer(), in.GetSemesterId())
	if errs != nil {
		utils.PrintError(*errs)
		result = GetScheduleResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*GetScheduleResponseData{}
	for _, v := range data.Data {
		resultData = append(resultData, &GetScheduleResponseData{
			Id:                 v.Id,
			IdNationalLecturer: v.IdNationalLecturer,
			Name:               v.Name,
			FrontTitle:         utils.NullStringScan(v.FrontTitle),
			BackDegree:         utils.NullStringScan(v.BackDegree),
			StudyProgramName:   v.StudyProgramName,
			SubjectName:        v.SubjectName,
			ClassName:          v.ClassName,
			TotalSubjectCredit: v.TotalSubjectCredit,
			LecturePlanDate:    v.LecturePlanDate.Format(constants.DateRFC),
			StartTime:          v.StartTime,
			EndTime:            v.EndTime,
			RoomName:           utils.NullStringScan(v.RoomName),
			TotalParticipant:   v.TotalParticipant,
		})
	}

	result = GetScheduleResponse{
		Meta: &Meta{
			Message: "Get Lecturer Schedule",
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

func (p lecturerHandler) GetAssignedClass(w http.ResponseWriter, r *http.Request) {
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
	defer p.AdminActivityLogService.Create(ctx, r, time.Now(), "Lecturer", "Get Assigned Class", nil)

	data, errs := p.LecturerService.GetAssignedClass(ctx, in.GetSemesterId(), in.GetLecturerId(), nil)
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
			TotalAttendance:      v.TotalAttendance,
			TotalLectureDone:     v.TotalLectureDone,
			AttendancePercentage: v.AttendancePercentage,
		})
	}

	result = GetAssignedClassResponse{
		Meta: &Meta{
			Message: "Get Lecturer Assigned Class",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: resultData,
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
