package location

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/service"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
	"github.com/sirupsen/logrus"
)

type locationHandler struct {
	*service.ServiceCtx
}

func (p locationHandler) GetListCountry(w http.ResponseWriter, r *http.Request) {
	var result GetListCountryResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetListCountryRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetListCountryResponse{
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
	data, errs := p.LocationService.GetListCountry(ctx, paginationData)
	if errs != nil {
		utils.PrintError(*errs)
		result = GetListCountryResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*GetListCountryResponseData{}
	for _, v := range data.Data {
		resultData = append(resultData, &GetListCountryResponseData{
			Id:   v.Id,
			Name: v.Name,
		})
	}

	result = GetListCountryResponse{
		Meta: &Meta{
			Message: "Get List Country",
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

func (p locationHandler) GetListProvince(w http.ResponseWriter, r *http.Request) {
	var result GetListProvinceResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetListProvinceRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetListProvinceResponse{
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
	data, errs := p.LocationService.GetListProvince(ctx, paginationData, in.GetCountryId())
	if errs != nil {
		utils.PrintError(*errs)
		result = GetListProvinceResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*GetListProvinceResponseData{}
	for _, v := range data.Data {
		resultData = append(resultData, &GetListProvinceResponseData{
			Id:   v.Id,
			Name: v.Name,
		})
	}

	result = GetListProvinceResponse{
		Meta: &Meta{
			Message: "Get List Province",
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

func (p locationHandler) GetListRegency(w http.ResponseWriter, r *http.Request) {
	var result GetListRegencyResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetListRegencyRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetListRegencyResponse{
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
	data, errs := p.LocationService.GetListRegency(ctx, paginationData, in.GetProvinceId())
	if errs != nil {
		utils.PrintError(*errs)
		result = GetListRegencyResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*GetListRegencyResponseData{}
	for _, v := range data.Data {
		resultData = append(resultData, &GetListRegencyResponseData{
			Id:   v.Id,
			Name: v.Name,
		})
	}

	result = GetListRegencyResponse{
		Meta: &Meta{
			Message: "Get List Regency",
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

func (p locationHandler) GetListDistrict(w http.ResponseWriter, r *http.Request) {
	var result GetListDistrictResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetListDistrictRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetListDistrictResponse{
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
	data, errs := p.LocationService.GetListDistrict(ctx, paginationData, in.GetRegencyId())
	if errs != nil {
		utils.PrintError(*errs)
		result = GetListDistrictResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*GetListDistrictResponseData{}
	for _, v := range data.Data {
		resultData = append(resultData, &GetListDistrictResponseData{
			Id:   v.Id,
			Name: v.Name,
		})
	}

	result = GetListDistrictResponse{
		Meta: &Meta{
			Message: "Get List District",
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

func (p locationHandler) GetListVillage(w http.ResponseWriter, r *http.Request) {
	var result GetListVillageResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetListVillageRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetListVillageResponse{
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
	data, errs := p.LocationService.GetListVillage(ctx, paginationData, in.GetDistrictId())
	if errs != nil {
		utils.PrintError(*errs)
		result = GetListVillageResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*GetListVillageResponseData{}
	for _, v := range data.Data {
		resultData = append(resultData, &GetListVillageResponseData{
			Id:   v.Id,
			Name: v.Name,
		})
	}

	result = GetListVillageResponse{
		Meta: &Meta{
			Message: "Get List Village",
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

func (a locationHandler) TempCreateData(w http.ResponseWriter, r *http.Request) {
	var result TempCreateDataResponse

	ctx := r.Context()
	var in TempCreateDataRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = TempCreateDataResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}

	data := objects.TempCreateData{
		Title: in.GetTitle(),
		Body:  in.GetBody(),
	}

	errs := a.LocationService.TempCreateData(ctx, data)
	if errs != nil {
		utils.PrintError(*errs)
		result = TempCreateDataResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = TempCreateDataResponse{
		Meta: &Meta{
			Message: "TempCreateData",
			Status:  http.StatusCreated,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusCreated, &result)
}

func (p locationHandler) TempGetData(w http.ResponseWriter, r *http.Request) {
	var result TempGetDataResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in TempGetDataRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = TempGetDataResponse{
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
		Limit: in.GetLimit(),
		Page:  in.GetPage(),
	}
	data, errs := p.LocationService.TempGetData(ctx, paginationData)
	if errs != nil {
		utils.PrintError(*errs)
		result = TempGetDataResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*TempGetDataResponseData{}
	for _, v := range data.Data {
		resultData = append(resultData, &TempGetDataResponseData{
			Id:    v.Id,
			Title: v.Title,
			Body:  v.Body,
		})
	}

	result = TempGetDataResponse{
		Meta: &Meta{
			Message: "Temp Get Data",
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
