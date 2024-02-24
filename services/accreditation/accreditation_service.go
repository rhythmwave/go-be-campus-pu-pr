package accreditation

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

type accreditationService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func mapGetList(accreditationData []models.GetAccreditation) []objects.GetAccreditation {
	results := []objects.GetAccreditation{}

	for _, v := range accreditationData {
		results = append(results, objects.GetAccreditation{
			Id:             v.Id,
			StudyProgramId: v.StudyProgramId,
			DecreeNumber:   v.DecreeNumber,
			DecreeDate:     v.DecreeDate,
			DecreeDueDate:  v.DecreeDueDate,
			Accreditation:  v.Accreditation,
			IsActive:       v.IsActive,
		})
	}

	return results
}

func (a accreditationService) GetList(ctx context.Context, paginationData common.PaginationRequest, studyProgramId string) (objects.AccreditationListWithPagination, *constants.ErrorResponse) {
	var result objects.AccreditationListWithPagination

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	_, errs = a.StudyProgramRepo.GetDetail(ctx, tx, studyProgramId, claims.Role, claims.ID)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	modelResult, paginationResult, errs := a.AccreditationRepo.GetList(ctx, tx, paginationData, studyProgramId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result = objects.AccreditationListWithPagination{
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

func (a accreditationService) Create(ctx context.Context, data objects.CreateAccreditation) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs = a.StudyProgramRepo.GetDetail(ctx, tx, data.StudyProgramId, claims.Role, claims.ID)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	createData := models.CreateAccreditation{
		StudyProgramId: data.StudyProgramId,
		DecreeNumber:   data.DecreeNumber,
		DecreeDate:     data.DecreeDate,
		DecreeDueDate:  data.DecreeDueDate,
		Accreditation:  data.Accreditation,
		IsActive:       data.IsActive,
		CreatedBy:      claims.ID,
	}
	errs = a.AccreditationRepo.Create(ctx, tx, createData)
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

func (a accreditationService) Update(ctx context.Context, data objects.UpdateAccreditation) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := a.AccreditationRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateAccreditation{
		Id:            data.Id,
		DecreeNumber:  data.DecreeNumber,
		DecreeDate:    data.DecreeDate,
		DecreeDueDate: data.DecreeDueDate,
		Accreditation: data.Accreditation,
		IsActive:      data.IsActive,
		UpdatedBy:     claims.ID,
	}
	errs = a.AccreditationRepo.Update(ctx, tx, updateData)
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

func (a accreditationService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := a.AccreditationRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = a.AccreditationRepo.Delete(ctx, tx, id)
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
