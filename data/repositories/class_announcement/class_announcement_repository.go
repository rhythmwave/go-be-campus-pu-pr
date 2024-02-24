package class_announcement

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

type classAnnouncementRepository struct {
	*db.DB
}

func mapQueryFilterGetList(search string, classIds, ids []string, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	if search != "" {
		searchParams := fmt.Sprintf("%%%s%%", search)
		filterArray = append(filterArray, "ca.title ILIKE $%d")
		filterParams = append(filterParams, searchParams)
	}
	if len(classIds) != 0 {
		filterArray = append(filterArray, "ca.class_id IN (SELECT UNNEST($%d::uuid[]))")
		filterParams = append(filterParams, pq.Array(classIds))
	}
	if len(ids) > 0 {
		filterArray = append(filterArray, "ca.id IN (SELECT UNNEST($%d::uuid[]))")
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

func (c classAnnouncementRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, classIds, ids []string) ([]models.GetClassAnnouncement, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetClassAnnouncement{}
	var paginationResult common.Pagination

	params := []interface{}{}
	filterQuery := mapQueryFilterGetList(pagination.Search, classIds, ids, &params)
	queryGet := fmt.Sprintf("%s %s", getListQuery, filterQuery)
	queryCount := fmt.Sprintf("%s %s", countListQuery, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"ca.title": constants.Ascending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (c classAnnouncementRepository) GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetClassAnnouncement, *constants.ErrorResponse) {
	results := []models.GetClassAnnouncement{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailQuery,
		id,
	)
	if err != nil {
		return models.GetClassAnnouncement{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetClassAnnouncement{}, utils.ErrDataNotFound("class announcement")
	}

	return results[0], nil
}

func (c classAnnouncementRepository) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateClassAnnouncement) *constants.ErrorResponse {
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

func (c classAnnouncementRepository) Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateClassAnnouncement) *constants.ErrorResponse {
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

func (c classAnnouncementRepository) Delete(ctx context.Context, tx *sqlx.Tx, ids []string) *constants.ErrorResponse {
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
