package shared_file

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	appConstants "github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
)

type sharedFileRepository struct {
	*db.DB
}

func mapQueryFilterGetList(search, appType, lecturerId string, isApproved *bool, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	if search != "" {
		searchParams := fmt.Sprintf("%%%s%%", search)
		filterArray = append(filterArray, "sf.title ILIKE $%d")
		filterParams = append(filterParams, searchParams)
	}
	if lecturerId != "" {
		filterArray = append(filterArray, "sf.lecturer_id = $%d")
		filterParams = append(filterParams, lecturerId)
	} else if appType != appConstants.AppTypeAdmin {
		filterArray = append(filterArray, "sf.is_approved IS true")
	}
	if isApproved != nil {
		filterArray = append(filterArray, "sf.is_approved = $%d")
		filterParams = append(filterParams, *isApproved)
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

func (f sharedFileRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, appType, lecturerId string, isApproved *bool) ([]models.GetSharedFile, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetSharedFile{}
	var paginationResult common.Pagination

	params := []interface{}{}
	filterQuery := mapQueryFilterGetList(pagination.Search, appType, lecturerId, isApproved, &params)
	queryGet := fmt.Sprintf("%s %s", getListQuery, filterQuery)
	queryCount := fmt.Sprintf("%s %s", countListQuery, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"sf.title": constants.Ascending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (f sharedFileRepository) GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetSharedFile, *constants.ErrorResponse) {
	results := []models.GetSharedFile{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailQuery,
		id,
	)
	if err != nil {
		return models.GetSharedFile{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetSharedFile{}, utils.ErrDataNotFound("document type")
	}

	return results[0], nil
}

func (f sharedFileRepository) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateSharedFile) *constants.ErrorResponse {
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

func (f sharedFileRepository) Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateSharedFile) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		updateQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (f sharedFileRepository) Approve(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		approveQuery,
		id,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (f sharedFileRepository) Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse {
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
