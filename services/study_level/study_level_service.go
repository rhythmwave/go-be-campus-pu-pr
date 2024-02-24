package study_level

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

type studyLevelService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (d studyLevelService) GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.StudyLevelListWithPagination, *constants.ErrorResponse) {
	var result objects.StudyLevelListWithPagination

	tx, err := d.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := d.StudyLevelRepo.GetList(ctx, tx, paginationData)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetStudyLevel{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetStudyLevel{
			Id:                    v.Id,
			Name:                  v.Name,
			ShortName:             v.ShortName,
			KkniQualification:     v.KkniQualification,
			AcceptanceRequirement: v.AcceptanceRequirement,
			FurtherEducationLevel: v.FurtherEducationLevel,
			ProfessionalStatus:    v.ProfessionalStatus,
			CourseLanguage:        v.CourseLanguage,
		})
	}

	result = objects.StudyLevelListWithPagination{
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

func (d studyLevelService) UpdateSkpi(ctx context.Context, data objects.UpdateStudyLevelSkpi) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := d.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs = d.StudyLevelRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateStudyLevelSkpi{
		Id:                    data.Id,
		KkniQualification:     utils.NewNullString(data.KkniQualification),
		AcceptanceRequirement: utils.NewNullString(data.AcceptanceRequirement),
		FurtherEducationLevel: utils.NewNullString(data.FurtherEducationLevel),
		ProfessionalStatus:    utils.NewNullString(data.ProfessionalStatus),
		CourseLanguage:        utils.NewNullString(data.CourseLanguage),
		UpdatedBy:             claims.ID,
	}
	errs = d.StudyLevelRepo.UpdateSkpi(ctx, tx, updateData)
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
