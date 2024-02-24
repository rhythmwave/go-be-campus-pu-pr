package auth

import (
	"encoding/json"
	"net/http"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/service"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/utils"
	"github.com/sirupsen/logrus"
)

type authHandler struct {
	*service.ServiceCtx
}

func (a authHandler) Login(w http.ResponseWriter, r *http.Request) {
	var result LoginResponse

	ctx := r.Context()
	var in LoginRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = LoginResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}

	data := objects.LoginRequest{
		Username: in.GetUsername(),
		Password: in.GetPassword(),
	}
	resultData, errs := a.AuthenticationService.Login(ctx, data)
	if errs != nil {
		utils.PrintError(*errs)
		result = LoginResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = LoginResponse{
		Meta: &Meta{
			Message: "Login",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &LoginResponseData{
			AccessToken:     resultData.AccessToken,
			RefreshToken:    resultData.RefreshToken,
			AppType:         resultData.AppType,
			ExpiryTime:      resultData.ExpiryTime.Format(constants.DateRFC),
			PermissionNames: resultData.PermissionNames,
			Name:            resultData.Name,
			Username:        resultData.Username,
			AdminRoleName:   resultData.AdminRoleName,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)

}

func (a authHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var result LoginResponse
	ctx := r.Context()

	resultData, errs := a.AuthenticationService.RefreshToken(ctx)
	if errs != nil {
		utils.PrintError(*errs)
		result = LoginResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = LoginResponse{
		Meta: &Meta{
			Message: "Refresh Token",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &LoginResponseData{
			AccessToken:     resultData.AccessToken,
			RefreshToken:    resultData.RefreshToken,
			AppType:         resultData.AppType,
			ExpiryTime:      resultData.ExpiryTime.Format(constants.DateRFC),
			PermissionNames: resultData.PermissionNames,
			Name:            resultData.Name,
			Username:        resultData.Username,
			AdminRoleName:   resultData.AdminRoleName,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a authHandler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
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

	data := objects.UpdatePasswordRequest{
		OldPassword: in.GetOldPassword(),
		NewPassword: in.GetNewPassword(),
	}

	errs := a.AuthenticationService.UpdatePassword(ctx, data)
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
			Message: "Update Password",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}

	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a authHandler) Logout(w http.ResponseWriter, r *http.Request) {
	var result LogoutResponse
	ctx := r.Context()

	errs := a.AuthenticationService.Logout(ctx)
	if errs != nil {
		utils.PrintError(*errs)
		result = LogoutResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = LogoutResponse{
		Meta: &Meta{
			Message: "Logout",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}

	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a authHandler) GetSsoAuth(w http.ResponseWriter, r *http.Request) {
	var result GetSsoAuthResponse
	ctx := r.Context()

	data, errs := a.AuthenticationService.GetSsoAuth(ctx)
	if errs != nil {
		utils.PrintError(*errs)
		result = GetSsoAuthResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = GetSsoAuthResponse{
		Meta: &Meta{
			Message: "Get Sso Auth",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &GetSsoAuthResponseData{
			Url:         data.Url,
			AppId:       data.AppId,
			FrontendUrl: data.FrontendUrl,
		},
	}

	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a authHandler) LoginWithSso(w http.ResponseWriter, r *http.Request) {
	var result LoginResponse

	ctx := r.Context()
	var in LoginWithSsoRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = LoginResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}

	resultData, errs := a.AuthenticationService.LoginWithSso(ctx, in.GetUsername(), in.GetSsoCode())
	if errs != nil {
		utils.PrintError(*errs)
		result = LoginResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = LoginResponse{
		Meta: &Meta{
			Message: "Login With SSO",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &LoginResponseData{
			AccessToken:     resultData.AccessToken,
			RefreshToken:    resultData.RefreshToken,
			AppType:         resultData.AppType,
			ExpiryTime:      resultData.ExpiryTime.Format(constants.DateRFC),
			PermissionNames: resultData.PermissionNames,
			Name:            resultData.Name,
			Username:        resultData.Username,
			AdminRoleName:   resultData.AdminRoleName,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
