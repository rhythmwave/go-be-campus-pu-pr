package authentication

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
)

type AuthenticationServiceInterface interface {
	Login(ctx context.Context, data objects.LoginRequest) (objects.LoginResponse, *constants.ErrorResponse)
	LoginWithSso(ctx context.Context, username, ssoCode string) (objects.LoginResponse, *constants.ErrorResponse)
	RefreshToken(ctx context.Context) (objects.LoginResponse, *constants.ErrorResponse)
	UpdatePassword(ctx context.Context, data objects.UpdatePasswordRequest) *constants.ErrorResponse
	Logout(ctx context.Context) *constants.ErrorResponse
	GetDetail(ctx context.Context, id string) (objects.GetAuthentication, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateAuthentication) (objects.CreateAuthenticationResponse, *constants.ErrorResponse)
	BulkCreate(ctx context.Context, data objects.BulkCreateAuthentication) ([]objects.BulkCreateAuthenticationResponse, *constants.ErrorResponse)
	Update(ctx context.Context, data objects.UpdateAuthentication) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
	UpdatePasswordByAdmin(ctx context.Context, id string) (string, *constants.ErrorResponse)
	GetSsoAuth(ctx context.Context) (objects.GetSsoAuthResponse, *constants.ErrorResponse)
}

func NewAuthenticationService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) AuthenticationServiceInterface {
	return &authenticationService{
		repoCtx,
		infraCtx,
	}
}
