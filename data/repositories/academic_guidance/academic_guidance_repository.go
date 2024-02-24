package academic_guidance

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

type academicGuidanceRepository struct {
	*db.DB
}

func mapQueryFilterGetListStudent(search string, params *[]interface{}) string {
	filterArray := []string{
		"ag.lecturer_id = $%d",
		"ag.semester_id = $%d",
	}
	filterParams := *params

	if search != "" {
		filterArray = append(filterArray, "s.name ILIKE $%d")
		filterParams = append(filterParams, fmt.Sprintf("%%%s%%", search))
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

func mapQueryFilterGetSession(academicGuidanceId, semesterId, lecturerId string, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	if academicGuidanceId != "" {
		filterArray = append(filterArray, "ag.id = $%d")
		filterParams = append(filterParams, academicGuidanceId)
	}
	if semesterId != "" {
		filterArray = append(filterArray, "ag.semester_id = $%d")
		filterParams = append(filterParams, semesterId)
	}
	if lecturerId != "" {
		filterArray = append(filterArray, "ag.lecturer_id = $%d")
		filterParams = append(filterParams, lecturerId)
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

func (a academicGuidanceRepository) GetListStudent(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, lecturerId, semesterId string) ([]models.GetAcademicGuidanceStudentList, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetAcademicGuidanceStudentList{}
	var paginationResult common.Pagination

	params := []interface{}{
		lecturerId,
		semesterId,
	}
	filterQuery := mapQueryFilterGetListStudent(pagination.Search, &params)
	queryGet := fmt.Sprintf("%s %s", getListStudentQuery, filterQuery)
	queryCount := fmt.Sprintf("%s %s", countListStudentQuery, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"s.name": constants.Ascending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (a academicGuidanceRepository) GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetAcademicGuidanceDetail, *constants.ErrorResponse) {
	results := []models.GetAcademicGuidanceDetail{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailQuery,
		id,
	)
	if err != nil {
		return models.GetAcademicGuidanceDetail{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetAcademicGuidanceDetail{}, utils.ErrDataNotFound("academic guidance")
	}

	return results[0], nil
}

func (a academicGuidanceRepository) GetDetailBySemesterIdLecturerId(ctx context.Context, tx *sqlx.Tx, semesterId, lecturerId string) (models.GetAcademicGuidanceDetail, *constants.ErrorResponse) {
	results := []models.GetAcademicGuidanceDetail{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailBySemesterIdLecturerIdQuery,
		semesterId,
		lecturerId,
	)
	if err != nil {
		return models.GetAcademicGuidanceDetail{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetAcademicGuidanceDetail{}, nil
	}

	return results[0], nil
}

func (a academicGuidanceRepository) GetDetailBySemesterIdStudentId(ctx context.Context, tx *sqlx.Tx, semesterId, studentId string) (models.GetAcademicGuidanceDetail, *constants.ErrorResponse) {
	results := []models.GetAcademicGuidanceDetail{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailBySemesterIdStudentIdQuery,
		semesterId,
		studentId,
	)
	if err != nil {
		return models.GetAcademicGuidanceDetail{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetAcademicGuidanceDetail{}, nil
	}

	return results[0], nil
}

func (a academicGuidanceRepository) Upsert(ctx context.Context, tx *sqlx.Tx, data models.UpsertAcademicGuidance) (string, *constants.ErrorResponse) {
	var result string
	err := tx.QueryRowContext(
		ctx,
		upsertQuery,
		data.SemesterId,
		data.LecturerId,
		data.CreatedBy,
	).Scan(&result)
	if err != nil {
		return result, constants.ErrorInternalServer(err.Error())
	}

	return result, nil
}

func (a academicGuidanceRepository) UpsertStudent(ctx context.Context, tx *sqlx.Tx, data []models.UpsertAcademicGuidanceStudent) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		upsertStudentQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (a academicGuidanceRepository) UpsertDecision(ctx context.Context, tx *sqlx.Tx, data models.UpsertDecisionAcademicGuidance) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		upsertDecisionQuery,
		data.LecturerId,
		data.SemesterId,
		data.DecisionNumber,
		data.DecisionDate,
		data.CreatedBy,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (a academicGuidanceRepository) GetSession(ctx context.Context, tx *sqlx.Tx, academicGuidanceId, semesterId, lecturerId string) ([]models.GetAcademicGuidanceSession, *constants.ErrorResponse) {
	var result []models.GetAcademicGuidanceSession

	var params []interface{}
	filterQuery := mapQueryFilterGetSession(academicGuidanceId, semesterId, lecturerId, &params)

	query := fmt.Sprintf("%s %s", getSessionQuery, filterQuery)

	err := tx.SelectContext(
		ctx,
		&result,
		query,
		params...,
	)
	if err != nil {
		return result, constants.ErrorInternalServer(err.Error())
	}

	return result, nil
}

func (a academicGuidanceRepository) GetSessionById(ctx context.Context, tx *sqlx.Tx, id string) (models.GetAcademicGuidanceSession, *constants.ErrorResponse) {
	var result []models.GetAcademicGuidanceSession

	err := tx.SelectContext(
		ctx,
		&result,
		getSessionByIdQuery,
		id,
	)
	if err != nil {
		return models.GetAcademicGuidanceSession{}, constants.ErrorInternalServer(err.Error())
	}
	if len(result) == 0 {
		return models.GetAcademicGuidanceSession{}, utils.ErrDataNotFound("academic guidance session")
	}

	return result[0], nil
}

func (a academicGuidanceRepository) GetSessionStudent(ctx context.Context, tx *sqlx.Tx, academicGuidanceSessionIds []string) ([]models.GetAcademicGuidanceSessionStudent, *constants.ErrorResponse) {
	var result []models.GetAcademicGuidanceSessionStudent

	err := tx.SelectContext(
		ctx,
		&result,
		getSessionStudentQuery,
		pq.Array(academicGuidanceSessionIds),
	)
	if err != nil {
		return result, constants.ErrorInternalServer(err.Error())
	}

	return result, nil
}

func (a academicGuidanceRepository) GetSessionFile(ctx context.Context, tx *sqlx.Tx, academicGuidanceSessionIds []string) ([]models.GetAcademicGuidanceSessionFile, *constants.ErrorResponse) {
	var result []models.GetAcademicGuidanceSessionFile

	err := tx.SelectContext(
		ctx,
		&result,
		getSessionFileQuery,
		pq.Array(academicGuidanceSessionIds),
	)
	if err != nil {
		return result, constants.ErrorInternalServer(err.Error())
	}

	return result, nil
}

func (a academicGuidanceRepository) UpsertSession(ctx context.Context, tx *sqlx.Tx, data models.UpsertAcademicGuidanceSession) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		upsertSessionQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (a academicGuidanceRepository) DeleteSessionStudentExcludingStudentIds(ctx context.Context, tx *sqlx.Tx, academicGuidanceSessionId string, studentIds []string) *constants.ErrorResponse {
	params := []interface{}{
		academicGuidanceSessionId,
	}

	var additionalFilter string
	if len(studentIds) != 0 {
		additionalFilter = `AND student_id NOT IN (SELECT UNNEST($2::uuid[]))`
		params = append(params, pq.Array(studentIds))
	}
	query := fmt.Sprintf(deleteSessionStudentExcludingStudentIdsQuery, additionalFilter)

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

func (a academicGuidanceRepository) DeleteSessionFileExcludingFilePaths(ctx context.Context, tx *sqlx.Tx, academicGuidanceSessionId string, filePaths []string) *constants.ErrorResponse {
	params := []interface{}{
		academicGuidanceSessionId,
	}

	var additionalFilter string
	if len(filePaths) != 0 {
		additionalFilter = `AND file_path NOT IN (SELECT UNNEST($2::text[]))`
		params = append(params, pq.Array(filePaths))
	}
	query := fmt.Sprintf(deleteSessionFileExcludingFilePathsQuery, additionalFilter)

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

func (a academicGuidanceRepository) UpsertSessionStudent(ctx context.Context, tx *sqlx.Tx, data []models.UpsertAcademicGuidanceSessionStudent) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		upsertSessionStudentQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (a academicGuidanceRepository) UpsertSessionFile(ctx context.Context, tx *sqlx.Tx, data []models.UpsertAcademicGuidanceSessionFile) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		upsertSessionFileQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (a academicGuidanceRepository) DeleteSession(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		deleteSessionQuery,
		id,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}
