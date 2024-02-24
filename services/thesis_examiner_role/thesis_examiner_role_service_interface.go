package thesis_examiner_role

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type ThesisExaminerRoleServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.ThesisExaminerRoleListWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateThesisExaminerRole) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateThesisExaminerRole) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewThesisExaminerRoleService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) ThesisExaminerRoleServiceInterface {
	return &thesisExaminerRoleService{
		repoCtx,
		infraCtx,
	}
}
