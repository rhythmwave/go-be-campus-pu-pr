package lecturer_resignation

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type LecturerResignationServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest, studyProgramId, idNationalLecturer, semesterId string) (objects.LecturerResignationListWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateLecturerResignation) *constants.ErrorResponse
}

func NewLecturerResignationService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) LecturerResignationServiceInterface {
	return &lecturerResignationService{
		repoCtx,
		infraCtx,
	}
}
