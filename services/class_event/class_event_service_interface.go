package class_event

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type ClassEventServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest, classId, frequency string, futureEventOnly bool, isActive *bool) (objects.ClassEventListWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateClassEvent) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateClassEvent) *constants.ErrorResponse
	BulkUpdateActivation(ctx context.Context, ids []string, isActive bool) *constants.ErrorResponse
	BulkDelete(ctx context.Context, ids []string) *constants.ErrorResponse
}

func NewClassEventService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) ClassEventServiceInterface {
	return &classEventService{
		repoCtx,
		infraCtx,
	}
}
