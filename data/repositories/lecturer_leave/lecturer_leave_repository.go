package lecturer_leave

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
)

type lecturerLeaveRepository struct {
	*db.DB
}

func mapQueryFilterGetList(search, studyProgramId, idNationalLecturer, semesterId string, isActive bool, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	filterArray = append(filterArray, "ll.is_active = $%d")
	filterParams = append(filterParams, isActive)

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
		filterArray = append(filterArray, "ll.semester_id = $%d")
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

func (l lecturerLeaveRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studyProgramId, idNationalLecturer, semesterId string, isActive bool) ([]models.GetLecturerLeave, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetLecturerLeave{}
	var paginationResult common.Pagination

	params := []interface{}{}
	filterQuery := mapQueryFilterGetList(pagination.Search, studyProgramId, idNationalLecturer, semesterId, isActive, &params)
	queryGet := fmt.Sprintf("%s %s", getListQuery, filterQuery)
	queryCount := fmt.Sprintf("%s %s", countListQuery, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"ll.start_date": constants.Descending, "l.name": constants.Ascending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (l lecturerLeaveRepository) GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetLecturerLeave, *constants.ErrorResponse) {
	results := []models.GetLecturerLeave{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailQuery,
		id,
	)
	if err != nil {
		return models.GetLecturerLeave{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetLecturerLeave{}, utils.ErrDataNotFound("lecturer leave")
	}

	return results[0], nil
}

func (l lecturerLeaveRepository) GetDetailByLecturerIdAndStartDate(ctx context.Context, tx *sqlx.Tx, lecturerId string, startDate time.Time) (models.GetLecturerLeave, *constants.ErrorResponse) {
	results := []models.GetLecturerLeave{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailByLecturerIdAndStartDateQuery,
		lecturerId,
		startDate,
	)
	if err != nil {
		return models.GetLecturerLeave{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetLecturerLeave{}, nil
	}

	return results[0], nil
}

func (l lecturerLeaveRepository) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateLecturerLeave) *constants.ErrorResponse {
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

func (l lecturerLeaveRepository) Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateLecturerLeave) *constants.ErrorResponse {
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

func (l lecturerLeaveRepository) End(ctx context.Context, tx *sqlx.Tx, id, adminId string) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		endQuery,
		id,
		adminId,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (l lecturerLeaveRepository) Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse {
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

func (l lecturerLeaveRepository) AutoSetActive(ctx context.Context, tx *sqlx.Tx) ([]models.LecturerId, *constants.ErrorResponse) {
	results := []models.LecturerId{}

	err := tx.QueryRowContext(
		ctx,
		autoSetActiveQuery,
	).Scan(&results)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}
