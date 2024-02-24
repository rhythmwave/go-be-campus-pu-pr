package graduation_session

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

type graduationSessionRepository struct {
	*db.DB
}

func mapQueryFilterGetList(search string, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	if search != "" {
		searchParams := fmt.Sprintf("%%%s%%", search)
		filterArray = append(filterArray, "(gs.session_year = $%d)")
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

func (f graduationSessionRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest) ([]models.GetGraduationSession, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetGraduationSession{}
	var paginationResult common.Pagination

	params := []interface{}{}
	filterQuery := mapQueryFilterGetList(pagination.Search, &params)
	queryGet := fmt.Sprintf("%s %s", getListQuery, filterQuery)
	queryCount := fmt.Sprintf("%s %s", countListQuery, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"gs.session_date": constants.Descending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (f graduationSessionRepository) GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetGraduationSession, *constants.ErrorResponse) {
	results := []models.GetGraduationSession{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailQuery,
		id,
	)
	if err != nil {
		return models.GetGraduationSession{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetGraduationSession{}, utils.ErrDataNotFound("graduation session")
	}

	return results[0], nil
}

func (f graduationSessionRepository) GetUpcoming(ctx context.Context, tx *sqlx.Tx) (models.GetGraduationSession, *constants.ErrorResponse) {
	results := []models.GetGraduationSession{}

	err := tx.SelectContext(
		ctx,
		&results,
		getUpcomingQuery,
	)
	if err != nil {
		return models.GetGraduationSession{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetGraduationSession{}, nil
	}

	return results[0], nil
}

func (f graduationSessionRepository) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateGraduationSession) *constants.ErrorResponse {
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

func (f graduationSessionRepository) Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateGraduationSession) *constants.ErrorResponse {
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

func (f graduationSessionRepository) Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse {
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
