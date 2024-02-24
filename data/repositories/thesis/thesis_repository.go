package thesis

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/sccicitb/pupr-backend/constants"
	appConstants "github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
)

type thesisRepository struct {
	*db.DB
}

func mapQueryFilterGetList(search string, studyProgramId string, nimNumber int64, startSemesterId, status string, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	if search != "" {
		searchParams := fmt.Sprintf("%%%s%%", search)
		filterArray = append(filterArray, "(t.title ILIKE $%d OR t.topic ILIKE $%d OR t.english_title ILIKE $%d)")
		filterParams = append(filterParams, searchParams, searchParams, searchParams)
	}
	if studyProgramId != "" {
		filterArray = append(filterArray, "s.study_program_id = $%d")
		filterParams = append(filterParams, studyProgramId)
	}
	if nimNumber != 0 {
		filterArray = append(filterArray, "s.nim_number = $%d")
		filterParams = append(filterParams, nimNumber)
	}
	if startSemesterId != "" {
		filterArray = append(filterArray, "t.start_semester_id = $%d")
		filterParams = append(filterParams, startSemesterId)
	}
	if status != "" {
		filterArray = append(filterArray, "t.status = $%d")
		filterParams = append(filterParams, status)
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

func (t thesisRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studyProgramId string, nimNumber int64, startSemesterId, status, supervisorLecturerId string) ([]models.GetListThesis, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetListThesis{}
	var paginationResult common.Pagination

	params := []interface{}{}

	var additionalJoin string
	if supervisorLecturerId != "" {
		errs := utils.IsValidUUID(supervisorLecturerId)
		if errs != nil {
			return resultData, paginationResult, errs
		}
		additionalJoin = fmt.Sprintf(`
			JOIN thesis_supervisors ts ON ts.thesis_id = t.id AND ts.lecturer_id = '%s'
		`, supervisorLecturerId)
	}
	queryGet := fmt.Sprintf(getListQuery, additionalJoin)
	queryCount := fmt.Sprintf(countListQuery, additionalJoin)

	filterQuery := mapQueryFilterGetList(pagination.Search, studyProgramId, nimNumber, startSemesterId, status, &params)
	queryGet = fmt.Sprintf("%s %s", queryGet, filterQuery)
	queryCount = fmt.Sprintf("%s %s", queryCount, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"t.created_at": constants.Descending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (t thesisRepository) GetById(ctx context.Context, tx *sqlx.Tx, id string) (models.GetDetailThesis, *constants.ErrorResponse) {
	var results []models.GetDetailThesis

	err := tx.SelectContext(
		ctx,
		&results,
		getByIdQuery,
		id,
	)
	if err != nil {
		return models.GetDetailThesis{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetDetailThesis{}, utils.ErrDataNotFound("thesis")
	}

	return results[0], nil
}

func (t thesisRepository) GetByStudentIdStatus(ctx context.Context, tx *sqlx.Tx, studentId, status string) (models.GetDetailThesis, *constants.ErrorResponse) {
	var results []models.GetDetailThesis

	err := tx.SelectContext(
		ctx,
		&results,
		getByStudentIdStatusQuery,
		studentId,
		status,
	)
	if err != nil {
		return models.GetDetailThesis{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetDetailThesis{}, nil
	}

	return results[0], nil
}

func (t thesisRepository) GetNonCancelled(ctx context.Context, tx *sqlx.Tx, studentId string) (models.GetDetailThesis, *constants.ErrorResponse) {
	var results []models.GetDetailThesis

	err := tx.SelectContext(
		ctx,
		&results,
		getNonCancelledQuery,
		studentId,
		appConstants.ThesisStatusDibatalkan,
	)
	if err != nil {
		return models.GetDetailThesis{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetDetailThesis{}, nil
	}

	return results[0], nil
}

func (t thesisRepository) GetFileByThesisId(ctx context.Context, tx *sqlx.Tx, thesisId string) ([]models.GetThesisFile, *constants.ErrorResponse) {
	var results []models.GetThesisFile

	err := tx.SelectContext(
		ctx,
		&results,
		getFileByThesisIdQuery,
		thesisId,
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (t thesisRepository) GetSupervisorByThesisId(ctx context.Context, tx *sqlx.Tx, thesisId string) ([]models.GetThesisSupervisor, *constants.ErrorResponse) {
	var results []models.GetThesisSupervisor

	err := tx.SelectContext(
		ctx,
		&results,
		getSupervisorByThesisIdQuery,
		thesisId,
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (t thesisRepository) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateThesis) *constants.ErrorResponse {
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

func (t thesisRepository) UpsertFile(ctx context.Context, tx *sqlx.Tx, data []models.UpsertThesisFile) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		upsertFileQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (t thesisRepository) UpsertSupervisor(ctx context.Context, tx *sqlx.Tx, data []models.UpsertThesisSupervisor) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		upsertSupervisorQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (t thesisRepository) DeleteFileExcludingPaths(ctx context.Context, tx *sqlx.Tx, thesisId string, excludedPaths []string) *constants.ErrorResponse {
	params := []interface{}{
		thesisId,
	}

	var additionalFilter string
	if len(excludedPaths) != 0 {
		additionalFilter = `AND file_path NOT IN (SELECT UNNEST($2::text[]))`
		params = append(params, pq.Array(excludedPaths))
	}
	query := fmt.Sprintf(deleteFileExcludingPathsQuery, additionalFilter)

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

func (t thesisRepository) DeleteSupervisorExcludingLecturerIds(ctx context.Context, tx *sqlx.Tx, thesisId string, excludedLecturerIds []string) *constants.ErrorResponse {
	params := []interface{}{
		thesisId,
	}

	var additionalFilter string
	if len(excludedLecturerIds) != 0 {
		additionalFilter = `AND lecturer_id NOT IN (SELECT UNNEST($2::uuid[]))`
		params = append(params, pq.Array(excludedLecturerIds))
	}
	query := fmt.Sprintf(deleteSupervisorExcludingLecturerIdQuery, additionalFilter)

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

func (t thesisRepository) Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateThesis) *constants.ErrorResponse {
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

func (t thesisRepository) Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse {
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

func (t thesisRepository) GetListDefenseRequest(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studyProgramId string, nimNumber int64, startSemesterId string) ([]models.GetThesisDefenseRequest, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetThesisDefenseRequest{}
	var paginationResult common.Pagination

	params := []interface{}{}
	filterQuery := mapQueryFilterGetList(pagination.Search, studyProgramId, nimNumber, startSemesterId, "", &params)
	queryGet := fmt.Sprintf("%s %s", getListDefenseRequestQuery, filterQuery)
	queryCount := fmt.Sprintf("%s %s", countListDefenseRequestQuery, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"tdr.created_at": constants.Descending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (t thesisRepository) CreateDefenseRequest(ctx context.Context, tx *sqlx.Tx, thesisId string) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		createDefenseRequestQuery,
		thesisId,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (t thesisRepository) GetActiveDefenseRequest(ctx context.Context, tx *sqlx.Tx, thesisId string) (models.GetThesisDefenseRequest, *constants.ErrorResponse) {
	var results []models.GetThesisDefenseRequest
	err := tx.SelectContext(
		ctx,
		&results,
		getActiveDefenseRequestQuery,
		thesisId,
	)
	if err != nil {
		return models.GetThesisDefenseRequest{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetThesisDefenseRequest{}, nil
	}

	return results[0], nil
}

func (t thesisRepository) CreateDefense(ctx context.Context, tx *sqlx.Tx, data models.CreateThesisDefense) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		createDefenseQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (t thesisRepository) DeleteExaminerExcludingLecturerIds(ctx context.Context, tx *sqlx.Tx, thesisDefenseId string, excludedLecturerIds []string) *constants.ErrorResponse {
	params := []interface{}{
		thesisDefenseId,
	}

	var additionalFilter string
	if len(excludedLecturerIds) != 0 {
		additionalFilter = `AND lecturer_id NOT IN (SELECT UNNEST($2::uuid[]))`
		params = append(params, pq.Array(excludedLecturerIds))
	}
	query := fmt.Sprintf(deleteExaminerExcludingLecturerIdQuery, additionalFilter)

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

func (t thesisRepository) UpsertDefenseExaminer(ctx context.Context, tx *sqlx.Tx, data []models.UpsertThesisDefenseExaminer) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		upsertDefenseExaminerQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (t thesisRepository) GetDefenseById(ctx context.Context, tx *sqlx.Tx, id string) (models.GetThesisDefense, *constants.ErrorResponse) {
	var results []models.GetThesisDefense
	err := tx.SelectContext(
		ctx,
		&results,
		getDefenseByIdQuery,
		id,
	)
	if err != nil {
		return models.GetThesisDefense{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetThesisDefense{}, nil
	}

	return results[0], nil
}

func (t thesisRepository) UpdateDefense(ctx context.Context, tx *sqlx.Tx, data models.UpdateThesisDefense) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		updateDefenseQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (t thesisRepository) FinishDefense(ctx context.Context, tx *sqlx.Tx, data models.FinishThesisDefense) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		finishDefenseQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (t thesisRepository) GetThesisDefenseExaminerByThesisDefenseIds(ctx context.Context, tx *sqlx.Tx, thesisDefenseIds []string) ([]models.GetThesisDefenseExaminer, *constants.ErrorResponse) {
	var result []models.GetThesisDefenseExaminer

	err := tx.SelectContext(
		ctx,
		&result,
		getThesisDefenseExaminerByThesisDefenseIdsQuery,
		pq.Array(thesisDefenseIds),
	)
	if err != nil {
		return result, constants.ErrorInternalServer(err.Error())
	}

	return result, nil
}

// GetActiveSemesterThesisSupervisorLog(ctx context.Context, tx *sqlx.Tx, lecturerIds []string) ([]models.GetThesisSupervisorLog, *constants.ErrorResponse)
func (t thesisRepository) GetActiveSemesterThesisSupervisorLog(ctx context.Context, tx *sqlx.Tx, lecturerIds []string) ([]models.GetThesisSupervisorLog, *constants.ErrorResponse) {
	var result []models.GetThesisSupervisorLog

	err := tx.SelectContext(
		ctx,
		&result,
		getActiveSemesterThesisSupervisorLogQuery,
		pq.Array(lecturerIds),
		appConstants.ThesisStatusSedangDikerjakan,
	)
	if err != nil {
		return result, constants.ErrorInternalServer(err.Error())
	}

	return result, nil
}

func (t thesisRepository) GetThesisSupervisorLog(ctx context.Context, tx *sqlx.Tx, semesterId string, lecturerIds []string) ([]models.GetThesisSupervisorLog, *constants.ErrorResponse) {
	var result []models.GetThesisSupervisorLog

	err := tx.SelectContext(
		ctx,
		&result,
		getThesisSupervisorLogQuery,
		semesterId,
		pq.Array(lecturerIds),
	)
	if err != nil {
		return result, constants.ErrorInternalServer(err.Error())
	}

	return result, nil
}
