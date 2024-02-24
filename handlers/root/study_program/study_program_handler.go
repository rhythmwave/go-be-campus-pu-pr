package study_program

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

type studyProgramHandler struct {
	*service.ServiceCtx
}

func (f studyProgramHandler) GetList(w http.ResponseWriter, r *http.Request) {
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
	defer f.AdminActivityLogService.Create(ctx, r, time.Now(), "Study Program", "Get List", nil)

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}
	data, errs := f.StudyProgramService.GetList(ctx, paginationData, in.GetMajorId(), true)
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
			Id:                    v.Id,
			Name:                  v.Name,
			StudyLevelShortName:   v.StudyLevelShortName,
			DiktiStudyProgramType: v.DiktiStudyProgramType,
		})
	}

	result = GetListResponse{
		Meta: &Meta{
			Message: "Get List Study Program",
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

func (f studyProgramHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
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
	defer f.AdminActivityLogService.Create(ctx, r, time.Now(), "Study Program", "Get Detail", nil)

	data, errs := f.StudyProgramService.GetDetail(ctx, in.GetId())
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
			Message: "Get Detail Study Program",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &GetDetailResponseData{
			Id:                            data.Id,
			DiktiStudyProgramId:           data.DiktiStudyProgramId,
			DiktiStudyProgramCode:         data.DiktiStudyProgramCode,
			DiktiStudyProgramName:         data.DiktiStudyProgramName,
			DiktiStudyProgramType:         data.DiktiStudyProgramType,
			StudyLevelShortName:           data.StudyLevelShortName,
			StudyLevelName:                data.StudyLevelName,
			Name:                          data.Name,
			EnglishName:                   utils.NullStringScan(data.EnglishName),
			ShortName:                     utils.NullStringScan(data.ShortName),
			EnglishShortName:              utils.NullStringScan(data.EnglishShortName),
			AdministrativeUnit:            utils.NullStringScan(data.AdministrativeUnit),
			FacultyId:                     data.FacultyId,
			FacultyName:                   data.FacultyName,
			MajorId:                       data.MajorId,
			MajorName:                     data.MajorName,
			Address:                       utils.NullStringScan(data.Address),
			PhoneNumber:                   utils.NullStringScan(data.PhoneNumber),
			Fax:                           utils.NullStringScan(data.Fax),
			Email:                         utils.NullStringScan(data.Email),
			Website:                       utils.NullStringScan(data.Website),
			ContactPerson:                 utils.NullStringScan(data.ContactPerson),
			CuriculumReviewFrequency:      data.CuriculumReviewFrequency,
			CuriculumReviewMethod:         data.CuriculumReviewMethod,
			EstablishmentDate:             data.EstablishmentDate.Format(constants.DateRFC),
			IsActive:                      data.IsActive,
			StartSemester:                 utils.NullStringScan(data.StartSemester),
			OperationalPermitNumber:       data.OperationalPermitNumber,
			OperationalPermitDate:         data.OperationalPermitDate.Format(constants.DateRFC),
			OperationalPermitDueDate:      data.OperationalPermitDueDate.Format(constants.DateRFC),
			HeadLecturerId:                utils.NullStringScan(data.HeadLecturerId),
			HeadLecturerName:              utils.NullStringScan(data.HeadLecturerName),
			HeadLecturerMobilePhoneNumber: utils.NullStringScan(data.HeadLecturerMobilePhoneNumber),
			OperatorName:                  utils.NullStringScan(data.OperatorName),
			OperatorPhoneNumber:           utils.NullStringScan(data.OperatorPhoneNumber),
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (f studyProgramHandler) Create(w http.ResponseWriter, r *http.Request) {
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
	defer f.AdminActivityLogService.Create(ctx, r, time.Now(), "Study Program", "Create", &in)

	establishmentDate, errs := utils.StringToTime(in.GetEstablishmentDate())
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
	operationalPermitDate, errs := utils.StringToTime(in.GetOperationalPermitDate())
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
	operationalPermitDueDate, errs := utils.StringToTime(in.GetOperationalPermitDueDate())
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

	data := objects.CreateStudyProgram{
		DiktiStudyProgramId:      in.GetDiktiStudyProgramId(),
		Name:                     in.GetName(),
		EnglishName:              in.GetEnglishName(),
		ShortName:                in.GetShortName(),
		EnglishShortName:         in.GetEnglishShortName(),
		AdministrativeUnit:       in.GetAdministrativeUnit(),
		MajorId:                  in.GetMajorId(),
		Address:                  in.GetAddress(),
		PhoneNumber:              in.GetPhoneNumber(),
		Fax:                      in.GetFax(),
		Email:                    in.GetEmail(),
		Website:                  in.GetWebsite(),
		ContactPerson:            in.GetContactPerson(),
		CuriculumReviewFrequency: in.GetCuriculumReviewFrequency(),
		CuriculumReviewMethod:    in.GetCuriculumReviewMethod(),
		EstablishmentDate:        establishmentDate,
		IsActive:                 in.GetIsActive(),
		StartSemester:            in.GetStartSemester(),
		OperationalPermitNumber:  in.GetOperationalPermitNumber(),
		OperationalPermitDate:    operationalPermitDate,
		OperationalPermitDueDate: operationalPermitDueDate,
		HeadLecturerId:           in.GetHeadLecturerId(),
		OperatorName:             in.GetOperatorName(),
		OperatorPhoneNumber:      in.GetOperatorPhoneNumber(),
	}

	errs = f.StudyProgramService.Create(ctx, data)
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
			Message: "Create Study Program",
			Status:  http.StatusCreated,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusCreated, &result)
}

func (f studyProgramHandler) Update(w http.ResponseWriter, r *http.Request) {
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
	defer f.AdminActivityLogService.Create(ctx, r, time.Now(), "Study Program", "Update", &in)

	establishmentDate, errs := utils.StringToTime(in.GetEstablishmentDate())
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
	operationalPermitDate, errs := utils.StringToTime(in.GetOperationalPermitDate())
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
	operationalPermitDueDate, errs := utils.StringToTime(in.GetOperationalPermitDueDate())
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

	data := objects.UpdateStudyProgram{
		Id:                       in.GetId(),
		DiktiStudyProgramId:      in.GetDiktiStudyProgramId(),
		Name:                     in.GetName(),
		EnglishName:              in.GetEnglishName(),
		ShortName:                in.GetShortName(),
		EnglishShortName:         in.GetEnglishShortName(),
		AdministrativeUnit:       in.GetAdministrativeUnit(),
		MajorId:                  in.GetMajorId(),
		Address:                  in.GetAddress(),
		PhoneNumber:              in.GetPhoneNumber(),
		Fax:                      in.GetFax(),
		Email:                    in.GetEmail(),
		Website:                  in.GetWebsite(),
		ContactPerson:            in.GetContactPerson(),
		CuriculumReviewFrequency: in.GetCuriculumReviewFrequency(),
		CuriculumReviewMethod:    in.GetCuriculumReviewMethod(),
		EstablishmentDate:        establishmentDate,
		IsActive:                 in.GetIsActive(),
		StartSemester:            in.GetStartSemester(),
		OperationalPermitNumber:  in.GetOperationalPermitNumber(),
		OperationalPermitDate:    operationalPermitDate,
		OperationalPermitDueDate: operationalPermitDueDate,
		HeadLecturerId:           in.GetHeadLecturerId(),
		OperatorName:             in.GetOperatorName(),
		OperatorPhoneNumber:      in.GetOperatorPhoneNumber(),
	}
	errs = f.StudyProgramService.Update(ctx, data)
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
			Message: "Update Study Program",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (f studyProgramHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
	defer f.AdminActivityLogService.Create(ctx, r, time.Now(), "Study Program", "Delete", nil)

	errs := f.StudyProgramService.Delete(ctx, in.GetId())
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
			Message: "Delete Study Program",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
