package exam_supervisor

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

type examSupervisorHandler struct {
	*service.ServiceCtx
}

func (p examSupervisorHandler) GetList(w http.ResponseWriter, r *http.Request) {
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

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}
	data, errs := p.ExamSupervisorService.GetList(ctx, paginationData, in.GetStudyProgramId())
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
			IdNationalLecturer: v.IdNationalLecturer,
			Name:               v.Name,
		})
	}

	result = GetListResponse{
		Meta: &Meta{
			Message: "Get List Exam Supervisor",
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

func (a examSupervisorHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Exam Supervisor", "Get Detail", nil)

	data, errs := a.ExamSupervisorService.GetDetail(ctx, in.GetId())
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
			Message: "Get Detail Exam Supervisor",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &GetDetailResponseData{
			Id:                 data.Id,
			IdNationalLecturer: data.IdNationalLecturer,
			Name:               data.Name,
			FrontTitle:         utils.NullStringScan(data.FrontTitle),
			BackDegree:         utils.NullStringScan(data.BackDegree),
			StudyProgramId:     utils.NullStringScan(data.StudyProgramId),
			StudyProgramName:   utils.NullStringScan(data.StudyProgramName),
			IdNumber:           utils.NullStringScan(data.IdNumber),
			BirthDate:          utils.SafetyDate(data.BirthDate),
			BirthRegencyId:     utils.NullUint32Scan(data.BirthRegencyId),
			BirthRegencyName:   utils.NullStringScan(data.BirthRegencyName),
			BirthCountryId:     utils.NullUint32Scan(data.BirthCountryId),
			BirthCountryName:   utils.NullStringScan(data.BirthCountryName),
			Sex:                utils.NullStringScan(data.Sex),
			BloodType:          utils.NullStringScan(data.BloodType),
			Religion:           utils.NullStringScan(data.Religion),
			MaritalStatus:      utils.NullStringScan(data.MaritalStatus),
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
			EmployeeType:       utils.NullStringScan(data.EmployeeType),
			EmployeeStatus:     utils.NullStringScan(data.EmployeeStatus),
			SkCpnsNumber:       utils.NullStringScan(data.SkCpnsNumber),
			SkCpnsDate:         utils.SafetyDate(data.SkCpnsDate),
			TmtCpnsDate:        utils.SafetyDate(data.TmtCpnsDate),
			CpnsCategory:       utils.NullStringScan(data.CpnsCategory),
			CpnsDurationMonth:  utils.NullUint32Scan(data.CpnsDurationMonth),
			PrePositionDate:    utils.SafetyDate(data.PrePositionDate),
			SkPnsNumber:        utils.NullStringScan(data.SkPnsNumber),
			SkPnsDate:          utils.SafetyDate(data.SkPnsDate),
			TmtPnsDate:         utils.SafetyDate(data.TmtPnsDate),
			PnsCategory:        utils.NullStringScan(data.PnsCategory),
			PnsOathDate:        utils.SafetyDate(data.PnsOathDate),
			JoinDate:           utils.SafetyDate(data.JoinDate),
			EndDate:            utils.SafetyDate(data.EndDate),
			TaspenNumber:       utils.NullStringScan(data.TaspenNumber),
			FormerInstance:     utils.NullStringScan(data.FormerInstance),
			Remarks:            utils.NullStringScan(data.Remarks),
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a examSupervisorHandler) Create(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Exam Supervisor", "Create", &in)

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
	data := objects.CreateExamSupervisor{
		IdNationalLecturer: in.GetIdNationalLecturer(),
		Name:               in.GetName(),
		FrontTitle:         in.GetFrontTitle(),
		BackDegree:         in.GetBackDegree(),
		StudyProgramId:     in.GetStudyProgramId(),
		IdNumber:           in.GetIdNumber(),
		BirthDate:          birthDate,
		BirthRegencyId:     in.GetBirthRegencyId(),
		Sex:                in.GetSex(),
		BloodType:          in.GetBloodType(),
		Religion:           in.GetReligion(),
		MaritalStatus:      in.GetMaritalStatus(),
		Address:            in.GetAddress(),
		RegencyId:          in.GetRegencyId(),
		PostalCode:         in.GetPostalCode(),
		PhoneNumber:        in.GetPhoneNumber(),
		Fax:                in.GetFax(),
		MobilePhoneNumber:  in.GetMobilePhoneNumber(),
		OfficePhoneNumber:  in.GetOfficePhoneNumber(),
		EmployeeType:       in.GetEmployeeType(),
		EmployeeStatus:     in.GetEmployeeStatus(),
		SkCpnsNumber:       in.GetSkCpnsNumber(),
		SkCpnsDate:         skCpnsDate,
		TmtCpnsDate:        tmtCpnsDate,
		CpnsCategory:       in.GetCpnsCategory(),
		CpnsDurationMonth:  in.GetCpnsDurationMonth(),
		PrePositionDate:    prePositionDate,
		SkPnsNumber:        in.GetSkPnsNumber(),
		SkPnsDate:          skPnsDate,
		TmtPnsDate:         tmtPnsDate,
		PnsCategory:        in.GetPnsCategory(),
		PnsOathDate:        pnsOathDate,
		JoinDate:           joinDate,
		EndDate:            endDate,
		TaspenNumber:       in.GetTaspenNumber(),
		FormerInstance:     in.GetFormerInstance(),
		Remarks:            in.GetRemarks(),
	}

	errs = a.ExamSupervisorService.Create(ctx, data)
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
			Message: "Create Exam Supervisor",
			Status:  http.StatusCreated,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusCreated, &result)
}

func (a examSupervisorHandler) Update(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Exam Supervisor", "Update", &in)

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
	data := objects.UpdateExamSupervisor{
		Id:                 in.GetId(),
		IdNationalLecturer: in.GetIdNationalLecturer(),
		Name:               in.GetName(),
		FrontTitle:         in.GetFrontTitle(),
		BackDegree:         in.GetBackDegree(),
		StudyProgramId:     in.GetStudyProgramId(),
		IdNumber:           in.GetIdNumber(),
		BirthDate:          birthDate,
		BirthRegencyId:     in.GetBirthRegencyId(),
		Sex:                in.GetSex(),
		BloodType:          in.GetBloodType(),
		Religion:           in.GetReligion(),
		MaritalStatus:      in.GetMaritalStatus(),
		Address:            in.GetAddress(),
		RegencyId:          in.GetRegencyId(),
		PostalCode:         in.GetPostalCode(),
		PhoneNumber:        in.GetPhoneNumber(),
		Fax:                in.GetFax(),
		MobilePhoneNumber:  in.GetMobilePhoneNumber(),
		OfficePhoneNumber:  in.GetOfficePhoneNumber(),
		EmployeeType:       in.GetEmployeeType(),
		EmployeeStatus:     in.GetEmployeeStatus(),
		SkCpnsNumber:       in.GetSkCpnsNumber(),
		SkCpnsDate:         skCpnsDate,
		TmtCpnsDate:        tmtCpnsDate,
		CpnsCategory:       in.GetCpnsCategory(),
		CpnsDurationMonth:  in.GetCpnsDurationMonth(),
		PrePositionDate:    prePositionDate,
		SkPnsNumber:        in.GetSkPnsNumber(),
		SkPnsDate:          skPnsDate,
		TmtPnsDate:         tmtPnsDate,
		PnsCategory:        in.GetPnsCategory(),
		PnsOathDate:        pnsOathDate,
		JoinDate:           joinDate,
		EndDate:            endDate,
		TaspenNumber:       in.GetTaspenNumber(),
		FormerInstance:     in.GetFormerInstance(),
		Remarks:            in.GetRemarks(),
	}
	errs = a.ExamSupervisorService.Update(ctx, data)
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
			Message: "Update Exam Supervisor",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a examSupervisorHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Exam Supervisor", "Delete", nil)

	errs := a.ExamSupervisorService.Delete(ctx, in.GetId())
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
			Message: "Delete Exam Supervisor",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
