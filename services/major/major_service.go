package major

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

type majorService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (m majorService) GetList(ctx context.Context, paginationData common.PaginationRequest, facultyId string) (objects.MajorListWithPagination, *constants.ErrorResponse) {
	var result objects.MajorListWithPagination

	tx, err := m.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := m.MajorRepo.GetList(ctx, tx, paginationData, facultyId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetMajor{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetMajor{
			Id:          v.Id,
			FacultyName: v.FacultyName,
			Name:        v.Name,
		})
	}

	result = objects.MajorListWithPagination{
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

func (m majorService) GetDetail(ctx context.Context, id string) (objects.GetMajorDetail, *constants.ErrorResponse) {
	var result objects.GetMajorDetail

	tx, err := m.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	resultData, errs := m.MajorRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result = objects.GetMajorDetail{
		Id:                        resultData.Id,
		FacultyId:                 resultData.FacultyId,
		FacultyName:               resultData.FacultyName,
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

func (m majorService) Create(ctx context.Context, data objects.CreateMajor) *constants.ErrorResponse {
	tx, err := m.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	createData := models.CreateMajor{
		FacultyId:                 data.FacultyId,
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
	errs = m.MajorRepo.Create(ctx, tx, createData)
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

func (m majorService) Update(ctx context.Context, data objects.UpdateMajor) *constants.ErrorResponse {
	tx, err := m.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := m.MajorRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateMajor{
		Id:                        data.Id,
		FacultyId:                 data.FacultyId,
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
	errs = m.MajorRepo.Update(ctx, tx, updateData)
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

func (m majorService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := m.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := m.MajorRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = m.MajorRepo.Delete(ctx, tx, id)
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
