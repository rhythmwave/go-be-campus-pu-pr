package yudicium_session

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
	appUtils "github.com/sccicitb/pupr-backend/utils"
)

type yudiciumSessionService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (f yudiciumSessionService) GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.YudiciumSessionListWithPagination, *constants.ErrorResponse) {
	var result objects.YudiciumSessionListWithPagination

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := f.YudiciumSessionRepo.GetList(ctx, tx, paginationData)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetYudiciumSession{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetYudiciumSession{
			Id:                 v.Id,
			SemesterId:         v.SemesterId,
			SemesterSchoolYear: appUtils.GenerateSchoolYear(v.SemesterStartYear),
			SemesterType:       v.SemesterType,
			Name:               v.Name,
			SessionDate:        v.SessionDate,
		})
	}

	result = objects.YudiciumSessionListWithPagination{
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

func (f yudiciumSessionService) Create(ctx context.Context, data objects.CreateYudiciumSession) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	createData := models.CreateYudiciumSession{
		SemesterId:  data.SemesterId,
		Name:        data.Name,
		SessionDate: data.SessionDate,
		CreatedBy:   claims.ID,
	}
	errs = f.YudiciumSessionRepo.Create(ctx, tx, createData)
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

func (f yudiciumSessionService) Update(ctx context.Context, data objects.UpdateYudiciumSession) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs = f.YudiciumSessionRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateYudiciumSession{
		Id:          data.Id,
		SemesterId:  data.SemesterId,
		Name:        data.Name,
		SessionDate: data.SessionDate,
		UpdatedBy:   claims.ID,
	}
	errs = f.YudiciumSessionRepo.Update(ctx, tx, updateData)
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

func (f yudiciumSessionService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := f.YudiciumSessionRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = f.YudiciumSessionRepo.Delete(ctx, tx, id)
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
