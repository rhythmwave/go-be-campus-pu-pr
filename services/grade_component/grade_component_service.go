package grade_component

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

type gradeComponentService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (f gradeComponentService) GetList(ctx context.Context, paginationData common.PaginationRequest, studyProgramId, subjectCategoryId string) (objects.GradeComponentListWithPagination, *constants.ErrorResponse) {
	var result objects.GradeComponentListWithPagination

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	_, errs = f.StudyProgramRepo.GetDetail(ctx, tx, studyProgramId, claims.Role, claims.ID)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	modelResult, paginationResult, errs := f.GradeComponentRepo.GetList(ctx, tx, paginationData, studyProgramId, subjectCategoryId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetGradeComponent{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetGradeComponent{
			Id:                  v.Id,
			StudyProgramId:      v.StudyProgramId,
			StudyProgramName:    v.StudyProgramName,
			SubjectCategoryId:   v.SubjectCategoryId,
			SubjectCategoryName: v.SubjectCategoryName,
			Name:                v.Name,
			IsActive:            v.IsActive,
			DefaultPercentage:   v.DefaultPercentage,
		})
	}

	result = objects.GradeComponentListWithPagination{
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

func (f gradeComponentService) Create(ctx context.Context, data objects.CreateGradeComponent) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs = f.StudyProgramRepo.GetDetail(ctx, tx, data.StudyProgramId, claims.Role, claims.ID)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	createData := models.CreateGradeComponent{
		StudyProgramId:    data.StudyProgramId,
		SubjectCategoryId: data.SubjectCategoryId,
		Name:              data.Name,
		IsActive:          data.IsActive,
		CreatedBy:         claims.ID,
	}
	errs = f.GradeComponentRepo.Create(ctx, tx, createData)
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

func (f gradeComponentService) Update(ctx context.Context, data objects.UpdateGradeComponent) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs = f.GradeComponentRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateGradeComponent{
		Id:                data.Id,
		SubjectCategoryId: data.SubjectCategoryId,
		Name:              data.Name,
		IsActive:          data.IsActive,
		UpdatedBy:         claims.ID,
	}
	errs = f.GradeComponentRepo.Update(ctx, tx, updateData)
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

func (f gradeComponentService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := f.GradeComponentRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = f.GradeComponentRepo.Delete(ctx, tx, id)
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

func (f gradeComponentService) GetListBySubjectCategory(ctx context.Context, paginationData common.PaginationRequest, studyProgramId string) (objects.GradeComponentBySubjectCategoryListWithPagination, *constants.ErrorResponse) {
	var result objects.GradeComponentBySubjectCategoryListWithPagination

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	_, errs = f.StudyProgramRepo.GetDetail(ctx, tx, studyProgramId, claims.Role, claims.ID)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	modelResult, paginationResult, errs := f.GradeComponentRepo.GetDistinctSubjectCategoryList(ctx, tx, paginationData, studyProgramId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}
	subjectCategoryIds := []string{}
	for _, v := range modelResult {
		subjectCategoryIds = append(subjectCategoryIds, v.SubjectCategoryId)
	}

	gradeComponentData, errs := f.GradeComponentRepo.GetPercentageBySubjectCategories(ctx, tx, studyProgramId, subjectCategoryIds)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result = objects.GradeComponentBySubjectCategoryListWithPagination{
		Pagination: paginationResult,
		Data:       mapGetListBySubjectCategory(modelResult, gradeComponentData),
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (f gradeComponentService) BulkUpdatePercentage(ctx context.Context, baseData objects.BulkUpdatePercentageGradeComponent, data []objects.BulkUpdatePercentageGradeComponentData) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	if len(data) == 0 {
		return utils.ErrEmptyValue("grade_components")
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs = f.StudyProgramRepo.GetDetail(ctx, tx, baseData.StudyProgramId, claims.Role, claims.ID)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	modelData := []models.BulkUpdateGradeComponentPercentage{}
	for _, v := range data {
		modelData = append(modelData, models.BulkUpdateGradeComponentPercentage{
			Id:                v.Id,
			StudyProgramId:    baseData.StudyProgramId,
			SubjectCategoryId: baseData.SubjectCategoryId,
			Name:              utils.RandomString(4),
			DefaultPercentage: v.DefaultPercentage,
			IsActive:          v.IsActive,
			UpdatedBy:         claims.ID,
		})
	}

	errs = f.GradeComponentRepo.BulkUpdatePercentage(ctx, tx, modelData)
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
