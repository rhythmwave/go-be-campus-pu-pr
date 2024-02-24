package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/sccicitb/pupr-backend/constants"
	appConstants "github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects/jwt_object"
	"github.com/sccicitb/pupr-backend/utils"
	"google.golang.org/grpc/codes"
)

type accessTokenMiddleware struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (a accessTokenMiddleware) extractAccessToken(ctx context.Context, r *http.Request) (*jwt_object.JWTClaims, *constants.ErrorResponse) {
	var claims *jwt_object.JWTClaims

	tokenHeader := r.Header.Get(constants.Authorization)
	if tokenHeader == "" {
		return claims, constants.Error(http.StatusForbidden, codes.PermissionDenied, constants.ErrUnknownCustomCode, "authorization header not exists")
	}

	claims, errs := a.Jwt.ExtractJWTClaims(ctx, tokenHeader, appConstants.AppName)
	if errs != nil {
		return claims, errs
	}

	return claims, nil
}

func (a accessTokenMiddleware) CareerAccessToken(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var apiKey string
		if a.Config.AppConfig[appConstants.CareerApiKeyConfig] != nil {
			apiKey = a.Config.AppConfig[appConstants.CareerApiKeyConfig].(string)
		}
		if apiKey == "" {
			errs := constants.ErrUnknown
			utils.PrintError(*errs)
			result := ErrorResponse{
				Meta: &Meta{
					Message: errs.Err.Error(),
					Status:  uint32(errs.HttpCode),
					Code:    errs.CustomCode,
				},
			}
			utils.JSONResponse(w, errs.HttpCode, &result)
			return
		}

		tokenHeader := r.Header.Get(constants.Authorization)
		if tokenHeader == "" {
			errs := constants.Error(http.StatusForbidden, codes.PermissionDenied, constants.ErrUnknownCustomCode, "authorization header not exists")
			utils.PrintError(*errs)
			result := ErrorResponse{
				Meta: &Meta{
					Message: errs.Err.Error(),
					Status:  uint32(errs.HttpCode),
					Code:    errs.CustomCode,
				},
			}
			utils.JSONResponse(w, errs.HttpCode, &result)
			return
		}

		splitToken := strings.Split(tokenHeader, constants.Bearer)
		if len(splitToken) != 2 {
			errs := constants.ErrTokenIsRequired
			utils.PrintError(*errs)
			result := ErrorResponse{
				Meta: &Meta{
					Message: errs.Err.Error(),
					Status:  uint32(errs.HttpCode),
					Code:    errs.CustomCode,
				},
			}
			utils.JSONResponse(w, errs.HttpCode, &result)
			return
		}

		reqToken := strings.TrimSpace(splitToken[1])

		if reqToken != apiKey {
			errs := constants.ErrUnauthenticated
			utils.PrintError(*errs)
			result := ErrorResponse{
				Meta: &Meta{
					Message: errs.Err.Error(),
					Status:  uint32(errs.HttpCode),
					Code:    errs.CustomCode,
				},
			}
			utils.JSONResponse(w, errs.HttpCode, &result)
			return
		}

		handlerFunc.ServeHTTP(w, r)
	})
}

func (a accessTokenMiddleware) PmbAccessToken(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var apiKey string
		if a.Config.AppConfig[appConstants.PmbApiKeyConfig] != nil {
			apiKey = a.Config.AppConfig[appConstants.PmbApiKeyConfig].(string)
		}
		if apiKey == "" {
			errs := constants.ErrUnknown
			utils.PrintError(*errs)
			result := ErrorResponse{
				Meta: &Meta{
					Message: errs.Err.Error(),
					Status:  uint32(errs.HttpCode),
					Code:    errs.CustomCode,
				},
			}
			utils.JSONResponse(w, errs.HttpCode, &result)
			return
		}

		tokenHeader := r.Header.Get(constants.Authorization)
		if tokenHeader == "" {
			errs := constants.Error(http.StatusForbidden, codes.PermissionDenied, constants.ErrUnknownCustomCode, "authorization header not exists")
			utils.PrintError(*errs)
			result := ErrorResponse{
				Meta: &Meta{
					Message: errs.Err.Error(),
					Status:  uint32(errs.HttpCode),
					Code:    errs.CustomCode,
				},
			}
			utils.JSONResponse(w, errs.HttpCode, &result)
			return
		}

		splitToken := strings.Split(tokenHeader, constants.Bearer)
		if len(splitToken) != 2 {
			errs := constants.ErrTokenIsRequired
			utils.PrintError(*errs)
			result := ErrorResponse{
				Meta: &Meta{
					Message: errs.Err.Error(),
					Status:  uint32(errs.HttpCode),
					Code:    errs.CustomCode,
				},
			}
			utils.JSONResponse(w, errs.HttpCode, &result)
			return
		}

		reqToken := strings.TrimSpace(splitToken[1])

		if reqToken != apiKey {
			errs := constants.ErrUnauthenticated
			utils.PrintError(*errs)
			result := ErrorResponse{
				Meta: &Meta{
					Message: errs.Err.Error(),
					Status:  uint32(errs.HttpCode),
					Code:    errs.CustomCode,
				},
			}
			utils.JSONResponse(w, errs.HttpCode, &result)
			return
		}

		handlerFunc.ServeHTTP(w, r)
	})
}

func (a accessTokenMiddleware) AdminAccessToken(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		claims, errs := a.extractAccessToken(ctx, r)
		if errs != nil {
			utils.PrintError(*errs)
			result := ErrorResponse{
				Meta: &Meta{
					Message: errs.Err.Error(),
					Status:  uint32(errs.HttpCode),
					Code:    errs.CustomCode,
				},
			}
			utils.JSONResponse(w, errs.HttpCode, &result)
			return
		}

		if claims.Role != appConstants.AppTypeAdmin {
			result := ErrorResponse{
				Meta: &Meta{
					Message: "you don't have permission to access this menu",
					Status:  uint32(http.StatusForbidden),
					Code:    constants.DefaultCustomErrorCode,
				},
			}
			utils.JSONResponse(w, http.StatusForbidden, &result)
			return
		}

		ctx = context.WithValue(ctx, constants.ClaimsContextKey, claims)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func (a accessTokenMiddleware) GeneralAccessToken(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		claims, errs := a.extractAccessToken(ctx, r)
		if errs != nil {
			utils.PrintError(*errs)
			result := ErrorResponse{
				Meta: &Meta{
					Message: errs.Err.Error(),
					Status:  uint32(errs.HttpCode),
					Code:    errs.CustomCode,
				},
			}
			utils.JSONResponse(w, errs.HttpCode, &result)
			return
		}

		ctx = context.WithValue(ctx, constants.ClaimsContextKey, claims)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func (a accessTokenMiddleware) LecturerAccessToken(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		claims, errs := a.extractAccessToken(ctx, r)
		if errs != nil {
			utils.PrintError(*errs)
			result := ErrorResponse{
				Meta: &Meta{
					Message: errs.Err.Error(),
					Status:  uint32(errs.HttpCode),
					Code:    errs.CustomCode,
				},
			}
			utils.JSONResponse(w, errs.HttpCode, &result)
			return
		}

		if claims.Role != appConstants.AppTypeLecturer {
			result := ErrorResponse{
				Meta: &Meta{
					Message: "you don't have permission to access this menu",
					Status:  uint32(http.StatusForbidden),
					Code:    constants.DefaultCustomErrorCode,
				},
			}
			utils.JSONResponse(w, http.StatusForbidden, &result)
			return
		}

		ctx = context.WithValue(ctx, constants.ClaimsContextKey, claims)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func (a accessTokenMiddleware) PermissionMiddleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		claims, errs := utils.GetJWTClaimsFromContext(ctx)
		if errs != nil {
			utils.PrintError(*errs)
			result := ErrorResponse{
				Meta: &Meta{
					Message: errs.Err.Error(),
					Status:  uint32(errs.HttpCode),
					Code:    errs.CustomCode,
				},
			}
			utils.JSONResponse(w, errs.HttpCode, &result)
			return
		}
		permission := fmt.Sprintf("%v", ctx.Value(appConstants.PermissionContextKey))

		hasPermission := utils.InArrayExist(permission, claims.Permissions)
		if !hasPermission {
			result := ErrorResponse{
				Meta: &Meta{
					Message: "you don't have permission to access this menu",
					Status:  uint32(http.StatusForbidden),
					Code:    constants.DefaultCustomErrorCode,
				},
			}
			utils.JSONResponse(w, http.StatusForbidden, &result)
			return
		}

		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func (a accessTokenMiddleware) RootAccessToken(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		claims, errs := a.extractAccessToken(ctx, r)
		if errs != nil {
			utils.PrintError(*errs)
			result := ErrorResponse{
				Meta: &Meta{
					Message: errs.Err.Error(),
					Status:  uint32(errs.HttpCode),
					Code:    errs.CustomCode,
				},
			}
			utils.JSONResponse(w, errs.HttpCode, &result)
			return
		}

		if claims.Role != appConstants.AppTypeRoot {
			result := ErrorResponse{
				Meta: &Meta{
					Message: "you don't have permission to access this menu",
					Status:  uint32(http.StatusForbidden),
					Code:    constants.DefaultCustomErrorCode,
				},
			}
			utils.JSONResponse(w, http.StatusForbidden, &result)
			return
		}

		ctx = context.WithValue(ctx, constants.ClaimsContextKey, claims)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

func (a accessTokenMiddleware) StudentAccessToken(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		claims, errs := a.extractAccessToken(ctx, r)
		if errs != nil {
			utils.PrintError(*errs)
			result := ErrorResponse{
				Meta: &Meta{
					Message: errs.Err.Error(),
					Status:  uint32(errs.HttpCode),
					Code:    errs.CustomCode,
				},
			}
			utils.JSONResponse(w, errs.HttpCode, &result)
			return
		}

		if claims.Role != appConstants.AppTypeStudent {
			result := ErrorResponse{
				Meta: &Meta{
					Message: "you don't have permission to access this menu",
					Status:  uint32(http.StatusForbidden),
					Code:    constants.DefaultCustomErrorCode,
				},
			}
			utils.JSONResponse(w, http.StatusForbidden, &result)
			return
		}

		ctx = context.WithValue(ctx, constants.ClaimsContextKey, claims)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}
