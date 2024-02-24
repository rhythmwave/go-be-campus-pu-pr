package shared_file

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

type sharedFileService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (s sharedFileService) GetList(ctx context.Context, paginationData common.PaginationRequest, appType string, onlyOwned bool, isApproved *bool) (objects.SharedFileListWithPagination, *constants.ErrorResponse) {
	var result objects.SharedFileListWithPagination

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	var lecturerId string
	if onlyOwned {
		lecturerId = claims.ID
	}

	modelResult, paginationResult, errs := s.SharedFileRepo.GetList(ctx, tx, paginationData, appType, lecturerId, isApproved)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetSharedFile{}
	for _, v := range modelResult {
		fileUrl, errs := s.Storage.GetURL(v.FilePath, v.FilePathType, nil)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}

		resultData = append(resultData, objects.GetSharedFile{
			Id:                 v.Id,
			LecturerId:         v.LecturerId,
			LecturerName:       v.LecturerName,
			LecturerFrontTitle: v.LecturerFrontTitle,
			LecturerBackDegree: v.LecturerBackDegree,
			Title:              v.Title,
			FilePath:           v.FilePath,
			FilePathType:       v.FilePathType,
			FileUrl:            fileUrl,
			Remarks:            v.Remarks,
			IsApproved:         v.IsApproved,
			CreatedAt:          v.CreatedAt,
		})
	}

	result = objects.SharedFileListWithPagination{
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

func (s sharedFileService) Create(ctx context.Context, data objects.CreateSharedFile) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	createData := models.CreateSharedFile{
		LecturerId:   claims.ID,
		Title:        data.Title,
		FilePath:     data.FilePath,
		FilePathType: data.FilePathType,
		Remarks:      utils.NewNullString(data.Remarks),
	}
	errs = s.SharedFileRepo.Create(ctx, tx, createData)
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

func (s sharedFileService) Update(ctx context.Context, data objects.UpdateSharedFile) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	sharedFileData, errs := s.SharedFileRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	if sharedFileData.LecturerId != claims.ID {
		_ = tx.Rollback()
		return appConstants.ErrUneditableSharedFile
	}

	updateData := models.UpdateSharedFile{
		Id:      data.Id,
		Title:   data.Title,
		Remarks: utils.NewNullString(data.Remarks),
	}
	errs = s.SharedFileRepo.Update(ctx, tx, updateData)
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

func (s sharedFileService) Approve(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := s.SharedFileRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = s.SharedFileRepo.Approve(ctx, tx, id)
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

func (s sharedFileService) Delete(ctx context.Context, id, appType string) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	sharedFileData, errs := s.SharedFileRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	if appType == appConstants.AppTypeLecturer && sharedFileData.LecturerId != claims.ID {
		_ = tx.Rollback()
		return appConstants.ErrUneditableSharedFile
	}

	errs = s.SharedFileRepo.Delete(ctx, tx, id)
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
