package announcement

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

type announcementRepository struct {
	*db.DB
}

func mapQueryFilterGetList(search, appType, announcementType string, ids []string, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	var startIndex int

	filterArray = append(filterArray, "a.id IN (SELECT UNNEST($%d::uuid[]))")
	filterParams = append(filterParams, pq.Array(ids))

	if search != "" {
		filterArray = append(filterArray, "a.title ILIKE $%d")
		filterParams = append(filterParams, fmt.Sprintf("%%%s%%", search))
	}
	if announcementType != "" {
		filterArray = append(filterArray, "a.type::text = $%d")
		filterParams = append(filterParams, announcementType)
	}
	switch appType {
	case appConstants.AppTypeLecturer:
		filterArray = append(filterArray, "a.for_lecturer IS true")
	case appConstants.AppTypeStudent:
		filterArray = append(filterArray, "a.for_student IS true")
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

func (a announcementRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, appType, announcementType string, ids []string) ([]models.GetAnnouncement, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetAnnouncement{}
	var paginationResult common.Pagination

	if len(ids) == 0 {
		pagination.GetPagination(0, pagination.Page, pagination.Limit)
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

	params := []interface{}{}

	filterQuery := mapQueryFilterGetList(pagination.Search, appType, announcementType, ids, &params)
	queryGet := fmt.Sprintf("%s %s", getListQuery, filterQuery)
	queryCount := fmt.Sprintf("%s %s", countListQuery, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"a.created_at": constants.Descending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (a announcementRepository) GetAnnouncementStudyProgramByAnnouncementIds(ctx context.Context, tx *sqlx.Tx, announcementIds []string) ([]models.GetAnnouncementStudyProgram, *constants.ErrorResponse) {
	results := []models.GetAnnouncementStudyProgram{}

	err := tx.SelectContext(
		ctx,
		&results,
		getAnnouncementStudyProgramByAnnouncementIdsQuery,
		pq.Array(announcementIds),
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (a announcementRepository) GetAnnouncementStudyProgramByStudyProgramIds(ctx context.Context, tx *sqlx.Tx, studyProgramIds []string) ([]models.GetAnnouncementStudyProgram, *constants.ErrorResponse) {
	results := []models.GetAnnouncementStudyProgram{}

	err := tx.SelectContext(
		ctx,
		&results,
		getAnnouncementStudyProgramByStudyProgramIdsQuery,
		pq.Array(studyProgramIds),
	)
	if err != nil {
		return results, constants.ErrorInternalServer(err.Error())
	}

	return results, nil
}

func (a announcementRepository) GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetAnnouncement, *constants.ErrorResponse) {
	results := []models.GetAnnouncement{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailQuery,
		id,
	)
	if err != nil {
		return models.GetAnnouncement{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetAnnouncement{}, utils.ErrDataNotFound("announcement")
	}

	return results[0], nil
}

func (a announcementRepository) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateAnnouncement) *constants.ErrorResponse {
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

func (a announcementRepository) Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateAnnouncement) *constants.ErrorResponse {
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

func (a announcementRepository) DeleteAnnouncementStudyProgramExcludingStudyProgramIds(ctx context.Context, tx *sqlx.Tx, announcementId string, excludedStudyProgramIds []string) *constants.ErrorResponse {
	params := []interface{}{
		announcementId,
	}
	query := fmt.Sprintf(deleteAnnouncementStudyProgramExcludingStudyProgramIdsQuery, "")
	if len(excludedStudyProgramIds) != 0 {
		additionalQuery := "AND study_program_id NOT IN (SELECT UNNEST($2::uuid[]))"

		params = append(params, pq.Array(excludedStudyProgramIds))
		query = fmt.Sprintf(deleteAnnouncementStudyProgramExcludingStudyProgramIdsQuery, additionalQuery)
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

func (a announcementRepository) UpsertAnnouncementStudyProgram(ctx context.Context, tx *sqlx.Tx, data []models.UpsertAnnouncementStudyProgram) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		upsertAnnouncementStudyProgramQuery,
		data,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (a announcementRepository) Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse {
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
