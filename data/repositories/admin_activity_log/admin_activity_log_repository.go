package admin_activity_log

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

type adminActivityLog struct {
	*db.DB
}

func mapQueryFilterGetList(search string, params *[]interface{}) string {
	filterArray := []string{
		"EXTRACT('year' FROM aal.created_at) = $1",
		"EXTRACT('month' FROM aal.created_at) = $2",
	}
	filterParams := *params

	if search != "" {
		filterArray = append(filterArray, "(aal.admin_name ILIKE $3 OR aal.admin_username ILIKE $3 OR aal.module ILIKE $3 OR aal.action ILIKE $3 OR aal.ip_address ILIKE $3 OR aal.user_agent ILIKE $3)")
		filterParams = append(filterParams, fmt.Sprintf("%%%s%%", search))
	}

	result := strings.Join(filterArray, " AND ")
	// args := []interface{}{}
	// for i := 0; i < len(filterParams); i++ {
	// 	args = append(args, i+1)
	// }
	// result = fmt.Sprintf(result, args...)
	if result != "" {
		result = fmt.Sprintf("WHERE %s", result)
	}

	*params = filterParams

	return result
}

func (a adminActivityLog) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, year, month uint32) ([]models.GetAdminActivityLog, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetAdminActivityLog{}
	var paginationResult common.Pagination

	params := []interface{}{
		year,
		month,
	}
	filterQuery := mapQueryFilterGetList(pagination.Search, &params)
	queryGet := fmt.Sprintf("%s %s", getListQuery, filterQuery)
	queryCount := fmt.Sprintf("%s %s", countListQuery, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"aal.created_at": constants.Descending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (a adminActivityLog) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateAdminActivityLog) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		createQuery,
		data.AdminId,
		data.AdminName,
		data.AdminUsername,
		data.Module,
		data.Action,
		data.IpAddress,
		data.UserAgent,
		data.ExecutionTime,
		data.MemoryUsage,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}
