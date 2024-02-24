package lesson_plan

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

type lessonPlanService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (f lessonPlanService) GetList(ctx context.Context, paginationData common.PaginationRequest, subjectId string) (objects.LessonPlanListWithPagination, *constants.ErrorResponse) {
	var result objects.LessonPlanListWithPagination

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := f.LessonPlanRepo.GetList(ctx, tx, paginationData, subjectId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetLessonPlan{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetLessonPlan{
			Id:            v.Id,
			MeetingOrder:  v.MeetingOrder,
			Lesson:        v.Lesson,
			EnglishLesson: v.EnglishLesson,
		})
	}

	result = objects.LessonPlanListWithPagination{
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

func (f lessonPlanService) Create(ctx context.Context, data objects.CreateLessonPlan) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	createData := models.CreateLessonPlan{
		SubjectId:     data.SubjectId,
		MeetingOrder:  data.MeetingOrder,
		Lesson:        data.Lesson,
		EnglishLesson: data.EnglishLesson,
		CreatedBy:     claims.ID,
	}
	errs = f.LessonPlanRepo.Create(ctx, tx, createData)
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

func (f lessonPlanService) Update(ctx context.Context, data objects.UpdateLessonPlan) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs = f.LessonPlanRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateLessonPlan{
		Id:            data.Id,
		MeetingOrder:  data.MeetingOrder,
		Lesson:        data.Lesson,
		EnglishLesson: data.EnglishLesson,
		UpdatedBy:     claims.ID,
	}
	errs = f.LessonPlanRepo.Update(ctx, tx, updateData)
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

func (f lessonPlanService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := f.LessonPlanRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = f.LessonPlanRepo.Delete(ctx, tx, id)
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
