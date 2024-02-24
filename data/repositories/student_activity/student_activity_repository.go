package student_activity

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

type studentActivityRepository struct {
	*db.DB
}

func mapQueryFilterGetList(search, activityType, studyProgramId, semesterId string, isMbkm bool, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	filterArray = append(filterArray, "sa.is_mbkm = $%d")
	filterParams = append(filterParams, isMbkm)
	if search != "" {
		searchParams := fmt.Sprintf("%%%s%%", search)
		filterArray = append(filterArray, "sa.title ILIKE $%d")
		filterParams = append(filterParams, searchParams)
	}
	if activityType != "" {
		filterArray = append(filterArray, "sa.activity_type = $%d")
		filterParams = append(filterParams, activityType)
	}
	if studyProgramId != "" {
		filterArray = append(filterArray, "sa.study_program_id = $%d")
		filterParams = append(filterParams, studyProgramId)
	}
	if semesterId != "" {
		filterArray = append(filterArray, "sa.semester_id = $%d")
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

func (f studentActivityRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, activityType, studyProgramId, semesterId string, isMbkm bool) ([]models.GetStudentActivity, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetStudentActivity{}
	var paginationResult common.Pagination

	params := []interface{}{}
	filterQuery := mapQueryFilterGetList(pagination.Search, activityType, studyProgramId, semesterId, isMbkm, &params)
	queryGet := fmt.Sprintf("%s %s", getListQuery, filterQuery)
	queryCount := fmt.Sprintf("%s %s", countListQuery, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"sa.title": constants.Ascending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (f studentActivityRepository) GetListParticipantByStudentActivityId(ctx context.Context, tx *sqlx.Tx, studentActivityId string) ([]models.GetStudentActivityParticipant, *constants.ErrorResponse) {
	results := []models.GetStudentActivityParticipant{}

	err := tx.SelectContext(
		ctx,
		&results,
		getListParticipantByStudentActivityIdQuery,
		studentActivityId,
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (f studentActivityRepository) GetListLecturerByStudentActivityId(ctx context.Context, tx *sqlx.Tx, studentActivityId string) ([]models.GetStudentActivityLecturer, *constants.ErrorResponse) {
	results := []models.GetStudentActivityLecturer{}

	err := tx.SelectContext(
		ctx,
		&results,
		getListLecturerByStudentActivityIdQuery,
		studentActivityId,
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (f studentActivityRepository) GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetStudentActivityDetail, *constants.ErrorResponse) {
	results := []models.GetStudentActivityDetail{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailQuery,
		id,
	)
	if err != nil {
		return models.GetStudentActivityDetail{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetStudentActivityDetail{}, utils.ErrDataNotFound("student activity")
	}

	return results[0], nil
}

func (f studentActivityRepository) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateStudentActivity) (string, *constants.ErrorResponse) {
	var result string

	err := tx.QueryRowContext(
		ctx,
		createQuery,
		data.StudyProgramId,
		data.SemesterId,
		data.ActivityType,
		data.Title,
		data.Location,
		data.DecisionNumber,
		data.DecisionDate,
		data.IsGroupActivity,
		data.Remarks,
		data.IsMbkm,
		data.CreatedBy,
	).Scan(&result)
	if err != nil {
		return result, constants.ErrorInternalServer(err.Error())
	}

	return result, nil
}

func (s studentActivityRepository) DeleteParticipantExcludingStudentIds(ctx context.Context, tx *sqlx.Tx, studentActivityId string, excludedStudentIds []string) *constants.ErrorResponse {
	query := fmt.Sprintf(deleteParticipantExcludingStudentIdsQuery, "")
	params := []interface{}{
		studentActivityId,
	}
	if len(excludedStudentIds) != 0 {
		additionalQuery := "AND student_id NOT IN (SELECT UNNEST($2::uuid[]))"
		query = fmt.Sprintf(deleteParticipantExcludingStudentIdsQuery, additionalQuery)
		params = append(params, pq.Array(excludedStudentIds))
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

func (s studentActivityRepository) DeleteLecturerExcludingLecturerIds(ctx context.Context, tx *sqlx.Tx, studentActivityId, role string, excludedLecturerIds []string) *constants.ErrorResponse {
	query := fmt.Sprintf(deleteLecturerExcludingLecturerIdsQuery, "")
	params := []interface{}{
		studentActivityId,
		role,
	}
	if len(excludedLecturerIds) != 0 {
		additionalQuery := "AND lecturer_id NOT IN (SELECT UNNEST($3::uuid[]))"
		query = fmt.Sprintf(deleteLecturerExcludingLecturerIdsQuery, additionalQuery)
		params = append(params, pq.Array(excludedLecturerIds))
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

func (s studentActivityRepository) UpsertParticipant(ctx context.Context, tx *sqlx.Tx, data []models.UpsertStudentActivityParticipant) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		upsertParticipantQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (s studentActivityRepository) UpsertLecturer(ctx context.Context, tx *sqlx.Tx, data []models.UpsertStudentActivityLecturer) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		upsertLecturerQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (f studentActivityRepository) Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateStudentActivity) *constants.ErrorResponse {
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

func (f studentActivityRepository) Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse {
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
