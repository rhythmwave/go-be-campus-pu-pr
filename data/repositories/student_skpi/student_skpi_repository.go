package student_skpi

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

type studentSkpiRepository struct {
	*db.DB
}

func mapQueryFilterGetList(paramsData objects.GetStudentSkpiRequest, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	if paramsData.StudyProgramId != "" {
		filterArray = append(filterArray, "s.study_program_id = $%d")
		filterParams = append(filterParams, paramsData.StudyProgramId)
	}
	if paramsData.Name != "" {
		searchParams := fmt.Sprintf("%%%s%%", paramsData.Name)
		filterArray = append(filterArray, "s.name ILIKE $%d")
		filterParams = append(filterParams, searchParams)
	}
	if paramsData.NimNumber != 0 {
		filterArray = append(filterArray, "s.nim_number = $%d")
		filterParams = append(filterParams, paramsData.NimNumber)
	}
	if paramsData.NimNumberFrom != 0 && paramsData.NimNumberTo != 0 {
		filterArray = append(filterArray, "(s.nim_number BETWEEN $%d AND $%d)")
		filterParams = append(filterParams, paramsData.NimNumberFrom, paramsData.NimNumberTo)
	}
	if paramsData.IsApproved != nil {
		filterArray = append(filterArray, "ss.is_approved = $%d")
		filterParams = append(filterParams, *paramsData.IsApproved)
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

func (s studentSkpiRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, paramsData objects.GetStudentSkpiRequest) ([]models.GetStudentSkpi, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetStudentSkpi{}
	var paginationResult common.Pagination

	params := []interface{}{}
	filterQuery := mapQueryFilterGetList(paramsData, &params)
	queryGet := fmt.Sprintf("%s %s", getListQuery, filterQuery)
	queryCount := fmt.Sprintf("%s %s", countListQuery, filterQuery)
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

func (s studentSkpiRepository) GetById(ctx context.Context, tx *sqlx.Tx, id string) (models.GetStudentSkpiDetail, *constants.ErrorResponse) {
	results := []models.GetStudentSkpiDetail{}

	err := tx.SelectContext(
		ctx,
		&results,
		getByIdQuery,
		id,
	)
	if err != nil {
		return models.GetStudentSkpiDetail{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetStudentSkpiDetail{}, utils.ErrDataNotFound("student skpi")
	}

	return results[0], nil
}

func (s studentSkpiRepository) GetByStudentId(ctx context.Context, tx *sqlx.Tx, studentId string) (models.GetStudentSkpiDetail, *constants.ErrorResponse) {
	results := []models.GetStudentSkpiDetail{}

	err := tx.SelectContext(
		ctx,
		&results,
		getByStudentIdQuery,
		studentId,
	)
	if err != nil {
		return models.GetStudentSkpiDetail{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetStudentSkpiDetail{}, utils.ErrDataNotFound("student skpi")
	}

	return results[0], nil
}

func (s studentSkpiRepository) GetStudentSkpiAchievementByStudentSkpiId(ctx context.Context, tx *sqlx.Tx, studentSkpiId string) ([]models.GetStudentSkpiAchievement, *constants.ErrorResponse) {
	results := []models.GetStudentSkpiAchievement{}

	err := tx.SelectContext(
		ctx,
		&results,
		getStudentSkpiAchievementByStudentSkpiIdQuery,
		studentSkpiId,
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (s studentSkpiRepository) GetStudentSkpiOrganizationByStudentSkpiId(ctx context.Context, tx *sqlx.Tx, studentSkpiId string) ([]models.GetStudentSkpiOrganization, *constants.ErrorResponse) {
	results := []models.GetStudentSkpiOrganization{}

	err := tx.SelectContext(
		ctx,
		&results,
		getStudentSkpiOrganizationByStudentSkpiIdQuery,
		studentSkpiId,
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (s studentSkpiRepository) GetStudentSkpiCertificateByStudentSkpiId(ctx context.Context, tx *sqlx.Tx, studentSkpiId string) ([]models.GetStudentSkpiCertificate, *constants.ErrorResponse) {
	results := []models.GetStudentSkpiCertificate{}

	err := tx.SelectContext(
		ctx,
		&results,
		getStudentSkpiCertificateByStudentSkpiIdQuery,
		studentSkpiId,
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (s studentSkpiRepository) GetStudentSkpiCharacterBuildingByStudentSkpiId(ctx context.Context, tx *sqlx.Tx, studentSkpiId string) ([]models.GetStudentSkpiCharacterBuilding, *constants.ErrorResponse) {
	results := []models.GetStudentSkpiCharacterBuilding{}

	err := tx.SelectContext(
		ctx,
		&results,
		getStudentSkpiCharacterBuildingByStudentSkpiIdQuery,
		studentSkpiId,
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (s studentSkpiRepository) GetStudentSkpiInternshipByStudentSkpiId(ctx context.Context, tx *sqlx.Tx, studentSkpiId string) ([]models.GetStudentSkpiInternship, *constants.ErrorResponse) {
	results := []models.GetStudentSkpiInternship{}

	err := tx.SelectContext(
		ctx,
		&results,
		getStudentSkpiInternshipByStudentSkpiIdQuery,
		studentSkpiId,
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (s studentSkpiRepository) GetStudentSkpiLanguageByStudentSkpiId(ctx context.Context, tx *sqlx.Tx, studentSkpiId string) ([]models.GetStudentSkpiLanguage, *constants.ErrorResponse) {
	results := []models.GetStudentSkpiLanguage{}

	err := tx.SelectContext(
		ctx,
		&results,
		getStudentSkpiLanguageByStudentSkpiIdQuery,
		studentSkpiId,
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (s studentSkpiRepository) UpsertStudentSkpi(ctx context.Context, tx *sqlx.Tx, data models.UpsertStudentSkpi) (string, *constants.ErrorResponse) {
	var result string

	err := tx.QueryRowContext(
		ctx,
		upsertStudentSkpiQuery,
		data.StudentId,
		data.AchievementPath,
		data.AchievementPathType,
		data.OrganizationPath,
		data.OrganizationPathType,
		data.CertificatePath,
		data.CertificatePathType,
		data.LanguagePath,
		data.LanguagePathType,
	).Scan(&result)
	if err != nil {
		return result, constants.ErrorInternalServer(err.Error())
	}

	return result, nil
}

func (s studentSkpiRepository) DeleteStudentSkpiAchievementExcludingName(ctx context.Context, tx *sqlx.Tx, studentSkpiId string, excludedName []string) *constants.ErrorResponse {
	query := fmt.Sprintf(deleteStudentSkpiAchievementExcludingNameQuery, "")
	params := []interface{}{
		studentSkpiId,
	}
	if len(excludedName) != 0 {
		additionalQuery := "AND name NOT IN (SELECT UNNEST($2::text[]))"
		query = fmt.Sprintf(deleteStudentSkpiAchievementExcludingNameQuery, additionalQuery)
		params = append(params, pq.Array(excludedName))
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

func (s studentSkpiRepository) DeleteStudentSkpiOrganizationExcludingName(ctx context.Context, tx *sqlx.Tx, studentSkpiId string, excludedName []string) *constants.ErrorResponse {
	query := fmt.Sprintf(deleteStudentSkpiOrganizationExcludingNameQuery, "")
	params := []interface{}{
		studentSkpiId,
	}
	if len(excludedName) != 0 {
		additionalQuery := "AND name NOT IN (SELECT UNNEST($2::text[]))"
		query = fmt.Sprintf(deleteStudentSkpiOrganizationExcludingNameQuery, additionalQuery)
		params = append(params, pq.Array(excludedName))
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

func (s studentSkpiRepository) DeleteStudentSkpiCertificateExcludingName(ctx context.Context, tx *sqlx.Tx, studentSkpiId string, excludedName []string) *constants.ErrorResponse {
	query := fmt.Sprintf(deleteStudentSkpiCertificateExcludingNameQuery, "")
	params := []interface{}{
		studentSkpiId,
	}
	if len(excludedName) != 0 {
		additionalQuery := "AND name NOT IN (SELECT UNNEST($2::text[]))"
		query = fmt.Sprintf(deleteStudentSkpiCertificateExcludingNameQuery, additionalQuery)
		params = append(params, pq.Array(excludedName))
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

func (s studentSkpiRepository) DeleteStudentSkpiCharacterBuildingExcludingName(ctx context.Context, tx *sqlx.Tx, studentSkpiId string, excludedName []string) *constants.ErrorResponse {
	query := fmt.Sprintf(deleteStudentSkpiCharacterBuildingExcludingNameQuery, "")
	params := []interface{}{
		studentSkpiId,
	}
	if len(excludedName) != 0 {
		additionalQuery := "AND name NOT IN (SELECT UNNEST($2::text[]))"
		query = fmt.Sprintf(deleteStudentSkpiCharacterBuildingExcludingNameQuery, additionalQuery)
		params = append(params, pq.Array(excludedName))
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

func (s studentSkpiRepository) DeleteStudentSkpiInternshipExcludingName(ctx context.Context, tx *sqlx.Tx, studentSkpiId string, excludedName []string) *constants.ErrorResponse {
	query := fmt.Sprintf(deleteStudentSkpiInternshipExcludingNameQuery, "")
	params := []interface{}{
		studentSkpiId,
	}
	if len(excludedName) != 0 {
		additionalQuery := "AND name NOT IN (SELECT UNNEST($2::text[]))"
		query = fmt.Sprintf(deleteStudentSkpiInternshipExcludingNameQuery, additionalQuery)
		params = append(params, pq.Array(excludedName))
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

func (s studentSkpiRepository) DeleteStudentSkpiLanguageExcludingName(ctx context.Context, tx *sqlx.Tx, studentSkpiId string, excludedName []string) *constants.ErrorResponse {
	query := fmt.Sprintf(deleteStudentSkpiLanguageExcludingNameQuery, "")
	params := []interface{}{
		studentSkpiId,
	}
	if len(excludedName) != 0 {
		additionalQuery := "AND name NOT IN (SELECT UNNEST($2::text[]))"
		query = fmt.Sprintf(deleteStudentSkpiLanguageExcludingNameQuery, additionalQuery)
		params = append(params, pq.Array(excludedName))
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

func (s studentSkpiRepository) UpsertStudentSkpiAchievement(ctx context.Context, tx *sqlx.Tx, data []models.UpsertStudentSkpiAchievement) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		upsertStudentSkpiAchievementQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (s studentSkpiRepository) UpsertStudentSkpiOrganization(ctx context.Context, tx *sqlx.Tx, data []models.UpsertStudentSkpiOrganization) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		upsertStudentSkpiOrganizationQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (s studentSkpiRepository) UpsertStudentSkpiCertificate(ctx context.Context, tx *sqlx.Tx, data []models.UpsertStudentSkpiCertificate) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		upsertStudentSkpiCertificateQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (s studentSkpiRepository) UpsertStudentSkpiCharacterBuilding(ctx context.Context, tx *sqlx.Tx, data []models.UpsertStudentSkpiCharacterBuilding) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		upsertStudentSkpiCharacterBuildingQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (s studentSkpiRepository) UpsertStudentSkpiInternship(ctx context.Context, tx *sqlx.Tx, data []models.UpsertStudentSkpiInternship) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		upsertStudentSkpiInternshipQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (s studentSkpiRepository) UpsertStudentSkpiLanguage(ctx context.Context, tx *sqlx.Tx, data []models.UpsertStudentSkpiLanguage) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		upsertStudentSkpiLanguageQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (s studentSkpiRepository) Approve(ctx context.Context, tx *sqlx.Tx, data models.ApproveStudentSkpi) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		approveQuery,
		data.SkpiNumber,
		data.Id,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (s studentSkpiRepository) Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse {
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
