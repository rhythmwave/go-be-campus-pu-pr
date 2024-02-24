package class_work

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

type classWorkRepository struct {
	*db.DB
}

func mapQueryFilterGetList(search string, classIds []string, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	if search != "" {
		searchParams := fmt.Sprintf("%%%s%%", search)
		filterArray = append(filterArray, "cw.title ILIKE $%d")
		filterParams = append(filterParams, searchParams)
	}
	if len(classIds) != 0 {
		filterArray = append(filterArray, "cw.class_id IN (SELECT UNNEST($%d::uuid[]))")
		filterParams = append(filterParams, pq.Array(classIds))
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

func mapQueryFilterGetSubmission(search string, ids []string, isSubmitted *bool, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	if search != "" {
		searchParams := fmt.Sprintf("%%%s%%", search)
		filterArray = append(filterArray, "(s.name ILIKE $%d OR s.nim_number::text ILIKE $%d)")
		filterParams = append(filterParams, searchParams, searchParams)
	}
	if len(ids) > 0 {
		filterArray = append(filterArray, "cws.id IN (SELECT UNNEST($%d::uuid[]))")
		filterParams = append(filterParams, pq.Array(ids))
	}
	if isSubmitted != nil {
		filterArray = append(filterArray, "(cws.id IS NOT NULL) = $%d")
		filterParams = append(filterParams, *isSubmitted)
	}

	result := strings.Join(filterArray, " AND ")
	args := []interface{}{}
	for i := 2; i < len(filterParams); i++ {
		args = append(args, i+1)
	}
	result = fmt.Sprintf(result, args...)
	if result != "" {
		result = fmt.Sprintf("WHERE %s", result)
	}

	*params = filterParams

	return result
}

func (c classWorkRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, classIds []string, studentId string) ([]models.GetClassWork, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetClassWork{}
	var paginationResult common.Pagination

	queryGet := fmt.Sprintf(getListQuery, "", "")
	queryCount := fmt.Sprintf(countListQuery, "")
	if studentId != "" {
		additionalSelect := `,
			cws.file_path AS submission_file_path,
			cws.file_path_type AS submission_file_path_type,
			cws.point AS submission_point
		`
		additionalJoin := fmt.Sprintf(`
			LEFT JOIN class_work_submissions cws ON cws.class_work_id = cw.id AND cws.student_id = '%s'
		`, studentId)

		queryGet = fmt.Sprintf(getListQuery, additionalSelect, additionalJoin)
		queryCount = fmt.Sprintf(countListQuery, additionalJoin)
	}

	params := []interface{}{}
	filterQuery := mapQueryFilterGetList(pagination.Search, classIds, &params)
	queryGet = fmt.Sprintf("%s %s", queryGet, filterQuery)
	queryCount = fmt.Sprintf("%s %s", queryCount, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"cw.title": constants.Ascending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (c classWorkRepository) GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetClassWork, *constants.ErrorResponse) {
	results := []models.GetClassWork{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailQuery,
		id,
	)
	if err != nil {
		return models.GetClassWork{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetClassWork{}, utils.ErrDataNotFound("class work")
	}

	return results[0], nil
}

func (c classWorkRepository) GetSubmission(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, classId, classWorkId string, ids []string, isSubmitted *bool) ([]models.GetClassWorkSubmission, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetClassWorkSubmission{}
	var paginationResult common.Pagination

	params := []interface{}{
		classId,
		classWorkId,
	}
	filterQuery := mapQueryFilterGetSubmission(pagination.Search, ids, isSubmitted, &params)
	queryGet := fmt.Sprintf("%s %s", getSubmissionQuery, filterQuery)
	queryCount := fmt.Sprintf("%s %s", countSubmissionQuery, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"s.nim_number": constants.Ascending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (c classWorkRepository) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateClassWork) *constants.ErrorResponse {
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

func (c classWorkRepository) Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateClassWork) *constants.ErrorResponse {
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

func (c classWorkRepository) Delete(ctx context.Context, tx *sqlx.Tx, ids []string) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		deleteQuery,
		pq.Array(ids),
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (c classWorkRepository) GradeSubmission(ctx context.Context, tx *sqlx.Tx, data []models.GradeClassWorkSubmission) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		gradeSubmissionQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (c classWorkRepository) Submit(ctx context.Context, tx *sqlx.Tx, classWorkId, studentId, filePath, filePathType string) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		submitQuery,
		classWorkId,
		studentId,
		filePath,
		filePathType,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}
