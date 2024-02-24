package student_skpi

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type StudentSkpiServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest, paramsData objects.GetStudentSkpiRequest) (objects.StudentSkpiListWithPagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, id string) (objects.GetStudentSkpiDetail, *constants.ErrorResponse)
	Upsert(ctx context.Context, data objects.UpsertStudentSkpi) *constants.ErrorResponse
	Approve(ctx context.Context, data objects.ApproveStudentSkpi) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewStudentSkpiService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) StudentSkpiServiceInterface {
	return &studentSkpiService{
		repoCtx,
		infraCtx,
	}
}
