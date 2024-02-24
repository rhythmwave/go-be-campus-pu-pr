package student

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
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
)

type studentRepository struct {
	*db.DB
}

func mapQueryFilterGetList(req objects.GetStudentRequest, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	if req.StudyProgramId != "" {
		filterArray = append(filterArray, "sp.id = $%d")
		filterParams = append(filterParams, req.StudyProgramId)
	}
	if req.StudentForceFrom != 0 && req.StudentForceTo != 0 {
		filterArray = append(filterArray, "(s.student_force BETWEEN $%d AND $%d)")
		filterParams = append(filterParams, req.StudentForceFrom, req.StudentForceTo)
	}
	if req.NimNumberFrom != 0 && req.NimNumberTo != 0 {
		filterArray = append(filterArray, "(s.nim_number BETWEEN $%d AND $%d)")
		filterParams = append(filterParams, req.NimNumberFrom, req.NimNumberTo)
	}
	if req.Name != "" {
		filterArray = append(filterArray, "s.name ILIKE $%d")
		filterParams = append(filterParams, fmt.Sprintf("%%%s%%", req.Name))
	}
	if req.Address != "" {
		filterArray = append(filterArray, "s.address ILIKE $%d")
		filterParams = append(filterParams, fmt.Sprintf("%%%s%%", req.Address))
	}
	if req.RegencyId != "" {
		filterArray = append(filterArray, "r.id = $%d")
		filterParams = append(filterParams, req.RegencyId)
	}
	if len(req.Status) != 0 {
		filterArray = append(filterArray, "s.status IN (SELECT UNNEST($%d::text[]))")
		filterParams = append(filterParams, pq.Array(req.Status))
	}
	if req.HasAuthentication != nil {
		filterArray = append(filterArray, "(a.id IS NOT NULL) = $%d")
		filterParams = append(filterParams, *req.HasAuthentication)
	}
	if req.StudyPlanSemesterId != "" {
		if req.StudyPlanIsSubmitted != nil {
			filterArray = append(filterArray, "spl.is_submitted = $%d")
			filterParams = append(filterParams, req.StudyPlanIsSubmitted)
		}

		if req.StudyPlanIsApproved != nil {
			studyPlanIsApproved := *req.StudyPlanIsApproved
			filterArray = append(filterArray, "spl.is_approved = $%d")
			filterParams = append(filterParams, studyPlanIsApproved)
		}

		if req.HasStudyPlan != nil {
			hasStudyPlan := *req.HasStudyPlan
			filterArray = append(filterArray, "(spl.id IS NOT NULL) = $%d")
			filterParams = append(filterParams, hasStudyPlan)
		}
	}
	if req.StatusSemesterId != "" {
		filterArray = append(filterArray, "ssl.semester_id = $%d")
		filterParams = append(filterParams, req.StatusSemesterId)
	}
	if req.IsRegistered != nil {
		filterArray = append(filterArray, "s.is_registered = $%d")
		filterParams = append(filterParams, *req.IsRegistered)
	}
	if req.HasPaid != nil {
		filterArray = append(filterArray, "s.has_paid = $%d")
		filterParams = append(filterParams, *req.HasPaid)
	}
	if req.IsGraduateEligible != nil {
		filterArray = append(filterArray, "(s.total_credit >= sp.minimum_graduation_credit) = $%d")
		filterParams = append(filterParams, *req.IsGraduateEligible)
	}
	if req.IsThesisEligible != nil {
		filterArray = append(filterArray, "(s.total_credit >= sp.minimum_thesis_credit) = $%d")
		filterParams = append(filterParams, *req.IsThesisEligible)
	}
	if req.YudiciumSessionId != "" {
		filterArray = append(filterArray, "ys.yudicium_session_id = $%d")
		filterParams = append(filterParams, req.YudiciumSessionId)
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

func mapQueryFilterGetStudentSubject(search, studentId string, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	if search != "" {
		searchParams := fmt.Sprintf("%%%s%%", search)
		filterArray = append(filterArray, "(s.code ILIKE $%d OR s.name ILIKE $%d)")
		filterParams = append(filterParams, searchParams, searchParams)
	}
	if studentId != "" {
		filterArray = append(filterArray, "ss.student_id = $%d")
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

func (a studentRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, requestData objects.GetStudentRequest) ([]models.GetStudent, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetStudent{}
	var paginationResult common.Pagination

	params := []interface{}{
		appConstants.ThesisStatusBerhasilDiselesaikan,
	}

	var additionalSelect string
	var additionalJoin string
	if requestData.GetAcademicGuidance {
		additionalSelect = `,
			agl.id AS academic_guidance_lecturer_id,
			agl.name AS academic_guidance_lecturer_name,
			agl.front_title AS academic_guidance_lecturer_front_title,
			agl.back_degree AS academic_guidance_lecturer_back_degree,
			agse.id AS academic_guidance_semester_id,
			agse.semester_start_year AS academic_guidance_semester_start_year
		`
		additionalJoin = `
			LEFT JOIN academic_guidance_students ags ON ags.student_id = s.id
			LEFT JOIN semesters agse ON agse.id = ags.semester_id AND agse.is_active IS true
			LEFT JOIN academic_guidances ag ON ag.id = ags.academic_guidance_id
			LEFT JOIN lecturers agl ON agl.id = ag.lecturer_id
		`
	}
	if requestData.StudyPlanSemesterId != "" {
		additionalStudyPlanSelect := `,
			spl.id AS study_plan_id,
			spl.total_mandatory_credit AS study_plan_total_mandatory_credit,
			spl.total_optional_credit AS study_plan_total_optional_credit,
			spl.maximum_credit AS study_plan_maximum_credit,
			spl.is_approved AS study_plan_is_approved
		`

		errs := utils.IsValidUUID(requestData.StudyPlanSemesterId)
		if errs != nil {
			return resultData, paginationResult, errs
		}

		additionalStudyPlanJoin := fmt.Sprintf(`
			LEFT JOIN study_plans spl ON spl.student_id = s.id AND spl.semester_id = '%s'::uuid
		`, requestData.StudyPlanSemesterId)
		additionalSelect = fmt.Sprintf("%s%s", additionalSelect, additionalStudyPlanSelect)
		additionalJoin = fmt.Sprintf("%s%s", additionalJoin, additionalStudyPlanJoin)
	}
	if requestData.StatusSemesterId != "" {
		additionalLeaveSelect := `,
			ssl.status AS status_log,
			ssl.semester_id AS status_semester_id,
			ssls.semester_start_year AS status_semester_start_year,
			ssls.semester_type AS status_semester_type,
			ssl.status_date AS status_date,
			ssl.reference_number AS status_reference_number,
			ssl.purpose AS status_purpose,
			ssl.remarks AS status_remarks
		`

		errs := utils.IsValidUUID(requestData.StatusSemesterId)
		if errs != nil {
			return resultData, paginationResult, errs
		}

		additionalLeaveJoin := fmt.Sprintf(`
			LEFT JOIN student_status_logs ssl ON ssl.student_id = s.id AND ssl.semester_id = '%s'
			LEFT JOIN semesters ssls ON ssls.id = ssl.semester_id
		`, requestData.StatusSemesterId)
		additionalSelect = fmt.Sprintf("%s%s", additionalSelect, additionalLeaveSelect)
		additionalJoin = fmt.Sprintf("%s%s", additionalJoin, additionalLeaveJoin)
	}
	if requestData.YudiciumSessionId != "" {
		additionalYudiciumJoin := `
			JOIN yudicium_students ys ON ys.student_id = s.id
		`
		additionalJoin = fmt.Sprintf("%s%s", additionalJoin, additionalYudiciumJoin)
	}

	queryGet := fmt.Sprintf(getListQuery, additionalSelect, additionalJoin)
	queryCount := fmt.Sprintf(countListQuery, additionalJoin)

	if len(requestData.Status) != 0 {
		for _, v := range requestData.Status {
			var isExist bool
			for _, w := range appConstants.ValidStudentStatus() {
				if v == w {
					isExist = true
					break
				}
			}
			if !isExist {
				return resultData, paginationResult, appConstants.ErrInvalidStudentStatus
			}
		}
	}

	filterQuery := mapQueryFilterGetList(requestData, &params)
	queryGet = fmt.Sprintf("%s %s", queryGet, filterQuery)
	queryCount = fmt.Sprintf("%s %s", queryCount, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"s.name": constants.Descending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (a studentRepository) GetDetail(ctx context.Context, tx *sqlx.Tx, id string, previousSemesterId string) (models.GetStudent, *constants.ErrorResponse) {
	results := []models.GetStudent{}

	query := fmt.Sprintf(getDetailQuery, "")
	params := []interface{}{
		appConstants.ThesisStatusBerhasilDiselesaikan,
		id,
	}
	if previousSemesterId != "" {
		params = append(params, previousSemesterId)
		additionalJoin := `
			LEFT JOIN study_plans spl ON spl.student_id = s.id AND spl.semester_id = $3
		`
		query = fmt.Sprintf(getDetailQuery, additionalJoin)
	}

	err := tx.SelectContext(
		ctx,
		&results,
		query,
		params...,
	)
	if err != nil {
		return models.GetStudent{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetStudent{}, utils.ErrDataNotFound("student")
	}

	return results[0], nil
}

func (a studentRepository) GetDetailByIds(ctx context.Context, tx *sqlx.Tx, ids []string) ([]models.GetStudent, *constants.ErrorResponse) {
	results := []models.GetStudent{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailByIdsQuery,
		appConstants.ThesisStatusBerhasilDiselesaikan,
		pq.Array(ids),
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (a studentRepository) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateStudent) *constants.ErrorResponse {
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

func (a studentRepository) Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateStudent) *constants.ErrorResponse {
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

func (a studentRepository) Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse {
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

func (a studentRepository) GetActive(ctx context.Context, tx *sqlx.Tx, previousSemesterId string, studentIds []string) ([]models.GetActiveStudent, *constants.ErrorResponse) {
	results := []models.GetActiveStudent{}

	query := fmt.Sprintf(getActiveQuery, "", "")
	params := []interface{}{
		appConstants.StudentStatusAktif,
		appConstants.StudentStatusMbkm,
	}
	if previousSemesterId != "" {
		params = append(params, previousSemesterId)
		additionalSelect := `,
			sp.grade_point AS previous_semester_grade_point
		`
		additionalJoin := `
			LEFT JOIN study_plans sp ON sp.student_id = s.id AND sp.semester_id = $3
		`
		query = fmt.Sprintf(getActiveQuery, additionalSelect, additionalJoin)
	}
	if len(studentIds) != 0 {
		paramsLen := len(params)
		additionalFilter := fmt.Sprintf(`AND s.id IN (SELECT UNNEST($%d::uuid[]))`, paramsLen+1)
		params = append(params, pq.Array(studentIds))
		query = fmt.Sprintf("%s %s", query, additionalFilter)
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

	return results, nil
}

func (a studentRepository) UpdateActiveSemesterPackage(ctx context.Context, tx *sqlx.Tx, studentIds []string) *constants.ErrorResponse {
	params := []interface{}{
		appConstants.StudentStatusAktif,
	}

	query := updateActiveSemesterPackageQuery
	if len(studentIds) != 0 {
		additionalFilter := `AND id IN (SELECT UNNEST($2::uuid[]))`
		params = append(params, pq.Array(studentIds))

		query = fmt.Sprintf("%s %s", updateActiveSemesterPackageQuery, additionalFilter)
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

func (a studentRepository) BulkUpdateStatus(ctx context.Context, tx *sqlx.Tx, data models.BulkUpdateStatusStudent) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		bulkUpdateStatusQuery,
		data.Status,
		data.StatusReferenceNumber,
		data.StatusDate,
		data.StatusPurpose,
		data.StatusRemarks,
		pq.Array(data.Ids),
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (a studentRepository) GetStatusSummary(ctx context.Context, tx *sqlx.Tx, studyProgramIds []string, semesterId string) ([]models.StudentStatusSummary, *constants.ErrorResponse) {
	results := []models.StudentStatusSummary{}

	err := tx.SelectContext(
		ctx,
		&results,
		getStatusSummaryQuery,
		pq.Array(appConstants.ValidStudentStatus()),
		pq.Array(studyProgramIds),
		semesterId,
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (a studentRepository) UpdateProfile(ctx context.Context, tx *sqlx.Tx, data models.UpdateStudentProfile) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		updateProfileQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (a studentRepository) UpdateParentProfile(ctx context.Context, tx *sqlx.Tx, data models.UpdateStudentParentProfile) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		updateParentProfileQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (a studentRepository) GetPreHighSchoolHistoryByStudentIds(ctx context.Context, tx *sqlx.Tx, studentIds []string) ([]models.GetStudentPreHighSchoolHistory, *constants.ErrorResponse) {
	results := []models.GetStudentPreHighSchoolHistory{}

	err := tx.SelectContext(
		ctx,
		&results,
		getPreHighSchoolHistoryByStudentIdsQuery,
		pq.Array(studentIds),
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (a studentRepository) UpdateSchoolProfile(ctx context.Context, tx *sqlx.Tx, data models.UpdateStudentSchoolProfile) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		updateSchoolProfileQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (a studentRepository) DeletePreHighSchoolHistoryExcludingLevel(ctx context.Context, tx *sqlx.Tx, studentId string, excludedLevel []string) *constants.ErrorResponse {
	query := fmt.Sprintf(deletePreHighSchoolHistoryExcludingLevelQuery, "")
	params := []interface{}{
		studentId,
	}
	if len(excludedLevel) != 0 {
		additionalQuery := "AND level NOT IN (SELECT UNNEST($2::text[]))"
		query = fmt.Sprintf(deletePreHighSchoolHistoryExcludingLevelQuery, additionalQuery)
		params = append(params, pq.Array(excludedLevel))
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

func (a studentRepository) UpsertPreHighSchoolHistory(ctx context.Context, tx *sqlx.Tx, data []models.UpsertStudentPreHighSchoolHistory) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		upsertPreHighSchoolHistoryQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (s studentRepository) GetStudentSubject(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, studentId string) ([]models.GetStudentSubject, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetStudentSubject{}
	var paginationResult common.Pagination

	params := []interface{}{}
	filterQuery := mapQueryFilterGetStudentSubject(pagination.Search, studentId, &params)
	queryGet := fmt.Sprintf("%s %s", getStudentSubjectQuery, filterQuery)
	queryCount := fmt.Sprintf("%s %s", countStudentSubjectQuery, filterQuery)
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

func (s studentRepository) UpdatePayment(ctx context.Context, tx *sqlx.Tx, studentIds []string, adminId string) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		updatePaymentQuery,
		adminId,
		pq.Array(studentIds),
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (s studentRepository) GetPaymentLog(ctx context.Context, tx *sqlx.Tx, studentId string) ([]models.GetStudentPaymentLog, *constants.ErrorResponse) {
	var result []models.GetStudentPaymentLog

	err := tx.SelectContext(
		ctx,
		&result,
		getPaymentLogQuery,
		studentId,
	)
	if err != nil {
		return result, constants.ErrorInternalServer(err.Error())
	}

	return result, nil
}

func (s studentRepository) BulkCreate(ctx context.Context, tx *sqlx.Tx, data []models.BulkCreateStudent) *constants.ErrorResponse {
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

func (s studentRepository) ConvertGrade(ctx context.Context, tx *sqlx.Tx, studentId string, subjectIds []string, data models.ConvertStudentGrade) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		convertGradeQuery,
		studentId,
		pq.Array(subjectIds),
		data.GradeSemesterId,
		data.GradePoint,
		data.GradeCode,
		data.MbkmSubjectId,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}
