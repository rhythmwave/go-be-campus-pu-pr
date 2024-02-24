package study_level

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type StudyLevelServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.StudyLevelListWithPagination, *constants.ErrorResponse)
	UpdateSkpi(ctx context.Context, data objects.UpdateStudyLevelSkpi) *constants.ErrorResponse
}

func NewStudyLevelService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) StudyLevelServiceInterface {
	return &studyLevelService{
		repoCtx,
		infraCtx,
	}
}
