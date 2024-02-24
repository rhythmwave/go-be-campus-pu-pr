package grade_type

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

type gradeTypeRepository struct {
	*db.DB
}

func mapQueryFilterGetList(search string, params *[]interface{}) string {
	filterArray := []string{
		"sl.id = $%d",
	}
	filterParams := *params

	if search != "" {
		searchParams := fmt.Sprintf("%%%s%%", search)
		filterArray = append(filterArray, "gt.code ILIKE $%d")
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

func (f gradeTypeRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studyLevelId string) ([]models.GetGradeType, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetGradeType{}
	var paginationResult common.Pagination

	params := []interface{}{
		studyLevelId,
	}
	filterQuery := mapQueryFilterGetList(pagination.Search, &params)
	queryGet := fmt.Sprintf("%s %s", getListQuery, filterQuery)
	queryCount := fmt.Sprintf("%s %s", countListQuery, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"sl.short_name": constants.Ascending, "gt.code": constants.Ascending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (f gradeTypeRepository) GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetGradeType, *constants.ErrorResponse) {
	results := []models.GetGradeType{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailQuery,
		id,
	)
	if err != nil {
		return models.GetGradeType{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetGradeType{}, utils.ErrDataNotFound("grade type")
	}

	return results[0], nil
}

func (f gradeTypeRepository) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateGradeType) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		createQuery,
		data.StudyLevelId,
		data.Code,
		data.GradePoint,
		data.MinimumGrade,
		data.MaximumGrade,
		data.GradeCategory,
		data.GradePointCategory,
		data.Label,
		data.EnglishLabel,
		data.StartDate,
		data.EndDate,
		data.CreatedBy,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (f gradeTypeRepository) Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateGradeType) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		updateQuery,
		data.Code,
		data.GradePoint,
		data.MinimumGrade,
		data.MaximumGrade,
		data.GradeCategory,
		data.GradePointCategory,
		data.Label,
		data.EnglishLabel,
		data.StartDate,
		data.EndDate,
		data.UpdatedBy,
		data.Id,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (f gradeTypeRepository) Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse {
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

func (f gradeTypeRepository) GetByGradeCode(ctx context.Context, tx *sqlx.Tx, studyLevelId string, gradeCode string) (models.GetGradeType, *constants.ErrorResponse) {
	results := []models.GetGradeType{}

	err := tx.SelectContext(
		ctx,
		&results,
		getByGradeCodeQuery,
		studyLevelId,
		gradeCode,
	)
	if err != nil {
		return models.GetGradeType{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetGradeType{}, utils.ErrDataNotFound("grade type")
	}

	return results[0], nil
}
