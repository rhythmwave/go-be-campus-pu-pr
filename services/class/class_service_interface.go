package class

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type ClassServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest, in objects.GetClassListRequest) (objects.ClassListWithPagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, id string) (objects.GetClassDetail, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateClass) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateClass) *constants.ErrorResponse
	UpdateActivation(ctx context.Context, id string, isActive bool) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
	Duplicate(ctx context.Context, fromSemesterId, toSemesterId string) *constants.ErrorResponse
	BulkUpdateMaximumParticipant(ctx context.Context, data []objects.UpdateClassMaximumParticipant) *constants.ErrorResponse
	GetClassParticipantList(ctx context.Context, paginationData common.PaginationRequest, classId, lectureId string, isGraded *bool, studentId string) (objects.ClassParticipantWithPagination, *constants.ErrorResponse)
}

func NewClassService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) ClassServiceInterface {
	return &classService{
		repoCtx,
		infraCtx,
	}
}
