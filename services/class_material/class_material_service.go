package class_material

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

type classMaterialService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (c classMaterialService) GetList(ctx context.Context, paginationData common.PaginationRequest, classId string, isActive *bool) (objects.ClassMaterialListWithPagination, *constants.ErrorResponse) {
	var result objects.ClassMaterialListWithPagination

	tx, err := c.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := c.ClassMaterialRepo.GetList(ctx, tx, paginationData, classId, isActive)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetClassMaterial{}
	for _, v := range modelResult {
		var fileUrl string
		if v.FilePath != nil && v.FilePathType != nil {
			fileUrl, errs = c.Storage.GetURL(*v.FilePath, *v.FilePathType, nil)
			if errs != nil {
				_ = tx.Rollback()
				return result, errs
			}
		}

		resultData = append(resultData, objects.GetClassMaterial{
			Id:                 v.Id,
			Title:              v.Title,
			Abstraction:        v.Abstraction,
			FilePath:           v.FilePath,
			FilePathType:       v.FilePathType,
			FileUrl:            fileUrl,
			LecturerId:         v.LecturerId,
			LecturerName:       v.LecturerName,
			LecturerFrontTitle: v.LecturerFrontTitle,
			LecturerBackDegree: v.LecturerBackDegree,
			IsActive:           v.IsActive,
			CreatedAt:          v.CreatedAt,
		})
	}

	result = objects.ClassMaterialListWithPagination{
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

func (c classMaterialService) Create(ctx context.Context, data objects.CreateClassMaterial) *constants.ErrorResponse {
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
		return appConstants.ErrUneditableClassMaterial
	}

	createData := models.CreateClassMaterial{
		LecturerId:   claims.ID,
		ClassId:      data.ClassId,
		Title:        data.Title,
		Abstraction:  utils.NewNullString(data.Abstraction),
		FilePath:     utils.NewNullString(data.FilePath),
		FilePathType: utils.NewNullString(data.FilePathType),
		IsActive:     data.IsActive,
	}
	errs = c.ClassMaterialRepo.Create(ctx, tx, createData)
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

func (c classMaterialService) Update(ctx context.Context, data objects.UpdateClassMaterial) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := c.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	classMaterialData, errs := c.ClassMaterialRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	if classMaterialData.LecturerId != claims.ID {
		_ = tx.Rollback()
		return appConstants.ErrUneditableClassMaterial
	}

	updateData := models.UpdateClassMaterial{
		Id:           data.Id,
		Title:        data.Title,
		Abstraction:  utils.NewNullString(data.Abstraction),
		FilePath:     utils.NewNullString(data.FilePath),
		FilePathType: utils.NewNullString(data.FilePathType),
		IsActive:     data.IsActive,
	}
	errs = c.ClassMaterialRepo.Update(ctx, tx, updateData)
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

func (c classMaterialService) BulkUpdateActivation(ctx context.Context, ids []string, isActive bool) *constants.ErrorResponse {
	tx, err := c.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	errs := c.ClassMaterialRepo.BulkUpdateActivation(ctx, tx, ids, isActive)
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

func (c classMaterialService) BulkDelete(ctx context.Context, ids []string) *constants.ErrorResponse {
	tx, err := c.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	errs := c.ClassMaterialRepo.BulkDelete(ctx, tx, ids)
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
