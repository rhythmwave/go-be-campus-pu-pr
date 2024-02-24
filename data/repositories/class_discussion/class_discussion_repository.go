package class_discussion

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

type classDiscussionRepository struct {
	*db.DB
}

func mapQueryFilterGetList(search, classId string, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	if search != "" {
		searchParams := fmt.Sprintf("%%%s%%", search)
		filterArray = append(filterArray, "cd.title ILIKE $%d")
		filterParams = append(filterParams, searchParams)
	}
	if classId != "" {
		filterArray = append(filterArray, "cd.class_id = $%d")
		filterParams = append(filterParams, classId)
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

func mapQueryFilterGetComment(classDiscussionId string, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	if classDiscussionId != "" {
		filterArray = append(filterArray, "cdc.class_discussion_id = $%d")
		filterParams = append(filterParams, classDiscussionId)
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

func (c classDiscussionRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, classId string) ([]models.GetClassDiscussion, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetClassDiscussion{}
	var paginationResult common.Pagination

	params := []interface{}{}
	filterQuery := mapQueryFilterGetList(pagination.Search, classId, &params)
	queryGet := fmt.Sprintf("%s %s", getListQuery, filterQuery)
	queryCount := fmt.Sprintf("%s %s", countListQuery, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"cd.title": constants.Ascending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (c classDiscussionRepository) GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetClassDiscussion, *constants.ErrorResponse) {
	results := []models.GetClassDiscussion{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailQuery,
		id,
	)
	if err != nil {
		return models.GetClassDiscussion{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetClassDiscussion{}, utils.ErrDataNotFound("class discussion")
	}

	return results[0], nil
}

func (c classDiscussionRepository) GetComment(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, classDiscussionId string) ([]models.GetClassDiscussionComment, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetClassDiscussionComment{}
	var paginationResult common.Pagination

	params := []interface{}{}
	filterQuery := mapQueryFilterGetComment(classDiscussionId, &params)
	queryGet := fmt.Sprintf("%s %s", getCommentQuery, filterQuery)
	queryCount := fmt.Sprintf("%s %s", countCommentQuery, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"cdc.created_at": constants.Ascending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (c classDiscussionRepository) GetDetailComment(ctx context.Context, tx *sqlx.Tx, id string) (models.GetClassDiscussionComment, *constants.ErrorResponse) {
	results := []models.GetClassDiscussionComment{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailCommentQuery,
		id,
	)
	if err != nil {
		return models.GetClassDiscussionComment{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetClassDiscussionComment{}, utils.ErrDataNotFound("class discussion")
	}

	return results[0], nil
}

func (c classDiscussionRepository) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateClassDiscussion) *constants.ErrorResponse {
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

func (c classDiscussionRepository) Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateClassDiscussion) *constants.ErrorResponse {
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

func (c classDiscussionRepository) Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse {
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

func (c classDiscussionRepository) CreateComment(ctx context.Context, tx *sqlx.Tx, data models.CreateClassDiscussionComment) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		createCommentQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (c classDiscussionRepository) DeleteComment(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		deleteCommentQuery,
		id,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}
