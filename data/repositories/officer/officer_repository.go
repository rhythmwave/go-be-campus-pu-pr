package officer

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

type officerRepository struct {
	*db.DB
}

func mapQueryFilterGetList(search string, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	if search != "" {
		searchParams := fmt.Sprintf("%%%s%%", search)
		filterArray = append(filterArray, "(o.name ILIKE $%d OR o.title ILIKE $%d OR o.english_title ILIKE $%d OR sp.name ILIKE $%d)")
		filterParams = append(filterParams, searchParams, searchParams, searchParams, searchParams)
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

func (a officerRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest) ([]models.GetOfficer, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetOfficer{}
	var paginationResult common.Pagination

	params := []interface{}{}
	filterQuery := mapQueryFilterGetList(pagination.Search, &params)
	queryGet := fmt.Sprintf("%s %s", getListQuery, filterQuery)
	queryCount := fmt.Sprintf("%s %s", countListQuery, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"o.name": constants.Ascending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (a officerRepository) GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetOfficer, *constants.ErrorResponse) {
	results := []models.GetOfficer{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailQuery,
		id,
	)
	if err != nil {
		return models.GetOfficer{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetOfficer{}, utils.ErrDataNotFound("officer")
	}

	return results[0], nil
}

func (a officerRepository) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateOfficer) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		createQuery,
		data.IdNationalLecturer,
		data.Name,
		data.Title,
		data.EnglishTitle,
		data.StudyProgramId,
		data.SignaturePath,
		data.SignaturePathType,
		data.ShowSignature,
		data.EmployeeNo,
		data.CreatedBy,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (a officerRepository) Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateOfficer) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		updateQuery,
		data.IdNationalLecturer,
		data.Name,
		data.Title,
		data.EnglishTitle,
		data.StudyProgramId,
		data.SignaturePath,
		data.SignaturePathType,
		data.ShowSignature,
		data.EmployeeNo,
		data.UpdatedBy,
		data.Id,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (a officerRepository) Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse {
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
