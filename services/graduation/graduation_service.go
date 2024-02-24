package graduation

import (
	"context"
	"time"

	"github.com/sccicitb/pupr-backend/constants"
	appConstants "github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
)

type graduationService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (y graduationService) Apply(ctx context.Context, data objects.ApplyGraduation) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := y.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	if claims.Role == appConstants.AppTypeStudent {
		data.StudentId = claims.ID
		data.ApplicationDate = time.Now()
	}

	createData := models.CreateGraduationStudent{
		StudentId:           data.StudentId,
		ApplicationDate:     data.ApplicationDate,
		GraduationSessionId: data.GraduationSessionId,
	}
	errs = y.GraduationStudentRepo.Create(ctx, tx, createData)
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

func (y graduationService) GetListStudent(ctx context.Context, pagination common.PaginationRequest, studyProgramId, graduationSessionId string) (objects.GetListStudentGraduationWithPagination, *constants.ErrorResponse) {
	var result objects.GetListStudentGraduationWithPagination

	tx, err := y.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := y.GraduationStudentRepo.GetList(ctx, tx, pagination, studyProgramId, graduationSessionId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetListStudentGraduation{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetListStudentGraduation{
			Id:                    v.Id,
			NimNumber:             v.NimNumber,
			Name:                  v.Name,
			DiktiStudyProgramCode: v.DiktiStudyProgramCode,
			StudyProgramName:      v.StudyProgramName,
			StudyLevelShortName:   v.StudyLevelShortName,
			DiktiStudyProgramType: v.DiktiStudyProgramType,
			ApplicationDate:       v.ApplicationDate,
		})
	}

	result = objects.GetListStudentGraduationWithPagination{
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
