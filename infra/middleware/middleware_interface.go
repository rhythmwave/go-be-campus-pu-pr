package middleware

import (
	"net/http"

	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
)

type Meta struct {
	Message string `json:"message"`
	Status  uint32 `json:"status"`
	Code    string `json:"code"`
}

type ErrorResponse struct {
	Meta *Meta       `json:"meta"`
	Data interface{} `json:"data"`
}

type MiddlewareInterface interface {
	AdminAccessToken(handlerFunc http.Handler) http.Handler
	CareerAccessToken(handlerFunc http.Handler) http.Handler
	GeneralAccessToken(handlerFunc http.Handler) http.Handler
	LecturerAccessToken(handlerFunc http.Handler) http.Handler
	PmbAccessToken(handlerFunc http.Handler) http.Handler
	PermissionMiddleware(handlerFunc http.Handler) http.Handler
	RootAccessToken(handlerFunc http.Handler) http.Handler
	StudentAccessToken(handlerFunc http.Handler) http.Handler
}

func AccessTokenMiddleware(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) MiddlewareInterface {
	return &accessTokenMiddleware{
		repoCtx,
		infraCtx,
	}
}
