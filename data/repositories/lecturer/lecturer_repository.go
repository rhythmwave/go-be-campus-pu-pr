package lecturer

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

type lecturerRepository struct {
	*db.DB
}

func mapQueryFilterGetList(search string, req objects.GetLecturerRequest, params *[]interface{}) string {
	var startIndex int
	filterArray := []string{}
	filterParams := *params

	if search != "" {
		filterArray = append(filterArray, "l.name ILIKE $%d")
		filterParams = append(filterParams, fmt.Sprintf("%%%s%%", search))
	}
	if req.StudyProgramId != "" {
		filterArray = append(filterArray, "sp.id = $%d")
		filterParams = append(filterParams, req.StudyProgramId)
	}
	if req.IdNationalLecturer != "" {
		filterArray = append(filterArray, "l.id_national_lecturer = $%d")
		filterParams = append(filterParams, req.IdNationalLecturer)
	}
	if req.EmploymentStatus != "" {
		filterArray = append(filterArray, "l.employment_status = $%d")
		filterParams = append(filterParams, req.EmploymentStatus)
	}
	if req.Status != "" {
		filterArray = append(filterArray, "l.status = $%d")
		filterParams = append(filterParams, req.Status)
	}
	if req.HasAuthentication != nil {
		filterArray = append(filterArray, "(a.id IS NOT NULL) = $%d")
		filterParams = append(filterParams, *req.HasAuthentication)
	}
	if req.AcademicGuidanceSemesterId != "" {
		startIndex = 1
	}
	if !req.ExcludeLectureDate.IsZero() && req.ExcludeStartTime != 0 && req.ExcludeEndTime != 0 {
		lectureFilter := "ls.id IS NULL"
		if req.ForceIncludeLectureId != "" {
			lectureFilter = "(ls.id IS NULL OR ls.id = $%d)"
			filterParams = append(filterParams, req.ForceIncludeLectureId)
		}
		filterArray = append(filterArray, lectureFilter)
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

func mapQueryFilterGetSchedule(search, studyProgramId, idNationalLecturer, semesterId string, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	if search != "" {
		filterArray = append(filterArray, "lr.name ILIKE $%d")
		filterParams = append(filterParams, fmt.Sprintf("%%%s%%", search))
	}
	if studyProgramId != "" {
		filterArray = append(filterArray, "sp.id = $%d")
		filterParams = append(filterParams, studyProgramId)
	}
	if idNationalLecturer != "" {
		filterArray = append(filterArray, "lr.id_national_lecturer = $%d")
		filterParams = append(filterParams, idNationalLecturer)
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

func mapQueryFilterGetAssignedClass(semesterId, classId string, classIsActive *bool, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	if semesterId != "" {
		filterArray = append(filterArray, "c.semester_id = $%d")
		filterParams = append(filterParams, semesterId)
	}
	if classId != "" {
		filterArray = append(filterArray, "c.id = $%d")
		filterParams = append(filterParams, classId)
	}
	if classIsActive != nil {
		filterArray = append(filterArray, "c.is_active = $%d")
		filterParams = append(filterParams, *classIsActive)
	}

	result := strings.Join(filterArray, " AND ")
	args := []interface{}{}
	for i := 1; i < len(filterParams); i++ {
		args = append(args, i+1)
	}
	result = fmt.Sprintf(result, args...)
	if result != "" {
		result = fmt.Sprintf("WHERE %s", result)
	}

	*params = filterParams

	return result
}

func (l lecturerRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, req objects.GetLecturerRequest) ([]models.GetLecturerList, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetLecturerList{}
	var paginationResult common.Pagination

	params := []interface{}{}
	var additionalSelect string
	var additionalJoin string
	if req.AcademicGuidanceSemesterId != "" {
		additionalSelect = `, 
			ag.id AS academic_guidance_id,
			ag.total_student AS academic_guidance_total_student,
			ag.decision_number AS academic_guidance_decision_number,
			ag.decision_date AS academic_guidance_decision_date
		`
		additionalJoin = `LEFT JOIN academic_guidances ag ON ag.lecturer_id = l.id AND ag.semester_id = $1`

		params = append(params, req.AcademicGuidanceSemesterId)
	}
	if req.ClassId != "" {
		errs := utils.IsValidUUID(req.ClassId)
		if errs != nil {
			return resultData, paginationResult, errs
		}

		classJoin := fmt.Sprintf(`
			JOIN class_lecturers cl ON cl.class_id = '%s' AND cl.lecturer_id = l.id
		`, req.ClassId)

		additionalJoin = fmt.Sprintf("%s %s", additionalJoin, classJoin)
	}
	if !req.ExcludeLectureDate.IsZero() && req.ExcludeStartTime != 0 && req.ExcludeEndTime != 0 {
		lectureJoin := fmt.Sprintf(`
			LEFT JOIN lectures ls ON ls.lecturer_id = l.id AND ls.lecture_plan_date = DATE('%s') AND %d < ls.lecture_plan_end_time AND %d > ls.lecture_plan_start_time AND ls.id != '9aa5575c-d759-48a2-8521-8c88835e6f98'
		`, req.ExcludeLectureDate.Format(constants.DateFormatStd), req.ExcludeStartTime, req.ExcludeEndTime)

		additionalJoin = fmt.Sprintf("%s %s", additionalJoin, lectureJoin)
	}

	queryGet := fmt.Sprintf(getListQuery, additionalSelect, additionalJoin)
	queryCount := fmt.Sprintf(countListQuery, additionalJoin)

	filterQuery := mapQueryFilterGetList(pagination.Search, req, &params)
	queryGet = fmt.Sprintf("%s %s", queryGet, filterQuery)
	queryCount = fmt.Sprintf("%s %s", queryCount, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"l.name": constants.Ascending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
		return resultData, paginationResult, err
	}

	utils.QueryLog(queryGet, params...)
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

func (l lecturerRepository) GetSchedule(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studyProgramId, idNationalLecturer, semesterId string) ([]models.GetLecturerSchedule, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetLecturerSchedule{}
	var paginationResult common.Pagination

	params := []interface{}{}
	filterQuery := mapQueryFilterGetSchedule(pagination.Search, studyProgramId, idNationalLecturer, semesterId, &params)
	queryGet := fmt.Sprintf("%s %s", getScheduleQuery, filterQuery)
	queryCount := fmt.Sprintf("%s %s", countScheduleQuery, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"lr.name": constants.Ascending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (l lecturerRepository) GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetLecturerDetail, *constants.ErrorResponse) {
	results := []models.GetLecturerDetail{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailQuery,
		id,
	)
	if err != nil {
		return models.GetLecturerDetail{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetLecturerDetail{}, utils.ErrDataNotFound("lecturer")
	}

	return results[0], nil
}

func (l lecturerRepository) GetDetailByIds(ctx context.Context, tx *sqlx.Tx, ids []string) ([]models.GetLecturerDetail, *constants.ErrorResponse) {
	results := []models.GetLecturerDetail{}

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

func (l lecturerRepository) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateLecturer) *constants.ErrorResponse {
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

func (l lecturerRepository) Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateLecturer) *constants.ErrorResponse {
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

func (l lecturerRepository) Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse {
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

func (l lecturerRepository) UpdateStatus(ctx context.Context, tx *sqlx.Tx, ids []string, status string) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		updateStatusQuery,
		status,
		pq.Array(ids),
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (l lecturerRepository) GetAssignedClass(ctx context.Context, tx *sqlx.Tx, lecturerId, semesterId, classId string, classIsActive *bool) ([]models.GetLecturerAssignedClass, *constants.ErrorResponse) {
	results := []models.GetLecturerAssignedClass{}

	params := []interface{}{
		lecturerId,
	}

	filterQuery := mapQueryFilterGetAssignedClass(semesterId, classId, classIsActive, &params)
	query := fmt.Sprintf(getAssignedClassQuery, filterQuery)

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
