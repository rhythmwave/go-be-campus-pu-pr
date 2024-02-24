package learning_achievement

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

type learningAchievementService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (f learningAchievementService) GetList(ctx context.Context, paginationData common.PaginationRequest, learningAchievementCategoryId string) (objects.LearningAchievementListWithPagination, *constants.ErrorResponse) {
	var result objects.LearningAchievementListWithPagination

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := f.LearningAchievementRepo.GetList(ctx, tx, paginationData, learningAchievementCategoryId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetLearningAchievement{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetLearningAchievement{
			Id:          v.Id,
			Name:        v.Name,
			EnglishName: v.EnglishName,
		})
	}

	result = objects.LearningAchievementListWithPagination{
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

func (f learningAchievementService) Create(ctx context.Context, data objects.CreateLearningAchievement) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	createData := models.CreateLearningAchievement{
		LearningAchievementCategoryId: data.LearningAchievementCategoryId,
		Name:                          data.Name,
		EnglishName:                   data.EnglishName,
		CreatedBy:                     claims.ID,
	}
	errs = f.LearningAchievementRepo.Create(ctx, tx, createData)
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

func (f learningAchievementService) Update(ctx context.Context, data objects.UpdateLearningAchievement) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs = f.LearningAchievementRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateLearningAchievement{
		Id:          data.Id,
		Name:        data.Name,
		EnglishName: data.EnglishName,
		UpdatedBy:   claims.ID,
	}
	errs = f.LearningAchievementRepo.Update(ctx, tx, updateData)
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

func (f learningAchievementService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := f.LearningAchievementRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = f.LearningAchievementRepo.Delete(ctx, tx, id)
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
