package grade_component

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

type gradeComponentRepository struct {
	*db.DB
}

func mapQueryFilterGetList(search, subjectCategoryId string, params *[]interface{}) string {
	filterArray := []string{
		"sp.id = $%d",
	}
	filterParams := *params

	if search != "" {
		searchParams := fmt.Sprintf("%%%s%%", search)
		filterArray = append(filterArray, "(sc.name ILIKE $%d OR gc.name ILIKE $%d)")
		filterParams = append(filterParams, searchParams, searchParams)
	}
	if subjectCategoryId != "" {
		filterArray = append(filterArray, "gc.subject_category_id = $%d")
		filterParams = append(filterParams, subjectCategoryId)
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

func mapQueryFilterGetDistinctSubjectCategoryList(search string, params *[]interface{}) string {
	filterArray := []string{
		"sp.id = $%d",
	}
	filterParams := *params

	if search != "" {
		searchParams := fmt.Sprintf("%%%s%%", search)
		filterArray = append(filterArray, "sc.name ILIKE $%d")
		filterParams = append(filterParams, searchParams)
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

func (f gradeComponentRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studyProgramId, subjectCategoryId string) ([]models.GetGradeComponent, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetGradeComponent{}
	var paginationResult common.Pagination

	params := []interface{}{
		studyProgramId,
	}
	filterQuery := mapQueryFilterGetList(pagination.Search, subjectCategoryId, &params)
	queryGet := fmt.Sprintf("%s %s", getListQuery, filterQuery)
	queryCount := fmt.Sprintf("%s %s", countListQuery, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"sc.name": constants.Ascending, "gc.name": constants.Ascending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (f gradeComponentRepository) GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetGradeComponent, *constants.ErrorResponse) {
	results := []models.GetGradeComponent{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailQuery,
		id,
	)
	if err != nil {
		return models.GetGradeComponent{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetGradeComponent{}, utils.ErrDataNotFound("grade component")
	}

	return results[0], nil
}

func (f gradeComponentRepository) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateGradeComponent) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		createQuery,
		data.StudyProgramId,
		data.SubjectCategoryId,
		data.Name,
		data.IsActive,
		data.CreatedBy,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (f gradeComponentRepository) Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateGradeComponent) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		updateQuery,
		data.SubjectCategoryId,
		data.Name,
		data.IsActive,
		data.UpdatedBy,
		data.Id,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (f gradeComponentRepository) Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse {
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

func (f gradeComponentRepository) GetDistinctSubjectCategoryList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studyProgramId string) ([]models.GetGradeComponentDistinctSubjectCategory, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetGradeComponentDistinctSubjectCategory{}
	var paginationResult common.Pagination

	params := []interface{}{
		studyProgramId,
	}
	filterQuery := mapQueryFilterGetDistinctSubjectCategoryList(pagination.Search, &params)
	queryGet := fmt.Sprintf("%s %s", getDistinctSubjectCategoryListQuery, filterQuery)
	queryCount := fmt.Sprintf(countDistinctSubjectCategoryListQuery, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"subject_category_name": constants.Ascending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (f gradeComponentRepository) GetPercentageBySubjectCategories(ctx context.Context, tx *sqlx.Tx, studyProgramId string, subjectCategoryIds []string) ([]models.GetPercentageBySubjectCategories, *constants.ErrorResponse) {
	results := []models.GetPercentageBySubjectCategories{}

	err := tx.SelectContext(
		ctx,
		&results,
		getPercentageBySubjectCategoriesQuery,
		studyProgramId,
		pq.Array(subjectCategoryIds),
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (f gradeComponentRepository) BulkUpdatePercentage(ctx context.Context, tx *sqlx.Tx, data []models.BulkUpdateGradeComponentPercentage) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		bulkUpdatePercentageQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}
