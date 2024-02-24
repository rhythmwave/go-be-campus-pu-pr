package shared_file

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type SharedFileServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest, appType string, onlyOwned bool, isApproved *bool) (objects.SharedFileListWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateSharedFile) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateSharedFile) *constants.ErrorResponse
	Approve(ctx context.Context, id string) *constants.ErrorResponse
	Delete(ctx context.Context, id, appType string) *constants.ErrorResponse
}

func NewSharedFileService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) SharedFileServiceInterface {
	return &sharedFileService{
		repoCtx,
		infraCtx,
	}
}
