package general

import (
	"net/http"
	"time"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/service"
	"github.com/sccicitb/pupr-backend/utils"
)

type lecturerGeneralHandler struct {
	*service.ServiceCtx
}

func (l lecturerGeneralHandler) GetSemesterSummary(w http.ResponseWriter, r *http.Request) {
	var result GetSemesterSummaryResponse

	ctx := r.Context()

	defer l.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "General", "Get Semester Summary", nil)

	data, errs := l.LecturerService.GetSemesterSummary(ctx)
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
			Message: "Get Semester Summary Lecturer",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &GetSemesterSummaryResponseData{
			StudyPlanApprovalStartDate:   data.StudyPlanApprovalStartDate.Format(constants.DateRFC),
			StudyPlanApprovalEndDate:     data.StudyPlanApprovalEndDate.Format(constants.DateRFC),
			AcademicGuidanceTotalStudent: data.AcademicGuidanceTotalStudent,
			TotalClass:                   data.TotalClass,
			SchoolYear:                   data.SchoolYear,
			SemesterType:                 data.SemesterType,
			GradingStartDate:             utils.SafetyDate(data.GradingStartDate),
			GradingEndDate:               utils.SafetyDate(data.GradingEndDate),
			SemesterId:                   data.SemesterId,
		},
	}

	utils.JSONResponse(w, http.StatusOK, &result)
}

func (l lecturerGeneralHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	var result GetProfileResponse

	ctx := r.Context()

	defer l.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "General", "Get Profile", nil)

	data, errs := l.LecturerService.GetProfile(ctx)
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
			Message: "Get Profile Lecturer",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &GetProfileResponseData{
			Id:                 data.Id,
			IdNationalLecturer: data.IdNationalLecturer,
			Name:               data.Name,
			FrontTitle:         utils.NullStringScan(data.FrontTitle),
			BackDegree:         utils.NullStringScan(data.BackDegree),
			StudyProgramId:     utils.NullStringScan(data.StudyProgramId),
			StudyProgramName:   utils.NullStringScan(data.StudyProgramName),
			BirthDate:          utils.SafetyDate(data.BirthDate),
			BirthRegencyId:     utils.NullUint32Scan(data.BirthRegencyId),
			BirthRegencyName:   utils.NullStringScan(data.BirthRegencyName),
			BirthCountryId:     utils.NullUint32Scan(data.BirthCountryId),
			BirthCountryName:   utils.NullStringScan(data.BirthCountryName),
			Sex:                utils.NullStringScan(data.Sex),
			Religion:           utils.NullStringScan(data.Religion),
			Address:            utils.NullStringScan(data.Address),
			RegencyId:          utils.NullUint32Scan(data.RegencyId),
			RegencyName:        utils.NullStringScan(data.RegencyName),
			CountryId:          utils.NullUint32Scan(data.CountryId),
			CountryName:        utils.NullStringScan(data.CountryName),
			PostalCode:         utils.NullStringScan(data.PostalCode),
			PhoneNumber:        utils.NullStringScan(data.PhoneNumber),
			Fax:                utils.NullStringScan(data.Fax),
			MobilePhoneNumber:  utils.NullStringScan(data.MobilePhoneNumber),
			OfficePhoneNumber:  utils.NullStringScan(data.OfficePhoneNumber),
			AcademicPosition:   utils.NullStringScan(data.AcademicPosition),
			EmploymentStatus:   utils.NullStringScan(data.EmploymentStatus),
			Status:             utils.NullStringScan(data.Status),
		},
	}

	utils.JSONResponse(w, http.StatusOK, &result)
}
