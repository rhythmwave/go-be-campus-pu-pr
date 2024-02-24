package student_leave

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type StudentLeaveServiceInterface interface {
	GetListRequest(ctx context.Context, paginationData common.PaginationRequest, appType, studyProgramId string, isApproved *bool) (objects.StudentLeaveRequestListWithPagination, *constants.ErrorResponse)
	GetList(ctx context.Context, paginationData common.PaginationRequest, studyProgramId, semesterId string) (objects.StudentLeaveListWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateStudentLeave) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateStudentLeave) *constants.ErrorResponse
	Approve(ctx context.Context, id string, isApproved bool) *constants.ErrorResponse
	End(ctx context.Context, id string) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewStudentLeaveService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) StudentLeaveServiceInterface {
	return &studentLeaveService{
		repoCtx,
		infraCtx,
	}
}
