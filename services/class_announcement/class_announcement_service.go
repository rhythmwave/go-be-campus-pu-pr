package class_announcement

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

type classAnnouncementService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (c classAnnouncementService) GetList(ctx context.Context, paginationData common.PaginationRequest, appType string, classIds []string) (objects.ClassAnnouncementListWithPagination, *constants.ErrorResponse) {
	var result objects.ClassAnnouncementListWithPagination

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := c.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	if appType == appConstants.AppTypeStudent {
		activeSemesterData, errs := c.SemesterRepo.GetActive(ctx, tx)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}

		classData, _, errs := c.StudentClassRepo.GetList(ctx, tx, common.PaginationRequest{Page: constants.DefaultPage, Limit: constants.DefaultUnlimited}, "", claims.ID, activeSemesterData.Id, nil)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}

		for _, v := range classData {
			classIds = append(classIds, v.ClassId)
		}
	}

	modelResult, paginationResult, errs := c.ClassAnnouncementRepo.GetList(ctx, tx, paginationData, classIds, nil)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetClassAnnouncement{}
	for _, v := range modelResult {
		var fileUrl string
		if v.FilePath != nil && v.FilePathType != nil {
			fileUrl, errs = c.Storage.GetURL(*v.FilePath, *v.FilePathType, nil)
			if errs != nil {
				_ = tx.Rollback()
				return result, errs
			}
		}

		resultData = append(resultData, objects.GetClassAnnouncement{
			Id:                 v.Id,
			Title:              v.Title,
			Content:            v.Content,
			LecturerId:         v.LecturerId,
			LecturerName:       v.LecturerName,
			LecturerFrontTitle: v.LecturerFrontTitle,
			LecturerBackDegree: v.LecturerBackDegree,
			FileUrl:            fileUrl,
			FilePath:           v.FilePath,
			FilePathType:       v.FilePathType,
			StartTime:          v.StartTime,
			EndTime:            v.EndTime,
		})
	}

	result = objects.ClassAnnouncementListWithPagination{
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

func (c classAnnouncementService) Create(ctx context.Context, data objects.CreateClassAnnouncement) *constants.ErrorResponse {
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
		return appConstants.ErrUneditableClassAnnouncement
	}

	createData := models.CreateClassAnnouncement{
		LecturerId:   claims.ID,
		ClassId:      data.ClassId,
		Title:        data.Title,
		Content:      data.Content,
		FilePath:     utils.NewNullString(data.FilePath),
		FilePathType: utils.NewNullString(data.FilePathType),
		StartTime:    utils.NewNullTime(data.StartTime),
		EndTime:      utils.NewNullTime(data.EndTime),
	}
	errs = c.ClassAnnouncementRepo.Create(ctx, tx, createData)
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

func (c classAnnouncementService) Update(ctx context.Context, data objects.UpdateClassAnnouncement) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := c.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	classAnnouncementData, errs := c.ClassAnnouncementRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	if classAnnouncementData.LecturerId != claims.ID {
		_ = tx.Rollback()
		return appConstants.ErrUneditableClassAnnouncement
	}

	updateData := models.UpdateClassAnnouncement{
		Id:           data.Id,
		Title:        data.Title,
		Content:      data.Content,
		FilePath:     utils.NewNullString(data.FilePath),
		FilePathType: utils.NewNullString(data.FilePathType),
		StartTime:    utils.NewNullTime(data.StartTime),
		EndTime:      utils.NewNullTime(data.EndTime),
	}
	errs = c.ClassAnnouncementRepo.Update(ctx, tx, updateData)
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

func (c classAnnouncementService) Delete(ctx context.Context, ids []string) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := c.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	paginationData := common.PaginationRequest{
		Page:  constants.DefaultPage,
		Limit: constants.DefaultUnlimited,
	}
	classAnnouncementData, _, errs := c.ClassAnnouncementRepo.GetList(ctx, tx, paginationData, nil, ids)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	for _, v := range classAnnouncementData {
		if v.LecturerId != claims.ID {
			_ = tx.Rollback()
			return appConstants.ErrUneditableClassAnnouncement
		}
	}

	errs = c.ClassAnnouncementRepo.Delete(ctx, tx, ids)
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
