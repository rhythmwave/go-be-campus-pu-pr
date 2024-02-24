package lecturer_leave

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type LecturerLeaveServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest, studyProgramId, idNationalLecturer, semesterId string, isActive bool) (objects.LecturerLeaveListWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateLecturerLeave) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateLecturerLeave) *constants.ErrorResponse
	End(ctx context.Context, id string) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewLecturerLeaveService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) LecturerLeaveServiceInterface {
	return &lecturerLeaveService{
		repoCtx,
		infraCtx,
	}
}
