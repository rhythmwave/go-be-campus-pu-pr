package class_discussion

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type ClassDiscussionServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest, classId string) (objects.ClassDiscussionListWithPagination, *constants.ErrorResponse)
	GetComment(ctx context.Context, paginationData common.PaginationRequest, classDiscussionId string) (objects.ClassDiscussionCommentWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateClassDiscussion) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateClassDiscussion) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
	CreateComment(ctx context.Context, data objects.CreateClassDiscussionComment) *constants.ErrorResponse
	DeleteComment(ctx context.Context, id string) *constants.ErrorResponse
}

func NewClassDiscussionService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) ClassDiscussionServiceInterface {
	return &classDiscussionService{
		repoCtx,
		infraCtx,
	}
}
