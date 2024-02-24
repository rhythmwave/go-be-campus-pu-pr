package officer_action

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

type officerActionService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (f officerActionService) GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.OfficerActionListWithPagination, *constants.ErrorResponse) {
	var result objects.OfficerActionListWithPagination

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := f.OfficerActionRepo.GetList(ctx, tx, paginationData)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetOfficerAction{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetOfficerAction{
			Id:                          v.Id,
			DocumentTypeId:              v.DocumentTypeId,
			DocumentTypeName:            v.DocumentTypeName,
			DocumentActionId:            v.DocumentActionId,
			DocumentActionAction:        v.DocumentActionAction,
			DocumentActionEnglishAction: v.DocumentActionEnglishAction,
			OfficerId:                   v.OfficerId,
			OfficerName:                 v.OfficerName,
			OfficerTitle:                v.OfficerTitle,
			OfficerEnglishTitle:         v.OfficerEnglishTitle,
			OfficerStudyProgramId:       v.OfficerStudyProgramId,
			OfficerStudyProgramName:     v.OfficerStudyProgramName,
		})
	}

	result = objects.OfficerActionListWithPagination{
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

func (f officerActionService) Create(ctx context.Context, data objects.CreateOfficerAction) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	createData := models.CreateOfficerAction{
		DocumentTypeId:   data.DocumentTypeId,
		DocumentActionId: data.DocumentActionId,
		OfficerId:        data.OfficerId,
		CreatedBy:        claims.ID,
	}
	errs = f.OfficerActionRepo.Create(ctx, tx, createData)
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

func (f officerActionService) Update(ctx context.Context, data objects.UpdateOfficerAction) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs = f.OfficerActionRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateOfficerAction{
		Id:               data.Id,
		DocumentTypeId:   data.DocumentTypeId,
		DocumentActionId: data.DocumentActionId,
		OfficerId:        data.OfficerId,
		UpdatedBy:        claims.ID,
	}
	errs = f.OfficerActionRepo.Update(ctx, tx, updateData)
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

func (f officerActionService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := f.OfficerActionRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = f.OfficerActionRepo.Delete(ctx, tx, id)
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
