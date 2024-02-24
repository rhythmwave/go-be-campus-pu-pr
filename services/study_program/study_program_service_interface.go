package study_program

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type StudyProgramServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest, majorId string, withAccessToken bool) (objects.StudyProgramListWithPagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, id string) (objects.GetStudyProgramDetail, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateStudyProgram) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateStudyProgram) *constants.ErrorResponse
	UpdateDegree(ctx context.Context, data objects.UpdateDegreeStudyProgram) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewStudyProgramService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) StudyProgramServiceInterface {
	return &studyProgramService{
		repoCtx,
		infraCtx,
	}
}
