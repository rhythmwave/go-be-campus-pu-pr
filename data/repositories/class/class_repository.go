package class

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

type classRepository struct {
	*db.DB
}

func mapQueryFilterGetList(search string, in objects.GetClassListRequest, params *[]interface{}) string {
	filterArray := []string{
		"cu.study_program_id = $%d",
	}
	filterParams := *params

	filterArray = append(filterArray, "s.is_mbkm = $%d")
	filterParams = append(filterParams, in.IsMbkm)
	if search != "" {
		searchParams := fmt.Sprintf("%%%s%%", search)
		filterArray = append(filterArray, "(c.name ILIKE $%d OR s.name ILIKE $%d OR s.code ILIKE $%d)")
		filterParams = append(filterParams, searchParams, searchParams, searchParams)
	}
	if in.SemesterId != "" {
		filterArray = append(filterArray, "c.semester_id = $%d")
		filterParams = append(filterParams, in.SemesterId)
	}
	if in.IsActive != nil {
		filterArray = append(filterArray, "c.is_active = $%d")
		filterParams = append(filterParams, *in.IsActive)
	}
	if in.ClassName != "" {
		className := fmt.Sprintf("%%%s%%", in.ClassName)
		filterArray = append(filterArray, "c.name ILIKE $%d")
		filterParams = append(filterParams, className)
	}
	if in.SubjectName != "" {
		subjectName := fmt.Sprintf("%%%s%%", in.SubjectName)
		filterArray = append(filterArray, "s.name ILIKE $%d")
		filterParams = append(filterParams, subjectName)
	}
	if in.SubjectId != "" {
		filterArray = append(filterArray, "s.id = $%d")
		filterParams = append(filterParams, in.SubjectId)
	}
	if in.ForOddSemester != nil {
		filterArray = append(filterArray, "(se.semester_type = $%d) = $%d")
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

func (a classRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, in objects.GetClassListRequest) ([]models.GetClass, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetClass{}
	var paginationResult common.Pagination

	params := []interface{}{
		in.StudyProgramId,
	}
	filterQuery := mapQueryFilterGetList(pagination.Search, in, &params)
	queryGet := fmt.Sprintf("%s %s", getListQuery, filterQuery)
	queryCount := fmt.Sprintf("%s %s", countListQuery, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{pagination.SortBy: pagination.Sort}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (a classRepository) GetBySubjectIdsSemesterId(ctx context.Context, tx *sqlx.Tx, subjectIds []string, semesterId string) ([]models.GetClass, *constants.ErrorResponse) {
	results := []models.GetClass{}

	err := tx.SelectContext(
		ctx,
		&results,
		getBySubjectIdsSemesterIdQuery,
		pq.Array(subjectIds),
		semesterId,
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (a classRepository) GetDetail(ctx context.Context, tx *sqlx.Tx, id, lecturerId string) (models.GetClassDetail, *constants.ErrorResponse) {
	results := []models.GetClassDetail{}

	query := fmt.Sprintf(getDetailQuery, "", "")
	params := []interface{}{
		id,
	}
	if lecturerId != "" {
		additionalSelect := `,
			cl.is_grading_responsible
		`
		additionalJoin := `
			LEFT JOIN class_lecturers cl ON cl.class_id = c.id AND cl.lecturer_id = $2
		`
		query = fmt.Sprintf(getDetailQuery, additionalSelect, additionalJoin)
		params = append(params, lecturerId)
	}
	err := tx.SelectContext(
		ctx,
		&results,
		query,
		params...,
	)
	if err != nil {
		return models.GetClassDetail{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetClassDetail{}, utils.ErrDataNotFound("class")
	}

	return results[0], nil
}

func (a classRepository) GetDetailByIds(ctx context.Context, tx *sqlx.Tx, ids []string, lecturerId string) ([]models.GetClassDetail, *constants.ErrorResponse) {
	results := []models.GetClassDetail{}

	query := fmt.Sprintf(getDetailByIdsQuery, "", "")
	params := []interface{}{
		pq.Array(ids),
	}
	if lecturerId != "" {
		additionalSelect := `,
			cl.is_grading_responsible
		`
		additionalJoin := `
			LEFT JOIN class_lecturers cl ON cl.class_id = c.id AND cl.lecturer_id = $2
		`
		query = fmt.Sprintf(getDetailByIdsQuery, additionalSelect, additionalJoin)
		params = append(params, lecturerId)
	}
	err := tx.SelectContext(
		ctx,
		&results,
		query,
		params...,
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return results, utils.ErrDataNotFound("class")
	}

	return results, nil
}

func (a classRepository) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateClass) (string, *constants.ErrorResponse) {
	var result string
	err := tx.QueryRowContext(
		ctx,
		createQuery,
		data.SubjectId,
		data.SemesterId,
		data.Name,
		data.Scope,
		data.IsOnline,
		data.IsOffline,
		data.MinimumParticipant,
		data.MaximumParticipant,
		data.Remarks,
		data.ApplicationDeadline,
		data.CreatedBy,
	).Scan(&result)
	if err != nil {
		return result, constants.ErrorInternalServer(err.Error())
	}

	return result, nil
}

func (a classRepository) Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateClass) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		updateQuery,
		data.SubjectId,
		data.Name,
		data.Scope,
		data.IsOnline,
		data.IsOffline,
		data.MinimumParticipant,
		data.MaximumParticipant,
		data.Remarks,
		data.ApplicationDeadline,
		data.UpdatedBy,
		data.Id,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (a classRepository) UpsertMaximumParticipant(ctx context.Context, tx *sqlx.Tx, data []models.CreateClass) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		upsertMaximumParticipantQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}
	return nil
}

func (a classRepository) UpdateActivation(ctx context.Context, tx *sqlx.Tx, id string, isActive bool) *constants.ErrorResponse {
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

func (a classRepository) Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse {
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

func (c classRepository) GetClassLecturerByClassIds(ctx context.Context, tx *sqlx.Tx, classIds []string) ([]models.GetClassLecturer, *constants.ErrorResponse) {
	results := []models.GetClassLecturer{}

	err := tx.SelectContext(
		ctx,
		&results,
		getClassLecturerByClassIdsQuery,
		pq.Array(classIds),
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (c classRepository) GetClassLecturersBySemesterIdLecturerId(ctx context.Context, tx *sqlx.Tx, semesterId, lecturerId string) ([]models.GetClassLecturer, *constants.ErrorResponse) {
	results := []models.GetClassLecturer{}

	err := tx.SelectContext(
		ctx,
		&results,
		getClassLecturersBySemesterIdLecturerIdQuery,
		semesterId,
		lecturerId,
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (c classRepository) DeleteClassLecturerExcludingLecturerIds(ctx context.Context, tx *sqlx.Tx, classId string, excludedLecturerIds []string) *constants.ErrorResponse {
	params := []interface{}{
		classId,
	}
	query := fmt.Sprintf(deleteClassLecturerExcludingLecturerIdsQuery, "")
	if len(excludedLecturerIds) != 0 {
		additionalQuery := "AND lecturer_id NOT IN (SELECT UNNEST($2::uuid[]))"

		params = append(params, pq.Array(excludedLecturerIds))
		query = fmt.Sprintf(deleteClassLecturerExcludingLecturerIdsQuery, additionalQuery)
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

func (c classRepository) UpsertClassLecturer(ctx context.Context, tx *sqlx.Tx, data []models.UpsertClassLecturer) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		upsertClassLecturerQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (c classRepository) Duplicate(ctx context.Context, tx *sqlx.Tx, fromSemesterId, toSemesterId, adminId string) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		duplicateQuery,
		fromSemesterId,
		toSemesterId,
		adminId,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (c classRepository) DuplicateLecturer(ctx context.Context, tx *sqlx.Tx, fromSemesterId, toSemesterId, adminId string) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		duplicateLecturerQuery,
		fromSemesterId,
		toSemesterId,
		adminId,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (c classRepository) GetActiveBySemesterId(ctx context.Context, tx *sqlx.Tx, semesterId string) ([]models.GetClass, *constants.ErrorResponse) {
	results := []models.GetClass{}

	err := tx.SelectContext(
		ctx,
		&results,
		getActiveBySemesterIdQuery,
		semesterId,
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (c classRepository) InactivateClasses(ctx context.Context, tx *sqlx.Tx, classIds []string) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		inactivateClassesQuery,
		pq.Array(classIds),
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (c classRepository) GetThesisClass(ctx context.Context, tx *sqlx.Tx, studentId string, semesterId string) (models.GetClass, *constants.ErrorResponse) {
	results := []models.GetClass{}

	err := tx.SelectContext(
		ctx,
		&results,
		getThesisClassQuery,
		studentId,
		semesterId,
	)
	if err != nil {
		return models.GetClass{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetClass{}, utils.ErrDataNotFound("class")
	}

	return results[0], nil
}

func (c classRepository) UpsertThesisClass(ctx context.Context, tx *sqlx.Tx, subjectId, lecturerId, semesterId, adminId string) (string, *constants.ErrorResponse) {
	var resultId string

	err := tx.QueryRowContext(
		ctx,
		upsertThesisClassQuery,
		subjectId,
		lecturerId,
		semesterId,
		adminId,
	).Scan(&resultId)
	if err != nil {
		return resultId, constants.ErrorInternalServer(err.Error())
	}

	_, err = tx.ExecContext(
		ctx,
		upsertThesisClassLecturer,
		resultId,
		lecturerId,
		adminId,
	)
	if err != nil {
		return resultId, constants.ErrorInternalServer(err.Error())
	}

	return resultId, nil
}
