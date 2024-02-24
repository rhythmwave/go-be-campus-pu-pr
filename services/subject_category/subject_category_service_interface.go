package subject_category

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type SubjectCategoryServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.SubjectCategoryListWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateSubjectCategory) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateSubjectCategory) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewSubjectCategoryService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) SubjectCategoryServiceInterface {
	return &subjectCategoryService{
		repoCtx,
		infraCtx,
	}
}
