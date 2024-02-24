package class_exam

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

type classExamRepository struct {
	*db.DB
}

func mapQueryFilterGetList(search string, classIds, ids []string, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	if search != "" {
		searchParams := fmt.Sprintf("%%%s%%", search)
		filterArray = append(filterArray, "ce.title ILIKE $%d")
		filterParams = append(filterParams, searchParams)
	}
	if len(classIds) != 0 {
		filterArray = append(filterArray, "ce.class_id IN (SELECT UNNEST($%d::uuid[]))")
		filterParams = append(filterParams, pq.Array(classIds))
	}
	if len(ids) > 0 {
		filterArray = append(filterArray, "ce.id IN (SELECT UNNEST($%d::uuid[]))")
		filterParams = append(filterParams, pq.Array(ids))
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
		filterArray = append(filterArray, "ces.id IN (SELECT UNNEST($%d::uuid[]))")
		filterParams = append(filterParams, pq.Array(ids))
	}
	if isSubmitted != nil {
		filterArray = append(filterArray, "(ces.id IS NOT NULL) = $%d")
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

func (c classExamRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, classIds []string, studentId string, ids []string) ([]models.GetClassExam, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetClassExam{}
	var paginationResult common.Pagination

	queryGet := fmt.Sprintf(getListQuery, "", "")
	queryCount := fmt.Sprintf(countListQuery, "")
	if studentId != "" {
		additionalSelect := `,
			ces.file_path AS submission_file_path,
			ces.file_path_type AS submission_file_path_type,
			ces.point AS submission_point
		`
		additionalJoin := fmt.Sprintf(`
			LEFT JOIN class_exam_submissions ces ON ces.class_exam_id = ce.id AND ces.student_id = '%s'
		`, studentId)

		queryGet = fmt.Sprintf(getListQuery, additionalSelect, additionalJoin)
		queryCount = fmt.Sprintf(countListQuery, additionalJoin)
	}

	params := []interface{}{}
	filterQuery := mapQueryFilterGetList(pagination.Search, classIds, ids, &params)
	queryGet = fmt.Sprintf("%s %s", queryGet, filterQuery)
	queryCount = fmt.Sprintf("%s %s", queryCount, filterQuery)
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

func (c classExamRepository) GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetClassExam, *constants.ErrorResponse) {
	results := []models.GetClassExam{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailQuery,
		id,
	)
	if err != nil {
		return models.GetClassExam{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetClassExam{}, utils.ErrDataNotFound("class exam")
	}

	return results[0], nil
}

func (c classExamRepository) GetSubmission(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, classId, classExamId string, ids []string, isSubmitted *bool) ([]models.GetClassExamSubmission, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetClassExamSubmission{}
	var paginationResult common.Pagination

	params := []interface{}{
		classId,
		classExamId,
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

func (c classExamRepository) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateClassExam) *constants.ErrorResponse {
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

func (c classExamRepository) Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateClassExam) *constants.ErrorResponse {
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

func (c classExamRepository) Delete(ctx context.Context, tx *sqlx.Tx, ids []string) *constants.ErrorResponse {
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

func (c classExamRepository) GradeSubmission(ctx context.Context, tx *sqlx.Tx, data []models.GradeClassExamSubmission) *constants.ErrorResponse {
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

func (c classExamRepository) Submit(ctx context.Context, tx *sqlx.Tx, classExamId, studentId, filePath, filePathType string) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		submitQuery,
		classExamId,
		studentId,
		filePath,
		filePathType,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}
