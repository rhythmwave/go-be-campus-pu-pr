package credit_quota

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
)

type creditQuotaRepository struct {
	*db.DB
}

func (f creditQuotaRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest) ([]models.GetCreditQuota, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetCreditQuota{}
	var paginationResult common.Pagination

	params := []interface{}{}
	queryGet := getListQuery
	queryCount := countListQuery
	if err := utils.QueryOperation(&queryGet, map[string]string{"cq.minimum_grade_point": constants.Descending, "cq.maximum_grade_point": constants.Descending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (f creditQuotaRepository) GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetCreditQuota, *constants.ErrorResponse) {
	results := []models.GetCreditQuota{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailQuery,
		id,
	)
	if err != nil {
		return models.GetCreditQuota{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetCreditQuota{}, utils.ErrDataNotFound("credit quota")
	}

	return results[0], nil
}

func (f creditQuotaRepository) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateCreditQuota) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		createQuery,
		data.MinimumGradePoint,
		data.MaximumGradePoint,
		data.MaximumCredit,
		data.CreatedBy,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (f creditQuotaRepository) Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateCreditQuota) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		updateQuery,
		data.MinimumGradePoint,
		data.MaximumGradePoint,
		data.MaximumCredit,
		data.UpdatedBy,
		data.Id,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (f creditQuotaRepository) Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse {
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
