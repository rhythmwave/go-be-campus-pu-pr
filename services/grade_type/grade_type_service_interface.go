package grade_type

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type GradeTypeServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest, studyLevelId string) (objects.GradeTypeListWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateGradeType) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateGradeType) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewGradeTypeService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) GradeTypeServiceInterface {
	return &gradeTypeService{
		repoCtx,
		infraCtx,
	}
}
