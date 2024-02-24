package authentication

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
)

type AuthenticationRepositoryInterface interface {
	GetByUsername(ctx context.Context, tx *sqlx.Tx, username string, isActive *bool) (models.GetAuthentication, *constants.ErrorResponse)
	GetById(ctx context.Context, tx *sqlx.Tx, id string) (models.GetAuthentication, *constants.ErrorResponse)
	GetByUserId(ctx context.Context, tx *sqlx.Tx, userId string, isActive *bool) (models.GetAuthentication, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data []models.CreateAuthentication) *constants.ErrorResponse
	UpdatePassword(ctx context.Context, tx *sqlx.Tx, hashedPassword, userId string) *constants.ErrorResponse
	UpdateActivation(ctx context.Context, tx *sqlx.Tx, userId string, isActive bool, suspensionRemarks string) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, userId string) *constants.ErrorResponse
	UpdateIdSso(ctx context.Context, tx *sqlx.Tx, id, idSso string) *constants.ErrorResponse
	UpdateSsoRefreshToken(ctx context.Context, tx *sqlx.Tx, id string, ssoRefreshToken sql.NullString) *constants.ErrorResponse
}

func NewAuthenticationRepository(db *db.DB) AuthenticationRepositoryInterface {
	return &authenticationRepository{
		db,
	}
}
