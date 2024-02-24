package class_event

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
)

type classEventRepository struct {
	*db.DB
}

func mapQueryFilterGetList(search, classId, frequency string, futureEventOnly bool, isActive *bool, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	if search != "" {
		searchParams := fmt.Sprintf("%%%s%%", search)
		filterArray = append(filterArray, "ce.title ILIKE $%d")
		filterParams = append(filterParams, searchParams)
	}
	if classId != "" {
		filterArray = append(filterArray, "ce.class_id = $%d")
		filterParams = append(filterParams, classId)
	}
	if frequency != "" {
		filterArray = append(filterArray, "ce.frequency = $%d")
		filterParams = append(filterParams, frequency)
	}
	if futureEventOnly {
		filterArray = append(filterArray, "ce.event_time >= now()")
	}
	if isActive != nil {
		filterArray = append(filterArray, "ce.is_active = $%d")
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

func (f classEventRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, classId, frequency string, futureEventOnly bool, isActive *bool) ([]models.GetClassEvent, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetClassEvent{}
	var paginationResult common.Pagination

	params := []interface{}{}
	filterQuery := mapQueryFilterGetList(pagination.Search, classId, frequency, futureEventOnly, isActive, &params)
	queryGet := fmt.Sprintf("%s %s", getListQuery, filterQuery)
	queryCount := fmt.Sprintf("%s %s", countListQuery, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"ce.title": constants.Ascending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (f classEventRepository) GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetClassEvent, *constants.ErrorResponse) {
	results := []models.GetClassEvent{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailQuery,
		id,
	)
	if err != nil {
		return models.GetClassEvent{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetClassEvent{}, utils.ErrDataNotFound("class event")
	}

	return results[0], nil
}

func (f classEventRepository) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateClassEvent) *constants.ErrorResponse {
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

func (f classEventRepository) Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateClassEvent) *constants.ErrorResponse {
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

func (f classEventRepository) BulkUpdateActivation(ctx context.Context, tx *sqlx.Tx, ids []string, isActive bool) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		bulkUpdateActivationQuery,
		pq.Array(ids),
		isActive,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (f classEventRepository) BulkDelete(ctx context.Context, tx *sqlx.Tx, ids []string) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		bulkDeleteQuery,
		pq.Array(ids),
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}
