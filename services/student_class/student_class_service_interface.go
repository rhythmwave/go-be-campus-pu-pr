package student_class

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type StudentClassServiceInterface interface {
	GetList(ctx context.Context, pagination common.PaginationRequest, studyPlanId, studentId, semesterId, appType string, isMbkm *bool) (objects.StudentClassListWithPagination, *constants.ErrorResponse)
	TransferStudentClass(ctx context.Context, data objects.TransferStudentClass) *constants.ErrorResponse
	ReshuffleStudentClass(ctx context.Context, data []objects.ReshuffleStudentClass) *constants.ErrorResponse
	MergeStudentClass(ctx context.Context, data objects.MergeStudentClass) *constants.ErrorResponse
	BulkGradeStudentClass(ctx context.Context, classId string, data []objects.GradeStudentClass) *constants.ErrorResponse
}

func NewStudentClassService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) StudentClassServiceInterface {
	return &studentClassService{
		repoCtx,
		infraCtx,
	}
}
