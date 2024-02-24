package student_leave

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

type studentLeaveRepository struct {
	*db.DB
}

func mapQueryFilterGetListRequest(search, studyProgramId, studentId string, isApproved *bool, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	if search != "" {
		filterArray = append(filterArray, "s.name ILIKE $%d")
		filterParams = append(filterParams, fmt.Sprintf("%%%s%%", search))
	}
	if studyProgramId != "" {
		filterArray = append(filterArray, "sp.id = $%d")
		filterParams = append(filterParams, studyProgramId)
	}
	if studentId != "" {
		filterArray = append(filterArray, "s.id = $%d")
		filterParams = append(filterParams, studentId)
	}
	if isApproved != nil {
		filterArray = append(filterArray, "(slr.is_approved IS NOT NULL) = $%d")
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

func mapQueryFilterGetList(search, studyProgramId, semesterId string, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	if search != "" {
		filterArray = append(filterArray, "s.name ILIKE $%d")
		filterParams = append(filterParams, fmt.Sprintf("%%%s%%", search))
	}
	if studyProgramId != "" {
		filterArray = append(filterArray, "sp.id = $%d")
		filterParams = append(filterParams, studyProgramId)
	}
	if semesterId != "" {
		filterArray = append(filterArray, "sl.semester_id = $%d")
		filterParams = append(filterParams, semesterId)
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

func (l studentLeaveRepository) GetListRequest(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studyProgramId, studentId string, isApproved *bool) ([]models.GetStudentLeaveRequest, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetStudentLeaveRequest{}
	var paginationResult common.Pagination

	params := []interface{}{}
	filterQuery := mapQueryFilterGetListRequest(pagination.Search, studyProgramId, studentId, isApproved, &params)
	queryGet := fmt.Sprintf("%s %s", getListRequestQuery, filterQuery)
	queryCount := fmt.Sprintf("%s %s", countListRequestQuery, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"slr.start_date": constants.Descending, "s.name": constants.Ascending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (l studentLeaveRepository) GetDetailRequest(ctx context.Context, tx *sqlx.Tx, id string) (models.GetStudentLeaveRequest, *constants.ErrorResponse) {
	results := []models.GetStudentLeaveRequest{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailRequestQuery,
		id,
	)
	if err != nil {
		return models.GetStudentLeaveRequest{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetStudentLeaveRequest{}, utils.ErrDataNotFound("student leave request")
	}

	return results[0], nil
}

func (l studentLeaveRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studyProgramId, semesterId string) ([]models.GetStudentLeave, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetStudentLeave{}
	var paginationResult common.Pagination

	params := []interface{}{}
	filterQuery := mapQueryFilterGetList(pagination.Search, studyProgramId, semesterId, &params)
	queryGet := fmt.Sprintf("%s %s", getListQuery, filterQuery)
	queryCount := fmt.Sprintf("%s %s", countListQuery, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"se.start_date": constants.Descending, "s.name": constants.Ascending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (l studentLeaveRepository) GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetStudentLeave, *constants.ErrorResponse) {
	results := []models.GetStudentLeave{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailQuery,
		id,
	)
	if err != nil {
		return models.GetStudentLeave{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetStudentLeave{}, utils.ErrDataNotFound("student leave")
	}

	return results[0], nil
}

func (l studentLeaveRepository) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateStudentLeave) *constants.ErrorResponse {
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

func (l studentLeaveRepository) Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateStudentLeave) *constants.ErrorResponse {
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

func (l studentLeaveRepository) Approve(ctx context.Context, tx *sqlx.Tx, id string, isApproved bool) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		approveQuery,
		id,
		isApproved,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (l studentLeaveRepository) End(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		endQuery,
		id,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (l studentLeaveRepository) Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse {
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
