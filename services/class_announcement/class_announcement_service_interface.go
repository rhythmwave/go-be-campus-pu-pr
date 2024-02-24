package class_announcement

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type ClassAnnouncementServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest, appType string, classIds []string) (objects.ClassAnnouncementListWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateClassAnnouncement) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateClassAnnouncement) *constants.ErrorResponse
	Delete(ctx context.Context, ids []string) *constants.ErrorResponse
}

func NewClassAnnouncementService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) ClassAnnouncementServiceInterface {
	return &classAnnouncementService{
		repoCtx,
		infraCtx,
	}
}
