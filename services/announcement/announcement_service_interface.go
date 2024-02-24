package announcement

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type AnnouncementServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest, appType string, announcementType string) (objects.AnnouncementListWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateAnnouncement) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateAnnouncement) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewAnnouncementService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) AnnouncementServiceInterface {
	return &announcementService{
		repoCtx,
		infraCtx,
	}
}
