package room

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
)

type roomRepository struct {
	*db.DB
}

func mapQueryFilterGetList(search string, req objects.GetRoomRequest, params *[]interface{}) string {
	filterArray := []string{
		"r.building_id = $%d",
	}
	filterParams := *params

	if search != "" {
		searchParams := fmt.Sprintf("%%%s%%", search)
		filterArray = append(filterArray, "(r.code ILIKE $%d OR r.name ILIKE $%d)")
		filterParams = append(filterParams, searchParams, searchParams)
	}
	if req.IsLaboratory != nil {
		filterArray = append(filterArray, "r.is_laboratory = $%d")
		filterParams = append(filterParams, *req.IsLaboratory)
	}
	if !req.ExcludeLectureDate.IsZero() && req.ExcludeStartTime != 0 && req.ExcludeEndTime != 0 {
		lectureFilter := "l.id IS NULL"
		if req.ForceIncludeLectureId != "" {
			lectureFilter = "(l.id IS NULL OR l.id = $%d)"
			filterParams = append(filterParams, req.ForceIncludeLectureId)
		}
		filterArray = append(filterArray, lectureFilter)
	}
	if req.MaximumParticipant != 0 {
		if req.ForExam {
			filterArray = append(filterArray, "r.exam_capacity >= $%d")
			filterParams = append(filterParams, req.MaximumParticipant)
		} else {
			filterArray = append(filterArray, "r.capacity >= $%d")
			filterParams = append(filterParams, req.MaximumParticipant)
		}
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

func mapQueryFilterGetSchedule(search string, roomIds []string, dayOfWeek uint32, semesterId string, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	filterArray = append(filterArray, "l.id IS NOT NULL")
	if search != "" {
		searchParams := fmt.Sprintf("%%%s%%", search)
		filterArray = append(filterArray, "r.name ILIKE $%d")
		filterParams = append(filterParams, searchParams)
	}
	if dayOfWeek != 0 {
		filterArray = append(filterArray, "l.lecture_plan_day_of_week = $%d")
		filterParams = append(filterParams, dayOfWeek)
	}
	if len(roomIds) != 0 {
		filterArray = append(filterArray, "r.id IN (SELECT UNNEST($%d::uuid[]))")
		filterParams = append(filterParams, pq.Array(roomIds))
	}
	if semesterId != "" {
		filterArray = append(filterArray, "c.semester_id = $%d")
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

func (a roomRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, req objects.GetRoomRequest) ([]models.GetRoom, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetRoom{}
	var paginationResult common.Pagination

	params := []interface{}{}

	var additionalJoin string
	if !req.ExcludeLectureDate.IsZero() && req.ExcludeStartTime != 0 && req.ExcludeEndTime != 0 {
		additionalJoin = fmt.Sprintf(`
			LEFT JOIN lectures l ON l.room_id = r.id AND l.lecture_plan_date = DATE('%s') AND %d < l.lecture_plan_end_time AND %d > l.lecture_plan_start_time
		`, req.ExcludeLectureDate.Format(constants.DateFormatStd), req.ExcludeStartTime, req.ExcludeEndTime)
	}
	params = append(params, req.BuildingId)
	queryGet := fmt.Sprintf(getListQuery, additionalJoin)
	queryCount := fmt.Sprintf(countListQuery, additionalJoin)

	filterQuery := mapQueryFilterGetList(pagination.Search, req, &params)
	queryGet = fmt.Sprintf("%s %s", queryGet, filterQuery)
	queryCount = fmt.Sprintf("%s %s", queryCount, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"r.code": constants.Ascending, "r.name": constants.Ascending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (a roomRepository) GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetRoomDetail, *constants.ErrorResponse) {
	results := []models.GetRoomDetail{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailQuery,
		id,
	)
	if err != nil {
		return models.GetRoomDetail{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetRoomDetail{}, utils.ErrDataNotFound("room")
	}

	return results[0], nil
}

func (a roomRepository) GetDetailByRoomIds(ctx context.Context, tx *sqlx.Tx, ids []string) ([]models.GetRoomDetail, *constants.ErrorResponse) {
	results := []models.GetRoomDetail{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailByRoomIdsQuery,
		pq.Array(ids),
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (a roomRepository) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateRoom) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		createQuery,
		data.BuildingId,
		data.Code,
		data.Name,
		data.Capacity,
		data.ExamCapacity,
		data.IsUsable,
		data.Area,
		data.PhoneNumber,
		data.Facility,
		data.Remarks,
		data.Purpose,
		data.Owner,
		data.Location,
		data.StudyProgramId,
		data.IsLaboratory,
		data.CreatedBy,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (a roomRepository) Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateRoom) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		updateQuery,
		data.Code,
		data.Name,
		data.Capacity,
		data.ExamCapacity,
		data.IsUsable,
		data.Area,
		data.PhoneNumber,
		data.Facility,
		data.Remarks,
		data.Purpose,
		data.Owner,
		data.Location,
		data.StudyProgramId,
		data.UpdatedBy,
		data.Id,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (a roomRepository) Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse {
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

func (r roomRepository) GetSchedule(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, roomId string, dayOfWeek uint32, semesterId string) ([]models.GetRoomSchedule, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetRoomSchedule{}
	var paginationResult common.Pagination

	params := []interface{}{}
	roomIds := []string{}
	if roomId != "" {
		roomIds = append(roomIds, roomId)
	}
	filterQuery := mapQueryFilterGetSchedule(pagination.Search, roomIds, dayOfWeek, semesterId, &params)
	queryGet := fmt.Sprintf("%s %s", getScheduleQuery, filterQuery)
	queryCount := fmt.Sprintf(countScheduleQuery, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"r.name": constants.Ascending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (r roomRepository) GetScheduleByRoomIds(ctx context.Context, tx *sqlx.Tx, roomIds []string, dayOfWeek uint32, semesterId string) ([]models.GetRoomScheduleDetail, *constants.ErrorResponse) {
	results := []models.GetRoomScheduleDetail{}

	params := []interface{}{}
	filterQuery := mapQueryFilterGetSchedule("", roomIds, dayOfWeek, semesterId, &params)
	query := fmt.Sprintf(getScheduleByRoomIdsQuery, filterQuery)
	err := tx.SelectContext(
		ctx,
		&results,
		query,
		params...,
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}
