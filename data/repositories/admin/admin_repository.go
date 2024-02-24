package admin

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
)

type adminRepository struct {
	*db.DB
}

func mapQueryFilterGetList(search string, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	if search != "" {
		filterArray = append(filterArray, "a.name ILIKE $%d")
		filterParams = append(filterParams, fmt.Sprintf("%%%s%%", search))
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

func (a adminRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest) ([]models.GetAdmin, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetAdmin{}
	var paginationResult common.Pagination

	params := []interface{}{}
	filterQuery := mapQueryFilterGetList(pagination.Search, &params)
	queryGet := fmt.Sprintf("%s %s", getListQuery, filterQuery)
	queryCount := fmt.Sprintf("%s %s", countListQuery, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"a.name": constants.Ascending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
		return resultData, paginationResult, err
	}

	err := tx.SelectContext(
		ctx,
		&resultData,
		queryGet,
		params...,
	)
	if err != nil {
		return resultData, paginationResult, constants.ErrorInternalServer(err.Error())
	}

	var count int
	err = tx.QueryRowContext(
		ctx,
		queryCount,
		params...,
	).Scan(&count)
	if err != nil {
		return resultData, paginationResult, constants.ErrorInternalServer(err.Error())
	}

	pagination.GetPagination(count, pagination.Page, pagination.Limit)
	next := int32(pagination.Next)
	prev := int32(pagination.Prev)
	totalPages := int32(pagination.TotalPages)
	totalRecords := int32(pagination.TotalRecords)
	paginationResult = common.Pagination{
		Page:         int32(pagination.Page),
		Limit:        int32(pagination.Limit),
		Sort:         pagination.Sort,
		SortBy:       pagination.SortBy,
		Next:         &next,
		Prev:         &prev,
		TotalPages:   &totalPages,
		TotalRecords: &totalRecords,
	}

	return resultData, paginationResult, nil
}

func (a adminRepository) GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetAdmin, *constants.ErrorResponse) {
	results := []models.GetAdmin{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailQuery,
		id,
	)
	if err != nil {
		return models.GetAdmin{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetAdmin{}, utils.ErrDataNotFound("admin")
	}

	return results[0], nil
}

func (a adminRepository) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateAdmin) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		createQuery,
		data.Username,
		data.Name,
		data.Password,
		data.RoleId,
		data.CreatedBy,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (a adminRepository) Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateAdmin) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		updateQuery,
		data.Username,
		data.Name,
		data.Password,
		data.RoleId,
		data.UpdatedBy,
		data.Id,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (a adminRepository) Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		deleteQuery,
		id,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (a adminRepository) GetSingleSuperAdmin(ctx context.Context, tx *sqlx.Tx) (models.GetAdmin, *constants.ErrorResponse) {
	results := []models.GetAdmin{}

	err := tx.SelectContext(
		ctx,
		&results,
		getSingleSuperAdminQuery,
	)
	if err != nil {
		return models.GetAdmin{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetAdmin{}, utils.ErrDataNotFound("admin")
	}

	return results[0], nil
}
