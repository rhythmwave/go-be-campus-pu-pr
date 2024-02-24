package dikti_study_program

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type DiktiStudyProgramServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.DiktiStudyProgramListWithPagination, *constants.ErrorResponse)
}

func NewDiktiStudyProgramService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) DiktiStudyProgramServiceInterface {
	return &diktiStudyProgramService{
		repoCtx,
		infraCtx,
	}
}
