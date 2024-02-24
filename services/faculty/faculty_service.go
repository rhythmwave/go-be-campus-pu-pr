package faculty

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
)

type facultyService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (f facultyService) GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.FacultyListWithPagination, *constants.ErrorResponse) {
	var result objects.FacultyListWithPagination

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := f.FacultyRepo.GetList(ctx, tx, paginationData)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetFaculty{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetFaculty{
			Id:        v.Id,
			Name:      v.Name,
			ShortName: v.ShortName,
		})
	}

	result = objects.FacultyListWithPagination{
		Pagination: paginationResult,
		Data:       resultData,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (f facultyService) GetDetail(ctx context.Context, id string) (objects.GetFacultyDetail, *constants.ErrorResponse) {
	var result objects.GetFacultyDetail

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	resultData, errs := f.FacultyRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result = objects.GetFacultyDetail{
		Id:                        resultData.Id,
		Name:                      resultData.Name,
		ShortName:                 resultData.ShortName,
		EnglishName:               resultData.EnglishName,
		EnglishShortName:          resultData.EnglishShortName,
		Address:                   resultData.Address,
		PhoneNumber:               resultData.PhoneNumber,
		Fax:                       resultData.Fax,
		Email:                     resultData.Email,
		ContactPerson:             resultData.ContactPerson,
		ExperimentBuildingArea:    resultData.ExperimentBuildingArea,
		LectureHallArea:           resultData.LectureHallArea,
		LectureHallCount:          resultData.LectureHallCount,
		LaboratoriumArea:          resultData.LaboratoriumArea,
		LaboratoriumCount:         resultData.LaboratoriumCount,
		PermanentLecturerRoomArea: resultData.PermanentLecturerRoomArea,
		AdministrationRoomArea:    resultData.AdministrationRoomArea,
		BookCount:                 resultData.BookCount,
		BookCopyCount:             resultData.BookCopyCount,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (f facultyService) Create(ctx context.Context, data objects.CreateFaculty) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	createData := models.CreateFaculty{
		Name:                      data.Name,
		ShortName:                 utils.NewNullString(data.ShortName),
		EnglishName:               utils.NewNullString(data.EnglishName),
		EnglishShortName:          utils.NewNullString(data.EnglishShortName),
		Address:                   data.Address,
		PhoneNumber:               utils.NewNullString(data.PhoneNumber),
		Fax:                       utils.NewNullString(data.Fax),
		Email:                     utils.NewNullString(data.Email),
		ContactPerson:             utils.NewNullString(data.ContactPerson),
		ExperimentBuildingArea:    utils.NewNullFloat64(&data.ExperimentBuildingArea),
		LectureHallArea:           utils.NewNullFloat64(&data.LectureHallArea),
		LectureHallCount:          utils.NewNullInt32(int32(data.LectureHallCount)),
		LaboratoriumArea:          utils.NewNullFloat64(&data.LaboratoriumArea),
		LaboratoriumCount:         utils.NewNullInt32(int32(data.LaboratoriumCount)),
		PermanentLecturerRoomArea: utils.NewNullFloat64(&data.PermanentLecturerRoomArea),
		AdministrationRoomArea:    utils.NewNullFloat64(&data.AdministrationRoomArea),
		BookCount:                 utils.NewNullInt32(int32(data.BookCount)),
		BookCopyCount:             utils.NewNullInt32(int32(data.BookCopyCount)),
		CreatedBy:                 claims.ID,
	}
	errs = f.FacultyRepo.Create(ctx, tx, createData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (f facultyService) Update(ctx context.Context, data objects.UpdateFaculty) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := f.FacultyRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateFaculty{
		Id:                        data.Id,
		Name:                      data.Name,
		ShortName:                 utils.NewNullString(data.ShortName),
		EnglishName:               utils.NewNullString(data.EnglishName),
		EnglishShortName:          utils.NewNullString(data.EnglishShortName),
		Address:                   data.Address,
		PhoneNumber:               utils.NewNullString(data.PhoneNumber),
		Fax:                       utils.NewNullString(data.Fax),
		Email:                     utils.NewNullString(data.Email),
		ContactPerson:             utils.NewNullString(data.ContactPerson),
		ExperimentBuildingArea:    utils.NewNullFloat64(&data.ExperimentBuildingArea),
		LectureHallArea:           utils.NewNullFloat64(&data.LectureHallArea),
		LectureHallCount:          utils.NewNullInt32(int32(data.LectureHallCount)),
		LaboratoriumArea:          utils.NewNullFloat64(&data.LaboratoriumArea),
		LaboratoriumCount:         utils.NewNullInt32(int32(data.LaboratoriumCount)),
		PermanentLecturerRoomArea: utils.NewNullFloat64(&data.PermanentLecturerRoomArea),
		AdministrationRoomArea:    utils.NewNullFloat64(&data.AdministrationRoomArea),
		BookCount:                 utils.NewNullInt32(int32(data.BookCount)),
		BookCopyCount:             utils.NewNullInt32(int32(data.BookCopyCount)),
		UpdatedBy:                 claims.ID,
	}
	errs = f.FacultyRepo.Update(ctx, tx, updateData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (f facultyService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := f.FacultyRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = f.FacultyRepo.Delete(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}
	return nil
}
