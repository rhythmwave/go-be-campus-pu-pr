package study_plan

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

type studyPlanRepository struct {
	*db.DB
}

func mapQueryFilterGetList(search, studentId, semesterId string, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	if search != "" {
		searchParams := fmt.Sprintf("%%%s%%", search)
		filterArray = append(filterArray, "s.name ILIKE $%d")
		filterParams = append(filterParams, searchParams)
	}
	if studentId != "" {
		filterArray = append(filterArray, "s.id = $%d")
		filterParams = append(filterParams, studentId)
	}
	if semesterId != "" {
		filterArray = append(filterArray, "se.id = $%d")
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

func (s studyPlanRepository) BulkCreate(ctx context.Context, tx *sqlx.Tx, data []models.CreateStudyPlan) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		bulkCreateQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (s studyPlanRepository) BulkApprove(ctx context.Context, tx *sqlx.Tx, studyPlanIds []string, isApproved bool) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		bulkApproveQuery,
		isApproved,
		pq.Array(studyPlanIds),
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (s studyPlanRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studentId, semesterId string) ([]models.GetStudyPlan, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetStudyPlan{}
	var paginationResult common.Pagination

	params := []interface{}{}
	filterQuery := mapQueryFilterGetList(pagination.Search, studentId, semesterId, &params)
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

func (s studyPlanRepository) GetApprovedByStudentId(ctx context.Context, tx *sqlx.Tx, studentId string) ([]models.GetStudyPlan, *constants.ErrorResponse) {
	results := []models.GetStudyPlan{}

	err := tx.SelectContext(
		ctx,
		&results,
		getApprovedByStudentIdQuery,
		studentId,
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (s studyPlanRepository) GetByStudentIdAndSemesterId(ctx context.Context, tx *sqlx.Tx, studentId, semesterId string) (models.GetStudyPlan, *constants.ErrorResponse) {
	results := []models.GetStudyPlan{}

	err := tx.SelectContext(
		ctx,
		&results,
		getByStudentIdAndSemesterIdQuery,
		studentId,
		semesterId,
	)
	if err != nil {
		return models.GetStudyPlan{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetStudyPlan{}, nil
	}

	return results[0], nil
}

func (s studyPlanRepository) GetByStudentIdsAndSemesterId(ctx context.Context, tx *sqlx.Tx, studentIds []string, semesterId string) ([]models.GetStudyPlan, *constants.ErrorResponse) {
	results := []models.GetStudyPlan{}

	err := tx.SelectContext(
		ctx,
		&results,
		getByStudentIdsAndSemesterIdQuery,
		pq.Array(studentIds),
		semesterId,
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}
