package authentication

import (
	"context"
	"time"

	"github.com/sccicitb/pupr-backend/constants"
	appConstants "github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	middleware "github.com/sccicitb/pupr-backend/infra/base_middleware"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/jwt_object"
	"github.com/sccicitb/pupr-backend/utils"
)

func (a authenticationService) generateToken(ctx context.Context, userData models.GetAuthentication, permissionData []models.GetPermissionByRoleIds) (objects.LoginResponse, *constants.ErrorResponse) {
	var result objects.LoginResponse

	var claimsId string
	var name string
	var role string
	if userData.AdminId != nil {
		claimsId = utils.NullStringScan(userData.AdminId)
		name = utils.NullStringScan(userData.AdminName)
		if userData.AdminRoleId != nil {
			role = appConstants.AppTypeAdmin
		} else {
			role = appConstants.AppTypeRoot
		}
	} else if userData.LecturerId != nil {
		claimsId = utils.NullStringScan(userData.LecturerId)
		name = utils.NullStringScan(userData.LecturerName)
		role = appConstants.AppTypeLecturer
	} else if userData.StudentId != nil {
		claimsId = utils.NullStringScan(userData.StudentId)
		name = utils.NullStringScan(userData.StudentName)
		role = appConstants.AppTypeStudent
	}

	permissions := []string{}
	for _, v := range permissionData {
		permissions = append(permissions, v.Name)
	}

	dataJwt := jwt_object.JWTRequest{
		ID:          claimsId,
		Name:        name,
		Email:       userData.Username,
		Role:        role,
		Permissions: permissions,
	}
	expireTimeAccessToken := time.Duration(a.Config.JWTConfig.ExpireAccessToken) * time.Minute
	expireTimeRefreshToken := time.Duration(a.Config.JWTConfig.ExpireRefreshToken) * time.Hour * 24 * 30

	accessToken, errs := a.Jwt.GenerateJWTToken(ctx, dataJwt, expireTimeAccessToken, appConstants.AppName, middleware.AuthKeyAccessToken, "")
	if errs != nil {
		return result, errs
	}
	refreshToken, errs := a.Jwt.GenerateJWTToken(ctx, dataJwt, expireTimeRefreshToken, appConstants.AppName, middleware.AuthKeyRefreshToken, "")
	if errs != nil {
		return result, errs
	}

	result = objects.LoginResponse{
		AccessToken:     accessToken,
		RefreshToken:    refreshToken,
		AppType:         role,
		ExpiryTime:      time.Now().Add(expireTimeAccessToken),
		PermissionNames: permissions,
		Name:            name,
		Username:        userData.Username,
		AdminRoleName:   utils.NullStringScan(userData.AdminRoleName),
	}

	return result, nil
}

func getAuthenticationType(resultData models.GetAuthentication) string {
	switch true {
	case resultData.AdminId != nil:
		return appConstants.AuthenticationAdmin
	case resultData.LecturerId != nil:
		return appConstants.AuthenticationLecturer
	case resultData.StudentId != nil:
		return appConstants.AuthenticationStudent
	}

	return ""
}
