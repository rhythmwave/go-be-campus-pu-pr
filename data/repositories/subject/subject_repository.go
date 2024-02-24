package subject

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

type subjectRepository struct {
	*db.DB
}

func mapQueryFilterGetList(search string, req objects.GetSubjectRequest, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	filterArray = append(filterArray, "s.is_mbkm = $%d")
	filterParams = append(filterParams, req.IsMbkm)
	if len(req.CurriculumIds) != 0 {
		filterArray = append(filterArray, "s.curriculum_id IN (SELECT UNNEST($%d::uuid[]))")
		filterParams = append(filterParams, pq.Array(req.CurriculumIds))
	}
	if search != "" {
		searchIlike := fmt.Sprintf("%%%s%%", search)
		filterArray = append(filterArray, "(s.code ILIKE $%d OR s.name ILIKE $%d)")
		filterParams = append(filterParams, searchIlike, searchIlike)
	}
	if req.PrerequisiteOfSubjectId != "" {
		filterArray = append(filterArray, "s.id != $%d")
		filterParams = append(filterParams, req.PrerequisiteOfSubjectId)
	}
	if req.EquivalentToCurriculumId != "" {
		filterArray = append(filterArray, "s.curriculum_id != $%d")
		filterParams = append(filterParams, req.EquivalentToCurriculumId)
	}
	if req.SemesterPackage != 0 {
		filterArray = append(filterArray, "s.semester_package = $%d")
		filterParams = append(filterParams, req.SemesterPackage)
	}
	if req.IsThesis != nil {
		filterArray = append(filterArray, "(s.is_thesis IS NOT NULL) = $%d")
		filterParams = append(filterParams, *req.IsThesis)
	}
	if req.StudyProgramId != "" {
		filterArray = append(filterArray, "sp.id = $%d")
		filterParams = append(filterParams, req.StudyProgramId)
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

func replaceQueryGetList(prerequisiteOfSubjectId, equivalentToCurriculumId string) (string, string, *constants.ErrorResponse) {
	var queryGet string
	var queryCount string

	var prerequisiteSelectQuery string
	var prerequisiteJoinQuery string
	var equivalentSelectQuery string
	var equivalentJoinQuery string
	if prerequisiteOfSubjectId != "" {
		errs := utils.IsValidUUID(prerequisiteOfSubjectId)
		if errs != nil {
			return queryGet, queryCount, errs
		}

		prerequisiteSelectQuery = `,
			spr.id AS subject_prerequisite_id,
			spr.prerequisite_type AS prerequisite_type,
			spr.minimum_grade_point AS prerequisite_minimum_grade_point
		`
		prerequisiteJoinQuery = fmt.Sprintf(`
			LEFT JOIN subject_prerequisites spr ON spr.subject_id = '%s' AND spr.prerequisite_subject_id = s.id
		`, prerequisiteOfSubjectId)
	}
	if equivalentToCurriculumId != "" {
		errs := utils.IsValidUUID(equivalentToCurriculumId)
		if errs != nil {
			return queryGet, queryCount, errs
		}

		equivalentSelectQuery = `,
			sesp.id AS equivalent_study_program_id,
			sesp.name AS equivalent_study_program_name,
			sec.id AS equivalent_curriculum_id,
			sec.name AS equivalent_curriculum_name,
			ses.id AS equivalent_subject_id,
			ses.code AS equivalent_subject_code,
			ses.name AS equivalent_subject_name
		`
		equivalentJoinQuery = fmt.Sprintf(`
			LEFT JOIN subject_equivalences se ON se.subject_id = s.id AND se.equivalent_curriculum_id = '%s'
			LEFT JOIN subjects ses ON ses.id = se.equivalent_subject_id
			LEFT JOIN curriculums sec ON sec.id = ses.curriculum_id
			LEFT JOIN study_programs sesp ON sesp.id = sec.study_program_id
		`, equivalentToCurriculumId)
	}

	queryGet = strings.Replace(getListQuery, "{{PREREQUISITE_SELECT}}", prerequisiteSelectQuery, 1)
	queryGet = strings.Replace(queryGet, "{{EQUIVALENT_SELECT}}", equivalentSelectQuery, 1)
	queryGet = strings.Replace(queryGet, "{{PREREQUISITE_JOIN}}", prerequisiteJoinQuery, 1)
	queryGet = strings.Replace(queryGet, "{{EQUIVALENT_JOIN}}", equivalentJoinQuery, 1)

	queryCount = strings.Replace(countListQuery, "{{PREREQUISITE_JOIN}}", prerequisiteJoinQuery, 1)
	queryCount = strings.Replace(queryCount, "{{EQUIVALENT_JOIN}}", equivalentJoinQuery, 1)

	return queryGet, queryCount, nil
}

func (a subjectRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, req objects.GetSubjectRequest) ([]models.GetSubject, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetSubject{}
	var paginationResult common.Pagination

	params := []interface{}{}

	queryGet, queryCount, errs := replaceQueryGetList(req.PrerequisiteOfSubjectId, req.EquivalentToCurriculumId)
	if errs != nil {
		return resultData, paginationResult, errs
	}

	filterQuery := mapQueryFilterGetList(pagination.Search, req, &params)
	queryGet = fmt.Sprintf("%s %s", queryGet, filterQuery)
	queryCount = fmt.Sprintf("%s %s", queryCount, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"s.code": constants.Ascending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (a subjectRepository) GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetSubjectDetail, *constants.ErrorResponse) {
	results := []models.GetSubjectDetail{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailQuery,
		id,
	)
	if err != nil {
		return models.GetSubjectDetail{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetSubjectDetail{}, utils.ErrDataNotFound("subject")
	}

	return results[0], nil
}

func (a subjectRepository) GetDetailByIds(ctx context.Context, tx *sqlx.Tx, ids []string) ([]models.GetSubjectDetail, *constants.ErrorResponse) {
	results := []models.GetSubjectDetail{}

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

func (a subjectRepository) GetThesisByCurriculumId(ctx context.Context, tx *sqlx.Tx, curriculumId string) (models.GetSubjectDetail, *constants.ErrorResponse) {
	results := []models.GetSubjectDetail{}

	err := tx.SelectContext(
		ctx,
		&results,
		getThesisByCurriculumIdQuery,
		curriculumId,
	)
	if err != nil {
		return models.GetSubjectDetail{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetSubjectDetail{}, utils.ErrDataNotFound("subject")
	}

	return results[0], nil
}

func (a subjectRepository) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateSubject) *constants.ErrorResponse {
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

func (a subjectRepository) Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateSubject) *constants.ErrorResponse {
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

func (a subjectRepository) Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse {
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
