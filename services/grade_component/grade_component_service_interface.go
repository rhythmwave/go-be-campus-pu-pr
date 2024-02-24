package grade_component

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type GradeComponentServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest, studyProgramId, SubjectCategoryId string) (objects.GradeComponentListWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateGradeComponent) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateGradeComponent) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
	GetListBySubjectCategory(ctx context.Context, paginationData common.PaginationRequest, studyProgramId string) (objects.GradeComponentBySubjectCategoryListWithPagination, *constants.ErrorResponse)
	BulkUpdatePercentage(ctx context.Context, baseData objects.BulkUpdatePercentageGradeComponent, data []objects.BulkUpdatePercentageGradeComponentData) *constants.ErrorResponse
}

func NewGradeComponentService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) GradeComponentServiceInterface {
	return &gradeComponentService{
		repoCtx,
		infraCtx,
	}
}
