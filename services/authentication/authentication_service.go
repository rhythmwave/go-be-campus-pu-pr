package authentication

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/sccicitb/pupr-backend/constants"
	appConstants "github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	middleware "github.com/sccicitb/pupr-backend/infra/base_middleware"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/utils"
)

type authenticationService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (a authenticationService) Login(ctx context.Context, data objects.LoginRequest) (objects.LoginResponse, *constants.ErrorResponse) {
	var result objects.LoginResponse

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	isActive := true
	userData, errs := a.AuthenticationRepo.GetByUsername(ctx, tx, data.Username, &isActive)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}
	isCheckPassword := utils.CheckPasswordHash(data.Password, userData.Password)
	if !isCheckPassword {
		_ = tx.Rollback()
		return result, constants.ErrEmailAndPasswordNotMatch
	}

	permissionData := []models.GetPermissionByRoleIds{}
	if userData.AdminRoleId != nil {
		permissionData, errs = a.PermissionRepo.GetByRoleIds(ctx, tx, []string{*userData.AdminRoleId})
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	}

	result, errs = a.generateToken(ctx, userData, permissionData)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (a authenticationService) LoginWithSso(ctx context.Context, username, ssoCode string) (objects.LoginResponse, *constants.ErrorResponse) {
	var result objects.LoginResponse

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	var ssoBaseUrl string
	if a.Config.AppConfig[appConstants.SsoApiUrlConfig] != nil {
		ssoBaseUrl = a.Config.AppConfig[appConstants.SsoApiUrlConfig].(string)
	}
	if ssoBaseUrl == "" {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	ssoSessionUrl := fmt.Sprintf("%s%s/%s", ssoBaseUrl, "/auth/sess", ssoCode)
	var sessionResult objects.SsoFindSession
	errs := utils.HttpClientDo(http.MethodPost, ssoSessionUrl, nil, nil, &sessionResult)
	if errs != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnauthenticated
	}

	ssoUserInfoUrl := fmt.Sprintf("%s%s", ssoBaseUrl, "/auth/user")
	refreshToken := sessionResult.RefreshToken
	ssoUrlInfoHeaders := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", sessionResult.AccessToken),
	}
	var userInfoResult objects.SsoUserInfo
	errs = utils.HttpClientDo(http.MethodPost, ssoUserInfoUrl, nil, ssoUrlInfoHeaders, &userInfoResult)
	if errs != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnauthenticated
	}

	if sessionResult.Username != username {
		_ = tx.Rollback()
		return result, constants.ErrUnauthenticated
	}

	isActive := true
	userData, errs := a.AuthenticationRepo.GetByUsername(ctx, tx, username, &isActive)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	if userData.IdSso != nil && utils.NullStringScan(userData.IdSso) != userInfoResult.Data.ID {
		_ = tx.Rollback()
		return result, constants.ErrUnauthenticated
	}
	if userData.IdSso == nil {
		errs = a.AuthenticationRepo.UpdateIdSso(ctx, tx, userData.Id, userInfoResult.Data.ID)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	}

	permissionData := []models.GetPermissionByRoleIds{}
	if userData.AdminRoleId != nil {
		permissionData, errs = a.PermissionRepo.GetByRoleIds(ctx, tx, []string{*userData.AdminRoleId})
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	}

	errs = a.AuthenticationRepo.UpdateSsoRefreshToken(ctx, tx, userData.Id, utils.NewNullString(refreshToken))
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result, errs = a.generateToken(ctx, userData, permissionData)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (a authenticationService) RefreshToken(ctx context.Context) (objects.LoginResponse, *constants.ErrorResponse) {
	var result objects.LoginResponse

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	isActive := true
	userData, errs := a.AuthenticationRepo.GetByUserId(ctx, tx, claims.ID, &isActive)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	permissionData := []models.GetPermissionByRoleIds{}
	if userData.AdminRoleId != nil {
		permissionData, errs = a.PermissionRepo.GetByRoleIds(ctx, tx, []string{*userData.AdminRoleId})
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	}

	result, errs = a.generateToken(ctx, userData, permissionData)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (a authenticationService) UpdatePassword(ctx context.Context, data objects.UpdatePasswordRequest) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	isActive := true
	userData, errs := a.AuthenticationRepo.GetByUserId(ctx, tx, claims.ID, &isActive)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	isCheckPassword := utils.CheckPasswordHash(data.OldPassword, userData.Password)
	if !isCheckPassword {
		_ = tx.Rollback()
		return constants.ErrEmailAndPasswordNotMatch
	}

	password, err := utils.HashPassword(data.NewPassword)
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrorInternalServer(err.Error())
	}
	errs = a.AuthenticationRepo.UpdatePassword(ctx, tx, password, claims.ID)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (a authenticationService) Logout(ctx context.Context) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	var ssoBaseUrl string
	var appId string
	if a.Config.AppConfig[appConstants.SsoApiUrlConfig] != nil {
		ssoBaseUrl = a.Config.AppConfig[appConstants.SsoApiUrlConfig].(string)
	}
	if a.Config.AppConfig[appConstants.SsoAppIdConfig] != nil {
		appId = a.Config.AppConfig[appConstants.SsoAppIdConfig].(string)
	}
	if ssoBaseUrl == "" || appId == "" {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	isActive := true
	userData, errs := a.AuthenticationRepo.GetByUserId(ctx, tx, claims.ID, &isActive)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	refreshToken := utils.NullStringScan(userData.SsoRefreshToken)

	if refreshToken != "" {
		// ssoLogoutUrl := fmt.Sprintf("%s%s", ssoBaseUrl, "/auth/logout")
		// body := url.Values{}
		// body.Add("refresh_token", refreshToken)
		// body.Add("clientId", appId)
		// var ssoLogout objects.SsoLogout
		// errs = utils.HttpClientDoUrlEncoded(http.MethodPost, ssoLogoutUrl, nil, nil, body, &ssoLogout)
		// if errs != nil {
		// 	_ = tx.Rollback()
		// 	return errs
		// }
	}

	errs = a.Jwt.DeleteTokenFromRedis(ctx, claims, middleware.AuthKeyAccessToken, appConstants.AppName)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	errs = a.Jwt.DeleteTokenFromRedis(ctx, claims, middleware.AuthKeyRefreshToken, appConstants.AppName)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = a.AuthenticationRepo.UpdateSsoRefreshToken(ctx, tx, userData.Id, utils.NewNullString(""))
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (a authenticationService) GetDetail(ctx context.Context, id string) (objects.GetAuthentication, *constants.ErrorResponse) {
	var result objects.GetAuthentication

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	resultData, errs := a.AuthenticationRepo.GetById(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result = objects.GetAuthentication{
		Id:                 resultData.Id,
		Username:           resultData.Username,
		AuthenticationType: getAuthenticationType(resultData),
		AdminId:            resultData.AdminId,
		AdminName:          resultData.AdminName,
		LecturerId:         resultData.LecturerId,
		LecturerName:       resultData.LecturerName,
		StudentId:          resultData.StudentId,
		StudentName:        resultData.StudentName,
		IsActive:           resultData.IsActive,
		SuspensionRemarks:  resultData.SuspensionRemarks,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (a authenticationService) Create(ctx context.Context, data objects.CreateAuthentication) (objects.CreateAuthenticationResponse, *constants.ErrorResponse) {
	var result objects.CreateAuthenticationResponse

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	var username string
	var lecturerId sql.NullString
	var studentId sql.NullString
	if data.AuthenticationType == appConstants.AuthenticationLecturer {
		lecturerData, errs := a.LecturerRepo.GetDetail(ctx, tx, data.UserId)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
		lecturerId = utils.NewNullString(lecturerData.Id)
		username = lecturerData.IdNationalLecturer
	} else if data.AuthenticationType == appConstants.AuthenticationStudent {
		studentData, errs := a.StudentRepo.GetDetail(ctx, tx, data.UserId, "")
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
		studentId = utils.NewNullString(studentData.Id)
		username = strconv.Itoa(int(studentData.NimNumber))
	} else {
		_ = tx.Rollback()
		return result, appConstants.ErrUneditableAuthentication
	}

	password := strconv.Itoa(utils.RandomInteger(6))
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrorInternalServer(err.Error())
	}

	createData := []models.CreateAuthentication{
		{
			Username:   username,
			Password:   hashedPassword,
			LecturerId: lecturerId,
			StudentId:  studentId,
		},
	}

	errs := a.AuthenticationRepo.Create(ctx, tx, createData)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result = objects.CreateAuthenticationResponse{
		Username: username,
		Password: password,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (a authenticationService) BulkCreate(ctx context.Context, data objects.BulkCreateAuthentication) ([]objects.BulkCreateAuthenticationResponse, *constants.ErrorResponse) {
	results := []objects.BulkCreateAuthenticationResponse{}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return results, constants.ErrUnknown
	}

	createData := []models.CreateAuthentication{}
	if data.AuthenticationType == appConstants.AuthenticationLecturer {
		lecturerData, errs := a.LecturerRepo.GetDetailByIds(ctx, tx, data.UserIds)
		if errs != nil {
			_ = tx.Rollback()
			return results, errs
		}
		if len(lecturerData) == 0 {
			return results, utils.ErrDataNotFound("lecturer")
		}
		for _, v := range lecturerData {
			password := strconv.Itoa(utils.RandomInteger(6))
			hashedPassword, err := utils.HashPassword(password)
			if err != nil {
				_ = tx.Rollback()
				return results, constants.ErrorInternalServer(err.Error())
			}

			results = append(results, objects.BulkCreateAuthenticationResponse{
				UserId:   v.Id,
				Name:     v.Name,
				Username: v.IdNationalLecturer,
				Password: password,
			})

			createData = append(createData, models.CreateAuthentication{
				Username:   v.IdNationalLecturer,
				Password:   hashedPassword,
				LecturerId: utils.NewNullString(v.Id),
			})
		}
	} else if data.AuthenticationType == appConstants.AuthenticationStudent {
		studentData, errs := a.StudentRepo.GetDetailByIds(ctx, tx, data.UserIds)
		if errs != nil {
			_ = tx.Rollback()
			return results, errs
		}
		if len(studentData) == 0 {
			return results, utils.ErrDataNotFound("student")
		}
		for _, v := range studentData {
			username := strconv.Itoa(int(v.NimNumber))
			password := strconv.Itoa(utils.RandomInteger(6))
			hashedPassword, err := utils.HashPassword(password)
			if err != nil {
				_ = tx.Rollback()
				return results, constants.ErrorInternalServer(err.Error())
			}

			results = append(results, objects.BulkCreateAuthenticationResponse{
				UserId:   v.Id,
				Name:     v.Name,
				Username: username,
				Password: password,
			})

			createData = append(createData, models.CreateAuthentication{
				Username:  username,
				Password:  hashedPassword,
				StudentId: utils.NewNullString(v.Id),
			})
		}
	} else {
		_ = tx.Rollback()
		return results, appConstants.ErrUneditableAuthentication
	}

	if len(createData) != 0 {
		errs := a.AuthenticationRepo.Create(ctx, tx, createData)
		if errs != nil {
			_ = tx.Rollback()
			return results, errs
		}
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return results, constants.ErrUnknown
	}

	return results, nil
}

func (a authenticationService) Update(ctx context.Context, data objects.UpdateAuthentication) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	authData, errs := a.AuthenticationRepo.GetById(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	var userId string
	if authData.LecturerId != nil {
		userId = utils.NullStringScan(authData.LecturerId)
	} else if authData.StudentId != nil {
		userId = utils.NullStringScan(authData.StudentId)
	} else {
		_ = tx.Rollback()
		return appConstants.ErrUneditableAuthentication
	}

	var suspensionRemarks string
	if !data.IsActive {
		suspensionRemarks = data.SuspensionRemarks
	}
	errs = a.AuthenticationRepo.UpdateActivation(ctx, tx, userId, data.IsActive, suspensionRemarks)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (a authenticationService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	authData, errs := a.AuthenticationRepo.GetById(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	var userId string
	if authData.LecturerId != nil {
		userId = utils.NullStringScan(authData.LecturerId)
	} else if authData.StudentId != nil {
		userId = utils.NullStringScan(authData.StudentId)
	} else {
		_ = tx.Rollback()
		return appConstants.ErrUneditableAuthentication
	}

	errs = a.AuthenticationRepo.Delete(ctx, tx, userId)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (a authenticationService) UpdatePasswordByAdmin(ctx context.Context, id string) (string, *constants.ErrorResponse) {
	var newPassword string

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return newPassword, constants.ErrUnknown
	}

	authData, errs := a.AuthenticationRepo.GetById(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return newPassword, errs
	}
	var userId string
	if authData.LecturerId != nil {
		userId = utils.NullStringScan(authData.LecturerId)
	} else if authData.StudentId != nil {
		userId = utils.NullStringScan(authData.StudentId)
	} else {
		_ = tx.Rollback()
		return newPassword, appConstants.ErrUneditableAuthentication
	}

	newPassword = strconv.Itoa(utils.RandomInteger(6))
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		_ = tx.Rollback()
		return newPassword, constants.ErrorInternalServer(err.Error())
	}

	errs = a.AuthenticationRepo.UpdatePassword(ctx, tx, hashedPassword, userId)
	if errs != nil {
		_ = tx.Rollback()
		return newPassword, errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return newPassword, constants.ErrUnknown
	}

	return newPassword, nil
}

func (a authenticationService) GetSsoAuth(ctx context.Context) (objects.GetSsoAuthResponse, *constants.ErrorResponse) {
	var result objects.GetSsoAuthResponse

	var url string
	var appId string
	var frontendUrl string
	if a.Config.AppConfig[appConstants.SsoUrlConfig] != nil {
		url = a.Config.AppConfig[appConstants.SsoUrlConfig].(string)
	}
	if a.Config.AppConfig[appConstants.SsoAppIdConfig] != nil {
		appId = a.Config.AppConfig[appConstants.SsoAppIdConfig].(string)
	}
	if a.Config.AppConfig[appConstants.SsoRedirectUrl] != nil {
		frontendUrl = a.Config.AppConfig[appConstants.SsoRedirectUrl].(string)
	}
	if url == "" || appId == "" || frontendUrl == "" {
		return result, constants.ErrUnknown
	}

	result = objects.GetSsoAuthResponse{
		Url:         url,
		AppId:       appId,
		FrontendUrl: frontendUrl,
	}

	return result, nil
}
