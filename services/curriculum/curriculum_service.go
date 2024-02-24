package curriculum

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

type curriculumService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func mapGetList(curriculumData []models.GetCurriculum) []objects.GetCurriculum {
	results := []objects.GetCurriculum{}

	for _, v := range curriculumData {
		results = append(results, objects.GetCurriculum{
			Id:                           v.Id,
			StudyProgramId:               v.StudyProgramId,
			StudyProgramName:             v.StudyProgramName,
			DiktiStudyProgramCode:        v.DiktiStudyProgramCode,
			Name:                         v.Name,
			Year:                         v.Year,
			IdealStudyPeriod:             v.IdealStudyPeriod,
			MaximumStudyPeriod:           v.MaximumStudyPeriod,
			IsActive:                     v.IsActive,
			TotalSubject:                 v.TotalSubject,
			TotalSubjectWithPrerequisite: v.TotalSubjectWithPrerequisite,
			TotalSubjectWithEquivalence:  v.TotalSubjectWithEquivalence,
		})
	}

	return results
}

func (a curriculumService) GetList(ctx context.Context, paginationData common.PaginationRequest, studyProgramId string) (objects.CurriculumListWithPagination, *constants.ErrorResponse) {
	var result objects.CurriculumListWithPagination

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	studyProgramIds := []string{}
	if studyProgramId != "" {
		_, errs = a.StudyProgramRepo.GetDetail(ctx, tx, studyProgramId, claims.Role, claims.ID)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
		studyProgramIds = append(studyProgramIds, studyProgramId)
	} else {
		studyProgramData, _, errs := a.StudyProgramRepo.GetList(ctx, tx, common.PaginationRequest{}, "", claims.Role, claims.ID)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
		for _, v := range studyProgramData {
			studyProgramIds = append(studyProgramIds, v.Id)
		}
	}

	modelResult, paginationResult, errs := a.CurriculumRepo.GetList(ctx, tx, paginationData, studyProgramIds)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result = objects.CurriculumListWithPagination{
		Pagination: paginationResult,
		Data:       mapGetList(modelResult),
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (f curriculumService) GetDetail(ctx context.Context, id string) (objects.GetCurriculumDetail, *constants.ErrorResponse) {
	var result objects.GetCurriculumDetail

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	resultData, errs := f.CurriculumRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result = objects.GetCurriculumDetail{
		Id:                    resultData.Id,
		StudyProgramId:        resultData.StudyProgramId,
		StudyProgramName:      resultData.StudyProgramName,
		StudyLevelShortName:   resultData.StudyLevelShortName,
		DiktiStudyProgramType: resultData.DiktiStudyProgramType,
		Name:                  resultData.Name,
		Year:                  resultData.Year,
		RectorDecisionNumber:  resultData.RectorDecisionNumber,
		RectorDecisionDate:    resultData.RectorDecisionDate,
		AggreeingParty:        resultData.AggreeingParty,
		AggreementDate:        resultData.AggreementDate,
		IdealStudyPeriod:      resultData.IdealStudyPeriod,
		MaximumStudyPeriod:    resultData.MaximumStudyPeriod,
		Remarks:               resultData.Remarks,
		IsActive:              resultData.IsActive,
		FinalScoreDeterminant: resultData.FinalScoreDeterminant,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (f curriculumService) GetActiveByStudyProgramId(ctx context.Context, studyProgramId string) (objects.GetCurriculumDetail, *constants.ErrorResponse) {
	var result objects.GetCurriculumDetail

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	resultData, errs := f.CurriculumRepo.GetActiveByStudyProgramId(ctx, tx, studyProgramId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result = objects.GetCurriculumDetail{
		Id:                    resultData.Id,
		StudyProgramId:        resultData.StudyProgramId,
		StudyProgramName:      resultData.StudyProgramName,
		StudyLevelShortName:   resultData.StudyLevelShortName,
		DiktiStudyProgramType: resultData.DiktiStudyProgramType,
		Name:                  resultData.Name,
		Year:                  resultData.Year,
		RectorDecisionNumber:  resultData.RectorDecisionNumber,
		RectorDecisionDate:    resultData.RectorDecisionDate,
		AggreeingParty:        resultData.AggreeingParty,
		AggreementDate:        resultData.AggreementDate,
		IdealStudyPeriod:      resultData.IdealStudyPeriod,
		MaximumStudyPeriod:    resultData.MaximumStudyPeriod,
		Remarks:               resultData.Remarks,
		IsActive:              resultData.IsActive,
		FinalScoreDeterminant: resultData.FinalScoreDeterminant,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (a curriculumService) Create(ctx context.Context, data objects.CreateCurriculum) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs = a.StudyProgramRepo.GetDetail(ctx, tx, data.StudyProgramId, claims.Role, claims.ID)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	createData := models.CreateCurriculum{
		StudyProgramId:        data.StudyProgramId,
		Name:                  data.Name,
		Year:                  data.Year,
		RectorDecisionNumber:  utils.NewNullString(data.RectorDecisionNumber),
		RectorDecisionDate:    utils.NewNullTime(data.RectorDecisionDate),
		AggreeingParty:        utils.NewNullString(data.AggreeingParty),
		AggreementDate:        utils.NewNullTime(data.AggreementDate),
		IdealStudyPeriod:      data.IdealStudyPeriod,
		MaximumStudyPeriod:    data.MaximumStudyPeriod,
		Remarks:               utils.NewNullString(data.Remarks),
		IsActive:              data.IsActive,
		FinalScoreDeterminant: data.FinalScoreDeterminant,
		CreatedBy:             claims.ID,
	}
	errs = a.CurriculumRepo.Create(ctx, tx, createData)
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

func (a curriculumService) Update(ctx context.Context, data objects.UpdateCurriculum) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := a.CurriculumRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateCurriculum{
		Id:                    data.Id,
		Name:                  data.Name,
		Year:                  data.Year,
		RectorDecisionNumber:  utils.NewNullString(data.RectorDecisionNumber),
		RectorDecisionDate:    utils.NewNullTime(data.RectorDecisionDate),
		AggreeingParty:        utils.NewNullString(data.AggreeingParty),
		AggreementDate:        utils.NewNullTime(data.AggreementDate),
		IdealStudyPeriod:      data.IdealStudyPeriod,
		MaximumStudyPeriod:    data.MaximumStudyPeriod,
		Remarks:               utils.NewNullString(data.Remarks),
		IsActive:              data.IsActive,
		FinalScoreDeterminant: data.FinalScoreDeterminant,
		UpdatedBy:             claims.ID,
	}
	errs = a.CurriculumRepo.Update(ctx, tx, updateData)
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

func (a curriculumService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := a.CurriculumRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = a.CurriculumRepo.Delete(ctx, tx, id)
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
