package semester

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

type semesterRepository struct {
	*db.DB
}

func mapQueryFilterGetList(studyProgramId, excludedId string, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	if studyProgramId != "" {
		filterArray = append(filterArray, "c.study_program_id = $%d")
		filterParams = append(filterParams, studyProgramId)
	}
	if excludedId != "" {
		filterArray = append(filterArray, "s.id != $%d")
		filterParams = append(filterParams, excludedId)
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

func (a semesterRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studyProgramId, excludedId string) ([]models.GetSemester, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetSemester{}
	var paginationResult common.Pagination

	params := []interface{}{}
	queryGet := getListQuery
	queryCount := countListQuery
	if studyProgramId != "" {
		additionalJoin := `
			JOIN semester_curriculum sc ON sc.semester_id = s.id 
			JOIN curriculums c ON c.id = sc.curriculum_id
		`
		queryGet = fmt.Sprintf("%s %s", getListQuery, additionalJoin)
		queryCount = fmt.Sprintf("%s %s", countListQuery, additionalJoin)
	}
	filterQuery := mapQueryFilterGetList(studyProgramId, excludedId, &params)
	queryGet = fmt.Sprintf("%s %s", queryGet, filterQuery)
	queryCount = fmt.Sprintf("%s %s", queryCount, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"s.start_date": constants.Descending, "s.is_active": constants.Descending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (a semesterRepository) GetById(ctx context.Context, tx *sqlx.Tx, id string) (models.GetSemesterDetail, *constants.ErrorResponse) {
	results := []models.GetSemesterDetail{}

	err := tx.SelectContext(
		ctx,
		&results,
		getByIdQuery,
		id,
	)
	if err != nil {
		return models.GetSemesterDetail{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetSemesterDetail{}, utils.ErrDataNotFound("semester")
	}

	return results[0], nil
}

func (a semesterRepository) GetActive(ctx context.Context, tx *sqlx.Tx) (models.GetSemesterDetail, *constants.ErrorResponse) {
	results := []models.GetSemesterDetail{}

	err := tx.SelectContext(
		ctx,
		&results,
		getActiveQuery,
	)
	if err != nil {
		return models.GetSemesterDetail{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetSemesterDetail{}, utils.ErrDataNotFound("semester")
	}

	return results[0], nil
}

func (a semesterRepository) GetCurriculumBySemesterIds(ctx context.Context, tx *sqlx.Tx, semesterIds []string) ([]models.GetSemesterCurriculum, *constants.ErrorResponse) {
	results := []models.GetSemesterCurriculum{}

	err := tx.SelectContext(
		ctx,
		&results,
		getCurriculumBySemesterIdsQuery,
		pq.Array(semesterIds),
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (a semesterRepository) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateSemester) (string, *constants.ErrorResponse) {
	var result string

	err := tx.QueryRowContext(
		ctx,
		createQuery,
		data.SemesterStartYear,
		data.SemesterType,
		data.StartDate,
		data.EndDate,
		data.MidtermStartDate,
		data.MidtermEndDate,
		data.EndtermStartDate,
		data.EndtermEndDate,
		data.StudyPlanInputStartDate,
		data.StudyPlanInputEndDate,
		data.StudyPlanApprovalStartDate,
		data.StudyPlanApprovalEndDate,
		data.ReferenceSemesterId,
		data.CheckMinimumGpa,
		data.CheckPassedCredit,
		data.DefaultCredit,
		data.GradingStartDate,
		data.GradingEndDate,
		data.CreatedBy,
	).Scan(&result)
	if err != nil {
		return result, constants.ErrorInternalServer(err.Error())
	}

	return result, nil
}

func (a semesterRepository) Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateSemester) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		updateQuery,
		data.SemesterStartYear,
		data.SemesterType,
		data.StartDate,
		data.EndDate,
		data.StudyPlanInputStartDate,
		data.StudyPlanInputEndDate,
		data.StudyPlanApprovalStartDate,
		data.StudyPlanApprovalEndDate,
		data.ReferenceSemesterId,
		data.CheckMinimumGpa,
		data.CheckPassedCredit,
		data.DefaultCredit,
		data.UpdatedBy,
		data.MidtermStartDate,
		data.MidtermEndDate,
		data.EndtermStartDate,
		data.EndtermEndDate,
		data.GradingStartDate,
		data.GradingEndDate,
		data.Id,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (a semesterRepository) DeleteCurriculumSemesterExcludingCurriculumId(ctx context.Context, tx *sqlx.Tx, semesterId string, excludedCurriculumIds []string) *constants.ErrorResponse {
	query := deleteCurriculumSemesterExcludingCurriculumIdQuery
	params := []interface{}{
		semesterId,
	}
	if len(excludedCurriculumIds) != 0 {
		additionalQuery := "AND curriculum_id NOT IN (SELECT UNNEST($2::uuid[]))"
		query = fmt.Sprintf("%s %s", deleteCurriculumSemesterExcludingCurriculumIdQuery, additionalQuery)
		params = append(params, pq.Array(excludedCurriculumIds))
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

func (a semesterRepository) UpsertCurriculum(ctx context.Context, tx *sqlx.Tx, data []models.UpsertSemesterCurriculum) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		upsertCurriculumQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (a semesterRepository) UpdateActivation(ctx context.Context, tx *sqlx.Tx, id string, isActive bool) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		updateActivationQuery,
		isActive,
		id,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (a semesterRepository) Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse {
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

func (s semesterRepository) AutoSetActive(ctx context.Context, tx *sqlx.Tx) (string, *constants.ErrorResponse) {
	var result string
	err := tx.QueryRowContext(
		ctx,
		autoSetActiveQuery,
	).Scan(&result)
	if err != nil {
		if err.Error() != constants.PGNoRows {
			return result, constants.ErrorInternalServer(err.Error())
		}
	}

	return result, nil
}

func (s semesterRepository) GetPreviousSemester(ctx context.Context, tx *sqlx.Tx, semesterId string) (models.GetSemesterDetail, *constants.ErrorResponse) {
	results := []models.GetSemesterDetail{}

	err := tx.SelectContext(
		ctx,
		&results,
		getPreviousSemesterQuery,
		semesterId,
	)
	if err != nil {
		return models.GetSemesterDetail{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetSemesterDetail{}, nil
	}

	return results[0], nil
}
