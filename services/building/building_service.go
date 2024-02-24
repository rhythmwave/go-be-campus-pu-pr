package building

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	appConstants "github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
)

type buildingService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func mapGetList(buildingData []models.GetBuilding) []objects.GetBuilding {
	results := []objects.GetBuilding{}

	for _, v := range buildingData {
		var level string
		var facultyId string
		var facultyName string
		if v.FacultyId != nil {
			facultyId = utils.NullStringScan(v.FacultyId)
			facultyName = utils.NullStringScan(v.FacultyName)
			level = appConstants.FacultyBuilding
		}
		if v.MajorId != nil {
			facultyId = utils.NullStringScan(v.MajorFacultyId)
			facultyName = utils.NullStringScan(v.MajorFacultyName)
			level = appConstants.MajorBuilding
		}

		results = append(results, objects.GetBuilding{
			Id:          v.Id,
			Level:       level,
			FacultyId:   facultyId,
			FacultyName: facultyName,
			MajorId:     utils.NullStringScan(v.MajorId),
			MajorName:   utils.NullStringScan(v.MajorName),
			Code:        v.Code,
			Name:        v.Name,
		})
	}

	return results
}

func (a buildingService) GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.BuildingListWithPagination, *constants.ErrorResponse) {
	var result objects.BuildingListWithPagination

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := a.BuildingRepo.GetList(ctx, tx, paginationData)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result = objects.BuildingListWithPagination{
		Pagination: paginationResult,
		Data:       mapGetList(modelResult),
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (a buildingService) GetDetail(ctx context.Context, id string) (objects.GetBuilding, *constants.ErrorResponse) {
	var result objects.GetBuilding

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	resultData, errs := a.BuildingRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	var level string
	var facultyId string
	var facultyName string
	if resultData.FacultyId != nil {
		facultyId = utils.NullStringScan(resultData.FacultyId)
		facultyName = utils.NullStringScan(resultData.FacultyName)
		level = appConstants.FacultyBuilding
	}
	if resultData.MajorId != nil {
		facultyId = utils.NullStringScan(resultData.MajorFacultyId)
		facultyName = utils.NullStringScan(resultData.MajorFacultyName)
		level = appConstants.MajorBuilding
	}

	result = objects.GetBuilding{
		Id:          resultData.Id,
		Level:       level,
		FacultyId:   facultyId,
		FacultyName: facultyName,
		MajorId:     utils.NullStringScan(resultData.MajorId),
		MajorName:   utils.NullStringScan(resultData.MajorName),
		Code:        resultData.Code,
		Name:        resultData.Name,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (a buildingService) Create(ctx context.Context, data objects.CreateBuilding) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	createData := models.CreateBuilding{
		FacultyId: utils.NewNullString(data.FacultyId),
		MajorId:   utils.NewNullString(data.MajorId),
		Code:      data.Code,
		Name:      data.Name,
		CreatedBy: claims.ID,
	}
	errs = a.BuildingRepo.Create(ctx, tx, createData)
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

func (a buildingService) Update(ctx context.Context, data objects.UpdateBuilding) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := a.BuildingRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateBuilding{
		Id:        data.Id,
		FacultyId: utils.NewNullString(data.FacultyId),
		MajorId:   utils.NewNullString(data.MajorId),
		Code:      data.Code,
		Name:      data.Name,
		UpdatedBy: claims.ID,
	}
	errs = a.BuildingRepo.Update(ctx, tx, updateData)
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

func (a buildingService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := a.BuildingRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = a.BuildingRepo.Delete(ctx, tx, id)
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
