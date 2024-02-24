package thesis_supervisor_role

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

type thesisSupervisorRoleService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (f thesisSupervisorRoleService) GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.ThesisSupervisorRoleListWithPagination, *constants.ErrorResponse) {
	var result objects.ThesisSupervisorRoleListWithPagination

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := f.ThesisSupervisorRoleRepo.GetList(ctx, tx, paginationData)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetThesisSupervisorRole{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetThesisSupervisorRole{
			Id:   v.Id,
			Name: v.Name,
			Sort: v.Sort,
		})
	}

	result = objects.ThesisSupervisorRoleListWithPagination{
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

func (f thesisSupervisorRoleService) Create(ctx context.Context, data objects.CreateThesisSupervisorRole) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	createData := models.CreateThesisSupervisorRole{
		Name:      data.Name,
		Sort:      data.Sort,
		CreatedBy: claims.ID,
	}
	errs = f.ThesisSupervisorRoleRepo.Create(ctx, tx, createData)
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

func (f thesisSupervisorRoleService) Update(ctx context.Context, data objects.UpdateThesisSupervisorRole) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs = f.ThesisSupervisorRoleRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateThesisSupervisorRole{
		Id:        data.Id,
		Name:      data.Name,
		Sort:      data.Sort,
		UpdatedBy: claims.ID,
	}
	errs = f.ThesisSupervisorRoleRepo.Update(ctx, tx, updateData)
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

func (f thesisSupervisorRoleService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := f.ThesisSupervisorRoleRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = f.ThesisSupervisorRoleRepo.Delete(ctx, tx, id)
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
