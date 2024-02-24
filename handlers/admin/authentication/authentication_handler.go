package authentication

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/schema"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/service"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/utils"
	"github.com/sirupsen/logrus"
)

type authenticationHandler struct {
	*service.ServiceCtx
}

func (a authenticationHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Authentication", "Get Detail", nil)

	data, errs := a.AuthenticationService.GetDetail(ctx, in.GetId())
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
			Message: "Get Detail Authentication",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &GetDetailResponseData{
			Id:                 data.Id,
			Username:           data.Username,
			AuthenticationType: data.AuthenticationType,
			AdminId:            utils.NullStringScan(data.AdminId),
			AdminName:          utils.NullStringScan(data.AdminName),
			LecturerId:         utils.NullStringScan(data.LecturerId),
			LecturerName:       utils.NullStringScan(data.LecturerName),
			StudentId:          utils.NullStringScan(data.StudentId),
			StudentName:        utils.NullStringScan(data.StudentName),
			IsActive:           data.IsActive,
			SuspensionRemarks:  utils.NullStringScan(data.SuspensionRemarks),
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a authenticationHandler) Create(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Authentication", "Create", &in)

	data := objects.CreateAuthentication{
		AuthenticationType: in.GetAuthenticationType(),
		UserId:             in.GetUserId(),
	}

	resultData, errs := a.AuthenticationService.Create(ctx, data)
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
			Message: "Create Authentication",
			Status:  http.StatusCreated,
			Code:    constants.SuccessCode,
		},
		Data: &CreateResponseData{
			Username: resultData.Username,
			Password: resultData.Password,
		},
	}
	utils.JSONResponse(w, http.StatusCreated, &result)
}

func (a authenticationHandler) BulkCreate(w http.ResponseWriter, r *http.Request) {
	var result BulkCreateResponse

	ctx := r.Context()
	var in BulkCreateRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = BulkCreateResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Authentication", "Bulk Create", &in)

	data := objects.BulkCreateAuthentication{
		AuthenticationType: in.GetAuthenticationType(),
		UserIds:            in.GetUserIds(),
	}

	resultData, errs := a.AuthenticationService.BulkCreate(ctx, data)
	if errs != nil {
		utils.PrintError(*errs)
		result = BulkCreateResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	results := []*BulkCreateResponseData{}
	for _, v := range resultData {
		results = append(results, &BulkCreateResponseData{
			UserId:   v.UserId,
			Name:     v.Name,
			Username: v.Username,
			Password: v.Password,
		})
	}

	result = BulkCreateResponse{
		Meta: &Meta{
			Message: "Bulk Create Authentication",
			Status:  http.StatusCreated,
			Code:    constants.SuccessCode,
		},
		Data: results,
	}
	utils.JSONResponse(w, http.StatusCreated, &result)
}

func (a authenticationHandler) Update(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Authentication", "Update", &in)

	data := objects.UpdateAuthentication{
		Id:                in.GetId(),
		IsActive:          in.GetIsActive(),
		SuspensionRemarks: in.GetSuspensionRemarks(),
	}
	errs := a.AuthenticationService.Update(ctx, data)
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
			Message: "Update Authentication",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a authenticationHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Authentication", "Delete", nil)

	errs := a.AuthenticationService.Delete(ctx, in.GetId())
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
			Message: "Delete Authentication",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a authenticationHandler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	var result UpdatePasswordResponse

	ctx := r.Context()
	var in UpdatePasswordRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = UpdatePasswordResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Authentication", "Update Password", nil)

	newPassword, errs := a.AuthenticationService.UpdatePasswordByAdmin(ctx, in.GetId())
	if errs != nil {
		utils.PrintError(*errs)
		result = UpdatePasswordResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = UpdatePasswordResponse{
		Meta: &Meta{
			Message: "Update Password Authentication",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &UpdatePasswordResponseData{
			NewPassword: newPassword,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
