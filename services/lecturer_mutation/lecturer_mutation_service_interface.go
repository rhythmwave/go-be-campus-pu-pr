package lecturer_mutation

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type LecturerMutationServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest, studyProgramId, idNationalLecturer, semesterId string) (objects.LecturerMutationListWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateLecturerMutation) *constants.ErrorResponse
}

func NewLecturerMutationService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) LecturerMutationServiceInterface {
	return &lecturerMutationService{
		repoCtx,
		infraCtx,
	}
}
