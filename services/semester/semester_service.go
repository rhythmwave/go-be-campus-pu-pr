package semester

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
	appUtils "github.com/sccicitb/pupr-backend/utils"
)

type semesterService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (a semesterService) GetList(ctx context.Context, paginationData common.PaginationRequest, studyProgramId, excludedId string) (objects.SemesterListWithPagination, *constants.ErrorResponse) {
	var result objects.SemesterListWithPagination

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	if studyProgramId != "" {
		_, errs = a.StudyProgramRepo.GetDetail(ctx, tx, studyProgramId, claims.Role, claims.ID)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	}

	modelResult, paginationResult, errs := a.SemesterRepo.GetList(ctx, tx, paginationData, studyProgramId, excludedId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	semesterIds := []string{}
	for _, v := range modelResult {
		semesterIds = append(semesterIds, v.Id)
	}

	curriculumData := []models.GetSemesterCurriculum{}
	if len(semesterIds) != 0 {
		curriculumData, errs = a.SemesterRepo.GetCurriculumBySemesterIds(ctx, tx, semesterIds)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	}

	curriculumMap := make(map[string][]objects.GetSemesterCurriculum)
	for _, v := range curriculumData {
		curriculumMap[v.SemesterId] = append(curriculumMap[v.SemesterId], objects.GetSemesterCurriculum{
			StudyProgramId:   v.StudyProgramId,
			StudyProgramName: v.StudyProgramName,
			CurriculumId:     v.CurriculumId,
			CurriculumName:   v.CurriculumName,
		})
	}

	resultData := []objects.GetSemester{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetSemester{
			Id:                v.Id,
			SemesterStartYear: v.SemesterStartYear,
			SchoolYear:        appUtils.GenerateSchoolYear(v.SemesterStartYear),
			SemesterType:      v.SemesterType,
			IsActive:          v.IsActive,
			StartDate:         v.StartDate,
			EndDate:           v.EndDate,
			MidtermStartDate:  v.MidtermStartDate,
			MidtermEndDate:    v.MidtermEndDate,
			EndtermStartDate:  v.EndtermStartDate,
			EndtermEndDate:    v.EndtermEndDate,
			Curriculums:       curriculumMap[v.Id],
		})
	}

	result = objects.SemesterListWithPagination{
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

func (a semesterService) GetDetail(ctx context.Context, id string) (objects.GetSemesterDetail, *constants.ErrorResponse) {
	var result objects.GetSemesterDetail

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	resultData, errs := a.SemesterRepo.GetById(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	curriculumData, errs := a.SemesterRepo.GetCurriculumBySemesterIds(ctx, tx, []string{id})
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	curriculums := []objects.GetSemesterDetailCurriculum{}
	for _, v := range curriculumData {
		curriculums = append(curriculums, objects.GetSemesterDetailCurriculum{
			StudyProgramId:   v.StudyProgramId,
			StudyProgramName: v.StudyProgramName,
			CurriculumId:     v.CurriculumId,
			CurriculumName:   v.CurriculumName,
		})
	}

	var referenceSchoolYear string
	if resultData.ReferenceSemesterStartYear != nil {
		referenceSchoolYear = appUtils.GenerateSchoolYear(*resultData.ReferenceSemesterStartYear)
	}
	result = objects.GetSemesterDetail{
		Id:                         resultData.Id,
		SemesterStartYear:          resultData.SemesterStartYear,
		SchoolYear:                 appUtils.GenerateSchoolYear(resultData.SemesterStartYear),
		SemesterType:               resultData.SemesterType,
		IsActive:                   resultData.IsActive,
		StartDate:                  resultData.StartDate,
		EndDate:                    resultData.EndDate,
		MidtermStartDate:           resultData.MidtermStartDate,
		MidtermEndDate:             resultData.MidtermEndDate,
		EndtermStartDate:           resultData.EndtermStartDate,
		EndtermEndDate:             resultData.EndtermEndDate,
		StudyPlanInputStartDate:    resultData.StudyPlanInputStartDate,
		StudyPlanInputEndDate:      resultData.StudyPlanInputEndDate,
		StudyPlanApprovalStartDate: resultData.StudyPlanApprovalStartDate,
		StudyPlanApprovalEndDate:   resultData.StudyPlanApprovalEndDate,
		ReferenceSemesterId:        resultData.ReferenceSemesterId,
		ReferenceSemesterStartYear: resultData.ReferenceSemesterStartYear,
		GradingStartDate:           resultData.GradingStartDate,
		GradingEndDate:             resultData.GradingEndDate,
		ReferenceSchoolYear:        &referenceSchoolYear,
		ReferenceSemesterType:      resultData.ReferenceSemesterType,
		CheckMinimumGpa:            resultData.CheckMinimumGpa,
		CheckPassedCredit:          resultData.CheckPassedCredit,
		DefaultCredit:              resultData.DefaultCredit,
		Curriculums:                curriculums,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (a semesterService) GetActive(ctx context.Context) (objects.GetSemesterDetail, *constants.ErrorResponse) {
	var result objects.GetSemesterDetail

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	resultData, errs := a.SemesterRepo.GetActive(ctx, tx)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	curriculumData, errs := a.SemesterRepo.GetCurriculumBySemesterIds(ctx, tx, []string{resultData.Id})
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	curriculums := []objects.GetSemesterDetailCurriculum{}
	for _, v := range curriculumData {
		curriculums = append(curriculums, objects.GetSemesterDetailCurriculum{
			StudyProgramId:   v.StudyProgramId,
			StudyProgramName: v.StudyProgramName,
			CurriculumId:     v.CurriculumId,
			CurriculumName:   v.CurriculumName,
		})
	}

	var referenceSchoolYear string
	if resultData.ReferenceSemesterStartYear != nil {
		referenceSchoolYear = appUtils.GenerateSchoolYear(*resultData.ReferenceSemesterStartYear)
	}
	result = objects.GetSemesterDetail{
		Id:                         resultData.Id,
		SemesterStartYear:          resultData.SemesterStartYear,
		SchoolYear:                 appUtils.GenerateSchoolYear(resultData.SemesterStartYear),
		SemesterType:               resultData.SemesterType,
		IsActive:                   resultData.IsActive,
		StartDate:                  resultData.StartDate,
		EndDate:                    resultData.EndDate,
		MidtermStartDate:           resultData.MidtermStartDate,
		MidtermEndDate:             resultData.MidtermEndDate,
		EndtermStartDate:           resultData.EndtermStartDate,
		EndtermEndDate:             resultData.EndtermEndDate,
		StudyPlanInputStartDate:    resultData.StudyPlanInputStartDate,
		StudyPlanInputEndDate:      resultData.StudyPlanInputEndDate,
		StudyPlanApprovalStartDate: resultData.StudyPlanApprovalStartDate,
		StudyPlanApprovalEndDate:   resultData.StudyPlanApprovalEndDate,
		ReferenceSemesterId:        resultData.ReferenceSemesterId,
		ReferenceSemesterStartYear: resultData.ReferenceSemesterStartYear,
		GradingStartDate:           resultData.GradingStartDate,
		GradingEndDate:             resultData.GradingEndDate,
		ReferenceSchoolYear:        &referenceSchoolYear,
		ReferenceSemesterType:      resultData.ReferenceSemesterType,
		CheckMinimumGpa:            resultData.CheckMinimumGpa,
		CheckPassedCredit:          resultData.CheckPassedCredit,
		DefaultCredit:              resultData.DefaultCredit,
		Curriculums:                curriculums,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (a semesterService) Create(ctx context.Context, data objects.CreateSemester) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	createData := models.CreateSemester{
		SemesterStartYear:          data.SemesterStartYear,
		SemesterType:               data.SemesterType,
		StartDate:                  data.StartDate,
		EndDate:                    data.EndDate,
		MidtermStartDate:           utils.NewNullTime(data.MidtermStartDate),
		MidtermEndDate:             utils.NewNullTime(data.MidtermEndDate),
		EndtermStartDate:           utils.NewNullTime(data.EndtermStartDate),
		EndtermEndDate:             utils.NewNullTime(data.EndtermEndDate),
		StudyPlanInputStartDate:    data.StudyPlanInputStartDate,
		StudyPlanInputEndDate:      data.StudyPlanInputEndDate,
		StudyPlanApprovalStartDate: data.StudyPlanApprovalStartDate,
		StudyPlanApprovalEndDate:   data.StudyPlanApprovalEndDate,
		GradingStartDate:           utils.NewNullTime(data.GradingStartDate),
		GradingEndDate:             utils.NewNullTime(data.GradingEndDate),
		ReferenceSemesterId:        utils.NewNullString(data.ReferenceSemesterId),
		CheckMinimumGpa:            data.CheckMinimumGpa,
		CheckPassedCredit:          data.CheckPassedCredit,
		DefaultCredit:              data.DefaultCredit,
		CreatedBy:                  claims.ID,
	}
	semesterId, errs := a.SemesterRepo.Create(ctx, tx, createData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	if len(data.Curriculums) != 0 {
		curriculums := []models.UpsertSemesterCurriculum{}
		for _, v := range data.Curriculums {
			curriculums = append(curriculums, models.UpsertSemesterCurriculum{
				SemesterId:   semesterId,
				CurriculumId: v.CurriculumId,
				CreatedBy:    claims.ID,
			})
		}
		errs = a.SemesterRepo.UpsertCurriculum(ctx, tx, curriculums)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (a semesterService) Update(ctx context.Context, data objects.UpdateSemester) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := a.SemesterRepo.GetById(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateSemester{
		Id:                         data.Id,
		SemesterStartYear:          data.SemesterStartYear,
		SemesterType:               data.SemesterType,
		StartDate:                  data.StartDate,
		EndDate:                    data.EndDate,
		MidtermStartDate:           utils.NewNullTime(data.MidtermStartDate),
		MidtermEndDate:             utils.NewNullTime(data.MidtermEndDate),
		EndtermStartDate:           utils.NewNullTime(data.EndtermStartDate),
		EndtermEndDate:             utils.NewNullTime(data.EndtermEndDate),
		StudyPlanInputStartDate:    data.StudyPlanInputStartDate,
		StudyPlanInputEndDate:      data.StudyPlanInputEndDate,
		StudyPlanApprovalStartDate: data.StudyPlanApprovalStartDate,
		StudyPlanApprovalEndDate:   data.StudyPlanApprovalEndDate,
		GradingStartDate:           utils.NewNullTime(data.GradingStartDate),
		GradingEndDate:             utils.NewNullTime(data.GradingEndDate),
		ReferenceSemesterId:        utils.NewNullString(data.ReferenceSemesterId),
		CheckMinimumGpa:            data.CheckMinimumGpa,
		CheckPassedCredit:          data.CheckPassedCredit,
		DefaultCredit:              data.DefaultCredit,
		UpdatedBy:                  claims.ID,
	}
	errs = a.SemesterRepo.Update(ctx, tx, updateData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	curriculums := []models.UpsertSemesterCurriculum{}
	curriculumIds := []string{}
	for _, v := range data.Curriculums {
		curriculums = append(curriculums, models.UpsertSemesterCurriculum{
			SemesterId:   data.Id,
			CurriculumId: v.CurriculumId,
			CreatedBy:    claims.ID,
		})
		curriculumIds = append(curriculumIds, v.CurriculumId)
	}
	errs = a.SemesterRepo.DeleteCurriculumSemesterExcludingCurriculumId(ctx, tx, data.Id, curriculumIds)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	if len(curriculums) != 0 {
		errs = a.SemesterRepo.UpsertCurriculum(ctx, tx, curriculums)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (a semesterService) UpdateActivation(ctx context.Context, id string, isActive bool) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := a.SemesterRepo.GetById(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = a.SemesterRepo.UpdateActivation(ctx, tx, id, isActive)
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

func (a semesterService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := a.SemesterRepo.GetById(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = a.SemesterRepo.Delete(ctx, tx, id)
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
