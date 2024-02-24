package study_program

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

type studyProgramRepository struct {
	*db.DB
}

func mapQueryFilterGetList(search string, majorId string, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	if search != "" {
		filterArray = append(filterArray, "sp.name ILIKE $%d")
		filterParams = append(filterParams, fmt.Sprintf("%%%s%%", search))
	}
	if majorId != "" {
		filterArray = append(filterArray, "sp.major_id = $%d")
		filterParams = append(filterParams, majorId)
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

func (f studyProgramRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, majorId, appType, userId string) ([]models.GetStudyProgramList, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetStudyProgramList{}
	var paginationResult common.Pagination

	queryGet := fmt.Sprintf(getListQuery, "")
	queryCount := fmt.Sprintf(countListQuery, "")
	if appType == appConstants.AppTypeAdmin {
		adminAdditionalQuery := fmt.Sprintf(`
			JOIN admins a ON a.id = '%s'
			JOIN role_study_program rsp ON rsp.role_id = a.role_id AND rsp.study_program_id = sp.id
		`, userId)
		queryGet = fmt.Sprintf(getListQuery, adminAdditionalQuery)
		queryCount = fmt.Sprintf(countListQuery, adminAdditionalQuery)
	}

	params := []interface{}{}
	filterQuery := mapQueryFilterGetList(pagination.Search, majorId, &params)
	queryGet = fmt.Sprintf("%s %s", queryGet, filterQuery)
	queryCount = fmt.Sprintf("%s %s", queryCount, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"sp.name": constants.Ascending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (f studyProgramRepository) GetDetail(ctx context.Context, tx *sqlx.Tx, id, appType, userId string) (models.GetStudyProgramDetail, *constants.ErrorResponse) {
	results := []models.GetStudyProgramDetail{}

	query := fmt.Sprintf(getDetailQuery, "")
	if appType == appConstants.AppTypeAdmin {
		adminAdditionalQuery := fmt.Sprintf(`
			JOIN admins a ON a.id = '%s'
			JOIN role_study_program rsp ON rsp.role_id = a.role_id AND rsp.study_program_id = sp.id
		`, userId)
		query = fmt.Sprintf(getDetailQuery, adminAdditionalQuery)
	}

	err := tx.SelectContext(
		ctx,
		&results,
		query,
		id,
	)
	if err != nil {
		return models.GetStudyProgramDetail{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetStudyProgramDetail{}, utils.ErrDataNotFound("study program")
	}

	return results[0], nil
}

func (f studyProgramRepository) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateStudyProgram) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		createQuery,
		data.DiktiStudyProgramId,
		data.Name,
		data.EnglishName,
		data.ShortName,
		data.EnglishShortName,
		data.AdministrativeUnit,
		data.MajorId,
		data.Address,
		data.PhoneNumber,
		data.Fax,
		data.Email,
		data.Website,
		data.ContactPerson,
		data.CuriculumReviewFrequency,
		data.CuriculumReviewMethod,
		data.EstablishmentDate,
		data.IsActive,
		data.StartSemester,
		data.OperationalPermitNumber,
		data.OperationalPermitDate,
		data.OperationalPermitDueDate,
		data.HeadLecturerId,
		data.OperatorName,
		data.OperatorPhoneNumber,
		data.CreatedBy,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (f studyProgramRepository) Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateStudyProgram) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		updateQuery,
		data.DiktiStudyProgramId,
		data.Name,
		data.EnglishName,
		data.ShortName,
		data.EnglishShortName,
		data.AdministrativeUnit,
		data.MajorId,
		data.Address,
		data.PhoneNumber,
		data.Fax,
		data.Email,
		data.Website,
		data.ContactPerson,
		data.CuriculumReviewFrequency,
		data.CuriculumReviewMethod,
		data.EstablishmentDate,
		data.IsActive,
		data.StartSemester,
		data.OperationalPermitNumber,
		data.OperationalPermitDate,
		data.OperationalPermitDueDate,
		data.HeadLecturerId,
		data.OperatorName,
		data.OperatorPhoneNumber,
		data.UpdatedBy,
		data.MinimumGraduationCredit,
		data.MinimumThesisCredit,
		data.Id,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (f studyProgramRepository) UpdateDegree(ctx context.Context, tx *sqlx.Tx, data models.UpdateDegreeStudyProgram) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		updateDegreeQuery,
		data.Degree,
		data.ShortDegree,
		data.EnglishDegree,
		data.UpdatedBy,
		data.Id,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (f studyProgramRepository) Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse {
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

func (s studyProgramRepository) GetByRoleIds(ctx context.Context, tx *sqlx.Tx, roleIds []string) ([]models.GetStudyProgramByRoleIds, *constants.ErrorResponse) {
	results := []models.GetStudyProgramByRoleIds{}

	err := tx.SelectContext(
		ctx,
		&results,
		getByRoleIdsQuery,
		pq.Array(roleIds),
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}
