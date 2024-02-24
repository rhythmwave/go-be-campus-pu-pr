package class_event

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	appConstants "github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
)

type classEventService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (c classEventService) GetList(ctx context.Context, paginationData common.PaginationRequest, classId, frequency string, futureEventOnly bool, isActive *bool) (objects.ClassEventListWithPagination, *constants.ErrorResponse) {
	var result objects.ClassEventListWithPagination

	tx, err := c.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := c.ClassEventRepo.GetList(ctx, tx, paginationData, classId, frequency, futureEventOnly, isActive)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetClassEvent{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetClassEvent{
			Id:                 v.Id,
			Title:              v.Title,
			Frequency:          v.Frequency,
			EventTime:          v.EventTime,
			LecturerId:         v.LecturerId,
			LecturerName:       v.LecturerName,
			LecturerFrontTitle: v.LecturerFrontTitle,
			LecturerBackDegree: v.LecturerBackDegree,
			Remarks:            v.Remarks,
			IsActive:           v.IsActive,
			CreatedAt:          v.CreatedAt,
		})
	}

	result = objects.ClassEventListWithPagination{
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

func (c classEventService) Create(ctx context.Context, data objects.CreateClassEvent) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := c.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	classLecturerData, errs := c.ClassLecturerRepo.GetByClassIdLecturerId(ctx, tx, data.ClassId, claims.ID)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	if classLecturerData.Id == "" {
		_ = tx.Rollback()
		return appConstants.ErrUneditableClassEvent
	}

	createData := models.CreateClassEvent{
		LecturerId: claims.ID,
		ClassId:    data.ClassId,
		Title:      data.Title,
		Frequency:  data.Frequency,
		EventTime:  data.EventTime,
		Remarks:    utils.NewNullString(data.Remarks),
		IsActive:   data.IsActive,
	}
	errs = c.ClassEventRepo.Create(ctx, tx, createData)
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

func (c classEventService) Update(ctx context.Context, data objects.UpdateClassEvent) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := c.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	classEventData, errs := c.ClassEventRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	if classEventData.LecturerId != claims.ID {
		_ = tx.Rollback()
		return appConstants.ErrUneditableClassEvent
	}

	updateData := models.UpdateClassEvent{
		Id:        data.Id,
		Title:     data.Title,
		Frequency: data.Frequency,
		EventTime: data.EventTime,
		Remarks:   utils.NewNullString(data.Remarks),
		IsActive:  data.IsActive,
	}
	errs = c.ClassEventRepo.Update(ctx, tx, updateData)
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

func (c classEventService) BulkUpdateActivation(ctx context.Context, ids []string, isActive bool) *constants.ErrorResponse {
	tx, err := c.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	errs := c.ClassEventRepo.BulkUpdateActivation(ctx, tx, ids, isActive)
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

func (c classEventService) BulkDelete(ctx context.Context, ids []string) *constants.ErrorResponse {
	tx, err := c.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	errs := c.ClassEventRepo.BulkDelete(ctx, tx, ids)
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
