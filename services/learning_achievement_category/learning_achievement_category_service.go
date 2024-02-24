package learning_achievement_category

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

type learningAchievementCategoryService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (f learningAchievementCategoryService) GetList(ctx context.Context, paginationData common.PaginationRequest, curriculumId string) (objects.LearningAchievementCategoryListWithPagination, *constants.ErrorResponse) {
	var result objects.LearningAchievementCategoryListWithPagination

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := f.LearningAchievementCategoryRepo.GetList(ctx, tx, paginationData, curriculumId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetLearningAchievementCategory{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetLearningAchievementCategory{
			Id:          v.Id,
			Name:        v.Name,
			EnglishName: v.EnglishName,
		})
	}

	result = objects.LearningAchievementCategoryListWithPagination{
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

func (f learningAchievementCategoryService) Create(ctx context.Context, data objects.CreateLearningAchievementCategory) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	createData := models.CreateLearningAchievementCategory{
		CurriculumId: data.CurriculumId,
		Name:         data.Name,
		EnglishName:  data.EnglishName,
		CreatedBy:    claims.ID,
	}
	errs = f.LearningAchievementCategoryRepo.Create(ctx, tx, createData)
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

func (f learningAchievementCategoryService) Update(ctx context.Context, data objects.UpdateLearningAchievementCategory) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs = f.LearningAchievementCategoryRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateLearningAchievementCategory{
		Id:          data.Id,
		Name:        data.Name,
		EnglishName: data.EnglishName,
		UpdatedBy:   claims.ID,
	}
	errs = f.LearningAchievementCategoryRepo.Update(ctx, tx, updateData)
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

func (f learningAchievementCategoryService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := f.LearningAchievementCategoryRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = f.LearningAchievementCategoryRepo.Delete(ctx, tx, id)
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
