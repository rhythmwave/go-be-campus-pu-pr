package learning_achievement_category

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

type learningAchievementCategoryRepository struct {
	*db.DB
}

func mapQueryFilterGetList(search string, params *[]interface{}) string {
	filterArray := []string{
		"lac.curriculum_id = $%d",
	}
	filterParams := *params

	if search != "" {
		filterArray = append(filterArray, "lac.name ILIKE $%d")
		filterParams = append(filterParams, fmt.Sprintf("%%%s%%", search))
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

func (f learningAchievementCategoryRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, curriculumId string) ([]models.GetLearningAchievementCategory, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetLearningAchievementCategory{}
	var paginationResult common.Pagination

	params := []interface{}{
		curriculumId,
	}
	filterQuery := mapQueryFilterGetList(pagination.Search, &params)
	queryGet := fmt.Sprintf("%s %s", getListQuery, filterQuery)
	queryCount := fmt.Sprintf("%s %s", countListQuery, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"lac.name": constants.Ascending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (f learningAchievementCategoryRepository) GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetLearningAchievementCategory, *constants.ErrorResponse) {
	results := []models.GetLearningAchievementCategory{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailQuery,
		id,
	)
	if err != nil {
		return models.GetLearningAchievementCategory{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetLearningAchievementCategory{}, utils.ErrDataNotFound("learning achievement category")
	}

	return results[0], nil
}

func (f learningAchievementCategoryRepository) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateLearningAchievementCategory) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		createQuery,
		data.CurriculumId,
		data.Name,
		data.EnglishName,
		data.CreatedBy,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (f learningAchievementCategoryRepository) Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateLearningAchievementCategory) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		updateQuery,
		data.Name,
		data.EnglishName,
		data.UpdatedBy,
		data.Id,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (f learningAchievementCategoryRepository) Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse {
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
