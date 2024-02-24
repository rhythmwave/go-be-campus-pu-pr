package grade_type

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

type gradeTypeService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (f gradeTypeService) GetList(ctx context.Context, paginationData common.PaginationRequest, studyLevelId string) (objects.GradeTypeListWithPagination, *constants.ErrorResponse) {
	var result objects.GradeTypeListWithPagination

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := f.GradeTypeRepo.GetList(ctx, tx, paginationData, studyLevelId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetGradeType{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetGradeType{
			Id:                  v.Id,
			StudyLevelId:        v.StudyLevelId,
			StudyLevelShortName: v.StudyLevelShortName,
			Code:                v.Code,
			GradePoint:          v.GradePoint,
			MinimumGrade:        v.MinimumGrade,
			MaximumGrade:        v.MaximumGrade,
			GradeCategory:       v.GradeCategory,
			GradePointCategory:  v.GradePointCategory,
			Label:               v.Label,
			EnglishLabel:        v.EnglishLabel,
			StartDate:           v.StartDate,
			EndDate:             v.EndDate,
		})
	}

	result = objects.GradeTypeListWithPagination{
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

func (f gradeTypeService) Create(ctx context.Context, data objects.CreateGradeType) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	createData := models.CreateGradeType{
		StudyLevelId:       data.StudyLevelId,
		Code:               data.Code,
		GradePoint:         data.GradePoint,
		MinimumGrade:       data.MinimumGrade,
		MaximumGrade:       data.MaximumGrade,
		GradeCategory:      data.GradeCategory,
		GradePointCategory: data.GradePointCategory,
		Label:              utils.NewNullString(data.Label),
		EnglishLabel:       utils.NewNullString(data.EnglishLabel),
		StartDate:          data.StartDate,
		EndDate:            data.EndDate,
		CreatedBy:          claims.ID,
	}
	errs = f.GradeTypeRepo.Create(ctx, tx, createData)
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

func (f gradeTypeService) Update(ctx context.Context, data objects.UpdateGradeType) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs = f.GradeTypeRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateGradeType{
		Id:                 data.Id,
		Code:               data.Code,
		GradePoint:         data.GradePoint,
		MinimumGrade:       data.MinimumGrade,
		MaximumGrade:       data.MaximumGrade,
		GradeCategory:      data.GradeCategory,
		GradePointCategory: data.GradePointCategory,
		Label:              utils.NewNullString(data.Label),
		EnglishLabel:       utils.NewNullString(data.EnglishLabel),
		StartDate:          data.StartDate,
		EndDate:            data.EndDate,
		UpdatedBy:          claims.ID,
	}
	errs = f.GradeTypeRepo.Update(ctx, tx, updateData)
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

func (f gradeTypeService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := f.GradeTypeRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = f.GradeTypeRepo.Delete(ctx, tx, id)
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
