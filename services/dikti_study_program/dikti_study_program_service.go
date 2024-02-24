package dikti_study_program

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type diktiStudyProgramService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (d diktiStudyProgramService) GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.DiktiStudyProgramListWithPagination, *constants.ErrorResponse) {
	var result objects.DiktiStudyProgramListWithPagination

	tx, err := d.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := d.DiktiStudyProgramRepo.GetList(ctx, tx, paginationData)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetDiktiStudyProgram{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetDiktiStudyProgram{
			Id:                  v.Id,
			Code:                v.Code,
			Name:                v.Name,
			StudyLevelShortName: v.StudyLevelShortName,
			StudyLevelName:      v.StudyLevelName,
		})
	}

	result = objects.DiktiStudyProgramListWithPagination{
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
