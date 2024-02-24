package lecturer_student_activity_log

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

type lecturerStudentActivityLog struct {
	*db.DB
}

func mapQueryFilterGetList(search string, params *[]interface{}) string {
	filterArray := []string{
		"EXTRACT('year' FROM lsal.created_at) = $1",
		"EXTRACT('month' FROM lsal.created_at) = $2",
	}
	filterParams := *params

	if search != "" {
		filterArray = append(filterArray, "(lsal.lecturer_name ILIKE $3 OR lsal.lecturer_username ILIKE $3 OR lsal.student_name ILIKE $3 OR lsal.student_username ILIKE $3 OR lsal.module ILIKE $3 OR lsal.action ILIKE $3 OR lsal.ip_address ILIKE $3 OR lsal.user_agent ILIKE $3)")
		filterParams = append(filterParams, fmt.Sprintf("%%%s%%", search))
	}

	result := strings.Join(filterArray, " AND ")
	if result != "" {
		result = fmt.Sprintf("WHERE %s", result)
	}

	*params = filterParams

	return result
}

func (a lecturerStudentActivityLog) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, year, month uint32) ([]models.GetLecturerStudentActivityLog, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetLecturerStudentActivityLog{}
	var paginationResult common.Pagination

	params := []interface{}{
		year,
		month,
	}
	filterQuery := mapQueryFilterGetList(pagination.Search, &params)
	queryGet := fmt.Sprintf("%s %s", getListQuery, filterQuery)
	queryCount := fmt.Sprintf("%s %s", countListQuery, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"lsal.created_at": constants.Descending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (a lecturerStudentActivityLog) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateLecturerStudentActivityLog) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		createQuery,
		data.LecturerId,
		data.LecturerName,
		data.LecturerUsername,
		data.StudentId,
		data.StudentName,
		data.StudentUsername,
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
