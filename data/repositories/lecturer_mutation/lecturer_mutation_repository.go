package lecturer_mutation

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

type lecturerMutationRepository struct {
	*db.DB
}

func mapQueryFilterGetList(search, studyProgramId, idNationalLecturer, semesterId string, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	if search != "" {
		filterArray = append(filterArray, "l.name ILIKE $%d")
		filterParams = append(filterParams, fmt.Sprintf("%%%s%%", search))
	}
	if studyProgramId != "" {
		filterArray = append(filterArray, "sp.id = $%d")
		filterParams = append(filterParams, studyProgramId)
	}
	if idNationalLecturer != "" {
		filterArray = append(filterArray, "l.id_national_lecturer = $%d")
		filterParams = append(filterParams, idNationalLecturer)
	}
	if semesterId != "" {
		filterArray = append(filterArray, "lm.semester_id = $%d")
		filterParams = append(filterParams, semesterId)
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

func (l lecturerMutationRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studyProgramId, idNationalLecturer, semesterId string) ([]models.GetLecturerMutation, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetLecturerMutation{}
	var paginationResult common.Pagination

	params := []interface{}{}
	filterQuery := mapQueryFilterGetList(pagination.Search, studyProgramId, idNationalLecturer, semesterId, &params)
	queryGet := fmt.Sprintf("%s %s", getListQuery, filterQuery)
	queryCount := fmt.Sprintf("%s %s", countListQuery, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"lm.mutation_date": constants.Descending, "l.name": constants.Ascending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (l lecturerMutationRepository) GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetLecturerMutation, *constants.ErrorResponse) {
	results := []models.GetLecturerMutation{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailQuery,
		id,
	)
	if err != nil {
		return models.GetLecturerMutation{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetLecturerMutation{}, utils.ErrDataNotFound("lecturer mutation")
	}

	return results[0], nil
}

func (l lecturerMutationRepository) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateLecturerMutation) *constants.ErrorResponse {
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
