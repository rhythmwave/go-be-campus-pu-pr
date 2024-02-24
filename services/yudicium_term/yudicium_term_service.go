package yudicium_term

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

type yudiciumTermService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (f yudiciumTermService) GetList(ctx context.Context, paginationData common.PaginationRequest, curriculumId string) (objects.YudiciumTermListWithPagination, *constants.ErrorResponse) {
	var result objects.YudiciumTermListWithPagination

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := f.YudiciumTermRepo.GetList(ctx, tx, paginationData, curriculumId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetYudiciumTerm{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetYudiciumTerm{
			Id:               v.Id,
			CurriculumId:     v.CurriculumId,
			CurriculumName:   v.CurriculumName,
			StudyProgramId:   v.StudyProgramId,
			StudyProgramName: v.StudyProgramName,
			Term:             v.Term,
			Remarks:          v.Remarks,
		})
	}

	result = objects.YudiciumTermListWithPagination{
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

func (f yudiciumTermService) Create(ctx context.Context, data objects.CreateYudiciumTerm) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	createData := models.CreateYudiciumTerm{
		CurriculumId: data.CurriculumId,
		Term:         data.Term,
		Remarks:      data.Remarks,
		CreatedBy:    claims.ID,
	}
	errs = f.YudiciumTermRepo.Create(ctx, tx, createData)
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

func (f yudiciumTermService) Update(ctx context.Context, data objects.UpdateYudiciumTerm) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs = f.YudiciumTermRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateYudiciumTerm{
		Id:        data.Id,
		Term:      data.Term,
		Remarks:   data.Remarks,
		UpdatedBy: claims.ID,
	}
	errs = f.YudiciumTermRepo.Update(ctx, tx, updateData)
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

func (f yudiciumTermService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := f.YudiciumTermRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = f.YudiciumTermRepo.Delete(ctx, tx, id)
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
