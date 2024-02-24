package exam_supervisor_role

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

type examSupervisorRoleService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (f examSupervisorRoleService) GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.ExamSupervisorRoleListWithPagination, *constants.ErrorResponse) {
	var result objects.ExamSupervisorRoleListWithPagination

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := f.ExamSupervisorRoleRepo.GetList(ctx, tx, paginationData)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetExamSupervisorRole{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetExamSupervisorRole{
			Id:   v.Id,
			Name: v.Name,
			Sort: v.Sort,
		})
	}

	result = objects.ExamSupervisorRoleListWithPagination{
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

func (f examSupervisorRoleService) Create(ctx context.Context, data objects.CreateExamSupervisorRole) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	createData := models.CreateExamSupervisorRole{
		Name:      data.Name,
		Sort:      data.Sort,
		CreatedBy: claims.ID,
	}
	errs = f.ExamSupervisorRoleRepo.Create(ctx, tx, createData)
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

func (f examSupervisorRoleService) Update(ctx context.Context, data objects.UpdateExamSupervisorRole) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs = f.ExamSupervisorRoleRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateExamSupervisorRole{
		Id:        data.Id,
		Name:      data.Name,
		Sort:      data.Sort,
		UpdatedBy: claims.ID,
	}
	errs = f.ExamSupervisorRoleRepo.Update(ctx, tx, updateData)
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

func (f examSupervisorRoleService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := f.ExamSupervisorRoleRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = f.ExamSupervisorRoleRepo.Delete(ctx, tx, id)
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
