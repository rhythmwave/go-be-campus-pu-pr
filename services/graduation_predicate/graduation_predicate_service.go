package graduation_predicate

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

type graduationPredicateService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (f graduationPredicateService) GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.GraduationPredicateListWithPagination, *constants.ErrorResponse) {
	var result objects.GraduationPredicateListWithPagination

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := f.GraduationPredicateRepo.GetList(ctx, tx, paginationData)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetGraduationPredicate{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetGraduationPredicate{
			Id:                          v.Id,
			Predicate:                   v.Predicate,
			MinimumGpa:                  v.MinimumGpa,
			MaximumStudySemester:        v.MaximumStudySemester,
			RepeatCourseLimit:           v.RepeatCourseLimit,
			BelowMinimumGradePointLimit: v.BelowMinimumGradePointLimit,
		})
	}

	result = objects.GraduationPredicateListWithPagination{
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

func (f graduationPredicateService) Create(ctx context.Context, data objects.CreateGraduationPredicate) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	createData := models.CreateGraduationPredicate{
		Predicate:                   data.Predicate,
		MinimumGpa:                  data.MinimumGpa,
		MaximumStudySemester:        data.MaximumStudySemester,
		RepeatCourseLimit:           data.RepeatCourseLimit,
		BelowMinimumGradePointLimit: data.BelowMinimumGradePointLimit,
		CreatedBy:                   claims.ID,
	}
	errs = f.GraduationPredicateRepo.Create(ctx, tx, createData)
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

func (f graduationPredicateService) Update(ctx context.Context, data objects.UpdateGraduationPredicate) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs = f.GraduationPredicateRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateGraduationPredicate{
		Id:                          data.Id,
		Predicate:                   data.Predicate,
		MinimumGpa:                  data.MinimumGpa,
		MaximumStudySemester:        data.MaximumStudySemester,
		RepeatCourseLimit:           data.RepeatCourseLimit,
		BelowMinimumGradePointLimit: data.BelowMinimumGradePointLimit,
		UpdatedBy:                   claims.ID,
	}
	errs = f.GraduationPredicateRepo.Update(ctx, tx, updateData)
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

func (f graduationPredicateService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := f.GraduationPredicateRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = f.GraduationPredicateRepo.Delete(ctx, tx, id)
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
