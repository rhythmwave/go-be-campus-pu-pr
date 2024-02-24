package document_action

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type DocumentActionServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.DocumentActionListWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateDocumentAction) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateDocumentAction) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewDocumentActionService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) DocumentActionServiceInterface {
	return &documentActionService{
		repoCtx,
		infraCtx,
	}
}
