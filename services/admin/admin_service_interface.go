package admin

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type AdminServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.AdminListWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateAdmin) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateAdmin) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewAdminService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) AdminServiceInterface {
	return &adminService{
		repoCtx,
		infraCtx,
	}
}
