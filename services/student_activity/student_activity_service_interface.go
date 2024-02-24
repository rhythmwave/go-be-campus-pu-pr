package student_activity

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type StudentActivityServiceInterface interface {
	GetList(ctx context.Context, pagination common.PaginationRequest, activityType, studyProgramId, semesterId string, isMbkm bool) (objects.StudentActivityListWithPagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, id string) (objects.GetStudentActivityDetail, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateStudentActivity) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateStudentActivity) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewStudentActivityService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) StudentActivityServiceInterface {
	return &studentActivityService{
		repoCtx,
		infraCtx,
	}
}
