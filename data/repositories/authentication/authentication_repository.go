package authentication

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/utils"
)

type authenticationRepository struct {
	*db.DB
}

func mapQueryFilterGetByUsername(isActive *bool, params *[]interface{}) string {
	filterArray := []string{
		"a.username ILIKE $%d",
	}
	filterParams := *params

	if isActive != nil {
		filterArray = append(filterArray, "a.is_active = $%d")
		filterParams = append(filterParams, *isActive)
	}

	result := strings.Join(filterArray, " AND ")
	args := []interface{}{}
	for i := 0; i < len(filterParams); i++ {
		args = append(args, i+1)
	}
	result = fmt.Sprintf(result, args...)
	if result != "" {
		result = fmt.Sprintf("WHERE %s", result)
	}

	*params = filterParams

	return result
}

func mapQueryFilterGetById(isActive *bool, params *[]interface{}) string {
	filterArray := []string{
		"(a.admin_id = $%d OR a.lecturer_id = $%d OR a.student_id = $%d)",
	}
	filterParams := *params

	if isActive != nil {
		filterArray = append(filterArray, "a.is_active = $%d")
		filterParams = append(filterParams, *isActive)
	}

	result := strings.Join(filterArray, " AND ")
	args := []interface{}{}
	for i := 0; i < len(filterParams); i++ {
		args = append(args, i+1)
	}
	result = fmt.Sprintf(result, args...)
	if result != "" {
		result = fmt.Sprintf("WHERE %s", result)
	}

	*params = filterParams

	return result
}

func (a authenticationRepository) GetByUsername(ctx context.Context, tx *sqlx.Tx, username string, isActive *bool) (models.GetAuthentication, *constants.ErrorResponse) {
	results := []models.GetAuthentication{}

	params := []interface{}{
		username,
	}
	query := fmt.Sprintf("%s %s", getByUsernameQuery, mapQueryFilterGetByUsername(isActive, &params))
	err := tx.SelectContext(
		ctx,
		&results,
		query,
		params...,
	)
	if err != nil {
		return models.GetAuthentication{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetAuthentication{}, utils.ErrDataNotFound("user")
	}

	return results[0], nil
}

func (a authenticationRepository) GetById(ctx context.Context, tx *sqlx.Tx, id string) (models.GetAuthentication, *constants.ErrorResponse) {
	results := []models.GetAuthentication{}

	err := tx.SelectContext(
		ctx,
		&results,
		getByIdQuery,
		id,
	)
	if err != nil {
		return models.GetAuthentication{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetAuthentication{}, utils.ErrDataNotFound("user")
	}

	return results[0], nil
}

func (a authenticationRepository) GetByUserId(ctx context.Context, tx *sqlx.Tx, userId string, isActive *bool) (models.GetAuthentication, *constants.ErrorResponse) {
	results := []models.GetAuthentication{}

	params := []interface{}{
		userId,
		userId,
		userId,
	}
	query := fmt.Sprintf("%s %s", getByUserIdQuery, mapQueryFilterGetById(isActive, &params))
	err := tx.SelectContext(
		ctx,
		&results,
		query,
		params...,
	)
	if err != nil {
		return models.GetAuthentication{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetAuthentication{}, utils.ErrDataNotFound("user")
	}

	return results[0], nil
}

func (a authenticationRepository) Create(ctx context.Context, tx *sqlx.Tx, data []models.CreateAuthentication) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		createQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (a authenticationRepository) UpdatePassword(ctx context.Context, tx *sqlx.Tx, hashedPassword, userId string) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		updatePasswordQuery,
		hashedPassword,
		userId,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (a authenticationRepository) UpdateActivation(ctx context.Context, tx *sqlx.Tx, userId string, isActive bool, suspensionRemarks string) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		updateActivationQuery,
		isActive,
		utils.NewNullString(suspensionRemarks),
		userId,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (a authenticationRepository) Delete(ctx context.Context, tx *sqlx.Tx, userId string) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		deleteQuery,
		userId,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (a authenticationRepository) UpdateIdSso(ctx context.Context, tx *sqlx.Tx, id, idSso string) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		updateIdSsoQuery,
		idSso,
		id,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (a authenticationRepository) UpdateSsoRefreshToken(ctx context.Context, tx *sqlx.Tx, id string, ssoRefreshToken sql.NullString) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		updateSsoRefreshTokenQuery,
		ssoRefreshToken,
		id,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}
