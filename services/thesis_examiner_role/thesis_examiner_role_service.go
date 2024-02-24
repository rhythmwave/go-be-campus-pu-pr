package thesis_examiner_role

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

type thesisExaminerRoleService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (f thesisExaminerRoleService) GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.ThesisExaminerRoleListWithPagination, *constants.ErrorResponse) {
	var result objects.ThesisExaminerRoleListWithPagination

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := f.ThesisExaminerRoleRepo.GetList(ctx, tx, paginationData)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetThesisExaminerRole{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetThesisExaminerRole{
			Id:      v.Id,
			Name:    v.Name,
			Remarks: v.Remarks,
			Sort:    v.Sort,
		})
	}

	result = objects.ThesisExaminerRoleListWithPagination{
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

func (f thesisExaminerRoleService) Create(ctx context.Context, data objects.CreateThesisExaminerRole) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	createData := models.CreateThesisExaminerRole{
		Name:      data.Name,
		Remarks:   data.Remarks,
		Sort:      data.Sort,
		CreatedBy: claims.ID,
	}
	errs = f.ThesisExaminerRoleRepo.Create(ctx, tx, createData)
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

func (f thesisExaminerRoleService) Update(ctx context.Context, data objects.UpdateThesisExaminerRole) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs = f.ThesisExaminerRoleRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateThesisExaminerRole{
		Id:        data.Id,
		Name:      data.Name,
		Remarks:   data.Remarks,
		Sort:      data.Sort,
		UpdatedBy: claims.ID,
	}
	errs = f.ThesisExaminerRoleRepo.Update(ctx, tx, updateData)
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

func (f thesisExaminerRoleService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := f.ThesisExaminerRoleRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = f.ThesisExaminerRoleRepo.Delete(ctx, tx, id)
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
