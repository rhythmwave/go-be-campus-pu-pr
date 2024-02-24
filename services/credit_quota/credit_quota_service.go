package credit_quota

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

type creditQuotaService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (f creditQuotaService) GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.CreditQuotaListWithPagination, *constants.ErrorResponse) {
	var result objects.CreditQuotaListWithPagination

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := f.CreditQuotaRepo.GetList(ctx, tx, paginationData)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetCreditQuota{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetCreditQuota{
			Id:                v.Id,
			MinimumGradePoint: v.MinimumGradePoint,
			MaximumGradePoint: v.MaximumGradePoint,
			MaximumCredit:     v.MaximumCredit,
		})
	}

	result = objects.CreditQuotaListWithPagination{
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

func (f creditQuotaService) Create(ctx context.Context, data objects.CreateCreditQuota) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	createData := models.CreateCreditQuota{
		MinimumGradePoint: data.MinimumGradePoint,
		MaximumGradePoint: data.MaximumGradePoint,
		MaximumCredit:     data.MaximumCredit,
		CreatedBy:         claims.ID,
	}
	errs = f.CreditQuotaRepo.Create(ctx, tx, createData)
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

func (f creditQuotaService) Update(ctx context.Context, data objects.UpdateCreditQuota) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs = f.CreditQuotaRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateCreditQuota{
		Id:                data.Id,
		MinimumGradePoint: data.MinimumGradePoint,
		MaximumGradePoint: data.MaximumGradePoint,
		MaximumCredit:     data.MaximumCredit,
		UpdatedBy:         claims.ID,
	}
	errs = f.CreditQuotaRepo.Update(ctx, tx, updateData)
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

func (f creditQuotaService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := f.CreditQuotaRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = f.CreditQuotaRepo.Delete(ctx, tx, id)
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
