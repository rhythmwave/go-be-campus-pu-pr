package major

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

type majorHandler struct {
	*service.ServiceCtx
}

func (m majorHandler) GetList(w http.ResponseWriter, r *http.Request) {
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
	defer m.AdminActivityLogService.Create(ctx, r, time.Now(), "Major", "Get List", nil)

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}
	data, errs := m.MajorService.GetList(ctx, paginationData, in.GetFacultyId())
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
			Id:          v.Id,
			FacultyName: v.FacultyName,
			Name:        v.Name,
		})
	}

	result = GetListResponse{
		Meta: &Meta{
			Message: "Get List Major",
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

func (m majorHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
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
	defer m.AdminActivityLogService.Create(ctx, r, time.Now(), "Major", "Get Detail", nil)

	data, errs := m.MajorService.GetDetail(ctx, in.GetId())
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
			Message: "Get Detail Major",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &GetDetailResponseData{
			Id:                        data.Id,
			FacultyId:                 data.FacultyId,
			FacultyName:               data.FacultyName,
			Name:                      data.Name,
			ShortName:                 utils.NullStringScan(data.ShortName),
			EnglishName:               utils.NullStringScan(data.EnglishName),
			EnglishShortName:          utils.NullStringScan(data.EnglishShortName),
			Address:                   data.Address,
			PhoneNumber:               utils.NullStringScan(data.PhoneNumber),
			Fax:                       utils.NullStringScan(data.Fax),
			Email:                     utils.NullStringScan(data.Email),
			ContactPerson:             utils.NullStringScan(data.ContactPerson),
			ExperimentBuildingArea:    utils.NullFloatScan(data.ExperimentBuildingArea),
			LectureHallArea:           utils.NullFloatScan(data.LectureHallArea),
			LectureHallCount:          utils.NullUint32Scan(data.LectureHallCount),
			LaboratoriumArea:          utils.NullFloatScan(data.LaboratoriumArea),
			LaboratoriumCount:         utils.NullUint32Scan(data.LaboratoriumCount),
			PermanentLecturerRoomArea: utils.NullFloatScan(data.PermanentLecturerRoomArea),
			AdministrationRoomArea:    utils.NullFloatScan(data.AdministrationRoomArea),
			BookCount:                 utils.NullUint32Scan(data.BookCount),
			BookCopyCount:             utils.NullUint32Scan(data.BookCopyCount),
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (m majorHandler) Create(w http.ResponseWriter, r *http.Request) {
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
	defer m.AdminActivityLogService.Create(ctx, r, time.Now(), "Major", "Create", &in)

	data := objects.CreateMajor{
		FacultyId:                 in.GetFacultyId(),
		Name:                      in.GetName(),
		ShortName:                 in.GetShortName(),
		EnglishName:               in.GetEnglishName(),
		EnglishShortName:          in.GetEnglishShortName(),
		Address:                   in.GetAddress(),
		PhoneNumber:               in.GetPhoneNumber(),
		Fax:                       in.GetFax(),
		Email:                     in.GetEmail(),
		ContactPerson:             in.GetContactPerson(),
		ExperimentBuildingArea:    in.GetExperimentBuildingArea(),
		LectureHallArea:           in.GetLectureHallArea(),
		LectureHallCount:          in.GetLectureHallCount(),
		LaboratoriumArea:          in.GetLaboratoriumArea(),
		LaboratoriumCount:         in.GetLaboratoriumCount(),
		PermanentLecturerRoomArea: in.GetPermanentLecturerRoomArea(),
		AdministrationRoomArea:    in.GetAdministrationRoomArea(),
		BookCount:                 in.GetBookCount(),
		BookCopyCount:             in.GetBookCopyCount(),
	}

	errs := m.MajorService.Create(ctx, data)
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
			Message: "Create Major",
			Status:  http.StatusCreated,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusCreated, &result)
}

func (m majorHandler) Update(w http.ResponseWriter, r *http.Request) {
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
	defer m.AdminActivityLogService.Create(ctx, r, time.Now(), "Major", "Update", &in)

	data := objects.UpdateMajor{
		Id:                        in.GetId(),
		FacultyId:                 in.GetFacultyId(),
		Name:                      in.GetName(),
		ShortName:                 in.GetShortName(),
		EnglishName:               in.GetEnglishName(),
		EnglishShortName:          in.GetEnglishShortName(),
		Address:                   in.GetAddress(),
		PhoneNumber:               in.GetPhoneNumber(),
		Fax:                       in.GetFax(),
		Email:                     in.GetEmail(),
		ContactPerson:             in.GetContactPerson(),
		ExperimentBuildingArea:    in.GetExperimentBuildingArea(),
		LectureHallArea:           in.GetLectureHallArea(),
		LectureHallCount:          in.GetLectureHallCount(),
		LaboratoriumArea:          in.GetLaboratoriumArea(),
		LaboratoriumCount:         in.GetLaboratoriumCount(),
		PermanentLecturerRoomArea: in.GetPermanentLecturerRoomArea(),
		AdministrationRoomArea:    in.GetAdministrationRoomArea(),
		BookCount:                 in.GetBookCount(),
		BookCopyCount:             in.GetBookCopyCount(),
	}
	errs := m.MajorService.Update(ctx, data)
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
			Message: "Update Major",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (m majorHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
	defer m.AdminActivityLogService.Create(ctx, r, time.Now(), "Major", "Delete", nil)

	errs := m.MajorService.Delete(ctx, in.GetId())
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
			Message: "Delete Major",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
