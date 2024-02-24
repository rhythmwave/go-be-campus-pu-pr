package lecture

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	appConstants "github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
)

type lectureRepository struct {
	*db.DB
}

func mapQueryFilterGetList(classId, semesterId string, hasActualLecture, isExam *bool, examType string, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	if classId != "" {
		filterArray = append(filterArray, "l.class_id = $%d")
		filterParams = append(filterParams, classId)
	}
	if semesterId != "" {
		filterArray = append(filterArray, "c.semester_id = $%d")
		filterParams = append(filterParams, semesterId)
	}
	if hasActualLecture != nil {
		filterArray = append(filterArray, "(l.lecture_actual_date IS NOT NULL) = $%d")
		filterParams = append(filterParams, *hasActualLecture)
	}
	if isExam != nil {
		filterArray = append(filterArray, "l.is_exam = $%d")
		filterParams = append(filterParams, *isExam)
	}
	switch examType {
	case appConstants.MidtermExam:
		filterArray = append(filterArray, "l.is_midterm_exam IS true")
	case appConstants.EndtermExam:
		filterArray = append(filterArray, "l.is_endterm_exam IS true")
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

func mapQueryFilterGetHistory(studentId string, startDate, endDate time.Time, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	filterArray = append(filterArray, "lp.is_attend IS true")
	if studentId != "" {
		filterArray = append(filterArray, "lp.student_id = $%d")
		filterParams = append(filterParams, studentId)
	}
	if !startDate.IsZero() && !endDate.IsZero() {
		filterArray = append(filterArray, "l.lecture_actual_date BETWEEN $%d AND $%d")
		filterParams = append(filterParams, startDate, endDate)
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

func (f lectureRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, classId, semesterId string, hasActualLecture, isExam *bool, examType string) ([]models.GetLectureList, common.Pagination, *appConstants.ErrorResponse) {
	resultData := []models.GetLectureList{}
	var paginationResult common.Pagination

	params := []interface{}{}
	filterQuery := mapQueryFilterGetList(classId, semesterId, hasActualLecture, isExam, examType, &params)
	queryGet := fmt.Sprintf("%s %s", getListQuery, filterQuery)
	queryCount := fmt.Sprintf("%s %s", countListQuery, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"l.lecture_plan_date": appConstants.Ascending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
		return resultData, paginationResult, err
	}

	err := tx.SelectContext(
		ctx,
		&resultData,
		queryGet,
		params...,
	)
	if err != nil {
		return resultData, paginationResult, appConstants.ErrorInternalServer(err.Error())
	}

	var count int
	err = tx.QueryRowContext(
		ctx,
		queryCount,
		params...,
	).Scan(&count)
	if err != nil {
		return resultData, paginationResult, appConstants.ErrorInternalServer(err.Error())
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

func (f lectureRepository) GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetLectureDetail, *appConstants.ErrorResponse) {
	results := []models.GetLectureDetail{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailQuery,
		id,
	)
	if err != nil {
		return models.GetLectureDetail{}, appConstants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetLectureDetail{}, utils.ErrDataNotFound("lecture")
	}

	return results[0], nil
}

func (f lectureRepository) GetByClassIds(ctx context.Context, tx *sqlx.Tx, classIds []string) ([]models.GetLectureDetail, *appConstants.ErrorResponse) {
	results := []models.GetLectureDetail{}

	err := tx.SelectContext(
		ctx,
		&results,
		getByClassIdsQuery,
		pq.Array(classIds),
	)
	if err != nil {
		return results, appConstants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (f lectureRepository) BulkCreate(ctx context.Context, tx *sqlx.Tx, data []models.CreateLecture) *appConstants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		bulkCreateQuery,
		data,
	)
	if err != nil {
		return appConstants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (f lectureRepository) Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateLecture) *appConstants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		updateQuery,
		data,
	)
	if err != nil {
		return appConstants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (f lectureRepository) Delete(ctx context.Context, tx *sqlx.Tx, id string) *appConstants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		deleteQuery,
		id,
	)
	if err != nil {
		return appConstants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (f lectureRepository) ResetParticipation(ctx context.Context, tx *sqlx.Tx, id string) *appConstants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		resetParticipationQuery,
		id,
	)
	if err != nil {
		return appConstants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (f lectureRepository) GetParticipantByLectureId(ctx context.Context, tx *sqlx.Tx, lectureId string) ([]models.GetLectureParticipant, *appConstants.ErrorResponse) {
	results := []models.GetLectureParticipant{}

	err := tx.SelectContext(
		ctx,
		&results,
		getParticipantByLectureIdQuery,
		lectureId,
	)
	if err != nil {
		return results, appConstants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (f lectureRepository) BulkUpdateParticipant(ctx context.Context, tx *sqlx.Tx, data []models.UpdateLectureParticipant) *appConstants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		bulkUpdateParticipantQuery,
		data,
	)
	if err != nil {
		return appConstants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (f lectureRepository) GetStudentParticipation(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, classId, studentId string) ([]models.GetLectureParticipation, common.Pagination, *appConstants.ErrorResponse) {
	resultData := []models.GetLectureParticipation{}
	var paginationResult common.Pagination

	queryGet := getStudentParticipationQuery
	if err := utils.QueryOperation(&queryGet, map[string]string{"l.lecture_plan_date": appConstants.Ascending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
		return resultData, paginationResult, err
	}

	err := tx.SelectContext(
		ctx,
		&resultData,
		queryGet,
		classId,
		studentId,
	)
	if err != nil {
		return resultData, paginationResult, appConstants.ErrorInternalServer(err.Error())
	}

	var count int
	err = tx.QueryRowContext(
		ctx,
		countStudentParticipationQuery,
		classId,
		studentId,
	).Scan(&count)
	if err != nil {
		return resultData, paginationResult, appConstants.ErrorInternalServer(err.Error())
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

func (f lectureRepository) AttendLecture(ctx context.Context, tx *sqlx.Tx, lectureId, studentId string) *appConstants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		attendLectureQuery,
		lectureId,
		studentId,
	)
	if err != nil {
		return appConstants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (f lectureRepository) GetLectureCalendar(ctx context.Context, tx *sqlx.Tx, req objects.GetLectureCalendarRequest) ([]models.GetLectureCalendar, *appConstants.ErrorResponse) {
	var results []models.GetLectureCalendar

	startDate := time.Date(int(req.Year), time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 1, -1)

	params := []interface{}{
		startDate,
		endDate,
	}
	var classFilter string
	var roomFilter string
	var lecturerFilter string
	if req.ClassId != "" {
		errs := utils.IsValidUUID(req.ClassId)
		if errs != nil {
			return results, errs
		}
		classFilter = fmt.Sprintf("AND c.id = '%s'", req.ClassId)
	}
	if req.RoomId != "" {
		errs := utils.IsValidUUID(req.RoomId)
		if errs != nil {
			return results, errs
		}
		roomFilter = fmt.Sprintf("AND r.id = '%s'", req.RoomId)
	}
	if req.LecturerId != "" {
		errs := utils.IsValidUUID(req.LecturerId)
		if errs != nil {
			return results, errs
		}
		lecturerFilter = fmt.Sprintf("AND lr.id = '%s'", req.LecturerId)
	}

	query := fmt.Sprintf(getLectureCalendarQuery, classFilter, roomFilter, lecturerFilter)

	err := tx.SelectContext(
		ctx,
		&results,
		query,
		params...,
	)
	if err != nil {
		return results, appConstants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (f lectureRepository) GetHistory(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studentId string, startDate, endDate time.Time) ([]models.GetLectureHistory, common.Pagination, *appConstants.ErrorResponse) {
	resultData := []models.GetLectureHistory{}
	var paginationResult common.Pagination

	params := []interface{}{}
	filterQuery := mapQueryFilterGetHistory(studentId, startDate, endDate, &params)
	queryGet := fmt.Sprintf("%s %s", getHistoryQuery, filterQuery)
	queryCount := fmt.Sprintf("%s %s", countHistoryQuery, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"l.lecture_actual_date": appConstants.Descending, "lp.updated_at": appConstants.Descending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
		return resultData, paginationResult, err
	}

	err := tx.SelectContext(
		ctx,
		&resultData,
		queryGet,
		params...,
	)
	if err != nil {
		return resultData, paginationResult, appConstants.ErrorInternalServer(err.Error())
	}

	var count int
	err = tx.QueryRowContext(
		ctx,
		queryCount,
		params...,
	).Scan(&count)
	if err != nil {
		return resultData, paginationResult, appConstants.ErrorInternalServer(err.Error())
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
