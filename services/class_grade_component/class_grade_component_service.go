package class_grade_component

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	appConstants "github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/utils"
)

type classGradeComponentService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (s classGradeComponentService) GetList(ctx context.Context, classId string) ([]objects.GetClassGradeComponent, *constants.ErrorResponse) {
	results := []objects.GetClassGradeComponent{}

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return results, constants.ErrUnknown
	}

	resultData, errs := s.ClassGradeComponentRepo.GetByClassId(ctx, tx, classId)
	if errs != nil {
		_ = tx.Rollback()
		return results, errs
	}

	for _, v := range resultData {
		results = append(results, objects.GetClassGradeComponent{
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

func (s classGradeComponentService) Set(ctx context.Context, classId string, data []objects.SetClassGradeComponent) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	if claims.Role == appConstants.AppTypeLecturer {
		isActive := true
		lecturerClassData, errs := s.LecturerRepo.GetAssignedClass(ctx, tx, claims.ID, "", classId, &isActive)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
		if len(lecturerClassData) == 0 {
			_ = tx.Rollback()
			return appConstants.ErrLecturerNotAssigned
		}
	}

	_, errs = s.ClassRepo.GetDetail(ctx, tx, classId, "")
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	if len(data) == 0 {
		errs = s.ClassGradeComponentRepo.DeleteAllByClassId(ctx, tx, classId)
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

	gradeComponentClassNames := []string{}
	upsertData := []models.CreateClassGradeComponent{}
	for _, v := range data {
		gradeComponentClassNames = append(gradeComponentClassNames, v.Name)
		upsertData = append(upsertData, models.CreateClassGradeComponent{
			ClassId:    classId,
			Name:       v.Name,
			Percentage: v.Percentage,
			IsActive:   v.IsActive,
		})
	}

	errs = s.ClassGradeComponentRepo.DeleteAllByClassIdExcludingNames(ctx, tx, classId, gradeComponentClassNames)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = s.ClassGradeComponentRepo.Upsert(ctx, tx, upsertData)
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
