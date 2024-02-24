package subject_grade_component

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/utils"
)

type subjectGradeComponentService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (s subjectGradeComponentService) GetList(ctx context.Context, subjectId string) ([]objects.GetSubjectGradeComponent, *constants.ErrorResponse) {
	results := []objects.GetSubjectGradeComponent{}

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return results, constants.ErrUnknown
	}

	resultData, errs := s.SubjectGradeComponentRepo.GetBySubjectId(ctx, tx, subjectId)
	if errs != nil {
		_ = tx.Rollback()
		return results, errs
	}

	for _, v := range resultData {
		results = append(results, objects.GetSubjectGradeComponent{
			Id:         v.Id,
			Name:       v.Name,
			Percentage: v.Percentage,
			IsActive:   v.IsActive,
		})
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return results, constants.ErrUnknown
	}

	return results, nil
}

func (s subjectGradeComponentService) Set(ctx context.Context, subjectId string, data []objects.SetSubjectGradeComponent) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs = s.SubjectRepo.GetDetail(ctx, tx, subjectId)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	if len(data) == 0 {
		errs = s.SubjectGradeComponentRepo.DeleteAllBySubjectId(ctx, tx, subjectId)
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

	gradeComponentSubjectNames := []string{}
	upsertData := []models.CreateSubjectGradeComponent{}
	for _, v := range data {
		gradeComponentSubjectNames = append(gradeComponentSubjectNames, v.Name)
		upsertData = append(upsertData, models.CreateSubjectGradeComponent{
			SubjectId:  subjectId,
			Name:       v.Name,
			Percentage: v.Percentage,
			IsActive:   v.IsActive,
			CreatedBy:  claims.ID,
		})
	}

	errs = s.SubjectGradeComponentRepo.DeleteAllBySubjectIdExcludingNames(ctx, tx, subjectId, gradeComponentSubjectNames)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = s.SubjectGradeComponentRepo.Upsert(ctx, tx, upsertData)
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
