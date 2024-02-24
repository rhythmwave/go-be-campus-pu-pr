package major

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type MajorServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest, facultyId string) (objects.MajorListWithPagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, id string) (objects.GetMajorDetail, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateMajor) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateMajor) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewMajorService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) MajorServiceInterface {
	return &majorService{
		repoCtx,
		infraCtx,
	}
}
