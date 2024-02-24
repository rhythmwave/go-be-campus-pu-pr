package document_type

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type DocumentTypeServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.DocumentTypeListWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateDocumentType) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateDocumentType) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewDocumentTypeService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) DocumentTypeServiceInterface {
	return &documentTypeService{
		repoCtx,
		infraCtx,
	}
}
