package exam_supervisor_role

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type ExamSupervisorRoleServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.ExamSupervisorRoleListWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateExamSupervisorRole) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateExamSupervisorRole) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewExamSupervisorRoleService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) ExamSupervisorRoleServiceInterface {
	return &examSupervisorRoleService{
		repoCtx,
		infraCtx,
	}
}
