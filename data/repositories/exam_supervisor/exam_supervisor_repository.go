package exam_supervisor

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

type examSupervisorRepository struct {
	*db.DB
}

func mapQueryFilterGetList(search, studyProgramId string, params *[]interface{}) string {
	var startIndex int
	filterArray := []string{}
	filterParams := *params

	if search != "" {
		filterArray = append(filterArray, "es.name ILIKE $%d")
		filterParams = append(filterParams, fmt.Sprintf("%%%s%%", search))
	}
	if studyProgramId != "" {
		filterArray = append(filterArray, "sp.id = $%d")
		filterParams = append(filterParams, studyProgramId)
	}

	result := strings.Join(filterArray, " AND ")
	args := []interface{}{}
	for i := startIndex; i < len(filterParams); i++ {
		args = append(args, i+1)
	}
	result = fmt.Sprintf(result, args...)
	if result != "" {
		result = fmt.Sprintf("WHERE %s", result)
	}

	*params = filterParams

	return result
}

func (l examSupervisorRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studyProgramId string) ([]models.GetExamSupervisorList, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetExamSupervisorList{}
	var paginationResult common.Pagination

	params := []interface{}{}

	filterQuery := mapQueryFilterGetList(pagination.Search, studyProgramId, &params)
	queryGet := fmt.Sprintf("%s %s", getListQuery, filterQuery)
	queryCount := fmt.Sprintf("%s %s", countListQuery, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"es.name": constants.Ascending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (l examSupervisorRepository) GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetExamSupervisorDetail, *constants.ErrorResponse) {
	results := []models.GetExamSupervisorDetail{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailQuery,
		id,
	)
	if err != nil {
		return models.GetExamSupervisorDetail{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetExamSupervisorDetail{}, utils.ErrDataNotFound("exam supervisor")
	}

	return results[0], nil
}

func (l examSupervisorRepository) GetDetailByIds(ctx context.Context, tx *sqlx.Tx, ids []string) ([]models.GetExamSupervisorDetail, *constants.ErrorResponse) {
	results := []models.GetExamSupervisorDetail{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailByIdsQuery,
		pq.Array(ids),
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (l examSupervisorRepository) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateExamSupervisor) *constants.ErrorResponse {
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

func (l examSupervisorRepository) Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateExamSupervisor) *constants.ErrorResponse {
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

func (l examSupervisorRepository) Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse {
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

func (e examSupervisorRepository) GetExamLectureSupervisorByLectureIds(ctx context.Context, tx *sqlx.Tx, lectureIds []string) ([]models.GetExamLectureSupervisor, *constants.ErrorResponse) {
	results := []models.GetExamLectureSupervisor{}

	err := tx.SelectContext(
		ctx,
		&results,
		getExamLectureSupervisorByLectureIdsQuery,
		pq.Array(lectureIds),
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (e examSupervisorRepository) DeleteExamLectureSupervisorExcludingExamLectureIds(ctx context.Context, tx *sqlx.Tx, lectureId string, excludedExamSupervisorIds []string) *constants.ErrorResponse {
	params := []interface{}{
		lectureId,
	}
	query := fmt.Sprintf(deleteExamLectureSupervisorExcludingExamSupervisorIdsQuery, "")
	if len(excludedExamSupervisorIds) != 0 {
		additionalQuery := "AND exam_supervisor_id NOT IN (SELECT UNNEST($2::uuid[]))"

		params = append(params, pq.Array(excludedExamSupervisorIds))
		query = fmt.Sprintf(deleteExamLectureSupervisorExcludingExamSupervisorIdsQuery, additionalQuery)
	}

	_, err := tx.ExecContext(
		ctx,
		query,
		params...,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (c examSupervisorRepository) UpsertExamLectureSupervisor(ctx context.Context, tx *sqlx.Tx, data []models.UpsertExamLectureSupervisor) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		upsertExamLectureSupervisorQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}
