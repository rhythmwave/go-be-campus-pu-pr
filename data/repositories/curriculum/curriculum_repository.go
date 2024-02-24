package curriculum

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

type curriculumRepository struct {
	*db.DB
}

func mapQueryFilterGetList(search string, studyProgramIds []string, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	if search != "" {
		searchIlike := fmt.Sprintf("%%%s%%", search)
		filterArray = append(filterArray, "(c.name ILIKE $%d OR c.rector_decision_number ILIKE $%d)")
		filterParams = append(filterParams, searchIlike, searchIlike)
	}
	if len(studyProgramIds) > 0 {
		filterArray = append(filterArray, "c.study_program_id IN (SELECT UNNEST($%d::uuid[]))")
		filterParams = append(filterParams, pq.Array(studyProgramIds))
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

func (a curriculumRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studyProgramIds []string) ([]models.GetCurriculum, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetCurriculum{}
	var paginationResult common.Pagination

	params := []interface{}{}
	filterQuery := mapQueryFilterGetList(pagination.Search, studyProgramIds, &params)
	queryGet := fmt.Sprintf("%s %s", getListQuery, filterQuery)
	queryCount := fmt.Sprintf("%s %s", countListQuery, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"c.year": constants.Descending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (a curriculumRepository) GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetCurriculumDetail, *constants.ErrorResponse) {
	results := []models.GetCurriculumDetail{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailQuery,
		id,
	)
	if err != nil {
		return models.GetCurriculumDetail{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetCurriculumDetail{}, utils.ErrDataNotFound("curriculum")
	}

	return results[0], nil
}

func (a curriculumRepository) GetActiveByStudyProgramId(ctx context.Context, tx *sqlx.Tx, studyProgramId string) (models.GetCurriculumDetail, *constants.ErrorResponse) {
	results := []models.GetCurriculumDetail{}

	err := tx.SelectContext(
		ctx,
		&results,
		getActiveByStudyProgramIdQuery,
		studyProgramId,
	)
	if err != nil {
		return models.GetCurriculumDetail{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetCurriculumDetail{}, utils.ErrDataNotFound("curriculum")
	}

	return results[0], nil
}

func (a curriculumRepository) GetActive(ctx context.Context, tx *sqlx.Tx) ([]models.GetCurriculumDetail, *constants.ErrorResponse) {
	results := []models.GetCurriculumDetail{}

	err := tx.SelectContext(
		ctx,
		&results,
		getActiveQuery,
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (a curriculumRepository) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateCurriculum) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		createQuery,
		data.StudyProgramId,
		data.Name,
		data.Year,
		data.RectorDecisionNumber,
		data.RectorDecisionDate,
		data.AggreeingParty,
		data.AggreementDate,
		data.IdealStudyPeriod,
		data.MaximumStudyPeriod,
		data.Remarks,
		data.IsActive,
		data.FinalScoreDeterminant,
		data.CreatedBy,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (a curriculumRepository) Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateCurriculum) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		updateQuery,
		data.Name,
		data.Year,
		data.RectorDecisionNumber,
		data.RectorDecisionDate,
		data.AggreeingParty,
		data.AggreementDate,
		data.IdealStudyPeriod,
		data.MaximumStudyPeriod,
		data.Remarks,
		data.IsActive,
		data.FinalScoreDeterminant,
		data.UpdatedBy,
		data.Id,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (a curriculumRepository) Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse {
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
