package yudicium

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

type yudiciumService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (y yudiciumService) Apply(ctx context.Context, data objects.ApplyYudicium) *constants.ErrorResponse {
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

	createData := models.CreateYudiciumStudent{
		StudentId:       data.StudentId,
		ApplicationDate: data.ApplicationDate,
		WithThesis:      data.WithThesis,
	}
	errs = y.YudiciumStudentRepo.Create(ctx, tx, createData)
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

func (y yudiciumService) GetListStudent(ctx context.Context, pagination common.PaginationRequest, req objects.GetListStudentYudiciumRequest) (objects.GetListStudentYudiciumWithPagination, *constants.ErrorResponse) {
	var result objects.GetListStudentYudiciumWithPagination

	tx, err := y.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := y.YudiciumStudentRepo.GetList(ctx, tx, pagination, req)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetListStudentYudicium{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetListStudentYudicium{
			Id:                    v.Id,
			NimNumber:             v.NimNumber,
			Name:                  v.Name,
			DiktiStudyProgramCode: v.DiktiStudyProgramCode,
			StudyProgramName:      v.StudyProgramName,
			StudyLevelShortName:   v.StudyLevelShortName,
			DiktiStudyProgramType: v.DiktiStudyProgramType,
			TotalCredit:           v.TotalCredit,
			Gpa:                   v.Gpa,
			Status:                v.Status,
			ApplicationDate:       v.ApplicationDate,
			DoneYudicium:          v.DoneYudicium,
		})
	}

	result = objects.GetListStudentYudiciumWithPagination{
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

func (y yudiciumService) Do(ctx context.Context, data objects.DoYudicium) *constants.ErrorResponse {
	tx, err := y.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	doData := models.DoYudicium{
		Id:             data.YudiciumSessionId,
		YudiciumNumber: data.YudiciumNumber,
		ActualDate:     data.YudiciumDate,
	}
	errs := y.YudiciumSessionRepo.Do(ctx, tx, doData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = y.YudiciumStudentRepo.Do(ctx, tx, data.YudiciumSessionId, data.StudentIds)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateStatusData := models.BulkUpdateStatusStudent{
		Ids:    data.StudentIds,
		Status: appConstants.StudentStatusLulus,
	}
	errs = y.StudentRepo.BulkUpdateStatus(ctx, tx, updateStatusData)
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
