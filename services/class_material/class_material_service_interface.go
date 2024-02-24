package class_material

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type ClassMaterialServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest, classId string, isActive *bool) (objects.ClassMaterialListWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateClassMaterial) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateClassMaterial) *constants.ErrorResponse
	BulkUpdateActivation(ctx context.Context, ids []string, isActive bool) *constants.ErrorResponse
	BulkDelete(ctx context.Context, ids []string) *constants.ErrorResponse
}

func NewClassMaterialService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) ClassMaterialServiceInterface {
	return &classMaterialService{
		repoCtx,
		infraCtx,
	}
}
