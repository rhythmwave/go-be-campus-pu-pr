package major

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
)

type majorRepository struct {
	*db.DB
}

func mapQueryFilterGetList(search string, facultyId string, params *[]interface{}) string {
	filterArray := []string{}
	filterParams := *params

	if search != "" {
		filterArray = append(filterArray, "m.name ILIKE $%d")
		filterParams = append(filterParams, fmt.Sprintf("%%%s%%", search))
	}
	if facultyId != "" {
		filterArray = append(filterArray, "m.faculty_id = $%d")
		filterParams = append(filterParams, facultyId)
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

func (f majorRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, facultyId string) ([]models.GetMajorList, common.Pagination, *constants.ErrorResponse) {
	resultData := []models.GetMajorList{}
	var paginationResult common.Pagination

	params := []interface{}{}
	filterQuery := mapQueryFilterGetList(pagination.Search, facultyId, &params)
	queryGet := fmt.Sprintf("%s %s", getListQuery, filterQuery)
	queryCount := fmt.Sprintf("%s %s", countListQuery, filterQuery)
	if err := utils.QueryOperation(&queryGet, map[string]string{"m.name": constants.Ascending}, "", uint32(pagination.Limit), uint32(pagination.Page)); err != nil {
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

func (f majorRepository) GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetMajorDetail, *constants.ErrorResponse) {
	results := []models.GetMajorDetail{}

	err := tx.SelectContext(
		ctx,
		&results,
		getDetailQuery,
		id,
	)
	if err != nil {
		return models.GetMajorDetail{}, constants.ErrorInternalServer(err.Error())
	}
	if len(results) == 0 {
		return models.GetMajorDetail{}, utils.ErrDataNotFound("major")
	}

	return results[0], nil
}

func (f majorRepository) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateMajor) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		createQuery,
		data.FacultyId,
		data.Name,
		data.ShortName,
		data.EnglishName,
		data.EnglishShortName,
		data.Address,
		data.PhoneNumber,
		data.Fax,
		data.Email,
		data.ContactPerson,
		data.ExperimentBuildingArea,
		data.LectureHallArea,
		data.LectureHallCount,
		data.LaboratoriumArea,
		data.LaboratoriumCount,
		data.PermanentLecturerRoomArea,
		data.AdministrationRoomArea,
		data.BookCount,
		data.BookCopyCount,
		data.CreatedBy,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (f majorRepository) Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateMajor) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		updateQuery,
		data.FacultyId,
		data.Name,
		data.ShortName,
		data.EnglishName,
		data.EnglishShortName,
		data.Address,
		data.PhoneNumber,
		data.Fax,
		data.Email,
		data.ContactPerson,
		data.ExperimentBuildingArea,
		data.LectureHallArea,
		data.LectureHallCount,
		data.LaboratoriumArea,
		data.LaboratoriumCount,
		data.PermanentLecturerRoomArea,
		data.AdministrationRoomArea,
		data.BookCount,
		data.BookCopyCount,
		data.UpdatedBy,
		data.Id,
	)
	if err != nil {
		return constants.ErrorInternalServer(err.Error())
	}

	return nil
}

func (f majorRepository) Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse {
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
