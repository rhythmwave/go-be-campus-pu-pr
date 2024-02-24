package graduation_session

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

type graduationSessionService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (f graduationSessionService) GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.GraduationSessionListWithPagination, *constants.ErrorResponse) {
	var result objects.GraduationSessionListWithPagination

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := f.GraduationSessionRepo.GetList(ctx, tx, paginationData)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetGraduationSession{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetGraduationSession{
			Id:                v.Id,
			SessionYear:       v.SessionYear,
			SessionSchoolYear: appUtils.GenerateSchoolYear(v.SessionYear),
			SessionNumber:     v.SessionNumber,
			SessionDate:       v.SessionDate,
		})
	}

	result = objects.GraduationSessionListWithPagination{
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

func (f graduationSessionService) Create(ctx context.Context, data objects.CreateGraduationSession) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	createData := models.CreateGraduationSession{
		SessionYear:   data.SessionYear,
		SessionNumber: data.SessionNumber,
		SessionDate:   data.SessionDate,
		CreatedBy:     claims.ID,
	}
	errs = f.GraduationSessionRepo.Create(ctx, tx, createData)
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

func (f graduationSessionService) Update(ctx context.Context, data objects.UpdateGraduationSession) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs = f.GraduationSessionRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateGraduationSession{
		Id:            data.Id,
		SessionYear:   data.SessionYear,
		SessionNumber: data.SessionNumber,
		SessionDate:   data.SessionDate,
		UpdatedBy:     claims.ID,
	}
	errs = f.GraduationSessionRepo.Update(ctx, tx, updateData)
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

func (f graduationSessionService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := f.GraduationSessionRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = f.GraduationSessionRepo.Delete(ctx, tx, id)
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
