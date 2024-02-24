package faculty

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type FacultyServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.FacultyListWithPagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, id string) (objects.GetFacultyDetail, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateFaculty) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateFaculty) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewFacultyService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) FacultyServiceInterface {
	return &facultyService{
		repoCtx,
		infraCtx,
	}
}
