package study_program

import (
	"net/http"

	"github.com/gorilla/schema"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/service"
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

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}
	data, errs := f.StudyProgramService.GetList(ctx, paginationData, "", false)
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
			StudyLevelName:        v.StudyLevelName,
			DiktiStudyProgramType: v.DiktiStudyProgramType,
			DiktiStudyProgramCode: v.DiktiStudyProgramCode,
			Accreditation:         utils.NullStringScan(v.Accreditation),
			ActiveCurriculumYear:  utils.NullStringScan(v.ActiveCurriculumYear),
			Degree:                utils.NullStringScan(v.Degree),
			ShortDegree:           utils.NullStringScan(v.ShortDegree),
			EnglishDegree:         utils.NullStringScan(v.EnglishDegree),
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
			MinimumGraduationCredit:       data.MinimumGraduationCredit,
			MinimumThesisCredit:           data.MinimumThesisCredit,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
