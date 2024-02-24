package student_class

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

type studentClassRepository struct {
	*db.DB
}

func mapQueryFilterGetList(search, studyPlanId, studentId, semesterId string, isMbkm *bool, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	if search != "" {
		searchParams := fmt.Sprintf("%%%s%%", search)
		filterArray = append(filterArray, "(s.name ILIKE $%d OR s.code ILIKE $%d)")
		filterParams = append(filterParams, searchParams, searchParams)
	}
	if studyPlanId != "" {
		filterArray = append(filterArray, "sc.study_plan_id = $%d")
		filterParams = append(filterParams, studyPlanId)
	}
	if studentId != "" {
		filterArray = append(filterArray, "sp.student_id = $%d")
		filterParams = append(filterParams, studentId)
	}
	if semesterId != "" {
		filterArray = append(filterArray, "sp.semester_id = $%d")
		filterParams = append(filterParams, semesterId)
	}
	if isMbkm != nil {
		filterArray = append(filterArray, "s.is_mbkm = $%d")
		filterParams = append(filterParams, *isMbkm)
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

func mapQueryFilterGetClassParticipant(search string, classIds []string, isGraded *bool, studentId string, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	if search != "" {
		searchParams := fmt.Sprintf("%%%s%%", search)
		filterArray = append(filterArray, "(s.name ILIKE $%d OR s.nim_nunmber::text ILIKE $%d)")
		filterParams = append(filterParams, searchParams, searchParams)
	}
	if len(classIds) != 0 {
		filterArray = append(filterArray, "sc.class_id IN (SELECT UNNEST($%d::uuid[]))")
		filterParams = append(filterParams, pq.Array(classIds))
	}
	if isGraded != nil {
		filterArray = append(filterArray, "(sc.graded_at IS NOT NULL) = $%d")
		filterParams = append(filterParams, *isGraded)
	}
	if studentId != "" {
		filterArray = append(filterArray, "sc.student_id = $%d")
		filterParams = append(filterParams, studentId)
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

func mapQueryGetStudentClassByStudentIdClassId(data []models.StudentIdClassId, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	for _, v := range data {
		filterArray = append(filterArray, "(sc.student_id = $%d AND sc.class_id = $%d)")
		filterParams = append(filterParams, v.StudentId, v.ClassId)
	}

	result := strings.Join(filterArray, " OR  ")
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

func (s studentClassRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studyPlanId, studentId, semesterId string, isMbkm *bool) ([]models.GetStudentClass, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetStudentClass{}
	var paginationResult common.Pagination

	params := []interface{}{}
	filterQuery := mapQueryFilterGetList(pagination.Search, studyPlanId, studentId, semesterId, isMbkm, &params)
	queryGet := fmt.Sprintf("%s %s", getListQuery, filterQuery)
	queryCount := fmt.Sprintf("%s %s", countListQuery, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"s.code": constants.Ascending, "s.name": constants.Ascending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (s studentClassRepository) GetClassParticipant(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, classIds []string, lectureId string, isGraded *bool, studentId string) ([]models.GetClassParticipant, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetClassParticipant{}
	var paginationResult common.Pagination

	params := []interface{}{}

	queryGet := fmt.Sprintf(getClassParticipantQuery, "", "")
	queryCount := fmt.Sprintf(countClassParticipantQuery, "")
	if lectureId != "" {
		errs := utils.IsValidUUID(lectureId)
		if errs != nil {
			return resultData, paginationResult, errs
		}
		additionalSelect := `,
			lp.is_attend,
			lp.is_sick,
			lp.is_leave,
			lp.is_awol
		`

		additionalJoin := fmt.Sprintf(`
			LEFT JOIN lecture_participants lp ON lp.student_id = s.id AND lp.lecture_id = '%s'
		`, lectureId)

		queryGet = fmt.Sprintf(getClassParticipantQuery, additionalSelect, additionalJoin)
		queryCount = fmt.Sprintf(countClassParticipantQuery, additionalJoin)
	}

	filterQuery := mapQueryFilterGetClassParticipant(pagination.Search, classIds, isGraded, studentId, &params)
	queryGet = fmt.Sprintf("%s %s", queryGet, filterQuery)
	queryCount = fmt.Sprintf("%s %s", queryCount, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"s.nim_number": constants.Ascending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (s studentClassRepository) BulkCreate(ctx context.Context, tx *sqlx.Tx, data []models.CreateStudentClass) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		bulkCreateQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (s studentClassRepository) GetStudentClassByStudentIdClassId(ctx context.Context, tx *sqlx.Tx, data []models.StudentIdClassId) ([]models.GetStudentClass, *constants.ErrorResponse) {
	results := []models.GetStudentClass{}

	params := []interface{}{}
	filterQuery := mapQueryGetStudentClassByStudentIdClassId(data, &params)
	queryGet := fmt.Sprintf("%s %s", getListQuery, filterQuery)
	err := tx.SelectContext(
		ctx,
		&results,
		queryGet,
		params...,
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (s studentClassRepository) BulkUpdateClass(ctx context.Context, tx *sqlx.Tx, data []models.CreateStudentClass) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		bulkUpdateClassQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (s studentClassRepository) BulkGradeStudentClass(ctx context.Context, tx *sqlx.Tx, data []models.GradeStudentClass) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		bulkGradeStudentClassQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (s studentClassRepository) GetStudentClassGradeByClassIdAndStudentIds(ctx context.Context, tx *sqlx.Tx, classId string, studentIds []string) ([]models.GetStudentClassGrade, *constants.ErrorResponse) {
	results := []models.GetStudentClassGrade{}
	err := tx.SelectContext(
		ctx,
		&results,
		getStudentClassGradeByClassIdAndStudentIdsQuery,
		classId,
		pq.Array(studentIds),
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (s studentClassRepository) BulkDeleteExcludingClassIds(ctx context.Context, tx *sqlx.Tx, data []models.DeleteStudentClassExcludingClassIds) *constants.ErrorResponse {
	filterArray := []string{}
	params := []interface{}{}

	for _, v := range data {
		filterArray = append(filterArray, "(study_plan_id = $%d AND class_id NOT IN (SELECT UNNEST($%d::uuid[])))")
		params = append(params, v.StudyPlanId, pq.Array(v.ExcludedClassIds))
	}
	filterResult := strings.Join(filterArray, " OR ")
	args := []interface{}{}
	for i := 0; i < len(params); i++ {
		args = append(args, i+1)
	}
	filterResult = fmt.Sprintf(filterResult, args...)

	query := fmt.Sprintf(bulkDeleteExcludingClassIdsQuery, filterResult)

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

func (s studentClassRepository) UpdateMbkmConvertedCredit(ctx context.Context, tx *sqlx.Tx, id string, convertedCredit uint32) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		updateMbkmConvertedCreditQuery,
		convertedCredit,
		id,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}
